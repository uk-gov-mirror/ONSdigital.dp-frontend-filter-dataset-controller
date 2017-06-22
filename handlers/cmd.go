package handlers

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/ONSdigital/dp-frontend-filter-dataset-controller/renderer"
	"github.com/ONSdigital/dp-frontend-models/model"
	"github.com/ONSdigital/dp-frontend-models/model/datasetpages/finishPage"
	"github.com/ONSdigital/dp-frontend-models/model/datasetpages/middlePage"
	"github.com/ONSdigital/dp-frontend-models/model/datasetpages/startPage"
	"github.com/ONSdigital/go-ns/log"
	"github.com/gorilla/mux"
)

// CMD represents the handlers for CMD
type CMD struct {
	r renderer.Renderer
}

// NewCMD creates a new instance of CMD
func NewCMD(r renderer.Renderer) *CMD {
	return &CMD{r: r}
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

// Landing handles the controller functionality for the landing page
func (c *CMD) Landing(w http.ResponseWriter, req *http.Request) {
	var p startPage.Page

	// Needs to be populated from API - this is stubbed data
	p.Metadata.Footer = getStubbedMetadataFooter()
	p.SearchDisabled = true

	pBytes, err := json.Marshal(p)
	if err != nil {
		log.Error(err, nil)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	b, err := c.r.Do("dataset/startpage", pBytes)
	if err != nil {
		log.Error(err, nil)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if _, err := w.Write(b); err != nil {
		log.Error(err, nil)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// CreateJobID controls the creating of a job idea when a new user journey is
// requested
func (c *CMD) CreateJobID(w http.ResponseWriter, req *http.Request) {
	// TODO: This is a stubbed job id - replace with real job id from api once
	// code has been written
	jobID := rand.Intn(100000000)
	jid := strconv.Itoa(jobID)

	log.Trace("created job id", log.Data{"job_id": jid})
	http.Redirect(w, req, "/jobs/"+jid, 301)
}

// Middle controls the rendering of a "middle" cmd page - this will be replaced
// by other handlers when further pages are defined by UX
func (c *CMD) Middle(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	var p middlePage.Page

	// Needs to be populated from API - this is stubbed data
	p.Metadata.Footer = getStubbedMetadataFooter()
	p.SearchDisabled = true
	p.Data.JobID = vars["jobID"]

	body, err := json.Marshal(p)
	if err != nil {
		log.Error(err, nil)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	b, err := c.r.Do("dataset/middlepage", body)
	if err != nil {
		log.Error(err, nil)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if _, err := w.Write(b); err != nil {
		log.Error(err, nil)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// PreviewAndDownload will control the rendering of the preview and download page
func (c *CMD) PreviewAndDownload(w http.ResponseWriter, req *http.Request) {
	var p finishPage.Page

	// Needs to be populated from API - this is stubbed data
	p.Metadata.Footer = getStubbedMetadataFooter()
	p.SearchDisabled = true

	pBytes, err := json.Marshal(p)
	if err != nil {
		log.Error(err, nil)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	b, err := c.r.Do("dataset/finishpage", pBytes)
	if err != nil {
		log.Error(err, nil)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if _, err := w.Write(b); err != nil {
		log.Error(err, nil)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
