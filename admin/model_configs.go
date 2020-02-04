/*
 * Paygate Admin API
 *
 * Paygate is ...
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package admin

// Configs struct for Configs
type Configs struct {
	CutoffTimes         []CutoffTime         `json:"CutoffTimes,omitempty"`
	FileTransferConfigs []FileTransferConfig `json:"FileTransferConfigs,omitempty"`
	FTPConfigs          []FtpConfig          `json:"FTPConfigs,omitempty"`
	SFTPConfigs         []SftpConfig         `json:"SFTPConfigs,omitempty"`
}