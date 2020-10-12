package recurly

import ()

type AddressCreate struct {
	Params `json:"-"`

	// First name
	FirstName *string `json:"first_name,omitempty"`

	// Last name
	LastName *string `json:"last_name,omitempty"`

	// Phone number
	Phone *string `json:"phone,omitempty"`

	// Street 1
	Street1 *string `json:"street1,omitempty"`

	// Street 2
	Street2 *string `json:"street2,omitempty"`

	// City
	City *string `json:"city,omitempty"`

	// State or province.
	Region *string `json:"region,omitempty"`

	// Zip or postal code.
	PostalCode *string `json:"postal_code,omitempty"`

	// Country, 2-letter ISO code.
	Country *string `json:"country,omitempty"`
}

func (attr *AddressCreate) toParams() *Params {
	return &Params{
		IdempotencyKey: attr.IdempotencyKey,
		Header:         attr.Header,
		Context:        attr.Context,
		Data:           attr,
	}
}
