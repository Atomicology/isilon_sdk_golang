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

type WormDomainCreateParams struct {

	// Specifies the autocommit time period for the domain in seconds.  After a file is in the domain without being modified for the specified time period, the file is automatically committed. If this parameter is set to null, there is no autocommit time, and files must be committed manually.
	AutocommitOffset int32 `json:"autocommit_offset,omitempty"`

	DefaultRetention string `json:"default_retention,omitempty"`

	MaxRetention string `json:"max_retention,omitempty"`

	MinRetention string `json:"min_retention,omitempty"`

	// Specifies the override retention date for the domain. If this date is later than the retention date for any committed file, the file will remain protected until the override retention date.
	OverrideDate int32 `json:"override_date,omitempty"`

	// When this value is set to 'on', files in this domain can be deleted through the privileged delete feature. If this value is set to 'disabled', privileged file deletes are permanently disabled and cannot be turned on again.
	PrivilegedDelete string `json:"privileged_delete,omitempty"`

	// Specifies whether the domain is an enterprise domain or a compliance domain. Compliance domains can not be created on enterprise clusters. Enterprise and compliance domains can be created on compliance clusters.
	Type_ string `json:"type,omitempty"`

	// Specifies the root path of this domain. Files in this directory and all sub-directories will be protected.
	Path string `json:"path"`
}