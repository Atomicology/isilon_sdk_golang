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

type StoragepoolStatusUnhealthyItemDiskpool struct {

	// The drives that are part of this disk pool.
	Drives []StoragepoolStatusUnprovisionedItem `json:"drives"`

	// The system ID given to the disk pool.
	Id int32 `json:"id"`

	// The disk pool name.
	Name string `json:"name"`

	// The system ID of the disk pool's node pool, if it is in a node pool.
	NodepoolId int32 `json:"nodepool_id,omitempty"`

	// The protection policy for the disk pool.
	ProtectionPolicy string `json:"protection_policy"`

	// The SSDs that are part of this disk pool.
	SsdDrives []StoragepoolStatusUnprovisionedItem `json:"ssd_drives"`
}