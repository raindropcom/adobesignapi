/*
 * AdobeSign MegaSigns APIs
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 6.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package megaSigns

type MegaSignCreationInfo struct {
	// A list of one or more CCs that will be copied in the megasign transaction. The CCs will each receive an email at the beginning of the transaction and also when the final document is signed. The email addresses will also receive a copy of the document, attached as a PDF file 
	Ccs []MegaSignCcInfo `json:"ccs,omitempty"`
	// Info corresponding to each child agreement of the megaSign 
	ChildAgreementsInfo *ChildAgreementsInfo `json:"childAgreementsInfo"`
	// Date when megasign was created. Format would be yyyy-MM-dd'T'HH:mm:ssZ. For example, e.g 2016-02-25T18:46:19Z represents UTC time
	CreatedDate string `json:"createdDate,omitempty"`
	// Time after which Agreement expires and needs to be signed before it. Format should be yyyy-MM-dd'T'HH:mm:ssZ. For example, e.g 2016-02-25T18:46:19Z represents UTC time. Should not be provided in offline agreement creation.
	ExpirationTime string `json:"expirationTime,omitempty"`
	// An arbitrary value from your system, which can be specified at sending time and then later returned or queried
	ExternalId *ExternalId `json:"externalId,omitempty"`
	// A list of one or more files (or references to files) that will be sent out for signature. If more than one file is provided, they will be combined into one PDF before being sent out. Note: Only one of the four parameters in every FileInfo object must be specified
	FileInfos []FileInfo `json:"fileInfos"`
	// Integer which specifies the delay in hours before sending the first reminder.<br>This is an optional field. The minimum value allowed is 1 hour and the maximum value can’t be more than the difference of agreement creation and expiry time of the agreement in hours.<br>If this is not specified but the reminder frequency is specified, then the first reminder will be sent based on frequency.<br>i.e. if the reminder is created with frequency specified as daily, the firstReminderDelay will be 24 hours
	FirstReminderDelay int32 `json:"firstReminderDelay,omitempty"`
	// The unique identifier of megasign 
	Id string `json:"id,omitempty"`
	// The locale associated with this agreement - specifies the language for the signing page and emails, for example en_US or fr_FR. If none specified, defaults to the language configured for the agreement sender
	Locale string `json:"locale,omitempty"`
	// An optional message to the participants, describing what is being sent or why their signature is required
	Message string `json:"message,omitempty"`
	// The name of the agreement that will be used to identify it, in emails, website and other places
	Name string `json:"name"`
	// URL and associated properties for the success page the user will be taken to after completing the signing process
	PostSignOption *PostSignOption `json:"postSignOption,omitempty"`
	// Optional parameter that sets how often you want to send reminders to the participants. If it is not specified, the default frequency set for the account will be used
	ReminderFrequency string `json:"reminderFrequency,omitempty"`
	// Optional security parameters for the megasign
	SecurityOption *MegaSignSecurityOption `json:"securityOption,omitempty"`
	// Email of agreement sender. Only provided in GET. Can not be provided in POST/PUT request. If provided in POST/PUT, it will be ignored
	SenderEmail string `json:"senderEmail,omitempty"`
	// Specifies the type of signature you would like to request - written or e-signature. The possible values are <br> ESIGN : Agreement needs to be signed electronically <br>, WRITTEN : Agreement will be signed using handwritten signature and signed document will be uploaded into the system
	SignatureType string `json:"signatureType"`
	// State of the Megasign
	State string `json:"state"`
	// Status of the Megasign
	Status string `json:"status"`
	// Vaulting properties that allows Adobe Sign to securely store documents with a vault provider
	VaultingInfo *VaultingInfo `json:"vaultingInfo,omitempty"`
}
