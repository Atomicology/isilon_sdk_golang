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

type SwiftAccountExtended struct {

	// Unique id of swift account
	Id string `json:"id,omitempty"`

	// Name of Swift account
	Name string `json:"name,omitempty"`

	// Path to root of Swift account
	Path string `json:"path,omitempty"`

	// Group with filesystem ownership of this account
	Swiftgroup string `json:"swiftgroup,omitempty"`

	// User with filesystem ownership of this account
	Swiftuser string `json:"swiftuser,omitempty"`

	// Users who are allowed to access Swift account
	Users []string `json:"users,omitempty"`

	// Name of access zone for account
	Zone string `json:"zone,omitempty"`
}
