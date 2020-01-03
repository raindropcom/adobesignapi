/*
 * AdobeSign Widgets APIs
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 6.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package widgets

type WidgetAgreements struct {
	// Pagination information for navigating through the response
	Page *PageInfo `json:"page"`
	// An array of UserAgreement items
	UserAgreementList []WidgetAgreement `json:"userAgreementList"`
}