# LibraryDocumentInfo

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedDate** | **string** | Date when library document was created. Format would be yyyy-MM-dd&#39;T&#39;HH:mm:ssZ. For example, e.g 2016-02-25T18:46:19Z represents UTC time | [optional] [default to null]
**CreatorEmail** | **string** | Email address of the library document creator. It will be ignored in POST call | [optional] [default to null]
**FileInfos** | [**[]FileInfo**](FileInfo.md) | A list of one or more files (or references to files) that will be used to create the template. If more than one file is provided, they will be combined into one PDF. Note: Only a single parameter in every FileInfo object must be specified | [default to null]
**Id** | **string** | The unique identifier that is used to refer to the library template. It will be ignored in POST call | [optional] [default to null]
**Name** | **string** | The name of the library template that will be used to identify it, in emails and on the website | [default to null]
**SharingMode** | **string** | Specifies who should have access to this library document. GLOBAL sharing mode is not applicable in POST/PUT calls | [default to null]
**State** | **string** | State of the library document. | [default to null]
**Status** | **string** | Status of the library document | [optional] [default to null]
**TemplateTypes** | **[]string** | A list of one or more library template types | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


