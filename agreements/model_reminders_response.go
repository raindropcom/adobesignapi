/*
 * AdobeSign Aggreements API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 6.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package agreements

type RemindersResponse struct {
	// A list of one or more reminders created on the agreement specified by the unique identifier agreementId by the user invoking the API.
	ReminderInfoList []ReminderInfo `json:"reminderInfoList"`
}
