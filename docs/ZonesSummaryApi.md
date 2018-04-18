# \ZonesSummaryApi

All URIs are relative to *https://YOUR_CLUSTER_HOSTNAME_OR_NODE_IP:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetZonesSummary**](ZonesSummaryApi.md#GetZonesSummary) | **Get** /platform/1/zones-summary | 
[**GetZonesSummaryZone**](ZonesSummaryApi.md#GetZonesSummaryZone) | **Get** /platform/1/zones-summary/{ZonesSummaryZone} | 


# **GetZonesSummary**
> ZonesSummaryExtended GetZonesSummary(ctx, optional)


Retrieve access zone summary information.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupnet** | **string**| Name of groupnet in which to list zones. | 

### Return type

[**ZonesSummaryExtended**](ZonesSummaryExtended.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetZonesSummaryZone**
> ZonesSummary GetZonesSummaryZone(ctx, zonesSummaryZone)


Retrieve non-privileged access zone information.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **zonesSummaryZone** | **int32**| Retrieve non-privileged access zone information. | 

### Return type

[**ZonesSummary**](ZonesSummary.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

