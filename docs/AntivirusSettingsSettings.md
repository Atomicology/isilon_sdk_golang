# AntivirusSettingsSettings

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FailOpen** | **bool** | Allow access when scanning fails. | [optional] [default to null]
**GlobFilters** | **[]string** | Glob patterns for leaf filenames. | [optional] [default to null]
**GlobFiltersEnabled** | **bool** | Enable glob filters. | [optional] [default to null]
**GlobFiltersInclude** | **bool** | If true, only scan files matching a glob filter. If false, only scan files that don&#39;t match a glob filter. | [optional] [default to null]
**PathPrefixes** | **[]string** | Paths to include in realtime scans. | [optional] [default to null]
**Quarantine** | **bool** | Try to quarantine files when threats are found. | [optional] [default to null]
**Repair** | **bool** | Try to repair files when threats are found. | [optional] [default to null]
**ReportExpiry** | **int32** | Amount of time in seconds until old reporting data is purged. | [optional] [default to null]
**ScanOnClose** | **bool** | Scan files when apps close them. | [optional] [default to null]
**ScanOnOpen** | **bool** | Scan files on access. | [optional] [default to null]
**ScanSizeMaximum** | **int32** | Skip scanning files larger than this. | [optional] [default to null]
**Service** | **bool** | Whether the antivirus service is enabled. | [optional] [default to null]
**Truncate** | **bool** | Try to truncate files when threats are found. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


