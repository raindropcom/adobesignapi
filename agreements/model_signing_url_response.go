/*
 * AdobeSign Aggreements API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 6.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package agreements

type SigningUrlResponse struct {
	// An array of urls for signer sets involved in this agreement.
	SigningUrlSetInfos []SigningUrlSetInfo `json:"signingUrlSetInfos"`
}
