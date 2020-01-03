/*
 * AdobeSign Aggreements API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 6.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package agreements

type ParticipantSecurityOption struct {
	// The authentication method for the participants to have access to view and sign the document
	AuthenticationMethod string `json:"authenticationMethod"`
	// The password required for the participant to view and sign the document. Note that AdobeSign will never show this password to anyone, so you will need to separately communicate it to any relevant parties. The password will not be returned in GET call. In case of PUT call, password associated with Agreement resource will remain unchanged if no password is specified but authentication method is provided as PASSWORD
	Password string `json:"password,omitempty"`
	// The phoneInfo required for the participant to view and sign the document
	PhoneInfo *PhoneInfo `json:"phoneInfo,omitempty"`
}
