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

type ChangelistLins struct {

	// 
	Atime *ChangelistLinsCtime `json:"atime,omitempty"`

	// 
	Ctime *ChangelistLinsCtime `json:"ctime,omitempty"`

	// The LIN number of the file associated with the entry.
	Id string `json:"id"`

	// 
	Mtime *ChangelistLinsCtime `json:"mtime,omitempty"`

	// The path to the file associated with the entry.
	Path string `json:"path"`

	// The size of the file associated with the entry.
	Size int32 `json:"size"`

	// Type of file.
	Type_ string `json:"type"`
}
