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

type ResultTopDirsDir struct {

	// Directory access time
	Atime int32 `json:"atime"`

	// Directory creation begin time.
	Btime int32 `json:"btime"`

	// Unix inode change time.
	Ctime int32 `json:"ctime"`

	// Relative directory path under /ifs/.
	Path string `json:"path"`
}
