# DetailedWidgetParticipantSetInfo

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier of the participant set. This cannot be changed as part of the PUT call. | [optional] [default to null]
**MemberInfos** | [**[]DetailedParticipantInfo**](DetailedParticipantInfo.md) | Array of ParticipantInfo objects, containing participant-specific data (e.g. email). All participants in the array belong to the same set | [default to null]
**Order** | **int32** | Index indicating sequential signing group (specified for hybrid routing). This cannot be changed as part of the PUT call. | [default to null]
**Role** | **string** | Role assumed by all participants in the set (signer, approver etc.). This cannot be changed as part of the PUT call. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


