# \SyncPoliciesApi

All URIs are relative to *https://YOUR_CLUSTER_HOSTNAME_OR_NODE_IP:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePolicyResetItem**](SyncPoliciesApi.md#CreatePolicyResetItem) | **Post** /platform/1/sync/policies/{Policy}/reset | 


# **CreatePolicyResetItem**
> CreateResponse CreatePolicyResetItem(ctx, policyResetItem, policy)


Reset a SyncIQ policy incremental state and force a full sync/copy.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **policyResetItem** | [**Empty**](Empty.md)|  | 
  **policy** | **string**|  | 

### Return type

[**CreateResponse**](CreateResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

