# \LicenseApi

All URIs are relative to *https://YOUR_CLUSTER_HOSTNAME_OR_NODE_IP:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLicenseLicense**](LicenseApi.md#CreateLicenseLicense) | **Post** /platform/5/license/licenses | 
[**GetLicenseGenerate**](LicenseApi.md#GetLicenseGenerate) | **Get** /platform/5/license/generate | 
[**GetLicenseLicense**](LicenseApi.md#GetLicenseLicense) | **Get** /platform/5/license/licenses/{LicenseLicenseId} | 
[**ListLicenseLicenses**](LicenseApi.md#ListLicenseLicenses) | **Get** /platform/5/license/licenses | 


# **CreateLicenseLicense**
> Empty CreateLicenseLicense(ctx, licenseLicense)


Install a new license file and/or activate evaluation licenses.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **licenseLicense** | [**LicenseLicenseCreateParams**](LicenseLicenseCreateParams.md)|  | 

### Return type

[**Empty**](Empty.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLicenseGenerate**
> LicenseGenerate GetLicenseGenerate(ctx, optional)


Generate license activation file.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **action** | **string**| enum: license_list_only (default), generate_activation, download_activation. Generate an activation file or return a list of activated licenses. If generating an activation file and no licenses are specified, the default configuration is to generate an activation file with the current set of licensed features. download_activation returns HTTP headers and the same XML content as seen in the response activation. | [default to license_list_only]
 **licensesToInclude** | **string**| Licenses to include in activation file. | 
 **licensesToExclude** | **string**| Licenses to omit from activation file. | 
 **onlyTheseLicenses** | **string**| Activate only the defined licenses. This setting overrides all other license activation settings. | 

### Return type

[**LicenseGenerate**](LicenseGenerate.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLicenseLicense**
> LicenseLicenses GetLicenseLicense(ctx, licenseLicenseId)


Retrieve license information for the feature.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **licenseLicenseId** | **string**| Retrieve license information for the feature. | 

### Return type

[**LicenseLicenses**](LicenseLicenses.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListLicenseLicenses**
> LicenseLicensesExtended ListLicenseLicenses(ctx, )


Retrieve license information for all licensable products.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**LicenseLicensesExtended**](LicenseLicensesExtended.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

