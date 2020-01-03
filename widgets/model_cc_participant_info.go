/*
 * AdobeSign Widgets APIs
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 6.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package widgets

type CcParticipantInfo struct {
	// Company of the CC participant, if available.
	Company string `json:"company"`
	// Email of the CC participant of the widget
	Email string `json:"email"`
	// Name of the CC participant, if available.
	Name string `json:"name"`
	//  The unique identifier of the CC participant of the widget.
	ParticipantId string `json:"participantId"`
}