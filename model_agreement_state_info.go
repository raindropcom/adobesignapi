/*
 * AdobeSign Aggreements API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 6.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package agreements

type AgreementStateInfo struct {
	// Cancellation information for the agreement
	AgreementCancellationInfo *AgreementCancellationInfo `json:"agreementCancellationInfo,omitempty"`
	// The state in which the agreement should land
	State string `json:"state"`
}