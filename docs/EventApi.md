# \EventApi

All URIs are relative to *https://YOUR_CLUSTER_HOSTNAME_OR_NODE_IP:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEventAlertCondition**](EventApi.md#CreateEventAlertCondition) | **Post** /platform/4/event/alert-conditions | 
[**CreateEventChannel**](EventApi.md#CreateEventChannel) | **Post** /platform/3/event/channels | 
[**CreateEventEvent**](EventApi.md#CreateEventEvent) | **Post** /platform/3/event/events | 
[**DeleteEventAlertCondition**](EventApi.md#DeleteEventAlertCondition) | **Delete** /platform/4/event/alert-conditions/{EventAlertConditionId} | 
[**DeleteEventAlertConditions**](EventApi.md#DeleteEventAlertConditions) | **Delete** /platform/4/event/alert-conditions | 
[**DeleteEventChannel**](EventApi.md#DeleteEventChannel) | **Delete** /platform/3/event/channels/{EventChannelId} | 
[**GetEventAlertCondition**](EventApi.md#GetEventAlertCondition) | **Get** /platform/4/event/alert-conditions/{EventAlertConditionId} | 
[**GetEventCategories**](EventApi.md#GetEventCategories) | **Get** /platform/3/event/categories | 
[**GetEventCategory**](EventApi.md#GetEventCategory) | **Get** /platform/3/event/categories/{EventCategoryId} | 
[**GetEventChannel**](EventApi.md#GetEventChannel) | **Get** /platform/3/event/channels/{EventChannelId} | 
[**GetEventEventgroupDefinition**](EventApi.md#GetEventEventgroupDefinition) | **Get** /platform/4/event/eventgroup-definitions/{EventEventgroupDefinitionId} | 
[**GetEventEventgroupDefinitions**](EventApi.md#GetEventEventgroupDefinitions) | **Get** /platform/4/event/eventgroup-definitions | 
[**GetEventEventgroupOccurrence**](EventApi.md#GetEventEventgroupOccurrence) | **Get** /platform/3/event/eventgroup-occurrences/{EventEventgroupOccurrenceId} | 
[**GetEventEventgroupOccurrences**](EventApi.md#GetEventEventgroupOccurrences) | **Get** /platform/3/event/eventgroup-occurrences | 
[**GetEventEventlist**](EventApi.md#GetEventEventlist) | **Get** /platform/3/event/eventlists/{EventEventlistId} | 
[**GetEventEventlists**](EventApi.md#GetEventEventlists) | **Get** /platform/3/event/eventlists | 
[**GetEventSettings**](EventApi.md#GetEventSettings) | **Get** /platform/3/event/settings | 
[**ListEventAlertConditions**](EventApi.md#ListEventAlertConditions) | **Get** /platform/4/event/alert-conditions | 
[**ListEventChannels**](EventApi.md#ListEventChannels) | **Get** /platform/3/event/channels | 
[**UpdateEventAlertCondition**](EventApi.md#UpdateEventAlertCondition) | **Put** /platform/4/event/alert-conditions/{EventAlertConditionId} | 
[**UpdateEventChannel**](EventApi.md#UpdateEventChannel) | **Put** /platform/3/event/channels/{EventChannelId} | 
[**UpdateEventEventgroupOccurrence**](EventApi.md#UpdateEventEventgroupOccurrence) | **Put** /platform/3/event/eventgroup-occurrences/{EventEventgroupOccurrenceId} | 
[**UpdateEventEventgroupOccurrences**](EventApi.md#UpdateEventEventgroupOccurrences) | **Put** /platform/3/event/eventgroup-occurrences | 
[**UpdateEventSettings**](EventApi.md#UpdateEventSettings) | **Put** /platform/3/event/settings | 


# **CreateEventAlertCondition**
> CreateResponse CreateEventAlertCondition(ctx, eventAlertCondition)


Create a new alert condition.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **eventAlertCondition** | [**EventAlertConditionCreateParams**](EventAlertConditionCreateParams.md)|  | 

### Return type

[**CreateResponse**](CreateResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateEventChannel**
> CreateResponse CreateEventChannel(ctx, eventChannel)


Create a new channel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **eventChannel** | [**EventChannelCreateParams**](EventChannelCreateParams.md)|  | 

### Return type

[**CreateResponse**](CreateResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateEventEvent**
> CreateQuotaReportResponse CreateEventEvent(ctx, eventEvent)


Create a test event.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **eventEvent** | [**EventEvent**](EventEvent.md)|  | 

### Return type

[**CreateQuotaReportResponse**](CreateQuotaReportResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteEventAlertCondition**
> DeleteEventAlertCondition(ctx, eventAlertConditionId)


Delete the alert-condition.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **eventAlertConditionId** | **string**| Delete the alert-condition. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteEventAlertConditions**
> DeleteEventAlertConditions(ctx, optional)


Bulk delete of alert conditions.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **channel** | **string**| Delete only conditions for this channel | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteEventChannel**
> DeleteEventChannel(ctx, eventChannelId)


Delete the channel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **eventChannelId** | **string**| Delete the channel. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEventAlertCondition**
> EventAlertConditions GetEventAlertCondition(ctx, eventAlertConditionId)


Retrieve the alert-condition.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **eventAlertConditionId** | **string**| Retrieve the alert-condition. | 

### Return type

[**EventAlertConditions**](EventAlertConditions.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEventCategories**
> EventCategoriesExtended GetEventCategories(ctx, optional)


List all eventgroup categories.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32**| Return no more than this many results at once (see resume). | 
 **resume** | **string**| Continue returning results from previous call using this token (token should come from the previous call, resume cannot be used with other options). | 

### Return type

[**EventCategoriesExtended**](EventCategoriesExtended.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEventCategory**
> EventCategories GetEventCategory(ctx, eventCategoryId)


Retrieve the eventgroup category.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **eventCategoryId** | **string**| Retrieve the eventgroup category. | 

### Return type

[**EventCategories**](EventCategories.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEventChannel**
> EventChannels GetEventChannel(ctx, eventChannelId)


Retrieve the channel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **eventChannelId** | **string**| Retrieve the channel. | 

### Return type

[**EventChannels**](EventChannels.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEventEventgroupDefinition**
> EventEventgroupDefinitions GetEventEventgroupDefinition(ctx, eventEventgroupDefinitionId)


Retrieve the eventgroup definition.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **eventEventgroupDefinitionId** | **string**| Retrieve the eventgroup definition. | 

### Return type

[**EventEventgroupDefinitions**](EventEventgroupDefinitions.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEventEventgroupDefinitions**
> EventEventgroupDefinitionsExtended GetEventEventgroupDefinitions(ctx, optional)


List all eventgroup definitions.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **category** | **int32**| Return eventgroups in the specified category | 
 **limit** | **int32**| Return no more than this many results at once (see resume). | 
 **resume** | **string**| Continue returning results from previous call using this token (token should come from the previous call, resume cannot be used with other options). | 

### Return type

[**EventEventgroupDefinitionsExtended**](EventEventgroupDefinitionsExtended.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEventEventgroupOccurrence**
> EventEventgroupOccurrences GetEventEventgroupOccurrence(ctx, eventEventgroupOccurrenceId)


Retrieve individual eventgroup occurrence.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **eventEventgroupOccurrenceId** | **string**| Retrieve individual eventgroup occurrence. | 

### Return type

[**EventEventgroupOccurrences**](EventEventgroupOccurrences.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEventEventgroupOccurrences**
> EventEventgroupOccurrencesExtended GetEventEventgroupOccurrences(ctx, optional)


List all eventgroup occurrences.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resolved** | **bool**| Filter for resolved eventgroups | 
 **sort** | **string**| The field that will be used for sorting. | 
 **begin** | **int32**| events that are in progress after this time | 
 **end** | **int32**| events that were in progress before this time | 
 **eventCount** | **int32**| events for which event_count &gt; this | 
 **resume** | **string**| Continue returning results from previous call using this token (token should come from the previous call, resume cannot be used with other options). | 
 **ignore** | **bool**| Filter for ignored eventgroups | 
 **limit** | **int32**| Return no more than this many results at once (see resume). | 
 **resolver** | **string**| Filter for eventgroup resolver | 
 **cause** | **string**| Filter for cause | 
 **dir** | **string**| The direction of the sort. | 

### Return type

[**EventEventgroupOccurrencesExtended**](EventEventgroupOccurrencesExtended.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEventEventlist**
> EventEventlists GetEventEventlist(ctx, eventEventlistId)


Retrieve the list of events for a eventgroup occureence.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **eventEventlistId** | **string**| Retrieve the list of events for a eventgroup occureence. | 

### Return type

[**EventEventlists**](EventEventlists.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEventEventlists**
> EventEventlistsExtended GetEventEventlists(ctx, optional)


List all event occurrences grouped by eventgroup occurrence.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eventInstance** | **string**| Return only this event occurrence | 
 **limit** | **int32**| Return no more than this many results at once (see resume). | 
 **resume** | **string**| Continue returning results from previous call using this token (token should come from the previous call, resume cannot be used with other options). | 

### Return type

[**EventEventlistsExtended**](EventEventlistsExtended.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEventSettings**
> EventSettings GetEventSettings(ctx, )


Retrieve the settings.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**EventSettings**](EventSettings.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListEventAlertConditions**
> EventAlertConditionsExtended ListEventAlertConditions(ctx, optional)


List all alert conditions.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **channels** | **string**| Return only conditions for the specified channel: | 
 **sort** | **string**| The field that will be used for sorting. | 
 **limit** | **int32**| Return no more than this many results at once (see resume). | 
 **dir** | **string**| The direction of the sort. | 
 **resume** | **string**| Continue returning results from previous call using this token (token should come from the previous call, resume cannot be used with other options). | 

### Return type

[**EventAlertConditionsExtended**](EventAlertConditionsExtended.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListEventChannels**
> EventChannelsExtended ListEventChannels(ctx, optional)


List all channels.

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

[**EventChannelsExtended**](EventChannelsExtended.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateEventAlertCondition**
> UpdateEventAlertCondition(ctx, eventAlertCondition, eventAlertConditionId)


Modify the alert-condition

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **eventAlertCondition** | [**EventAlertCondition**](EventAlertCondition.md)|  | 
  **eventAlertConditionId** | **string**| Modify the alert-condition | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateEventChannel**
> UpdateEventChannel(ctx, eventChannel, eventChannelId)


Modify the channel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **eventChannel** | [**EventChannel**](EventChannel.md)|  | 
  **eventChannelId** | **string**| Modify the channel. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateEventEventgroupOccurrence**
> UpdateEventEventgroupOccurrence(ctx, eventEventgroupOccurrence, eventEventgroupOccurrenceId)


modify eventgroup occurrence.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **eventEventgroupOccurrence** | [**EventEventgroupOccurrence**](EventEventgroupOccurrence.md)|  | 
  **eventEventgroupOccurrenceId** | **string**| modify eventgroup occurrence. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateEventEventgroupOccurrences**
> UpdateEventEventgroupOccurrences(ctx, eventEventgroupOccurrences)


Modify all eventgroup occurrences, resolve or ignore all

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **eventEventgroupOccurrences** | [**EventEventgroupOccurrence**](EventEventgroupOccurrence.md)|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateEventSettings**
> UpdateEventSettings(ctx, eventSettings)


Update settings

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **eventSettings** | [**EventSettings**](EventSettings.md)|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

