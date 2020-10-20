package recurly

import (
	"net/http"
)

type ExportDates struct {
	recurlyResponse *ResponseMetadata

	// Object type
	Object string `json:"object,omitempty"`

	// An array of dates that have available exports.
	Dates []string `json:"dates,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *ExportDates) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *ExportDates) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type exportDatesList struct {
	ListMetadata
	Data            []ExportDates `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *exportDatesList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *exportDatesList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// ExportDatesList allows you to paginate ExportDates objects
type ExportDatesList struct {
	client       HttpCaller
	nextPagePath string

	HasMore bool
	Data    []ExportDates
}

func NewExportDatesList(client HttpCaller, nextPagePath string) *ExportDatesList {
	return &ExportDatesList{
		client:       client,
		nextPagePath: nextPagePath,
		HasMore:      true,
	}
}

// Fetch fetches the next page of data into the `Data` property
func (list *ExportDatesList) Fetch() error {
	resources := &exportDatesList{}
	err := list.client.Call(http.MethodGet, list.nextPagePath, nil, resources)
	if err != nil {
		return err
	}
	// copy over properties from the response
	list.nextPagePath = resources.Next
	list.HasMore = resources.HasMore
	list.Data = resources.Data
	return nil
}

// Count returns the count of items on the server that match this pager
func (list *ExportDatesList) Count() (*int64, error) {
	resources := &exportDatesList{}
	err := list.client.Call(http.MethodHead, list.nextPagePath, nil, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}