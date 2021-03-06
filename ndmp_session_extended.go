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

type NdmpSessionExtended struct {

	// Bytes transferred to/from the filesystem
	DataBytesTransferred int32 `json:"data_bytes_transferred"`

	// State of the NDMP Data Service
	DataState string `json:"data_state"`

	// The path being recovered to
	DestPath string `json:"dest_path"`

	// IP address of the DMA
	DmaIpAddr string `json:"dma_ip_addr"`

	// Number of seconds elapsed since the backup was started
	ElapsedTime int32 `json:"elapsed_time"`

	// Unique display ID.
	Id string `json:"id"`

	// Bytes transferred to/from tape or remote writer
	MoverBytesTransferred int32 `json:"mover_bytes_transferred"`

	// State of the NDMP Mover Service
	MoverState string `json:"mover_state"`

	// The type of backup session
	Operation string `json:"operation"`

	// IP address of the remote NDMP participant
	RemoteIpAddr string `json:"remote_ip_addr"`

	// Name of the media changer device used if any
	ScsiDevice string `json:"scsi_device"`

	// Session ID in form <lnn>.<pid>.
	Session string `json:"session"`

	// The path being backed up
	SourcePath string `json:"source_path"`

	// Time backup was started in seconds since epoch
	StartTime int32 `json:"start_time"`

	// Name of the tape device used if any
	TapeDevice string `json:"tape_device"`

	// Describes the mode in which the tape is opened
	TapeOpenMode string `json:"tape_open_mode"`

	// The throughput in MB/s
	Throughput int32 `json:"throughput"`
}
