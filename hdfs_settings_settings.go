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

type HdfsSettingsSettings struct {

	// Ambari metrics collector
	AmbariMetricsCollector string `json:"ambari_metrics_collector,omitempty"`

	// NameNode of Ambari server
	AmbariNamenode string `json:"ambari_namenode,omitempty"`

	// Ambari server
	AmbariServer string `json:"ambari_server,omitempty"`

	// Type of authentication for HDFS protocol.
	AuthenticationMode string `json:"authentication_mode,omitempty"`

	// Encryption algorithm to use for data transfer (if any)
	DataTransferCipher string `json:"data_transfer_cipher,omitempty"`

	// Block size (size=2**value) reported by HDFS server.
	DefaultBlockSize int32 `json:"default_block_size,omitempty"`

	// Checksum type reported by HDFS server.
	DefaultChecksumType string `json:"default_checksum_type,omitempty"`

	// ODP stack repository version number
	OdpVersion string `json:"odp_version,omitempty"`

	// HDFS root directory
	RootDirectory string `json:"root_directory,omitempty"`

	// Enable or disable the HDFS service.
	Service bool `json:"service,omitempty"`

	// Enable or disable WebHDFS
	WebhdfsEnabled bool `json:"webhdfs_enabled,omitempty"`
}
