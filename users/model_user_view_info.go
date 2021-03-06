/*
 * AdobeSign Users APIs
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 6.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package users

type UserViewInfo struct {
	// Common view configuration for all the available views
	CommonViewConfiguration *CommonViewConfiguration `json:"commonViewConfiguration,omitempty"`
	// Name of the requested user view
	Name string `json:"name"`
}
