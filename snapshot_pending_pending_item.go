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

type SnapshotPendingPendingItem struct {

	// The system supplied unique ID used for sorting and paging.
	Id string `json:"id"`

	// The /ifs path that will snapshotted.
	Path string `json:"path"`

	// The name of the schedule used to create this snapshot.
	Schedule string `json:"schedule"`

	// The system snapshot name formed from the schedule formate.
	Snapshot string `json:"snapshot"`

	// The Unix Epoch time the snapshot will be created.
	Time int32 `json:"time"`
}
