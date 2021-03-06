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

type ClusterPatchPatchesPatchFile struct {

	// The md5 checksum of the file.
	Md5 string `json:"md5,omitempty"`

	// The path of the file.
	Path string `json:"path,omitempty"`
}
