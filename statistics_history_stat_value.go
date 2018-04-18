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

type StatisticsHistoryStatValue struct {

	// Unix Epoch time in seconds that statistic was collected.
	Time int32 `json:"time"`

	// Key dependent value.
	Value string `json:"value"`
}
