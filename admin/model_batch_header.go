/*
 * Paygate Admin API
 *
 * Paygate is a RESTful API enabling Automated Clearing House ([ACH](https://en.wikipedia.org/wiki/Automated_Clearing_House)) transactions to be submitted and received without a deep understanding of a full NACHA file specification.
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package admin

// BatchHeader struct for BatchHeader
type BatchHeader struct {
	// Batch Header ID
	ID string `json:"ID,omitempty"`
	// Service Class Code - ACH Credits Only 220 and ACH Debits Only 225
	ServiceClassCode int32 `json:"serviceClassCode"`
	// Company originating the entries in the batch
	CompanyName string `json:"companyName"`
	// The 9 digit FEIN number (proceeded by a predetermined alpha or numeric character) of the entity in the company name field
	CompanyDiscretionaryData string `json:"companyDiscretionaryData,omitempty"`
	// Identifies the payment type (product) found within an ACH batch-using a 3-character code.
	StandardEntryClassCode string `json:"standardEntryClassCode,omitempty"`
	// A description of the entries contained in the batch. The Originator establishes the value of this field to provide a description of the purpose of the entry to be displayed back to the receive For example, \"GAS BILL,\" \"REG. SALARY,\" \"INS. PREM,\", \"SOC. SEC.,\" \"DTC,\" \"TRADE PAY,\" \"PURCHASE,\" etc. This field must contain the word \"REVERSAL\" (left justified) when the batch contains reversing entries. This field must contain the word \"RECLAIM\" (left justified) when the batch contains reclamation entries. This field must contain the word \"NONSETTLED\" (left justified) when the batch contains entries which could not settle.
	CompanyEntryDescription string `json:"companyEntryDescription,omitempty"`
	// The Originator establishes this field as the date it would like to see displayed to the receiver for descriptive purposes. This field is never used to control timing of any computer or manual operation. It is solely for descriptive purposes. The RDFI should not assume any specific format.
	CompanyDescriptiveDate string `json:"companyDescriptiveDate,omitempty"`
	// Date on which the entries are to settle. Format YYMMDD (Y=Year, M=Month, D=Day)
	EffectiveEntryDate string `json:"effectiveEntryDate,omitempty"`
	// ODFI initiating the Entry. 0 ADV File prepared by an ACH Operator. 1 This code identifies the Originator as a depository financial institution. 2 This code identifies the Originator as a Federal Government entity or agency.
	OriginatorStatusCode int32 `json:"originatorStatusCode,omitempty"`
	// First 8 digits of the originating DFI transit routing number
	ODFIIdentification string `json:"ODFIIdentification"`
	// BatchNumber is assigned in ascending sequence to each batch by the ODFI or its Sending Point in a given file of entries. Since the batch number in the Batch Header Record and the Batch Control Record is the same, the ascending sequence number should be assigned by batch and not by record.
	BatchNumber string `json:"batchNumber,omitempty"`
}
