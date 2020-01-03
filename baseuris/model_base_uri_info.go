/*
 * AdobeSign BaseURIs APIs
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 6.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package baseuris

type BaseUriInfo struct {
	// The access point from where other APIs need to be accessed. In case other APIs are accessed from a different end point, it will be considered an invalid request
	ApiAccessPoint string `json:"apiAccessPoint"`
	// The access point from where Adobe Sign website can be be accessed
	WebAccessPoint string `json:"webAccessPoint"`
}
