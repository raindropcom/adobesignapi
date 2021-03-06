/*
 * AdobeSign Library Documents APIs
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 6.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package libraryDocuments

type LibraryDocuments struct {
	// An array of document library items
	LibraryDocumentList []LibraryDocument `json:"libraryDocumentList"`
	// Pagination information for navigating through the response
	Page *PageInfo `json:"page"`
}
