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

type NfsNlmLocksLock struct {

	// Specifies the client host name and IP address.
	Client string `json:"client,omitempty"`

	// Specifies the client ID.
	ClientId string `json:"client_id,omitempty"`

	// Specifies the UNIX EPoch time that the lock was created.
	Created int32 `json:"created,omitempty"`

	// Specifies the system-assigned ID given to the lock. This value is returned when the lock is created with the POST method.
	Id string `json:"id,omitempty"`

	// Specifies the LIN in /ifs that is locked.
	Lin string `json:"lin,omitempty"`

	// Specifies the lock type.
	LockType string `json:"lock_type,omitempty"`

	// Specifies the path under /ifs that is locked.
	Path string `json:"path,omitempty"`

	// Specifies the byte range within the locked file.
	Range_ []int32 `json:"range,omitempty"`
}
