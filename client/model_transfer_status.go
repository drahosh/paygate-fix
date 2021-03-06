/*
 * Paygate API
 *
 * Paygate is a RESTful API enabling Automated Clearing House ([ACH](https://en.wikipedia.org/wiki/Automated_Clearing_House)) transactions to be submitted and received without a deep understanding of a full NACHA file specification.
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client

// TransferStatus Defines the state of the Transfer
type TransferStatus string

// List of TransferStatus
const (
	CANCELED   TransferStatus = "canceled"
	FAILED     TransferStatus = "failed"
	REVIEWABLE TransferStatus = "reviewable"
	PENDING    TransferStatus = "pending"
	PROCESSED  TransferStatus = "processed"
)
