package handlers

import (
	"github.com/ONSdigital/go-ns/clients/codelist"
	"github.com/ONSdigital/go-ns/clients/dataset"
	"github.com/ONSdigital/go-ns/clients/filter"
	"github.com/ONSdigital/go-ns/clients/hierarchy"
	"github.com/ONSdigital/go-ns/healthcheck"
)

// ClientError implements error interface with additional code method
type ClientError interface {
	error
	Code() int
}

// FilterClient contains the methods expected for a filter client
type FilterClient interface {
	healthcheck.Client
	GetDimensions(filterID string, cfg ...filter.Config) (dims []filter.Dimension, err error)
	GetDimensionOptions(filterID, name string, cfg ...filter.Config) (fdv []filter.DimensionOption, err error)
	GetJobState(filterID string, cfg ...filter.Config) (f filter.Model, err error)
	GetOutput(filterOutputID string, cfg ...filter.Config) (f filter.Model, err error)
	GetDimension(filterID, name string, cfg ...filter.Config) (dim filter.Dimension, err error)
	AddDimensionValue(filterID, name, value string, cfg ...filter.Config) error
	RemoveDimensionValue(filterID, name, value string, cfg ...filter.Config) error
	RemoveDimension(filterID, name string, cfg ...filter.Config) (err error)
	AddDimension(filterID, name string, cfg ...filter.Config) (err error)
	AddDimensionValues(filterID, name string, options []string, cfg ...filter.Config) error
	UpdateBlueprint(m filter.Model, doSubmit bool, cfg ...filter.Config) (filter.Model, error)
	CreateBlueprint(string, []string, ...filter.Config) (string, error)
	GetPreview(string, ...filter.Config) (filter.Preview, error)
}

// DatasetClient is an interface with methods required for a dataset client
type DatasetClient interface {
	healthcheck.Client
	Get(id string, cfg ...dataset.Config) (m dataset.Model, err error)
	GetEditions(id string, cfg ...dataset.Config) (m []dataset.Edition, err error)
	GetVersions(id, edition string, cfg ...dataset.Config) (m []dataset.Version, err error)
	GetVersion(id, edition, version string, cfg ...dataset.Config) (m dataset.Version, err error)
	GetDimensions(id, edition, version string, cfg ...dataset.Config) (m dataset.Dimensions, err error)
	GetOptions(id, edition, version, dimension string, cfg ...dataset.Config) (m dataset.Options, err error)
	GetVersionMetadata(id, edition, version string, cfg ...dataset.Config) (m dataset.Metadata, err error)
}

// CodelistClient contains methods expected for a codelist client
type CodelistClient interface {
	healthcheck.Client
	GetValues(id string) (vals codelist.DimensionValues, err error)
	GetIDNameMap(id string) (map[string]string, error)
}

// HierarchyClient contains methods expected for a heirarchy client
type HierarchyClient interface {
	healthcheck.Client
	GetRoot(instanceID, name string) (hierarchy.Model, error)
	GetChild(instanceID, name, code string) (hierarchy.Model, error)
}

// Renderer provides an interface for a service template renderer
type Renderer interface {
	healthcheck.Client
	Do(path string, b []byte) ([]byte, error)
}
