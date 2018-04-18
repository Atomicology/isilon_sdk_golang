# \FilepoolApi

All URIs are relative to *https://YOUR_CLUSTER_HOSTNAME_OR_NODE_IP:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFilepoolPolicy**](FilepoolApi.md#CreateFilepoolPolicy) | **Post** /platform/4/filepool/policies | 
[**DeleteFilepoolPolicy**](FilepoolApi.md#DeleteFilepoolPolicy) | **Delete** /platform/4/filepool/policies/{FilepoolPolicyId} | 
[**GetFilepoolDefaultPolicy**](FilepoolApi.md#GetFilepoolDefaultPolicy) | **Get** /platform/4/filepool/default-policy | 
[**GetFilepoolPolicy**](FilepoolApi.md#GetFilepoolPolicy) | **Get** /platform/4/filepool/policies/{FilepoolPolicyId} | 
[**GetFilepoolTemplate**](FilepoolApi.md#GetFilepoolTemplate) | **Get** /platform/4/filepool/templates/{FilepoolTemplateId} | 
[**GetFilepoolTemplates**](FilepoolApi.md#GetFilepoolTemplates) | **Get** /platform/4/filepool/templates | 
[**ListFilepoolPolicies**](FilepoolApi.md#ListFilepoolPolicies) | **Get** /platform/4/filepool/policies | 
[**UpdateFilepoolDefaultPolicy**](FilepoolApi.md#UpdateFilepoolDefaultPolicy) | **Put** /platform/4/filepool/default-policy | 
[**UpdateFilepoolPolicy**](FilepoolApi.md#UpdateFilepoolPolicy) | **Put** /platform/4/filepool/policies/{FilepoolPolicyId} | 


# **CreateFilepoolPolicy**
> CreateFilepoolPolicyResponse CreateFilepoolPolicy(ctx, filepoolPolicy)


Create a new policy.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **filepoolPolicy** | [**FilepoolPolicyCreateParams**](FilepoolPolicyCreateParams.md)|  | 

### Return type

[**CreateFilepoolPolicyResponse**](CreateFilepoolPolicyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteFilepoolPolicy**
> DeleteFilepoolPolicy(ctx, filepoolPolicyId)


Delete file pool policy.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **filepoolPolicyId** | **string**| Delete file pool policy. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFilepoolDefaultPolicy**
> FilepoolDefaultPolicy GetFilepoolDefaultPolicy(ctx, )


List default file pool policy.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**FilepoolDefaultPolicy**](FilepoolDefaultPolicy.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFilepoolPolicy**
> FilepoolPolicies GetFilepoolPolicy(ctx, filepoolPolicyId)


Retrieve file pool policy information.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **filepoolPolicyId** | **string**| Retrieve file pool policy information. | 

### Return type

[**FilepoolPolicies**](FilepoolPolicies.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFilepoolTemplate**
> FilepoolTemplates GetFilepoolTemplate(ctx, filepoolTemplateId)


List all templates.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **filepoolTemplateId** | **string**| List all templates. | 

### Return type

[**FilepoolTemplates**](FilepoolTemplates.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFilepoolTemplates**
> FilepoolTemplates GetFilepoolTemplates(ctx, )


List all templates.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**FilepoolTemplates**](FilepoolTemplates.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListFilepoolPolicies**
> FilepoolPolicies ListFilepoolPolicies(ctx, )


List all policies.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**FilepoolPolicies**](FilepoolPolicies.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateFilepoolDefaultPolicy**
> UpdateFilepoolDefaultPolicy(ctx, filepoolDefaultPolicy)


Set default file pool policy.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **filepoolDefaultPolicy** | [**FilepoolDefaultPolicyExtended**](FilepoolDefaultPolicyExtended.md)|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateFilepoolPolicy**
> UpdateFilepoolPolicy(ctx, filepoolPolicy, filepoolPolicyId)


Modify file pool policy. All input fields are optional, but one or more must be supplied.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **filepoolPolicy** | [**FilepoolPolicy**](FilepoolPolicy.md)|  | 
  **filepoolPolicyId** | **string**| Modify file pool policy. All input fields are optional, but one or more must be supplied. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

