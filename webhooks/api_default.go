/*
 * AdobeSign WebHooks APIs
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 6.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package webhooks

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"fmt"
	"github.com/antihax/optional"
)

// Linger please
var (
	_ context.Context
)

type DefaultApiService service

/* 
DefaultApiService Creates a webhook.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param authorization An &lt;a href&#x3D;\&quot;#\&quot; onclick&#x3D;\&quot;this.href&#x3D;oauthDoc()\&quot; oncontextmenu&#x3D;\&quot;this.href&#x3D;oauthDoc()\&quot; target&#x3D;\&quot;oauthDoc\&quot;&gt;OAuth Access Token&lt;/a&gt; with scopes:&lt;ul&gt;&lt;li style&#x3D;&#39;list-style-type: square&#39;&gt;&lt;a href&#x3D;\&quot;#\&quot; onclick&#x3D;\&quot;this.href&#x3D;oauthDoc(&#39;webhook_write&#39;)\&quot; oncontextmenu&#x3D;\&quot;this.href&#x3D;oauthDoc(&#39;webhook_write&#39;)\&quot; target&#x3D;\&quot;oauthDoc\&quot;&gt;webhook_write&lt;/a&gt;&lt;/li&gt;&lt;/ul&gt;in the format &lt;b&gt;&#39;Bearer {accessToken}&#39;.
 * @param webhookInfo Information about the webhook that you want to create
 * @param optional nil or *CreateWebhookOpts - Optional Parameters:
     * @param "XApiUser" (optional.String) -  The userId or email of API caller using the account or group token in the format &lt;b&gt;userid:{userId} OR email:{email}.&lt;/b&gt; If it is not specified, then the caller is inferred from the token.
     * @param "XOnBehalfOfUser" (optional.String) -  The userId or email in the format &lt;b&gt;userid:{userId} OR email:{email}.&lt;/b&gt; of the user that has shared his/her account

@return WebhookCreationResponse
*/

type CreateWebhookOpts struct { 
	XApiUser optional.String
	XOnBehalfOfUser optional.String
}

func (a *DefaultApiService) CreateWebhook(ctx context.Context, authorization string, webhookInfo WebhookInfo, localVarOptionals *CreateWebhookOpts) (WebhookCreationResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue WebhookCreationResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/webhooks"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	localVarHeaderParams["Authorization"] = parameterToString(authorization, "")
	if localVarOptionals != nil && localVarOptionals.XApiUser.IsSet() {
		localVarHeaderParams["x-api-user"] = parameterToString(localVarOptionals.XApiUser.Value(), "")
	}
	if localVarOptionals != nil && localVarOptionals.XOnBehalfOfUser.IsSet() {
		localVarHeaderParams["x-on-behalf-of-user"] = parameterToString(localVarOptionals.XOnBehalfOfUser.Value(), "")
	}
	// body params
	localVarPostBody = &webhookInfo
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		
		if localVarHttpResponse.StatusCode == 200 {
			var v WebhookCreationResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/* 
DefaultApiService Deletes a webhook.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param authorization An &lt;a href&#x3D;\&quot;#\&quot; onclick&#x3D;\&quot;this.href&#x3D;oauthDoc()\&quot; oncontextmenu&#x3D;\&quot;this.href&#x3D;oauthDoc()\&quot; target&#x3D;\&quot;oauthDoc\&quot;&gt;OAuth Access Token&lt;/a&gt; with scopes:&lt;ul&gt;&lt;li style&#x3D;&#39;list-style-type: square&#39;&gt;&lt;a href&#x3D;\&quot;#\&quot; onclick&#x3D;\&quot;this.href&#x3D;oauthDoc(&#39;webhook_retention&#39;)\&quot; oncontextmenu&#x3D;\&quot;this.href&#x3D;oauthDoc(&#39;webhook_retention&#39;)\&quot; target&#x3D;\&quot;oauthDoc\&quot;&gt;webhook_retention&lt;/a&gt;&lt;/li&gt;&lt;/ul&gt;in the format &lt;b&gt;&#39;Bearer {accessToken}&#39;.
 * @param ifMatch The server will only update the resource if it matches the listed ETag otherwise error RESOURCE_MODIFIED(412) is returned.
 * @param webhookId The webhook identifier, as returned by the webhook creation API or retrieved from the API to fetch webhooks.
 * @param optional nil or *DeleteWebhookOpts - Optional Parameters:
     * @param "XApiUser" (optional.String) -  The userId or email of API caller using the account or group token in the format &lt;b&gt;userid:{userId} OR email:{email}.&lt;/b&gt; If it is not specified, then the caller is inferred from the token.
     * @param "XOnBehalfOfUser" (optional.String) -  The userId or email in the format &lt;b&gt;userid:{userId} OR email:{email}.&lt;/b&gt; of the user that has shared his/her account


*/

type DeleteWebhookOpts struct { 
	XApiUser optional.String
	XOnBehalfOfUser optional.String
}

func (a *DefaultApiService) DeleteWebhook(ctx context.Context, authorization string, ifMatch string, webhookId string, localVarOptionals *DeleteWebhookOpts) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/webhooks/{webhookId}"
	localVarPath = strings.Replace(localVarPath, "{"+"webhookId"+"}", fmt.Sprintf("%v", webhookId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	localVarHeaderParams["Authorization"] = parameterToString(authorization, "")
	if localVarOptionals != nil && localVarOptionals.XApiUser.IsSet() {
		localVarHeaderParams["x-api-user"] = parameterToString(localVarOptionals.XApiUser.Value(), "")
	}
	if localVarOptionals != nil && localVarOptionals.XOnBehalfOfUser.IsSet() {
		localVarHeaderParams["x-on-behalf-of-user"] = parameterToString(localVarOptionals.XOnBehalfOfUser.Value(), "")
	}
	localVarHeaderParams["If-Match"] = parameterToString(ifMatch, "")
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarHttpResponse, err
	}


	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}

/* 
DefaultApiService Retrieves the details of a webhook.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param authorization An &lt;a href&#x3D;\&quot;#\&quot; onclick&#x3D;\&quot;this.href&#x3D;oauthDoc()\&quot; oncontextmenu&#x3D;\&quot;this.href&#x3D;oauthDoc()\&quot; target&#x3D;\&quot;oauthDoc\&quot;&gt;OAuth Access Token&lt;/a&gt; with scopes:&lt;ul&gt;&lt;li style&#x3D;&#39;list-style-type: square&#39;&gt;&lt;a href&#x3D;\&quot;#\&quot; onclick&#x3D;\&quot;this.href&#x3D;oauthDoc(&#39;webhook_read&#39;)\&quot; oncontextmenu&#x3D;\&quot;this.href&#x3D;oauthDoc(&#39;webhook_read&#39;)\&quot; target&#x3D;\&quot;oauthDoc\&quot;&gt;webhook_read&lt;/a&gt;&lt;/li&gt;&lt;/ul&gt;in the format &lt;b&gt;&#39;Bearer {accessToken}&#39;.
 * @param webhookId The webhook identifier, as returned by the webhook creation API or retrieved from the API to fetch webhooks.
 * @param optional nil or *GetWebhookInfoOpts - Optional Parameters:
     * @param "XApiUser" (optional.String) -  The userId or email of API caller using the account or group token in the format &lt;b&gt;userid:{userId} OR email:{email}.&lt;/b&gt; If it is not specified, then the caller is inferred from the token.
     * @param "XOnBehalfOfUser" (optional.String) -  The userId or email in the format &lt;b&gt;userid:{userId} OR email:{email}.&lt;/b&gt; of the user that has shared his/her account
     * @param "IfNoneMatch" (optional.String) -  Pass the value of the e-tag header obtained from the previous response to the same request to get a RESOURCE_NOT_MODIFIED(304) if the resource hasn&#39;t changed.

@return WebhookInfo
*/

type GetWebhookInfoOpts struct { 
	XApiUser optional.String
	XOnBehalfOfUser optional.String
	IfNoneMatch optional.String
}

func (a *DefaultApiService) GetWebhookInfo(ctx context.Context, authorization string, webhookId string, localVarOptionals *GetWebhookInfoOpts) (WebhookInfo, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue WebhookInfo
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/webhooks/{webhookId}"
	localVarPath = strings.Replace(localVarPath, "{"+"webhookId"+"}", fmt.Sprintf("%v", webhookId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	localVarHeaderParams["Authorization"] = parameterToString(authorization, "")
	if localVarOptionals != nil && localVarOptionals.XApiUser.IsSet() {
		localVarHeaderParams["x-api-user"] = parameterToString(localVarOptionals.XApiUser.Value(), "")
	}
	if localVarOptionals != nil && localVarOptionals.XOnBehalfOfUser.IsSet() {
		localVarHeaderParams["x-on-behalf-of-user"] = parameterToString(localVarOptionals.XOnBehalfOfUser.Value(), "")
	}
	if localVarOptionals != nil && localVarOptionals.IfNoneMatch.IsSet() {
		localVarHeaderParams["If-None-Match"] = parameterToString(localVarOptionals.IfNoneMatch.Value(), "")
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		
		if localVarHttpResponse.StatusCode == 200 {
			var v WebhookInfo
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/* 
DefaultApiService Retrieves webhooks for a user.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param authorization An &lt;a href&#x3D;\&quot;#\&quot; onclick&#x3D;\&quot;this.href&#x3D;oauthDoc()\&quot; oncontextmenu&#x3D;\&quot;this.href&#x3D;oauthDoc()\&quot; target&#x3D;\&quot;oauthDoc\&quot;&gt;OAuth Access Token&lt;/a&gt; with scopes:&lt;ul&gt;&lt;li style&#x3D;&#39;list-style-type: square&#39;&gt;&lt;a href&#x3D;\&quot;#\&quot; onclick&#x3D;\&quot;this.href&#x3D;oauthDoc(&#39;webhook_read&#39;)\&quot; oncontextmenu&#x3D;\&quot;this.href&#x3D;oauthDoc(&#39;webhook_read&#39;)\&quot; target&#x3D;\&quot;oauthDoc\&quot;&gt;webhook_read&lt;/a&gt;&lt;/li&gt;&lt;/ul&gt;in the format &lt;b&gt;&#39;Bearer {accessToken}&#39;.
 * @param optional nil or *GetWebhooksOpts - Optional Parameters:
     * @param "XApiUser" (optional.String) -  The userId or email of API caller using the account or group token in the format &lt;b&gt;userid:{userId} OR email:{email}.&lt;/b&gt; If it is not specified, then the caller is inferred from the token.
     * @param "XOnBehalfOfUser" (optional.String) -  The userId or email in the format &lt;b&gt;userid:{userId} OR email:{email}.&lt;/b&gt; of the user that has shared his/her account
     * @param "ShowInActiveWebhooks" (optional.Bool) -  A query parameter to fetch all the inactive webhooks along with the active webhooks.
     * @param "Scope" (optional.String) -  Scope of webhook. The possible values are ACCOUNT, GROUP, USER or RESOURCE
     * @param "ResourceType" (optional.String) -  The type of resource on which webhook was created. The possible values are AGREEMENT, WIDGET and MEGASIGN.
     * @param "Cursor" (optional.String) -  Used to navigate through the pages. If not provided, returns the first page.
     * @param "PageSize" (optional.Int32) -  Number of intended items in the response page.

@return UserWebhooks
*/

type GetWebhooksOpts struct { 
	XApiUser optional.String
	XOnBehalfOfUser optional.String
	ShowInActiveWebhooks optional.Bool
	Scope optional.String
	ResourceType optional.String
	Cursor optional.String
	PageSize optional.Int32
}

func (a *DefaultApiService) GetWebhooks(ctx context.Context, authorization string, localVarOptionals *GetWebhooksOpts) (UserWebhooks, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue UserWebhooks
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/webhooks"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.ShowInActiveWebhooks.IsSet() {
		localVarQueryParams.Add("showInActiveWebhooks", parameterToString(localVarOptionals.ShowInActiveWebhooks.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Scope.IsSet() {
		localVarQueryParams.Add("scope", parameterToString(localVarOptionals.Scope.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ResourceType.IsSet() {
		localVarQueryParams.Add("resourceType", parameterToString(localVarOptionals.ResourceType.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Cursor.IsSet() {
		localVarQueryParams.Add("cursor", parameterToString(localVarOptionals.Cursor.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PageSize.IsSet() {
		localVarQueryParams.Add("pageSize", parameterToString(localVarOptionals.PageSize.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	localVarHeaderParams["Authorization"] = parameterToString(authorization, "")
	if localVarOptionals != nil && localVarOptionals.XApiUser.IsSet() {
		localVarHeaderParams["x-api-user"] = parameterToString(localVarOptionals.XApiUser.Value(), "")
	}
	if localVarOptionals != nil && localVarOptionals.XOnBehalfOfUser.IsSet() {
		localVarHeaderParams["x-on-behalf-of-user"] = parameterToString(localVarOptionals.XOnBehalfOfUser.Value(), "")
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		
		if localVarHttpResponse.StatusCode == 200 {
			var v UserWebhooks
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/* 
DefaultApiService Updates a webhook.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param authorization An &lt;a href&#x3D;\&quot;#\&quot; onclick&#x3D;\&quot;this.href&#x3D;oauthDoc()\&quot; oncontextmenu&#x3D;\&quot;this.href&#x3D;oauthDoc()\&quot; target&#x3D;\&quot;oauthDoc\&quot;&gt;OAuth Access Token&lt;/a&gt; with scopes:&lt;ul&gt;&lt;li style&#x3D;&#39;list-style-type: square&#39;&gt;&lt;a href&#x3D;\&quot;#\&quot; onclick&#x3D;\&quot;this.href&#x3D;oauthDoc(&#39;webhook_write&#39;)\&quot; oncontextmenu&#x3D;\&quot;this.href&#x3D;oauthDoc(&#39;webhook_write&#39;)\&quot; target&#x3D;\&quot;oauthDoc\&quot;&gt;webhook_write&lt;/a&gt;&lt;/li&gt;&lt;/ul&gt;in the format &lt;b&gt;&#39;Bearer {accessToken}&#39;.
 * @param ifMatch The server will only update the resource if it matches the listed ETag otherwise error RESOURCE_MODIFIED(412) is returned.
 * @param webhookId The webhook identifier, as returned by the webhook creation API or retrieved from the API to fetch webhooks.
 * @param webhookInfo Information necessary to update a webhook
 * @param optional nil or *UpdateWebhookOpts - Optional Parameters:
     * @param "XApiUser" (optional.String) -  The userId or email of API caller using the account or group token in the format &lt;b&gt;userid:{userId} OR email:{email}.&lt;/b&gt; If it is not specified, then the caller is inferred from the token.
     * @param "XOnBehalfOfUser" (optional.String) -  The userId or email in the format &lt;b&gt;userid:{userId} OR email:{email}.&lt;/b&gt; of the user that has shared his/her account


*/

type UpdateWebhookOpts struct { 
	XApiUser optional.String
	XOnBehalfOfUser optional.String
}

func (a *DefaultApiService) UpdateWebhook(ctx context.Context, authorization string, ifMatch string, webhookId string, webhookInfo WebhookInfo, localVarOptionals *UpdateWebhookOpts) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/webhooks/{webhookId}"
	localVarPath = strings.Replace(localVarPath, "{"+"webhookId"+"}", fmt.Sprintf("%v", webhookId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	localVarHeaderParams["Authorization"] = parameterToString(authorization, "")
	if localVarOptionals != nil && localVarOptionals.XApiUser.IsSet() {
		localVarHeaderParams["x-api-user"] = parameterToString(localVarOptionals.XApiUser.Value(), "")
	}
	if localVarOptionals != nil && localVarOptionals.XOnBehalfOfUser.IsSet() {
		localVarHeaderParams["x-on-behalf-of-user"] = parameterToString(localVarOptionals.XOnBehalfOfUser.Value(), "")
	}
	localVarHeaderParams["If-Match"] = parameterToString(ifMatch, "")
	// body params
	localVarPostBody = &webhookInfo
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarHttpResponse, err
	}


	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}

/* 
DefaultApiService Updates the state of a webhook identified by webhookId in the path.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param authorization An &lt;a href&#x3D;\&quot;#\&quot; onclick&#x3D;\&quot;this.href&#x3D;oauthDoc()\&quot; oncontextmenu&#x3D;\&quot;this.href&#x3D;oauthDoc()\&quot; target&#x3D;\&quot;oauthDoc\&quot;&gt;OAuth Access Token&lt;/a&gt; with scopes:&lt;ul&gt;&lt;li style&#x3D;&#39;list-style-type: square&#39;&gt;&lt;a href&#x3D;\&quot;#\&quot; onclick&#x3D;\&quot;this.href&#x3D;oauthDoc(&#39;webhook_write&#39;)\&quot; oncontextmenu&#x3D;\&quot;this.href&#x3D;oauthDoc(&#39;webhook_write&#39;)\&quot; target&#x3D;\&quot;oauthDoc\&quot;&gt;webhook_write&lt;/a&gt;&lt;/li&gt;&lt;/ul&gt;in the format &lt;b&gt;&#39;Bearer {accessToken}&#39;.
 * @param ifMatch The server will only update the resource if it matches the listed ETag otherwise error RESOURCE_MODIFIED(412) is returned.
 * @param webhookId The webhook identifier, as returned by the webhook creation API or retrieved from the API to fetch webhooks.
 * @param webhookStateInfo
 * @param optional nil or *UpdateWebhookStateOpts - Optional Parameters:
     * @param "XApiUser" (optional.String) -  The userId or email of API caller using the account or group token in the format &lt;b&gt;userid:{userId} OR email:{email}.&lt;/b&gt; If it is not specified, then the caller is inferred from the token.
     * @param "XOnBehalfOfUser" (optional.String) -  The userId or email in the format &lt;b&gt;userid:{userId} OR email:{email}.&lt;/b&gt; of the user that has shared his/her account


*/

type UpdateWebhookStateOpts struct { 
	XApiUser optional.String
	XOnBehalfOfUser optional.String
}

func (a *DefaultApiService) UpdateWebhookState(ctx context.Context, authorization string, ifMatch string, webhookId string, webhookStateInfo WebhookStateInfo, localVarOptionals *UpdateWebhookStateOpts) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/webhooks/{webhookId}/state"
	localVarPath = strings.Replace(localVarPath, "{"+"webhookId"+"}", fmt.Sprintf("%v", webhookId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	localVarHeaderParams["Authorization"] = parameterToString(authorization, "")
	if localVarOptionals != nil && localVarOptionals.XApiUser.IsSet() {
		localVarHeaderParams["x-api-user"] = parameterToString(localVarOptionals.XApiUser.Value(), "")
	}
	if localVarOptionals != nil && localVarOptionals.XOnBehalfOfUser.IsSet() {
		localVarHeaderParams["x-on-behalf-of-user"] = parameterToString(localVarOptionals.XOnBehalfOfUser.Value(), "")
	}
	localVarHeaderParams["If-Match"] = parameterToString(ifMatch, "")
	// body params
	localVarPostBody = &webhookStateInfo
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarHttpResponse, err
	}


	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}
