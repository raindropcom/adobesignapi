/*
 * AdobeSign Transient Documents APIs
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 6.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package transientDocuments

type TransientDocumentResponse struct {
	// The unique identifier of the uploaded document that can be used in an agreement or a megaSign or widget creation call
	TransientDocumentId string `json:"transientDocumentId"`
}
