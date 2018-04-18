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

type CloudProxyCreateParams struct {

	// A host name or network address for connecting to this proxy
	Host string `json:"host"`

	// A unique friendly name for this proxy configuration
	Name string `json:"name"`

	// The password to connect to this proxy if required (write-only)
	Password string `json:"password,omitempty"`

	// The port used to connect to this proxy
	Port int32 `json:"port"`

	// The type of connection used to connect to this proxy
	Type_ string `json:"type"`

	// The username to connect to this proxy if required
	Username string `json:"username,omitempty"`
}
