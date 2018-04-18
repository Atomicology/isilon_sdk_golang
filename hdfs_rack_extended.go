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

type HdfsRackExtended struct {

	// Array of IP ranges. Clients from one of these IP ranges are served by corresponding nodes from ip_pools array.
	ClientIpRanges []string `json:"client_ip_ranges,omitempty"`

	// Array of IP pool names to use for serving clients from client_ip_ranges.
	IpPools []string `json:"ip_pools,omitempty"`

	// Name of the rack
	Name string `json:"name,omitempty"`

	// The ID of the rack.
	Id string `json:"id,omitempty"`
}