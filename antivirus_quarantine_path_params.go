/*
 * Isilon SDK
 *
 * Isilon SDK - Language bindings for the OneFS API
 *
 * API version: 5
 * Contact: sdk@isilon.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package isi_sdk_8_1_0

// The quarantine status of a file in /ifs.
type AntivirusQuarantinePathParams struct {

	// If true, this file is quarantined.  If false, the file is not quarantined.
	Quarantined bool `json:"quarantined,omitempty"`
}