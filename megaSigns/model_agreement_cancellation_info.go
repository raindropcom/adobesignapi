/*
 * AdobeSign MegaSigns APIs
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 6.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package megaSigns

type AgreementCancellationInfo struct {
	// An optional comment describing to the recipients why you want to cancel the transaction
	Comment string `json:"comment,omitempty"`
	// Whether or not you would like the recipients to be notified that the transaction has been cancelled. The default value is false
	NotifyOthers bool `json:"notifyOthers,omitempty"`
}
