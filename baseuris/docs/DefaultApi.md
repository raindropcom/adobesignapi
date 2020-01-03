# \DefaultApi

All URIs are relative to *https://localhost/api/rest/v6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBaseUris**](DefaultApi.md#GetBaseUris) | **Get** /baseUris | Gets the base uri to access other APIs. In case other APIs are accessed from a different end point, it will be considered an invalid request.


# **GetBaseUris**
> BaseUriInfo GetBaseUris(ctx, authorization)
Gets the base uri to access other APIs. In case other APIs are accessed from a different end point, it will be considered an invalid request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authorization** | **string**| An &lt;a href&#x3D;\&quot;#\&quot; onclick&#x3D;\&quot;this.href&#x3D;oauthDoc()\&quot; oncontextmenu&#x3D;\&quot;this.href&#x3D;oauthDoc()\&quot; target&#x3D;\&quot;oauthDoc\&quot;&gt;OAuth Access Token&lt;/a&gt; with any of the valid scopes&lt;ul&gt;&lt;/ul&gt;in the format &lt;b&gt;&#39;Bearer {accessToken}&#39;. | 

### Return type

[**BaseUriInfo**](BaseUriInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

