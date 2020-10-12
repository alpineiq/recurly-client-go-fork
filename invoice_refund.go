package recurly

import ()

type InvoiceRefund struct {
	Params `json:"-"`

	// The type of refund. Amount and line items cannot both be specified in the request.
	Type *string `json:"type,omitempty"`

	// The amount to be refunded. The amount will be split between the line items.
	// If no amount is specified, it will default to refunding the total refundable amount on the invoice.
	Amount *float64 `json:"amount,omitempty"`

	// The line items to be refunded. This is required when `type=line_items`.
	LineItems []LineItemRefund `json:"line_items,omitempty"`

	// Indicates how the invoice should be refunded when both a credit and transaction are present on the invoice:
	// - `transaction_first` – Refunds the transaction first, then any amount is issued as credit back to the account. Default value when Credit Invoices feature is enabled.
	// - `credit_first` – Issues credit back to the account first, then refunds any remaining amount back to the transaction. Default value when Credit Invoices feature is not enabled.
	// - `all_credit` – Issues credit to the account for the entire amount of the refund. Only available when the Credit Invoices feature is enabled.
	// - `all_transaction` – Refunds the entire amount back to transactions, using transactions from previous invoices if necessary. Only available when the Credit Invoices feature is enabled.
	RefundMethod *string `json:"refund_method,omitempty"`

	// Used as the Customer Notes on the credit invoice.
	// This field can only be include when the Credit Invoices feature is enabled.
	CreditCustomerNotes *string `json:"credit_customer_notes,omitempty"`

	// Indicates that the refund was settled outside of Recurly, and a manual transaction should be created to track it in Recurly.
	// Required when:
	// - refunding a manually collected charge invoice, and `refund_method` is not `all_credit`
	// - refunding a credit invoice that refunded manually collecting invoices
	// - refunding a credit invoice for a partial amount
	// This field can only be included when the Credit Invoices feature is enabled.
	ExternalRefund *ExternalRefund `json:"external_refund,omitempty"`
}

func (attr *InvoiceRefund) toParams() *Params {
	return &Params{
		IdempotencyKey: attr.IdempotencyKey,
		Header:         attr.Header,
		Context:        attr.Context,
		Data:           attr,
	}
}
