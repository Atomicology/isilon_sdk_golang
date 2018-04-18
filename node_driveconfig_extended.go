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

// An object containing a node's drive subsystem XML configuration file.
type NodeDriveconfigExtended struct {

	// Configuration settings for automatic replacement recognition (ARR).
	AutomaticReplacementRecognition *NodeDriveconfigNodeAutomaticReplacementRecognition `json:"automatic_replacement_recognition,omitempty"`
}
