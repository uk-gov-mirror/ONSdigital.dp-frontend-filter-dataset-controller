package hierarchy

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/ONSdigital/dp-frontend-filter-dataset-controller/data"
)

// ErrInvalidHierarchyAPIResponse is returned when the hierarchy api does not respond
// with a valid status
type ErrInvalidHierarchyAPIResponse struct {
	expectedCode int
	actualCode   int
	uri          string
}

func (e ErrInvalidHierarchyAPIResponse) Error() string {
	return fmt.Sprintf("invalid response from hierarchy api - should be: %d, got: %d, path: %s",
		e.expectedCode,
		e.actualCode,
		e.uri,
	)
}

var _ error = ErrInvalidHierarchyAPIResponse{}

// Client is a hierarchy api client which can be used to make requests to the server
type Client struct {
	cli *http.Client
	url string
}

// New creates a new instance of Client with a given filter api url
func New(hierarchyAPIURL string) *Client {
	return &Client{
		cli: &http.Client{Timeout: 5 * time.Second},
		url: hierarchyAPIURL,
	}
}

// GetHierarchy returns the hierarchy for the requested path
func (c Client) GetHierarchy(path string) (h data.Hierarchy, err error) {
	uri := fmt.Sprintf("%s/hierarchies/%s", c.url, path)

	resp, err := c.cli.Get(uri)
	if err != nil {
		return
	}

	if resp.StatusCode != http.StatusOK {
		err = &ErrInvalidHierarchyAPIResponse{http.StatusOK, resp.StatusCode, uri}
		return
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	err = json.Unmarshal(b, &h)
	return
}
