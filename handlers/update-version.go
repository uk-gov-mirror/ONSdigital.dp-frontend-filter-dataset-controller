package handlers

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/ONSdigital/dp-frontend-filter-dataset-controller/helpers"
	"github.com/gorilla/mux"
)

// UseLatest will create a new filter job for the same dataset with the
// latest version
func (f *Filter) UseLatest(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	filterID := vars["filterID"]

	req = forwardFlorenceTokenIfRequired(req)

	oldJob, err := f.FilterClient.GetJobState(req.Context(), filterID)
	if err != nil {
		setStatusCode(req, w, err)
		return
	}

	dims, err := f.FilterClient.GetDimensions(req.Context(), filterID)
	if err != nil {
		setStatusCode(req, w, err)
		return
	}

	versionURL, err := url.Parse(oldJob.Links.Version.HRef)
	if err != nil {
		setStatusCode(req, w, err)
		return
	}
	datasetID, _, _, err := helpers.ExtractDatasetInfoFromPath(versionURL.Path)
	if err != nil {
		setStatusCode(req, w, err)
		return
	}

	dst, err := f.DatasetClient.Get(req.Context(), datasetID)
	if err != nil {
		setStatusCode(req, w, err)
		return
	}

	latestVersionURL, err := url.Parse(dst.Links.LatestVersion.URL)
	if err != nil {
		setStatusCode(req, w, err)
		return
	}
	_, edition, version, err := helpers.ExtractDatasetInfoFromPath(latestVersionURL.Path)
	if err != nil {
		setStatusCode(req, w, err)
		return
	}

	newFilterID, err := f.FilterClient.CreateBlueprint(req.Context(), datasetID, edition, version, []string{})
	if err != nil {
		setStatusCode(req, w, err)
		return
	}

	for _, dim := range dims {
		if err := f.FilterClient.AddDimension(req.Context(), newFilterID, dim.Name); err != nil {
			setStatusCode(req, w, err)
			return
		}

		dimValues, err := f.FilterClient.GetDimensionOptions(req.Context(), filterID, dim.Name)
		if err != nil {
			setStatusCode(req, w, err)
			return
		}

		var vals []string
		for _, val := range dimValues {
			vals = append(vals, val.Option)
		}

		if err := f.FilterClient.AddDimensionValues(req.Context(), newFilterID, dim.Name, vals); err != nil {
			setStatusCode(req, w, err)
			return
		}
	}

	redirectURL := fmt.Sprintf("/filters/%s/dimensions", newFilterID)
	http.Redirect(w, req, redirectURL, 302)

}
