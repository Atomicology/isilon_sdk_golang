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

type HdfsRangerPluginSettingsSettings struct {

	// Enable or disable the HDFS ranger plugin
	Enabled bool `json:"enabled,omitempty"`

	// The scheme, hostname, and port of the Apache Ranger server (e.g. http://ranger.com:6080)
	PolicyManagerUrl string `json:"policy_manager_url,omitempty"`

	// The HDFS repository name that is registered with Apache Ranger server
	RepositoryName string `json:"repository_name,omitempty"`
}
