/*
 * Paygate Admin API
 *
 * Paygate is a RESTful API enabling Automated Clearing House ([ACH](https://en.wikipedia.org/wiki/Automated_Clearing_House)) transactions to be submitted and received without a deep understanding of a full NACHA file specification.
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package admin

// ReturnCode struct for ReturnCode
type ReturnCode struct {
	// Optional NACHA return code for this Transfer
	Code string `json:"code,omitempty"`
	// Short NACHA description of return code
	Reason string `json:"reason,omitempty"`
	// Long form explanation of return code
	Description string `json:"description,omitempty"`
}
