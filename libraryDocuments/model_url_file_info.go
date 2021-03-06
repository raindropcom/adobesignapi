/*
 * AdobeSign Library Documents APIs
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 6.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package libraryDocuments

type UrlFileInfo struct {
	// The mime type of the referenced file, used to determine if the file can be accepted and the necessary conversion steps can be performed
	MimeType string `json:"mimeType"`
	// The original system file name of the document being sent
	Name string `json:"name"`
	// A publicly accessible URL for retrieving the raw file content
	Url string `json:"url"`
}
