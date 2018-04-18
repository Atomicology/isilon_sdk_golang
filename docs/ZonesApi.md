# \ZonesApi

All URIs are relative to *https://YOUR_CLUSTER_HOSTNAME_OR_NODE_IP:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateZone**](ZonesApi.md#CreateZone) | **Post** /platform/3/zones | 
[**DeleteZone**](ZonesApi.md#DeleteZone) | **Delete** /platform/3/zones/{ZoneId} | 
[**GetZone**](ZonesApi.md#GetZone) | **Get** /platform/3/zones/{ZoneId} | 
[**ListZones**](ZonesApi.md#ListZones) | **Get** /platform/3/zones | 
[**UpdateZone**](ZonesApi.md#UpdateZone) | **Put** /platform/3/zones/{ZoneId} | 


# **CreateZone**
> CreateResponse CreateZone(ctx, zone)


Create a new access zone.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **zone** | [**ZoneCreateParams**](ZoneCreateParams.md)|  | 

### Return type

[**CreateResponse**](CreateResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteZone**
> DeleteZone(ctx, zoneId)


Delete the access zone.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **zoneId** | **int32**| Delete the access zone. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetZone**
> Zones GetZone(ctx, zoneId)


Retrieve the access zone information.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **zoneId** | **int32**| Retrieve the access zone information. | 

### Return type

[**Zones**](Zones.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListZones**
> ZonesExtended ListZones(ctx, )


List all access zones.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ZonesExtended**](ZonesExtended.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateZone**
> UpdateZone(ctx, zone, zoneId)


Modify the access zone. All input fields are optional, but one or more must be supplied.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **zone** | [**Zone**](Zone.md)|  | 
  **zoneId** | **int32**| Modify the access zone. All input fields are optional, but one or more must be supplied. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

