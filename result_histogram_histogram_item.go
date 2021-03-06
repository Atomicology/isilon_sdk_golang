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

type ResultHistogramHistogramItem struct {

	// Bucket for holding file counts within a range.
	Bucket int32 `json:"bucket"`

	// Number of files within the bucket.
	Value int32 `json:"value"`
}
