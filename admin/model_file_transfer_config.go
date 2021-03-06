/*
 * Paygate Admin API
 *
 * Paygate is a RESTful API enabling Automated Clearing House ([ACH](https://en.wikipedia.org/wiki/Automated_Clearing_House)) transactions to be submitted and received without a deep understanding of a full NACHA file specification.
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package admin

// FileTransferConfig struct for FileTransferConfig
type FileTransferConfig struct {
	// Filepath for inbound files
	InboundPath string `json:"inboundPath,omitempty"`
	// Filepath for outbound files
	OutboundPath string `json:"outboundPath,omitempty"`
	// Filepath for return files
	ReturnPath string `json:"returnPath,omitempty"`
	// Go template for uploaded ACH filenames. Refer to our documentation on filename templates for more details. https://docs.moov.io/paygate/ach/#filename-templates
	OutboundFilenameTemplate string `json:"outboundFilenameTemplate,omitempty"`
	// CIDR range or IP address for allowed remote server
	AllowedIPs string `json:"allowedIPs,omitempty"`
}
