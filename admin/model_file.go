/*
 * Paygate Admin API
 *
 * Paygate is a RESTful API enabling Automated Clearing House ([ACH](https://en.wikipedia.org/wiki/Automated_Clearing_House)) transactions to be submitted and received without a deep understanding of a full NACHA file specification.
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package admin

// File struct for File
type File struct {
	// File ID
	ID                   string      `json:"ID,omitempty"`
	FileHeader           FileHeader  `json:"fileHeader,omitempty"`
	Batches              []Batch     `json:"batches,omitempty"`
	IATBatches           []IatBatch  `json:"IATBatches,omitempty"`
	FileControl          FileControl `json:"fileControl,omitempty"`
	NotificationOfChange []Batch     `json:"NotificationOfChange,omitempty"`
	ReturnEntries        []Batch     `json:"ReturnEntries,omitempty"`
}
