package handlers

// Filter represents the handlers for Filtering
type Filter struct {
	Renderer        Renderer
	FilterClient    FilterClient
	DatasetClient   DatasetClient
	CodeListClient  CodelistClient
	HierarchyClient HierarchyClient
	val             Validator
}

// NewFilter creates a new instance of Filter
func NewFilter(r Renderer, fc FilterClient, dc DatasetClient, clc CodelistClient, hc HierarchyClient, val Validator) *Filter {
	return &Filter{
		Renderer:        r,
		FilterClient:    fc,
		DatasetClient:   dc,
		CodeListClient:  clc,
		HierarchyClient: hc,
		val:             val,
	}
}
