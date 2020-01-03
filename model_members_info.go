/*
 * AdobeSign Aggreements API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 6.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package agreements

type MembersInfo struct {
	// Information of CC participants of the agreement.
	CcsInfo []CcParticipantInfo `json:"ccsInfo"`
	// Information of next participant sets.
	NextParticipantSets []DetailedParticipantSetInfo `json:"nextParticipantSets"`
	// Information about the participant Sets.
	ParticipantSets []DetailedParticipantSetInfo `json:"participantSets"`
	// Information of the sender of the agreement.
	SenderInfo *SenderInfo `json:"senderInfo"`
	// Information of the participants with whom the agreement has been shared.
	SharesInfo []ShareParticipantInfo `json:"sharesInfo"`
}
