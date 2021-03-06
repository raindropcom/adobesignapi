/*
 * AdobeSign Widgets APIs
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 6.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package widgets

type SecurityOption struct {
	// The secondary password that will be used to secure the PDF document. Note that AdobeSign will never show this password to anyone, so you will need to separately communicate it to any relevant parties
	OpenPassword string `json:"openPassword,omitempty"`
}
