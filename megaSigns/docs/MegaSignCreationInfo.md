# MegaSignCreationInfo

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ccs** | [**[]MegaSignCcInfo**](MegaSignCcInfo.md) | A list of one or more CCs that will be copied in the megasign transaction. The CCs will each receive an email at the beginning of the transaction and also when the final document is signed. The email addresses will also receive a copy of the document, attached as a PDF file  | [optional] [default to null]
**ChildAgreementsInfo** | [***ChildAgreementsInfo**](ChildAgreementsInfo.md) | Info corresponding to each child agreement of the megaSign  | [default to null]
**CreatedDate** | **string** | Date when megasign was created. Format would be yyyy-MM-dd&#39;T&#39;HH:mm:ssZ. For example, e.g 2016-02-25T18:46:19Z represents UTC time | [optional] [default to null]
**ExpirationTime** | **string** | Time after which Agreement expires and needs to be signed before it. Format should be yyyy-MM-dd&#39;T&#39;HH:mm:ssZ. For example, e.g 2016-02-25T18:46:19Z represents UTC time. Should not be provided in offline agreement creation. | [optional] [default to null]
**ExternalId** | [***ExternalId**](ExternalId.md) | An arbitrary value from your system, which can be specified at sending time and then later returned or queried | [optional] [default to null]
**FileInfos** | [**[]FileInfo**](FileInfo.md) | A list of one or more files (or references to files) that will be sent out for signature. If more than one file is provided, they will be combined into one PDF before being sent out. Note: Only one of the four parameters in every FileInfo object must be specified | [default to null]
**FirstReminderDelay** | **int32** | Integer which specifies the delay in hours before sending the first reminder.&lt;br&gt;This is an optional field. The minimum value allowed is 1 hour and the maximum value can’t be more than the difference of agreement creation and expiry time of the agreement in hours.&lt;br&gt;If this is not specified but the reminder frequency is specified, then the first reminder will be sent based on frequency.&lt;br&gt;i.e. if the reminder is created with frequency specified as daily, the firstReminderDelay will be 24 hours | [optional] [default to null]
**Id** | **string** | The unique identifier of megasign  | [optional] [default to null]
**Locale** | **string** | The locale associated with this agreement - specifies the language for the signing page and emails, for example en_US or fr_FR. If none specified, defaults to the language configured for the agreement sender | [optional] [default to null]
**Message** | **string** | An optional message to the participants, describing what is being sent or why their signature is required | [optional] [default to null]
**Name** | **string** | The name of the agreement that will be used to identify it, in emails, website and other places | [default to null]
**PostSignOption** | [***PostSignOption**](PostSignOption.md) | URL and associated properties for the success page the user will be taken to after completing the signing process | [optional] [default to null]
**ReminderFrequency** | **string** | Optional parameter that sets how often you want to send reminders to the participants. If it is not specified, the default frequency set for the account will be used | [optional] [default to null]
**SecurityOption** | [***MegaSignSecurityOption**](MegaSignSecurityOption.md) | Optional security parameters for the megasign | [optional] [default to null]
**SenderEmail** | **string** | Email of agreement sender. Only provided in GET. Can not be provided in POST/PUT request. If provided in POST/PUT, it will be ignored | [optional] [default to null]
**SignatureType** | **string** | Specifies the type of signature you would like to request - written or e-signature. The possible values are &lt;br&gt; ESIGN : Agreement needs to be signed electronically &lt;br&gt;, WRITTEN : Agreement will be signed using handwritten signature and signed document will be uploaded into the system | [default to null]
**State** | **string** | State of the Megasign | [default to null]
**Status** | **string** | Status of the Megasign | [default to null]
**VaultingInfo** | [***VaultingInfo**](VaultingInfo.md) | Vaulting properties that allows Adobe Sign to securely store documents with a vault provider | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


