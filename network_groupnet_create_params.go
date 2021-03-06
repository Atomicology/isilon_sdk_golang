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

type NetworkGroupnetCreateParams struct {

	// A description of the groupnet.
	Description string `json:"description,omitempty"`

	// DNS caching is enabled or disabled.
	DnsCacheEnabled bool `json:"dns_cache_enabled,omitempty"`

	// List of DNS resolver options.
	DnsOptions []string `json:"dns_options,omitempty"`

	// List of DNS search suffixes.
	DnsSearch []string `json:"dns_search,omitempty"`

	// List of Domain Name Server IP addresses.
	DnsServers []string `json:"dns_servers,omitempty"`

	// The name of the groupnet.
	Name string `json:"name"`

	// Enable or disable appending nodes DNS search  list to client DNS inquiries directed at SmartConnect service IP.
	ServerSideDnsSearch bool `json:"server_side_dns_search,omitempty"`
}
