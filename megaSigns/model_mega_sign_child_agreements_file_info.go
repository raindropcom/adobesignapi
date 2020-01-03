/*
 * AdobeSign MegaSigns APIs
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 6.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package megaSigns

type MegaSignChildAgreementsFileInfo struct {
	// id of the file containg information about the existing childAgreementsInfo associated with the megaSign. Will be ignored in POST call and in case of GET call, this is the only thing that will be returned. The content of the file can be fetched through GET /megaSigns/{megaSignId}/childAgreementsInfo/{childAgreementsInfoFileId} endpoint.
	ChildAgreementsInfoFileId string `json:"childAgreementsInfoFileId,omitempty"`
	// Input type through which participantSetsInfos will be provided. Whichever input type is provided, the values should be provided in its corresponding value object. Currently we are supporting CSV file format for providing megaSIgn child recipients.
	FileType string `json:"fileType"`
	// Transient id of the input file which contains participantSetsInfos. Currently only csv format is suppported. More details about CSV format <a href='https://www.adobe.com/go/documentcloud_megasigncsv'>here</a>  
	TransientDocumentId string `json:"transientDocumentId"`
}