/*
 * Paygate API
 *
 * Paygate is a RESTful API enabling Automated Clearing House ([ACH](https://en.wikipedia.org/wiki/Automated_Clearing_House)) transactions to be submitted and received without a deep understanding of a full NACHA file specification.
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client

// Addenda99 struct for Addenda99
type Addenda99 struct {
	// Client defined string used as a reference to this record.
	Id string `json:"id,omitempty"`
	// 99 - NACHA regulations
	TypeCode string `json:"typeCode,omitempty"`
	// Standard code used by an ACH Operator or RDFI to describe the reason for returning an Entry.
	ReturnCode string `json:"returnCode,omitempty"`
	// OriginalTrace This field contains the Trace Number as originally included on the forward Entry or Prenotification. The RDFI must include the Original Entry Trace Number in the Addenda Record of an Entry being returned to an ODFI, in the Addenda Record of an 98, within an Acknowledgment Entry, or with an RDFI request for a copy of an authorization.
	OriginalTrace string `json:"originalTrace,omitempty"`
	// DateOfDeath The field date of death is to be supplied on Entries being returned for reason of death (return reason codes R14 and R15). Format YYMMDD (Y=Year, M=Month, D=Day)
	DateOfDeath string `json:"dateOfDeath,omitempty"`
	// OriginalDFI field contains the Receiving DFI Identification (addenda.RDFIIdentification) as originally included on the forward Entry or Prenotification that the RDFI is returning or correcting.
	OriginalDFI string `json:"originalDFI,omitempty"`
	// Information related to the return
	AddendaInformation string `json:"addendaInformation,omitempty"`
	// Matches the Entry Detail Trace Number of the entry being returned.
	TraceNumber string `json:"traceNumber,omitempty"`
}
