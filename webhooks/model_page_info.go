/*
 * AdobeSign WebHooks APIs
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 6.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package webhooks

type PageInfo struct {
	// Used to navigate to the next page. If not returned, there are no further pages.
	NextCursor string `json:"nextCursor"`
}
