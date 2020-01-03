/*
 * AdobeSign Widgets APIs
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 6.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package widgets

type Document struct {
	// ID of the document. In case of PUT call, this is the only field that is accepted in Document structure. Name and mimeType are ignored in case of PUT call
	Id string `json:"id"`
	// Label of the document
	Label string `json:"label"`
	// mimeType of the original file. This is returned in GET but not accepted back in PUT
	MimeType string `json:"mimeType,omitempty"`
	// Name of the original document uploaded. This is returned in GET but not accepted back in PUT
	Name string `json:"name,omitempty"`
	// Number of pages in the document
	NumPages int32 `json:"numPages"`
}
