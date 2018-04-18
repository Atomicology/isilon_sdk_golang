# LicenseLicenseCreateParams

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Evaluation** | **[]string** | A list of evaluation licenses to enable on the cluster. | [optional] [default to null]
**LicenseFileContent** | **string** | License file string content. The license file is obtained from EMC&#39;s SLC web portal. Do not use with the license_file_path option. | [optional] [default to null]
**LicenseFilePath** | **string** | Path to new license file, must be under /ifs. The license file is obtained from EMC&#39;s SLC web portal. Do not include the path when only enabling evaluation licenses. Do not use with the license_file_content option. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


