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

type CompatibilitiesClassActiveActiveItem struct {

	// The first class in an active compatibility
	Class1 string `json:"class_1"`

	// The second class in an active compatibility
	Class2 string `json:"class_2"`

	// The id of this active compatibility
	Id int32 `json:"id"`
}
