package recurly

import (
	"net/http"
	"time"
)

type AccountAcquisition struct {
	recurlyResponse *ResponseMetadata

	// Account balance
	Cost AccountAcquisitionCost `json:"cost,omitempty"`

	// The channel through which the account was acquired.
	Channel string `json:"channel,omitempty"`

	// An arbitrary subchannel string representing a distinction/subcategory within a broader channel.
	Subchannel string `json:"subchannel,omitempty"`

	// An arbitrary identifier for the marketing campaign that led to the acquisition of this account.
	Campaign string `json:"campaign,omitempty"`

	Id string `json:"id,omitempty"`

	// Object type
	Object string `json:"object,omitempty"`

	// Account mini details
	Account AccountMini `json:"account,omitempty"`

	// When the account acquisition data was created.
	CreatedAt time.Time `json:"created_at,omitempty"`

	// When the account acquisition data was last changed.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *AccountAcquisition) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *AccountAcquisition) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type accountAcquisitionList struct {
	ListMetadata
	Data            []AccountAcquisition `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *accountAcquisitionList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *accountAcquisitionList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// AccountAcquisitionList allows you to paginate AccountAcquisition objects
type AccountAcquisitionList struct {
	client       HttpCaller
	nextPagePath string

	HasMore bool
	Data    []AccountAcquisition
}

func NewAccountAcquisitionList(client HttpCaller, nextPagePath string) *AccountAcquisitionList {
	return &AccountAcquisitionList{
		client:       client,
		nextPagePath: nextPagePath,
		HasMore:      true,
	}
}

// Fetch fetches the next page of data into the `Data` property
func (list *AccountAcquisitionList) Fetch() error {
	resources := &accountAcquisitionList{}
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
func (list *AccountAcquisitionList) Count() (*int64, error) {
	resources := &accountAcquisitionList{}
	err := list.client.Call(http.MethodHead, list.nextPagePath, nil, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}
