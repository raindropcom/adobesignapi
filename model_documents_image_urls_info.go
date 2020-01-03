/*
 * AdobeSign Aggreements API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 6.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package agreements

type DocumentsImageUrlsInfo struct {
	// A list of original document image URLs info.
	OriginalDocumentsImageUrlsInfo []DocumentImageUrlsInfo `json:"originalDocumentsImageUrlsInfo"`
	// A list of supporting document image URLs info.
	SupportingDocumentsImageUrlsInfo []DocumentImageUrlsInfo `json:"supportingDocumentsImageUrlsInfo,omitempty"`
}
