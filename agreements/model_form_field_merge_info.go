/*
 * AdobeSign Aggreements API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 6.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package agreements

// Data for updating the default values of existing form fields
type FormFieldMergeInfo struct {
	// A mapping of field names to default values
	FieldMergeInfos []MergefieldInfo `json:"fieldMergeInfos,omitempty"`
}
