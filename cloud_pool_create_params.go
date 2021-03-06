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

type CloudPoolCreateParams struct {

	// A list of valid names for the accounts in this pool.  There is currently only one account allowed per pool.
	Accounts []string `json:"accounts"`

	// The guid of the cluster where this pool was created
	BirthClusterId string `json:"birth_cluster_id,omitempty"`

	// A brief description of this pool
	Description string `json:"description,omitempty"`

	// A unique name for this pool
	Name string `json:"name"`

	// A string identifier of the cloud services vendor
	Vendor string `json:"vendor,omitempty"`

	// The type of cloud protocol required.  E.g., \"isilon\" for EMC Isilon, \"ecs\" for EMC ECS Appliance, \"virtustream\" for Virtustream Storage Cloud, \"azure\" for Microsoft Azure and \"s3\" for Amazon S3
	Type_ string `json:"type"`
}
