# NodeDriveconfigNodeStall

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClearTime** | **int32** | The amount of time in seconds with no stalls before ignoring previous stalls. | [optional] [default to null]
**DiskscrubStripes** | **int32** | Number of stripes to read during a diskscrub. | [optional] [default to null]
**MaxErrorFrequency** | **int32** | The number of errors during stalled drive disk exercises to cause the drive to be softfailed. | [optional] [default to null]
**MaxSlowAccess** | **int32** | The number of slow accesses during stalled drive disk exercises to cause the drive to be softfailed. | [optional] [default to null]
**MaxSlowFrequency** | **int32** | The number of slow frequency triggers during stalled drive disk exercises to cause the drive to be softfailed. | [optional] [default to null]
**MaxTotalStallTime** | **int32** | The maximum amount of time, in seconds, to remain stalled before softfailing the drive. | [optional] [default to null]
**ScanMaxEccDelay** | **int32** | Maximum delay in seconds after an ECC correction during a scan. | [optional] [default to null]
**ScanSize** | **int32** | Total bytes of error-free reads to complete a scan. | [optional] [default to null]
**Sleep** | **int32** | Delay in seconds between evaluations. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


