package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"sync"

	"github.com/ONSdigital/dp-api-clients-go/filter"
	"github.com/ONSdigital/dp-api-clients-go/headers"
	"github.com/ONSdigital/dp-api-clients-go/hierarchy"
	"github.com/ONSdigital/dp-frontend-filter-dataset-controller/helpers"
	"github.com/ONSdigital/dp-frontend-filter-dataset-controller/mapper"
	"github.com/ONSdigital/log.go/log"
	"github.com/gorilla/mux"
)

// these form vars are not regular input fields, but transmit meta form info
var specialFormVars = map[string]bool{
	"save-and-return": true,
	":uri":            true,
	"q":               true,
}

// HierarchyUpdate controls the updating of a hierarchy job
func (f *Filter) HierarchyUpdate(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	filterID := vars["filterID"]
	name := vars["name"]
	code := vars["code"]

	ctx := req.Context()
	collectionID := getCollectionIDFromContext(ctx)
	userAccessToken, err := headers.GetUserAuthToken(req)
	if err != nil {
		if headers.IsNotErrNotFound(err) {
			log.Event(ctx, "failed to retrieve auth header", log.Error(err))
		}
	}

	if err := req.ParseForm(); err != nil {
		log.Event(ctx, "failed to parse request", log.Error(err))
		return
	}

	var redirectURI string
	if len(req.Form["save-and-return"]) > 0 {
		redirectURI = fmt.Sprintf("/filters/%s/dimensions", filterID)
	} else {
		if len(code) > 0 {
			redirectURI = fmt.Sprintf("/filters/%s/dimensions/%s/%s", filterID, name, code)
		} else {
			redirectURI = fmt.Sprintf("/filters/%s/dimensions/%s", filterID, name)
		}
	}

	fil, err := f.FilterClient.GetJobState(req.Context(), userAccessToken, "", "", collectionID, filterID)
	if err != nil {
		log.Event(ctx, "failed to get job state", log.Error(err), log.Data{"filter_id": filterID})
		setStatusCode(req, w, err)
		return
	}

	if len(req.Form["add-all"]) > 0 {
		f.addAllHierarchyLevel(w, req, fil, name, code, redirectURI)
		return
	}

	if len(req.Form["remove-all"]) > 0 {
		f.removeAllHierarchyLevel(w, req, fil, name, code, redirectURI)
		return
	}

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()

		var h hierarchy.Model
		var err error
		if len(code) > 0 {
			h, err = f.HierarchyClient.GetChild(ctx, fil.InstanceID, name, code)
		} else {
			if name == "geography" {
				h, err = f.flattenGeographyTopLevel(ctx, fil.InstanceID)
			} else {
				h, err = f.HierarchyClient.GetRoot(ctx, fil.InstanceID, name)
			}

			// We include the value on the root as a selectable item, so append
			// the value on the root to the child to see if it has been removed by
			// the user
			h.Children = append(h.Children, hierarchy.Child{
				Links: h.Links,
			})
		}
		if err != nil {
			log.Event(ctx, "failed to get hierarchy node", log.Error(err), log.Data{"filter_id": filterID, "dimension": name, "code": code})
			setStatusCode(req, w, err)
			return
		}

		f.removeDimensionValues(ctx, req, userAccessToken, collectionID, filterID, name, h.Children)

	}()

	if formRedirectURI := f.hierarchyUpdateProcessRequestForm(req.Context(), userAccessToken, collectionID, filterID, name, req.Form, 3); formRedirectURI != "" {
		redirectURI = formRedirectURI
	}

	wg.Wait()
	http.Redirect(w, req, redirectURI, http.StatusFound)
}

// removeDimensionValues removes the dimension values that are present in the request form and match any child self link ID.
// It does it in O(NumChildren + NumOpts*log(NumChildren)*log(NumReqForm)) time.
func (f *Filter) removeDimensionValues(ctx context.Context, req *http.Request, userAccessToken, collectionID, filterID, name string, children []hierarchy.Child) {

	// Get dimension options from filterAPI
	opts, err := f.FilterClient.GetDimensionOptions(req.Context(), userAccessToken, "", collectionID, filterID, name)
	if err != nil {
		log.Event(ctx, "failed to get dimension options", log.Error(err))
	}

	// Create children map by self link ID
	childrenMap := map[string]*hierarchy.Child{}
	for _, child := range children {
		childrenMap[child.Links.Self.ID] = &child
	}

	// Iterate options, and remove any dimension value not provided in the form, and present in the children.
	for _, opt := range opts {
		if _, childFound := childrenMap[opt.Option]; childFound {
			if _, reqFound := req.Form[opt.Option]; !reqFound {
				if err := f.FilterClient.RemoveDimensionValue(req.Context(), userAccessToken, "", collectionID, filterID, name, opt.Option); err != nil {
					log.Event(ctx, "failed to remove dimension value", log.Error(err))
				}
			}
		}
	}
}

// hierarchyUpdateProcessRequestForm iterates the provided url.Values, and Adds the dimension values in batches. If a redirectURI is found, it will be returned.
func (f *Filter) hierarchyUpdateProcessRequestForm(ctx context.Context, userAccessToken, collectionID, filterID, name string, form url.Values, maxOptionsBatchSize int) (redirectURI string) {

	options := []string{}

	for k := range form {
		if _, foundSpecial := specialFormVars[k]; foundSpecial {
			continue
		}

		if strings.Contains(k, "redirect:") {
			redirectReg := regexp.MustCompile(`^redirect:(.+)$`)
			redirectSubs := redirectReg.FindStringSubmatch(k)
			redirectURI = redirectSubs[1]
			continue
		}

		options = append(options, k)

		// If the batch size is complete, send it and reset the batch
		if len(options) == maxOptionsBatchSize {
			if err := f.FilterClient.AddDimensionValues(ctx, userAccessToken, "", collectionID, filterID, name, options); err != nil {
				log.Event(ctx, "failed to add dimension values", log.Error(err))
			}
			options = []string{}
		}

	}

	// Send any remaining Dimension value
	if len(options) > 0 {
		if err := f.FilterClient.AddDimensionValues(ctx, userAccessToken, "", collectionID, filterID, name, options); err != nil {
			log.Event(ctx, "failed to add dimension values", log.Error(err))
		}
	}

	return redirectURI
}

func (f *Filter) addAllHierarchyLevel(w http.ResponseWriter, req *http.Request, fil filter.Model, name, code, redirectURI string) {

	ctx := req.Context()
	collectionID := getCollectionIDFromContext(ctx)
	userAccessToken, err := headers.GetUserAuthToken(req)
	if err != nil {
		if headers.IsNotErrNotFound(err) {
			log.Event(ctx, "failed to retrieve auth header", log.Error(err))
		}
	}

	var h hierarchy.Model
	if len(code) > 0 {
		h, err = f.HierarchyClient.GetChild(ctx, fil.InstanceID, name, code)
	} else {
		if name == "geography" {
			h, err = f.flattenGeographyTopLevel(ctx, fil.InstanceID)
		} else {
			h, err = f.HierarchyClient.GetRoot(ctx, fil.InstanceID, name)
		}
	}
	if err != nil {
		log.Event(ctx, "failed to get hierarchy node", log.Error(err), log.Data{"filter_id": fil.FilterID, "dimension": name, "code": code})
		setStatusCode(req, w, err)
		return
	}

	var options []string
	for _, child := range h.Children {
		options = append(options, child.Links.Self.ID)
	}
	if err := f.FilterClient.AddDimensionValues(req.Context(), userAccessToken, "", collectionID, fil.FilterID, name, options); err != nil {
		log.Event(ctx, "failed to add dimension values", log.Error(err))
	}

	http.Redirect(w, req, redirectURI, 302)
}

func (f *Filter) removeAllHierarchyLevel(w http.ResponseWriter, req *http.Request, fil filter.Model, name, code, redirectURI string) {
	ctx := req.Context()
	collectionID := getCollectionIDFromContext(ctx)
	userAccessToken, err := headers.GetUserAuthToken(req)
	if err != nil {
		if headers.IsNotErrNotFound(err) {
			log.Event(ctx, "failed to retrieve auth header", log.Error(err))
		}
	}

	var h hierarchy.Model

	if len(code) > 0 {
		h, err = f.HierarchyClient.GetChild(ctx, fil.InstanceID, name, code)
	} else {
		if name == "geography" {
			h, err = f.flattenGeographyTopLevel(ctx, fil.InstanceID)
		} else {
			h, err = f.HierarchyClient.GetRoot(ctx, fil.InstanceID, name)
		}
	}
	if err != nil {
		log.Event(ctx, "failed to get hierarchy node", log.Error(err), log.Data{"filter_id": fil.FilterID, "dimension": name, "code": code})
		setStatusCode(req, w, err)
		return
	}

	for _, child := range h.Children {
		if err := f.FilterClient.RemoveDimensionValue(req.Context(), userAccessToken, "", collectionID, fil.FilterID, name, child.Links.Self.ID); err != nil {
			log.Event(ctx, "failed to remove dimension value", log.Error(err))
		}
	}

	http.Redirect(w, req, redirectURI, 302)
}

func (f *Filter) Hierarchy(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	filterID := vars["filterID"]
	name := vars["name"]
	code := vars["code"]
	ctx := req.Context()

	collectionID := getCollectionIDFromContext(ctx)
	userAccessToken, err := headers.GetUserAuthToken(req)
	if err != nil {
		if headers.IsNotErrNotFound(err) {
			log.Event(ctx, "failed to retrieve auth header", log.Error(err))
		}
	}

	fil, err := f.FilterClient.GetJobState(req.Context(), userAccessToken, "", "", collectionID, filterID)
	if err != nil {
		log.Event(ctx, "failed to get job state", log.Error(err), log.Data{"filter_id": filterID})
		setStatusCode(req, w, err)
		return
	}

	var h hierarchy.Model
	if len(code) > 0 {
		h, err = f.HierarchyClient.GetChild(ctx, fil.InstanceID, name, code)
	} else {
		if name == "geography" {
			h, err = f.flattenGeographyTopLevel(ctx, fil.InstanceID)
		} else {
			h, err = f.HierarchyClient.GetRoot(ctx, fil.InstanceID, name)
		}
	}
	if err != nil {
		log.Event(ctx, "failed to get hierarchy node", log.Error(err), log.Data{"filter_id": filterID, "dimension": name, "code": code})
		setStatusCode(req, w, err)
		return
	}

	selVals, err := f.FilterClient.GetDimensionOptions(req.Context(), userAccessToken, "", collectionID, filterID, name)
	if err != nil {
		log.Event(ctx, "failed to get options from filter client", log.Error(err), log.Data{"filter_id": filterID, "dimension": name})
		setStatusCode(req, w, err)
		return
	}

	versionURL, err := url.Parse(fil.Links.Version.HRef)
	if err != nil {
		log.Event(ctx, "failed to parse version href", log.Error(err), log.Data{"filter_id": filterID})
		setStatusCode(req, w, err)
		return
	}
	datasetID, edition, version, err := helpers.ExtractDatasetInfoFromPath(ctx, versionURL.Path)
	if err != nil {
		log.Event(ctx, "failed to extract dataset info from path", log.Error(err), log.Data{"filter_id": filterID, "path": versionURL})
		setStatusCode(req, w, err)
		return
	}

	d, err := f.DatasetClient.Get(req.Context(), userAccessToken, "", collectionID, datasetID)
	if err != nil {
		log.Event(req.Context(), "failed to get dataset", log.Error(err), log.Data{"dataset_id": datasetID})
		setStatusCode(req, w, err)
		return
	}
	ver, err := f.DatasetClient.GetVersion(req.Context(), userAccessToken, "", "", collectionID, datasetID, edition, version)
	if err != nil {
		log.Event(req.Context(), "failed to get version", log.Error(err), log.Data{"dataset_id": datasetID, "edition": edition, "version": version})
		setStatusCode(req, w, err)
		return
	}

	allVals, err := f.DatasetClient.GetOptions(req.Context(), userAccessToken, "", collectionID, datasetID, edition, version, name)
	if err != nil {
		log.Event(ctx, "failed to get options from dataset client", log.Error(err),
			log.Data{"dataset_id": datasetID, "edition": edition, "version": version})
		setStatusCode(req, w, err)
		return
	}

	dims, err := f.DatasetClient.GetDimensions(req.Context(), userAccessToken, "", collectionID, datasetID, edition, version)
	if err != nil {
		log.Event(ctx, "failed to get dimensions", log.Error(err),
			log.Data{"dataset_id": datasetID, "edition": edition, "version": version})
		setStatusCode(req, w, err)
		return
	}

	p := mapper.CreateHierarchyPage(req, h, d, fil, selVals, allVals, dims, name, req.URL.Path, datasetID, ver.ReleaseDate)

	b, err := json.Marshal(p)
	if err != nil {
		log.Event(req.Context(), "failed to marshal json", log.Error(err), log.Data{"filter_id": filterID})
		setStatusCode(req, w, err)
		return
	}

	templateBytes, err := f.Renderer.Do("dataset-filter/hierarchy", b)
	if err != nil {
		log.Event(req.Context(), "failed to render", log.Error(err), log.Data{"filter_id": filterID})
		setStatusCode(req, w, err)
		return
	}

	w.Write(templateBytes)

}

type flatNodes struct {
	list  []*hierarchy.Child
	order []string
}

func (n flatNodes) addWithoutChildren(val hierarchy.Child, i int) {
	if !val.HasData {
		return
	}

	n.list[i] = &hierarchy.Child{
		Label:   val.Label,
		Links:   val.Links,
		HasData: val.HasData,
	}
}

func (n flatNodes) addWithChildren(val hierarchy.Child, i int) {
	n.list[i] = &hierarchy.Child{
		Label:            val.Label,
		Links:            val.Links,
		HasData:          val.HasData,
		NumberofChildren: val.NumberofChildren,
	}
}

// Flatten the geography hierarchy - please note this will only work for this particular hierarchy,
// need helper functions for other geog hierarchies too.
func (f *Filter) flattenGeographyTopLevel(ctx context.Context, instanceID string) (h hierarchy.Model, err error) {
	root, err := f.HierarchyClient.GetRoot(ctx, instanceID, "geography")
	if err != nil {
		return
	}

	if root.HasData {
		h.Label = root.Label
		h.Links = root.Links
		h.HasData = root.HasData
	}

	// Order: Great Britain, England and Wales, England, Northern Ireland, Scotland, Wales
	nodes := flatNodes{
		list:  make([]*hierarchy.Child, 6),
		order: []string{"K03000001", "K04000001", "E92000001", "N92000002", "S92000003", "W92000004"},
	}

	for _, val := range root.Children {
		if val.Links.Code.ID == nodes.order[0] {
			nodes.addWithoutChildren(val, 0)

			child, err := f.HierarchyClient.GetChild(ctx, instanceID, "geography", val.Links.Code.ID)
			if err != nil {
				return h, err
			}

			for _, childVal := range child.Children {

				if childVal.Links.Code.ID == nodes.order[1] {
					nodes.addWithoutChildren(childVal, 1)

					grandChild, err := f.HierarchyClient.GetChild(ctx, instanceID, "geography", childVal.Links.Code.ID)
					if err != nil {
						return h, err
					}

					for _, grandChildVal := range grandChild.Children {
						if grandChildVal.Links.Code.ID == nodes.order[2] {
							nodes.addWithChildren(grandChildVal, 2)
						}

						if grandChildVal.Links.Code.ID == nodes.order[5] {
							nodes.addWithChildren(grandChildVal, 5)
						}

					}
				}

				if childVal.Links.Code.ID == nodes.order[4] {
					nodes.addWithChildren(childVal, 4)
				}
			}
		}

		if val.Links.Code.ID == nodes.order[3] {
			nodes.addWithChildren(val, 3)
		}
	}

	//remove nil elements from list
	children := []hierarchy.Child{}
	for _, c := range nodes.list {
		if c != nil {
			children = append(children, *c)
		}
	}

	if len(children) == 0 {
		children = root.Children
	}

	h.Children = children
	return h, err
}
