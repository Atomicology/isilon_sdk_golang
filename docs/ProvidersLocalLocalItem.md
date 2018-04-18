# ProvidersLocalLocalItem

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Authentication** | **bool** | If true, enables authentication and identity management through the authentication provider. | [optional] [default to null]
**CreateHomeDirectory** | **bool** | Automatically creates the home directory on the first login. | [optional] [default to null]
**HomeDirectoryTemplate** | **string** | Specifies the path to the home directory template. | [optional] [default to null]
**Id** | **string** | Specifies the local provider ID. | [optional] [default to null]
**LockoutDuration** | **int32** | Specifies the length of time in seconds that an account will be inaccessible after multiple failed login attempts. | [optional] [default to null]
**LockoutThreshold** | **int32** | Specifies the number of failed login attempts necessary before an account is locked. | [optional] [default to null]
**LockoutWindow** | **int32** | Specifies the duration of time in seconds in which the number of failed attempts set in the &#39;lockout_threshold&#39; parameter must be made before an account is locked. | [optional] [default to null]
**LoginShell** | **string** | Specifies the login shell path. | [optional] [default to null]
**MachineName** | **string** | Specifies the domain for this provider through which users and groups are qualified. | [optional] [default to null]
**MaxPasswordAge** | **int32** | Specifies the maximum password age in seconds. | [optional] [default to null]
**MinPasswordAge** | **int32** | Specifies the minimum password age in seconds. | [optional] [default to null]
**MinPasswordLength** | **int32** | Specifies the minimum password length. | [optional] [default to null]
**Name** | **string** | Specifies the local provider name. | [optional] [default to null]
**PasswordComplexity** | **[]string** | Specifies the conditions required for a password. | [optional] [default to null]
**PasswordHistoryLength** | **int32** | Specifies the number of previous passwords to store. | [optional] [default to null]
**PasswordPromptTime** | **int32** | Specifies the time in seconds remaining before a user will be prompted for a password change. | [optional] [default to null]
**Status** | **string** | Specifies the status of the provider. | [optional] [default to null]
**System** | **bool** | If true, indicates that this provider instance was created by OneFS and cannot be removed. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


