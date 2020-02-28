/*
 * Paygate Admin API
 *
 * Paygate is a RESTful API enabling Automated Clearing House ([ACH](https://en.wikipedia.org/wiki/Automated_Clearing_House)) transactions to be submitted and received without a deep understanding of a full NACHA file specification.
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package admin

// FtpConfig struct for FtpConfig
type FtpConfig struct {
	// DNS or IP address for FTP server
	Hostname string `json:"hostname"`
	// username for authentication
	Username string `json:"username"`
	// password for authentication
	Password string `json:"password,omitempty"`
}
