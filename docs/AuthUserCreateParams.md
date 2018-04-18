# AuthUserCreateParams

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | Specifies an email address for the user. | [optional] [default to null]
**Enabled** | **bool** | If true, the authenticated user is enabled. | [optional] [default to null]
**Expiry** | **int32** | Specifies the Unix Epoch time when the auth user will expire. | [optional] [default to null]
**Gecos** | **string** | Specifies the GECOS value, which is usually the full name. | [optional] [default to null]
**HomeDirectory** | **string** | Specifies a home directory for the user. | [optional] [default to null]
**Password** | **string** | Changes the password for the user. | [optional] [default to null]
**PasswordExpires** | **bool** | If true, the password should expire. | [optional] [default to null]
**PrimaryGroup** | [***AuthAccessAccessItemFileGroup**](AuthAccessAccessItemFileGroup.md) | Specifies properties for a persona, which consists of either a &#39;type&#39; and a &#39;name&#39; or an &#39;ID&#39;. | [optional] [default to null]
**PromptPasswordChange** | **bool** | If true, prompts the user to change their password at the next login. | [optional] [default to null]
**Shell** | **string** | Specifies the shell for the user. | [optional] [default to null]
**Sid** | **string** | Specifies a security identifier. | [optional] [default to null]
**Uid** | **int32** | Specifies a numeric user identifier. | [optional] [default to null]
**Unlock** | **bool** | If true, the user account should be unlocked. | [optional] [default to null]
**Name** | **string** | Specifies a user name. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


