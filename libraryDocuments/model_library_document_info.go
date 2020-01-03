/*
 * AdobeSign Library Documents APIs
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 6.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package libraryDocuments

type LibraryDocumentInfo struct {
	// Date when library document was created. Format would be yyyy-MM-dd'T'HH:mm:ssZ. For example, e.g 2016-02-25T18:46:19Z represents UTC time
	CreatedDate string `json:"createdDate,omitempty"`
	// Email address of the library document creator. It will be ignored in POST call
	CreatorEmail string `json:"creatorEmail,omitempty"`
	// A list of one or more files (or references to files) that will be used to create the template. If more than one file is provided, they will be combined into one PDF. Note: Only a single parameter in every FileInfo object must be specified
	FileInfos []FileInfo `json:"fileInfos"`
	// The unique identifier that is used to refer to the library template. It will be ignored in POST call
	Id string `json:"id,omitempty"`
	// The name of the library template that will be used to identify it, in emails and on the website
	Name string `json:"name"`
	// Specifies who should have access to this library document. GLOBAL sharing mode is not applicable in POST/PUT calls
	SharingMode string `json:"sharingMode"`
	// State of the library document.
	State string `json:"state"`
	// Status of the library document
	Status string `json:"status,omitempty"`
	// A list of one or more library template types
	TemplateTypes []string `json:"templateTypes"`
}
