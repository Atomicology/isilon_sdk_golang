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

type JobStatisticsJobNodeCpu struct {

	// The average CPU utilization of the job on this node.
	Average float32 `json:"average,omitempty"`

	// The current CPU utilization of the job on this node.
	Current float32 `json:"current"`

	// The maximum CPU utilization of the job on this node.
	Maximum float32 `json:"maximum"`

	// The minimum CPU utilization of the job on this node.
	Minimum float32 `json:"minimum"`
}
