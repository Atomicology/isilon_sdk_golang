# CloudAccountCreateParams

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | (S3 only) The user id of the S3 account | [optional] [default to null]
**AccountUsername** | **string** | The username required to authenticate against the cloud service | [default to null]
**BirthClusterId** | **string** | The guid of the cluster where this account was created | [optional] [default to null]
**Enabled** | **bool** | Whether this account is explicitly enabled or disabled by a user | [optional] [default to null]
**Key** | **string** | A valid authentication key for connecting to the cloud | [default to null]
**Name** | **string** | A unique name for this account | [default to null]
**Proxy** | **string** | The id or name of a proxy to be used by this account | [optional] [default to null]
**SkipSslValidation** | **bool** | Indicates whether to skip SSL certificate validation when connecting to the cloud | [optional] [default to null]
**StorageRegion** | **string** | (S3 only) An appropriate region for the S3 account.  For example, faster access times may be gained by referencing a nearby region | [optional] [default to null]
**TelemetryBucket** | **string** | (S3 only) The name of the bucket into which generated metrics reports are placed by the cloud service provider | [optional] [default to null]
**Type_** | **string** | The type of cloud protocol required.  E.g., \&quot;isilon\&quot; for EMC Isilon, \&quot;ecs\&quot; for EMC ECS Appliance, \&quot;virtustream\&quot; for Virtustream Storage Cloud, \&quot;azure\&quot; for Microsoft Azure and \&quot;s3\&quot; for Amazon S3 | [default to null]
**Uri** | **string** | A valid URI pointing to the location of the cloud storage | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


