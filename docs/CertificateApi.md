# \CertificateApi

All URIs are relative to *https://YOUR_CLUSTER_HOSTNAME_OR_NODE_IP:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCertificateServerItem**](CertificateApi.md#CreateCertificateServerItem) | **Post** /platform/4/certificate/server | 
[**DeleteCertificateServerById**](CertificateApi.md#DeleteCertificateServerById) | **Delete** /platform/4/certificate/server/{CertificateServerId} | 
[**GetCertificateServerById**](CertificateApi.md#GetCertificateServerById) | **Get** /platform/4/certificate/server/{CertificateServerId} | 
[**ListCertificateServer**](CertificateApi.md#ListCertificateServer) | **Get** /platform/4/certificate/server | 
[**UpdateCertificateServerById**](CertificateApi.md#UpdateCertificateServerById) | **Put** /platform/4/certificate/server/{CertificateServerId} | 


# **CreateCertificateServerItem**
> CreateResponse CreateCertificateServerItem(ctx, certificateServerItem)


Import a TLS server certificate.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **certificateServerItem** | [**CertificateServerItem**](CertificateServerItem.md)|  | 

### Return type

[**CreateResponse**](CreateResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCertificateServerById**
> DeleteCertificateServerById(ctx, certificateServerId)


Delete a TLS server certificate.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **certificateServerId** | **string**| Delete a TLS server certificate. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCertificateServerById**
> CertificateServer GetCertificateServerById(ctx, certificateServerId)


Retrieve a single TLS server certificate.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **certificateServerId** | **string**| Retrieve a single TLS server certificate. | 

### Return type

[**CertificateServer**](CertificateServer.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListCertificateServer**
> CertificateServerExtended ListCertificateServer(ctx, optional)


Retrieve a list of all configured TLS server certificates.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sort** | **string**| The field that will be used for sorting. | 
 **limit** | **int32**| Return no more than this many results at once (see resume). | 
 **dir** | **string**| The direction of the sort. | 
 **resume** | **string**| Continue returning results from previous call using this token (token should come from the previous call, resume cannot be used with other options). | 

### Return type

[**CertificateServerExtended**](CertificateServerExtended.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCertificateServerById**
> UpdateCertificateServerById(ctx, certificateServerIdParams, certificateServerId)


Modify a TLS server certificate.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **certificateServerIdParams** | [**CertificateServerIdParams**](CertificateServerIdParams.md)|  | 
  **certificateServerId** | **string**| Modify a TLS server certificate. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

