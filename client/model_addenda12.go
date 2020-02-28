/*
 * Paygate API
 *
 * Paygate is a RESTful API enabling Automated Clearing House ([ACH](https://en.wikipedia.org/wiki/Automated_Clearing_House)) transactions to be submitted and received without a deep understanding of a full NACHA file specification.
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client

// Addenda12 struct for Addenda12
type Addenda12 struct {
	// Client defined string used as a reference to this record.
	Id string `json:"id,omitempty"`
	// 10 - NACHA regulations
	TypeCode string `json:"typeCode,omitempty"`
	// Originator City & State / Province Data elements City and State / Province  should be separated with an asterisk (*) as a delimiter and the field should end with a backslash (\\).
	OriginatorCityStateProvince string `json:"originatorCityStateProvince,omitempty"`
	// Originator Country & Postal Code Data elements must be separated by an asterisk (*) and must end with a backslash (\\)
	OriginatorCountryPostalCode string `json:"originatorCountryPostalCode,omitempty"`
	// EntryDetailSequenceNumber contains the ascending sequence number section of the Entry Detail or Corporate Entry Detail Record's trace number This number is the same as the last seven digits of the trace number of the related Entry Detail Record or Corporate Entry Detail Record.
	EntryDetailSequenceNumber float32 `json:"entryDetailSequenceNumber,omitempty"`
}
