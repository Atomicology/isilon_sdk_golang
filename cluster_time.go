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

type ClusterTime struct {

	// A list of errors encountered by the individual nodes involved in this request, or an empty list if there were no errors.
	Errors []NodeDrivesPurposelistError `json:"errors,omitempty"`

	// The responses from the individual nodes involved in this request.
	Nodes []ClusterTimeNode `json:"nodes,omitempty"`

	// The total number of nodes responding.
	Total int32 `json:"total,omitempty"`
}
