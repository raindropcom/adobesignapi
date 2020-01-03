/*
 * AdobeSign Users APIs
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 6.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package users

type UsersInfo struct {
	// Pagination information for navigating through the response
	Page *PageInfo `json:"page"`
	// The list of users in the account.
	UserInfoList []UserInfo `json:"userInfoList"`
}
