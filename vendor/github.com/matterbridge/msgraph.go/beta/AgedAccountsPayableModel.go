// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// AgedAccountsPayable undocumented
type AgedAccountsPayable struct {
	// Entity is the base model of AgedAccountsPayable
	Entity
	// VendorNumber undocumented
	VendorNumber *string `json:"vendorNumber,omitempty"`
	// Name undocumented
	Name *string `json:"name,omitempty"`
	// CurrencyCode undocumented
	CurrencyCode *string `json:"currencyCode,omitempty"`
	// BalanceDue undocumented
	BalanceDue *int `json:"balanceDue,omitempty"`
	// CurrentAmount undocumented
	CurrentAmount *int `json:"currentAmount,omitempty"`
	// Period1Amount undocumented
	Period1Amount *int `json:"period1Amount,omitempty"`
	// Period2Amount undocumented
	Period2Amount *int `json:"period2Amount,omitempty"`
	// Period3Amount undocumented
	Period3Amount *int `json:"period3Amount,omitempty"`
	// AgedAsOfDate undocumented
	AgedAsOfDate *time.Time `json:"agedAsOfDate,omitempty"`
	// PeriodLengthFilter undocumented
	PeriodLengthFilter *string `json:"periodLengthFilter,omitempty"`
}
