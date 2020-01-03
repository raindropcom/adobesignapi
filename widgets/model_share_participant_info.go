/*
 * AdobeSign Widgets APIs
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 6.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package widgets

type ShareParticipantInfo struct {
	// Company of the sharee participant of the widget, if available.
	Company string `json:"company"`
	// Email of the sharee participant of the widget.
	Email string `json:"email"`
	// Name of the sharee participant of the widget, if available.
	Name string `json:"name"`
	// The unique identifier of the sharee participant of the widget.
	ParticipantId string `json:"participantId"`
	// The unique identifier of the participant who shared the widget.
	SharerParticipantId string `json:"sharerParticipantId"`
}
