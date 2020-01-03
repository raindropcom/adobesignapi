/*
 * AdobeSign Aggreements API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 6.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package agreements

type DocumentImageUrlsInfo struct {
	// Id of the document
	DocumentId string `json:"documentId"`
	// A list of documents image URLs.
	DocumentImageUrlsList []DocumentImageUrls `json:"documentImageUrlsList"`
}
