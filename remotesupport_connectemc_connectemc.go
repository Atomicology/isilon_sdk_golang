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

type RemotesupportConnectemcConnectemc struct {

	// Email the customer if all transmission methods fail.
	EmailCustomerOnFailure bool `json:"email_customer_on_failure,omitempty"`

	// Enable ConnectEMC.
	Enabled bool `json:"enabled,omitempty"`

	// List of network pools that are able to connect to the ESRS gateway.  Necessary to enable ConnectEMC.
	GatewayAccessPools []string `json:"gateway_access_pools,omitempty"`

	// Primary ESRS Gateway. Necessary to enable ConnectEMC.
	PrimaryEsrsGateway string `json:"primary_esrs_gateway,omitempty"`

	// Secondary ESRS Gateway. Used if Primary is unavailable.
	SecondaryEsrsGateway string `json:"secondary_esrs_gateway,omitempty"`

	// Use SMPT if primary and secondary gateways are unavailable.
	UseSmtpFailover bool `json:"use_smtp_failover,omitempty"`
}
