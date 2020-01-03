/*
 * AdobeSign Library Documents APIs
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 6.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package libraryDocuments

type LibraryDocumentEvent struct {
	// Email address of the user that created the event
	ActingUserEmail string `json:"actingUserEmail"`
	// The IP address of the user that created the event
	ActingUserIpAddress string `json:"actingUserIpAddress"`
	// The name of the acting user
	ActingUserName string `json:"actingUserName"`
	// The event comment. For RECALLED or REJECTED, the reason given by the user that initiates the event. For DELEGATE or SHARE, the message from the acting user to the participant
	Comment string `json:"comment,omitempty"`
	// The date of the audit event. Format would be yyyy-MM-dd'T'HH:mm:ssZ. For example, e.g 2016-02-25T18:46:19Z represents UTC time
	Date string `json:"date"`
	// A description of the audit event
	Description string `json:"description"`
	// Location of the device that generated the event (This value may be null due to limited privileges)
	DeviceLocation *LibDocEventDeviceLocation `json:"deviceLocation"`
	// Phone number from the device used when the participation is completed on a mobile phone
	DevicePhoneNumber string `json:"devicePhoneNumber"`
	// This is present for ESIGNED events when the participation is signed digitally
	DigitalSignatureInfo *DigitalSignatureInfo `json:"digitalSignatureInfo"`
	// Email address of the user that initiated the event on behalf of the acting user when the account is shared. Will be empty if there is no account sharing in effect
	InitiatingUserEmail string `json:"initiatingUserEmail"`
	// Full name of the user that initiated the event on behalf of the acting user when the account is shared. Will be empty if there is no account sharing in effect
	InitiatingUserName string `json:"initiatingUserName"`
	// Email address of the user that is the participant for the event. This may be different than the acting user for certain event types. For example, for a DELEGATION event, this is the user who was delegated to
	ParticipantEmail string `json:"participantEmail"`
	// The unique identifier of the participant for the event. This may be different than the acting user for certain event types. For example, for a DELEGATION event, this is the user who was delegated to
	ParticipantId string `json:"participantId"`
	// Role assumed by all participants in the participant set the participant belongs to (signer, approver etc.).
	ParticipantRole string `json:"participantRole"`
	// A unique identifier linking offline events to synchronization events (specified for offline signing events and synchronization events, else null)
	SynchronizationId string `json:"synchronizationId"`
	// Type of library document event
	Type_ string `json:"type"`
	// The identifier assigned by the vault provider for the vault event (if vaulted, otherwise null)
	VaultEventId string `json:"vaultEventId"`
	// Name of the vault provider for the vault event (if vaulted, otherwise null)
	VaultProviderName string `json:"vaultProviderName"`
	// An ID which uniquely identifies the version of the document associated with this audit event
	VersionId string `json:"versionId"`
}
