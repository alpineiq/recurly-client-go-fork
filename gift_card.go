// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
	"context"
	"net/http"
	"time"
)

type GiftCard struct {
	recurlyResponse *ResponseMetadata

	// Gift card ID
	Id string `json:"id,omitempty"`

	// Object type
	Object string `json:"object,omitempty"`

	// The ID of the account that purchased the gift card.
	GifterAccountId string `json:"gifter_account_id,omitempty"`

	// The ID of the account that redeemed the gift card redemption code.  Does not have a value until gift card is redeemed.
	RecipientAccountId string `json:"recipient_account_id,omitempty"`

	// The ID of the invoice for the gift card purchase made by the gifter.
	PurchaseInvoiceId string `json:"purchase_invoice_id,omitempty"`

	// The ID of the invoice for the gift card redemption made by the recipient.  Does not have a value until gift card is redeemed.
	RedemptionInvoiceId string `json:"redemption_invoice_id,omitempty"`

	// The unique redemption code for the gift card, generated by Recurly. Will be 16 characters, alphanumeric, displayed uppercase, but accepted in any case at redemption. Used by the recipient account to create a credit in the amount of the `unit_amount` on their account.
	RedemptionCode string `json:"redemption_code,omitempty"`

	// The remaining credit on the recipient account associated with this gift card. Only has a value once the gift card has been redeemed. Can be used to create gift card balance displays for your customers.
	Balance float64 `json:"balance,omitempty"`

	// The product code or SKU of the gift card product.
	ProductCode string `json:"product_code,omitempty"`

	// The amount of the gift card, which is the amount of the charge to the gifter account and the amount of credit that is applied to the recipient account upon successful redemption.
	UnitAmount float64 `json:"unit_amount,omitempty"`

	// 3-letter ISO 4217 currency code.
	Currency string `json:"currency,omitempty"`

	// The delivery details for the gift card.
	Delivery GiftCardDelivery `json:"delivery,omitempty"`

	// The ID of a performance obligation. Performance obligations are
	// only accessible as a part of the Recurly RevRec Standard and
	// Recurly RevRec Advanced features.
	PerformanceObligationId string `json:"performance_obligation_id,omitempty"`

	// The ID of a general ledger account. General ledger accounts are
	// only accessible as a part of the Recurly RevRec Standard and
	// Recurly RevRec Advanced features.
	LiabilityGlAccountId string `json:"liability_gl_account_id,omitempty"`

	// The ID of a general ledger account. General ledger accounts are
	// only accessible as a part of the Recurly RevRec Standard and
	// Recurly RevRec Advanced features.
	RevenueGlAccountId string `json:"revenue_gl_account_id,omitempty"`

	// Created at
	CreatedAt time.Time `json:"created_at,omitempty"`

	// Last updated at
	UpdatedAt time.Time `json:"updated_at,omitempty"`

	// When the gift card was sent to the recipient by Recurly via email, if method was email and the "Gift Card Delivery" email template was enabled. This will be empty for post delivery or email delivery where the email template was disabled.
	DeliveredAt time.Time `json:"delivered_at,omitempty"`

	// When the gift card was redeemed by the recipient.
	RedeemedAt time.Time `json:"redeemed_at,omitempty"`

	// When the gift card was canceled.
	CanceledAt time.Time `json:"canceled_at,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *GiftCard) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *GiftCard) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type giftCardList struct {
	ListMetadata
	Data            []GiftCard `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *giftCardList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *giftCardList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// GiftCardList allows you to paginate GiftCard objects
type GiftCardList struct {
	client         HTTPCaller
	requestOptions *RequestOptions
	nextPagePath   string
	hasMore        bool
	data           []GiftCard
}

func NewGiftCardList(client HTTPCaller, nextPagePath string, requestOptions *RequestOptions) *GiftCardList {
	return &GiftCardList{
		client:         client,
		requestOptions: requestOptions,
		nextPagePath:   nextPagePath,
		hasMore:        true,
	}
}

type GiftCardLister interface {
	Fetch() error
	FetchWithContext(ctx context.Context) error
	Count() (*int64, error)
	CountWithContext(ctx context.Context) (*int64, error)
	Data() []GiftCard
	HasMore() bool
	Next() string
}

func (list *GiftCardList) HasMore() bool {
	return list.hasMore
}

func (list *GiftCardList) Next() string {
	return list.nextPagePath
}

func (list *GiftCardList) Data() []GiftCard {
	return list.data
}

// Fetch fetches the next page of data into the `Data` property
func (list *GiftCardList) FetchWithContext(ctx context.Context) error {
	resources := &giftCardList{}
	err := list.client.Call(ctx, http.MethodGet, list.nextPagePath, nil, nil, list.requestOptions, resources)
	if err != nil {
		return err
	}
	// copy over properties from the response
	list.nextPagePath = resources.Next
	list.hasMore = resources.HasMore
	list.data = resources.Data
	return nil
}

// Fetch fetches the next page of data into the `Data` property
func (list *GiftCardList) Fetch() error {
	return list.FetchWithContext(context.Background())
}

// Count returns the count of items on the server that match this pager
func (list *GiftCardList) CountWithContext(ctx context.Context) (*int64, error) {
	resources := &giftCardList{}
	err := list.client.Call(ctx, http.MethodHead, list.nextPagePath, nil, nil, list.requestOptions, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}

// Count returns the count of items on the server that match this pager
func (list *GiftCardList) Count() (*int64, error) {
	return list.CountWithContext(context.Background())
}
