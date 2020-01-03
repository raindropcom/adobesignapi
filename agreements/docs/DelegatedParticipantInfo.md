# DelegatedParticipantInfo

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | Email of the participant. In case of modifying a participant set (PUT) this is a required field. In case of GET, this is the required field and will always be returned unless it is a fax workflow (legacy agreements) that were created using fax as input | [default to null]
**SecurityOption** | [***DelegatedParticipantSecurityOption**](DelegatedParticipantSecurityOption.md) | Security options that apply to the participant. This cannot be changed as part of the PUT call | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


