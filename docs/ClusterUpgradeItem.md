# ClusterUpgradeItem

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstallImagePath** | **string** | The location (path) of the upgrade image which must be within /ifs. | [optional] [default to null]
**NodesToRollingUpgrade** | **[]int32** | The nodes (to be) scheduled for upgrade ordered by queue position number. Null if the cluster_state is &#39;partially upgraded&#39; or upgrade_type is &#39;simultaneous&#39;. One of the following values: [&lt;lnn-1&gt;, &lt;lnn-2&gt;, ... ], &#39;All&#39;, null | [optional] [default to null]
**SkipOptional** | **bool** | Used to indicate that the pre-upgrade check should be skipped | [optional] [default to null]
**UpgradeType** | **string** | The type of upgrade to perform. One of the following values: &#39;rolling&#39;, &#39;simultaneous&#39; | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


