# MegasignEvent

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActingUserEmail** | **string** | Email address of the user that created the event | [default to null]
**ActingUserIpAddress** | **string** | The IP address of the user that created the event | [default to null]
**ActingUserName** | **string** | The name of the acting user | [default to null]
**Comment** | **string** | The event comment. For RECALLED or REJECTED, the reason given by the user that initiates the event. For DELEGATE or SHARE, the message from the acting user to the participant | [optional] [default to null]
**Date** | **string** | The date of the audit event. Format would be yyyy-MM-dd&#39;T&#39;HH:mm:ssZ. For example, e.g 2016-02-25T18:46:19Z represents UTC time | [default to null]
**Description** | **string** | A description of the audit event | [default to null]
**DeviceLocation** | [***DeviceLocation**](DeviceLocation.md) | Location of the device that generated the event (This value may be null due to limited privileges) | [default to null]
**DevicePhoneNumber** | **string** | Phone number from the device used when the participation is completed on a mobile phone | [default to null]
**DigitalSignatureInfo** | [***DigitalSignatureInfo**](DigitalSignatureInfo.md) | This is present for ESIGNED events when the participation is signed digitally | [default to null]
**InitiatingUserEmail** | **string** | Email address of the user that initiated the event on behalf of the acting user when the account is shared. Will be empty if there is no account sharing in effect | [default to null]
**InitiatingUserName** | **string** | Full name of the user that initiated the event on behalf of the acting user when the account is shared. Will be empty if there is no account sharing in effect | [default to null]
**ParticipantEmail** | **string** | Email address of the user that is the participant for the event. This may be different than the acting user for certain event types. For example, for a DELEGATION event, this is the user who was delegated to | [default to null]
**ParticipantId** | **string** | The unique identifier of the participant for the event. This may be different than the acting user for certain event types. For example, for a DELEGATION event, this is the user who was delegated to | [default to null]
**ParticipantRole** | **string** | Role assumed by all participants in the participant set the participant belongs to (signer, approver etc.). | [default to null]
**SynchronizationId** | **string** | A unique identifier linking offline events to synchronization events (specified for offline signing events and synchronization events, else null) | [default to null]
**Type_** | **string** | Type of MegaSign event | [default to null]
**VaultEventId** | **string** | The identifier assigned by the vault provider for the vault event (if vaulted, otherwise null) | [default to null]
**VaultProviderName** | **string** | Name of the vault provider for the vault event (if vaulted, otherwise null) | [default to null]
**VersionId** | **string** | An ID which uniquely identifies the version of the document associated with this audit event | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


