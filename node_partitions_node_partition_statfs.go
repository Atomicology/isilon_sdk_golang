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

type NodePartitionsNodePartitionStatfs struct {

	// Free blocks available to non-superuser on this partition.
	FBavail int32 `json:"f_bavail,omitempty"`

	// Free blocks on this partition.
	FBfree int32 `json:"f_bfree,omitempty"`

	// Total data blocks on this partition.
	FBlocks int32 `json:"f_blocks,omitempty"`

	// Filesystem fragment size; block size in OneFS.
	FBsize int32 `json:"f_bsize,omitempty"`

	// Free file nodes avail to non-superuser.
	FFfree int32 `json:"f_ffree,omitempty"`

	// Total file nodes in filesystem.
	FFiles int32 `json:"f_files,omitempty"`

	// Mount exported flags.
	FFlags int32 `json:"f_flags,omitempty"`

	// File system type name.
	FFstypename string `json:"f_fstypename,omitempty"`

	// Optimal transfer block size.
	FIosize int32 `json:"f_iosize,omitempty"`

	// Names of devices this partition is mounted from.
	FMntfromname string `json:"f_mntfromname,omitempty"`

	// Directory this partition is mounted to.
	FMntonname string `json:"f_mntonname,omitempty"`

	// Maximum filename length.
	FNamemax int32 `json:"f_namemax,omitempty"`

	// UID of user that mounted the filesystem.
	FOwner int32 `json:"f_owner,omitempty"`

	// Type of filesystem.
	FType int32 `json:"f_type,omitempty"`

	// statfs() structure version number.
	FVersion int32 `json:"f_version,omitempty"`
}
