package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"time"

	"github.com/ONSdigital/dp-frontend-filter-dataset-controller/data"
	"github.com/ONSdigital/dp-frontend-filter-dataset-controller/dates"
	"github.com/ONSdigital/dp-frontend-filter-dataset-controller/mapper"
	"github.com/ONSdigital/dp-frontend-filter-dataset-controller/renderer"
	"github.com/ONSdigital/dp-frontend-models/model"
	"github.com/ONSdigital/dp-frontend-models/model/dataset-filter/geography"
	"github.com/ONSdigital/dp-frontend-models/model/dataset-filter/listSelector"
	"github.com/ONSdigital/go-ns/log"
	"github.com/gorilla/mux"
)

// Filter represents the handlers for Filtering
type Filter struct {
	r   renderer.Renderer
	fc  FilterClient
	dc  DatasetClient
	clc CodelistClient
	val Validator
}

// NewFilter creates a new instance of Filter
func NewFilter(r renderer.Renderer, fc FilterClient, dc DatasetClient, clc CodelistClient, val Validator) *Filter {
	return &Filter{
		r:   r,
		fc:  fc,
		dc:  dc,
		clc: clc,
		val: val,
	}
}

func getStubbedMetadataFooter() model.Footer {
	return model.Footer{
		Enabled:     true,
		Contact:     "Matt Rout",
		ReleaseDate: "11 November 2016",
		NextRelease: "11 November 2017",
		DatasetID:   "MR",
	}
}

// PreviewPage controls the rendering of the preview and download page
func (f *Filter) PreviewPage(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	dimensions := []data.Dimension{
		{
			Name:   "Geography",
			Values: []string{"Wales", "Cardiff", "Caerphilly", "Newport", "Pontypridd", "Merthyr Tydfil"},
		},
		{
			Name:   "Population",
			Values: []string{"30000000", "284384", "37238", "428219", "32161", "281994"},
		},
	}

	dataset := data.Dataset{
		ID:          "849209",
		ReleaseDate: "17 January 2017",
		NextRelease: "17 February 2017",
		Contact: data.Contact{
			Name:      "Matt Rout",
			Telephone: "07984593234",
			Email:     "matt@gmail.com",
		},
		Title: "Small Area Population Estimates",
	}

	filter := data.Filter{
		FilterID: vars["filterID"],
		Edition:  "12345",
		Dataset:  "849209",
		Version:  "2017",
		Downloads: map[string]data.Download{
			"csv": {
				Size: "362783",
				URL:  "/",
			},
			"xls": {
				Size: "373929",
				URL:  "/",
			},
		},
	}

	p := mapper.CreatePreviewPage(dimensions, filter, dataset, vars["filterID"])

	body, err := json.Marshal(p)
	if err != nil {
		log.ErrorR(req, err, nil)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	b, err := f.r.Do("dataset-filter/preview-page", body)
	if err != nil {
		log.ErrorR(req, err, nil)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if _, err := w.Write(b); err != nil {
		log.ErrorR(req, err, nil)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// Geography ...
func (f *Filter) Geography(w http.ResponseWriter, req *http.Request) {
	p := geography.Page{
		FilterID: "12345",
		Data: geography.Geography{
			SaveAndReturn: geography.Link{
				URL: "/filters/12345/dimensions",
			},
			Cancel: geography.Link{
				URL: "/filters/12345/dimensions",
			},
			FiltersAmount: 2,
			FiltersAdded: []geography.Filter{
				{
					RemoveURL: "/remove-this/",
					Label:     "All ages",
				},
				{
					RemoveURL: "/remove-this-2/",
					Label:     "43",
				},
				{
					RemoveURL: "/remove-this-3/",
					Label:     "18",
				},
			},
			FilterList: []geography.List{
				{
					Location: "United Kingdom",
				},
				{
					Location: "England",
					SubNum:   10,
					SubType:  "Regions",
					SubURL:   "/regions/",
				},
				{
					Location: "Wales",
					SubNum:   5,
					SubType:  "Regions",
					SubURL:   "/regions/",
				},
			},
			RemoveAll: geography.Link{
				URL: "/remove-all/",
			},
			AddAllFilters: geography.AddAll{
				URL:    "/add-all/",
				Amount: 3,
			},
			GoBack: geography.Link{
				URL: "/back/",
			},
			Parent: "Wales: Counties",
		},
	}

	p.Breadcrumb = []model.TaxonomyNode{
		{
			Title: "Title of dataset",
			URI:   "/",
		},
		{
			Title: "Filter this dataset",
			URI:   "/",
		},
	}

	p.SearchDisabled = true

	p.Metadata.Footer = getStubbedMetadataFooter()

	b, err := json.Marshal(p)
	if err != nil {
		log.Error(err, nil)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	templateBytes, err := f.r.Do("dataset-filter/geography", b)
	if err != nil {
		log.Error(err, nil)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(templateBytes)
}

// FilterOverview controls the render of the filter overview template
// Contains stubbed data for now - page to be populated by the API
func (f *Filter) FilterOverview(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	filterID := vars["filterID"]

	dims, err := f.fc.GetDimensions(filterID)
	if err != nil {
		log.ErrorR(req, err, nil)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var dimensions []data.Dimension
	for _, dim := range dims {
		var vals data.DimensionValues
		vals, err = f.fc.GetDimensionOptions(filterID, dim.Name)
		if err != nil {
			log.ErrorR(req, err, nil)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		var values []string
		for _, val := range vals.Items {
			values = append(values, val.Name)
		}

		dimensions = append(dimensions, data.Dimension{
			Name:   dim.Name,
			Values: values,
		})
	}

	filter, err := f.fc.GetJobState(filterID)
	if err != nil {
		log.ErrorR(req, err, nil)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	dataset, err := f.dc.GetDataset(filterID, filter.Edition, filter.Version)
	if err != nil {
		log.ErrorR(req, err, nil)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	p := mapper.CreateFilterOverview(dimensions, filter, dataset, filterID)

	b, err := json.Marshal(p)
	if err != nil {
		log.Error(err, nil)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	templateBytes, err := f.r.Do("dataset-filter/filter-overview", b)
	if err != nil {
		log.Error(err, nil)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(templateBytes)
}

// DimensionSelector controls the render of the range selector template
// Contains stubbed data for now - page to be populated by the API
func (f *Filter) DimensionSelector(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	name := vars["name"]
	filterID := vars["filterID"]

	filter, err := f.fc.GetJobState(filterID)
	if err != nil {
		log.ErrorR(req, err, nil)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	selectedValues, err := f.fc.GetDimensionOptions(filterID, name)
	if err != nil {
		log.ErrorR(req, err, nil)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	dataset, err := f.dc.GetDataset(filterID, filter.Edition, filter.Version)
	if err != nil {
		log.ErrorR(req, err, nil)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	dim, err := f.fc.GetDimension(filterID, name)
	if err != nil {
		log.ErrorR(req, err, nil)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	codeID := getCodeIDFromURI(dim.URI)
	if codeID == "" {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	allValues, err := f.clc.GetValues(codeID)
	if err != nil {
		log.ErrorR(req, err, nil)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	selectorType := req.URL.Query().Get("selectorType")
	if selectorType == "list" {
		f.listSelector(w, req)
	} else {
		f.rangeSelector(w, req, name, selectedValues, allValues, filter, dataset)
	}
}

func (f *Filter) rangeSelector(w http.ResponseWriter, req *http.Request, name string, selectedValues, allValues data.DimensionValues, filter data.Filter, dataset data.Dataset) {
	ns := req.URL.Query().Get("nSelectors")

	var nSelectors int
	var err error
	if ns == "" {
		nSelectors = 1
	} else {
		nSelectors, err = strconv.Atoi(ns)
		if err != nil {
			nSelectors = 1
		}
	}

	p := mapper.CreateRangeSelectorPage(name, selectedValues, allValues, filter, dataset, nSelectors)

	b, err := json.Marshal(p)
	if err != nil {
		log.Error(err, nil)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	templateBytes, err := f.r.Do("dataset-filter/range-selector", b)
	if err != nil {
		log.Error(err, nil)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(templateBytes)
}

// ListSelector controls the render of the age selector list template
// Contains stubbed data for now - page to be populated by the API
func (f *Filter) listSelector(w http.ResponseWriter, req *http.Request) {
	p := listSelector.Page{
		FilterID: "12345",
		Data: listSelector.ListSelector{
			AddFromRange: listSelector.Link{
				Label: "add age range",
				URL:   "/filters/12345/dimensions/age-range",
			},
			SaveAndReturn: listSelector.Link{
				URL: "/filters/12345/dimensions",
			},
			AddAllInRange: listSelector.Link{
				Label: "All ages",
			},
			Cancel: listSelector.Link{
				URL: "/filters/12345/dimensions",
			},
			FiltersAdded: []listSelector.Filter{
				{
					RemoveURL: "/remove-this/",
					Label:     "All ages",
				},
				{
					RemoveURL: "/remove-this-2/",
					Label:     "43",
				},
				{
					RemoveURL: "/remove-this-3/",
					Label:     "18",
				},
			},
			RemoveAll: listSelector.Link{
				URL: "/remove-all/",
			},
			FiltersAmount: 2,
			RangeData: listSelector.Range{
				StartNum:     30,
				EndNum:       90,
				StartLabel:   "Youngest",
				EndLabel:     "Oldest",
				AppendString: "and over",
			},
		},
	}

	p.Breadcrumb = []model.TaxonomyNode{
		{
			Title: "Title of dataset",
			URI:   "/",
		},
		{
			Title: "Filter this dataset",
			URI:   "/",
		},
	}

	p.SearchDisabled = true

	p.Metadata.Footer = getStubbedMetadataFooter()

	b, err := json.Marshal(p)
	if err != nil {
		log.Error(err, nil)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	templateBytes, err := f.r.Do("dataset-filter/list-selector", b)
	if err != nil {
		log.Error(err, nil)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(templateBytes)
}

type Range struct {
	Start string `schema:"start"`
	End   string `schema:"end"`
}

// AddRange will add a range of values to a filter job
func (f *Filter) AddRange(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	name := vars["name"]
	ns := req.URL.Query().Get("nSelectors")
	filterID := vars["filterID"]

	var r Range

	if err := f.val.Validate(req, &r); err != nil {
		log.ErrorR(req, err, nil)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var nSelectors int
	var err error
	if ns == "" {
		nSelectors = 1
	} else {
		nSelectors, err = strconv.Atoi(ns)
		if err != nil {
			nSelectors = 1
		}
	}

	dim, err := f.fc.GetDimension(filterID, name)
	if err != nil {
		log.ErrorR(req, err, nil)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	codeID := getCodeIDFromURI(dim.URI)
	if codeID == "" {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	allValues, err := f.clc.GetValues(codeID)
	if err != nil {
		log.ErrorR(req, err, nil)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var values []string
	labelIDMap := make(map[string]string)
	for _, val := range allValues.Items {
		values = append(values, val.Label)
		labelIDMap[val.Label] = val.ID
	}

	if name == "month" {
		dats, err := dates.ConvertToReadable(values)
		if err != nil {
			log.ErrorR(req, err, nil)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		start, err := time.Parse("01 January 2006", fmt.Sprintf("01 %s", r.Start))
		if err != nil {
			log.ErrorR(req, err, nil)
			w.WriteHeader(http.StatusInternalServerError)
			redirectURL := fmt.Sprintf("/filters/%s/dimensions/%s?nSelectors=%d", filterID, name, nSelectors)
			http.Redirect(w, req, redirectURL, 301)
		}

		end, err := time.Parse("01 January 2006", fmt.Sprintf("01 %s", r.End))
		if err != nil {
			log.ErrorR(req, err, nil)
			w.WriteHeader(http.StatusInternalServerError)
			redirectURL := fmt.Sprintf("/filters/%s/dimensions/%s?nSelectors=%d", filterID, name, nSelectors)
			http.Redirect(w, req, redirectURL, 301)
		}

		if end.Before(start) {
			log.Info("end date before start date", log.Data{"start": start, "end": end})
			w.WriteHeader(http.StatusInternalServerError)
			redirectURL := fmt.Sprintf("/filters/%s/dimensions/%s?nSelectors=%d", filterID, name, nSelectors)
			http.Redirect(w, req, redirectURL, 301)
		}

		dats = dates.Sort(dats)
		values = dates.ConvertToCoded(dats)
		for i, dat := range dats {
			if dat.Equal(start) || dat.After(start) && dat.Before(end) || dat.Equal(end) {
				if err := f.fc.AddDimensionValue(filterID, name, labelIDMap[values[i]]); err != nil {
					log.TraceR(req, err.Error(), nil)
					continue
				}
			}
		}
	}

	redirectURL := fmt.Sprintf("/filters/%s/dimensions/%s?nSelectors=%d", filterID, name, nSelectors+1)
	http.Redirect(w, req, redirectURL, 301)
}

func (f *Filter) RemoveRange(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	name := vars["name"]
	ns := req.URL.Query().Get("nSelectors")
	filterID := vars["filterID"]

	var nSelectors int
	var err error
	if ns == "" {
		nSelectors = 1
	} else {
		nSelectors, err = strconv.Atoi(ns)
		if err != nil {
			nSelectors = 1
		}
	}

	url := fmt.Sprintf("/filters/%s/dimensions/%s?nSelectors=%d", filterID, name, nSelectors-1)
	http.Redirect(w, req, url, 301)
}

func getCodeIDFromURI(uri string) string {
	codeReg := regexp.MustCompile(`^\/code-lists\/(.+)\/codes$`)
	subs := codeReg.FindStringSubmatch(uri)

	if len(subs) == 2 {
		return subs[1]
	}

	log.Info("could not extract codeID from uri", nil)
	return ""
}
