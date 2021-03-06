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

type DebugStatsUnknown struct {

	// The number of calls.
	Calls int32 `json:"calls,omitempty"`

	// The number of errors.
	Errors int32 `json:"errors,omitempty"`

	// The total amount of time spent in this method.
	Time float32 `json:"time,omitempty"`
}
