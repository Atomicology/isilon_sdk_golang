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

type JobJobPrepairParams struct {

	// Type of permissions; not accepted with mode=clone or mode=inherit.
	MappingType string `json:"mapping_type,omitempty"`

	// Type of PermissionRepair operation.
	Mode string `json:"mode"`

	// IFS file or directory to use as a template; required with mode=clone and mode=inherit, not accepted with mode=convert.
	Template string `json:"template,omitempty"`

	// Authentication zone; not accepted with mode=clone or mode=inherit.
	Zone string `json:"zone,omitempty"`
}
