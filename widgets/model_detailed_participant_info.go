/*
 * AdobeSign Widgets APIs
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 6.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package widgets

type DetailedParticipantInfo struct {
	// The company of the participant, if available. This cannot be changed as part of the PUT call.
	Company string `json:"company,omitempty"`
	// Email of the participant. In case of modifying a participant set (PUT) this is a required field. In case of GET, this is the required field and will always be returned unless it is a fax workflow (legacy agreements) that were created using fax as input
	Email string `json:"email"`
	// Fax of the participant. New Agreements can not be created with fax option. This is only returned for legacy agreements created with fax as participants
	Fax string `json:"fax,omitempty"`
	// The unique identifier of the participant. This will be returned as part of Get call but is not mandatory to be passed as part of PUT call for agreements/{id}/members/participantSets/{id}.
	Id string `json:"id,omitempty"`
	// The name of the participant, if available. This cannot be changed as part of the PUT call.
	Name string `json:"name,omitempty"`
	// The private message of the participant, if available. This cannot be changed as part of the PUT call.
	PrivateMessage string `json:"privateMessage,omitempty"`
	// Security options that apply to the participant. This cannot be changed as part of the PUT call
	SecurityOption *ParticipantSecurityOption `json:"securityOption"`
	// True if this participant is the same user that is calling the API. Returned as part of Get. Ignored (not required) if modifying a participant set (PUT).
	Self bool `json:"self,omitempty"`
	// The status of the participant. This cannot be changed as part of the PUT call.
	Status string `json:"status,omitempty"`
}
