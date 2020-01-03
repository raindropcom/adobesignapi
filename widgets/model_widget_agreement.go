/*
 * AdobeSign Widgets APIs
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 6.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package widgets

type WidgetAgreement struct {
	// The display date for the agreement. Format would be yyyy-MM-dd'T'HH:mm:ssZ. For example, e.g 2016-02-25T18:46:19Z represents UTC time
	DisplayDate string `json:"displayDate"`
	// The most relevant current user set for the agreement. It is typically the next signer if the agreement is from the current user, or the sender if received from another user
	DisplayParticipantSetInfos []DisplayWidgetParticipantSetInfo `json:"displayParticipantSetInfos"`
	// True if this is an e-sign document
	Esign bool `json:"esign"`
	// True if agreement is hidden for the user
	Hidden bool `json:"hidden"`
	// The unique identifier of the agreement.If provided in POST, it will simply be ignored
	Id string `json:"id,omitempty"`
	// A version ID which uniquely identifies the current version of the agreement
	LatestVersionId string `json:"latestVersionId"`
	// Name of the Agreement
	Name string `json:"name"`
	// The current status of the document from the perspective of the originator
	Status string `json:"status"`
}
