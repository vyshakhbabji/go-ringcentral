/*
 * RingCentral Connect Platform API Explorer
 *
 * <p>This is a beta interactive API explorer for the RingCentral Connect Platform. To use this service, you will need to have an account with the proper credentials to generate an OAuth2 access token.</p><p><h2>Quick Start</h2></p><ol><li>1) Go to <b>Authentication > /oauth/token</b></li><li>2) Enter <b>app_key, app_secret, username, password</b> fields and then click \"Try it out!\"</li><li>3) Upon success, your access_token is loaded and you can access any form requiring authorization.</li></ol><h2>Links</h2><ul><li><a href=\"https://github.com/ringcentral\" target=\"_blank\">RingCentral SDKs on Github</a></li><li><a href=\"mailto:devsupport@ringcentral.com\">RingCentral Developer Support Email</a></li></ul>
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package ringcentral

import (
	"io/ioutil"
	"net/url"
	"net/http"
	"strings"
	"golang.org/x/net/context"
	"encoding/json"
	"fmt"
)

// Linger please
var (
	_ context.Context
)

type ReportingApiService service


/* ReportingApiService Get Reporting Settings
 &lt;p style&#x3D;&#39;font-style:italic;&#39;&gt;Since 1.0.17 (Release 7.2)&lt;/p&gt;&lt;p&gt;Returns user-defined configuration of CFA (Customer Facing Analytics) reports.&lt;/p&gt;&lt;h4&gt;Required Permissions&lt;/h4&gt;&lt;table class&#x3D;&#39;fullwidth&#39;&gt;&lt;thead&gt;&lt;tr&gt;&lt;th&gt;Permission&lt;/th&gt;&lt;th&gt;Description&lt;/th&gt;&lt;/tr&gt;&lt;/thead&gt;&lt;tbody&gt;&lt;tr&gt;&lt;td class&#x3D;&#39;code&#39;&gt;ReadAccounts&lt;/td&gt;&lt;td&gt;Viewing user account info (including name, business name, address and phone number/account number)&lt;/td&gt;&lt;/tr&gt;&lt;/tbody&gt;&lt;/table&gt;&lt;h4&gt;API Group&lt;/h4&gt;&lt;p&gt;Light&lt;/p&gt;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param accountId Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session
 @param extensionId Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session
 @return ReportingSettings*/
func (a *ReportingApiService) LoadReportingSettings(ctx context.Context, accountId string, extensionId string) (ReportingSettings,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  ReportingSettings
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/restapi/v1.0/account/{accountId}/extension/{extensionId}/reporting/settings"
	localVarPath = strings.Replace(localVarPath, "{"+"accountId"+"}", fmt.Sprintf("%v", accountId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"extensionId"+"}", fmt.Sprintf("%v", extensionId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* ReportingApiService Update Reporting Settings
 &lt;p style&#x3D;&#39;font-style:italic;&#39;&gt;&lt;/p&gt;&lt;p&gt;Sets user-defined configuration of CFA (Customer Facing Analytics) reports.&lt;/p&gt;&lt;h4&gt;Required Permissions&lt;/h4&gt;&lt;table class&#x3D;&#39;fullwidth&#39;&gt;&lt;thead&gt;&lt;tr&gt;&lt;th&gt;Permission&lt;/th&gt;&lt;th&gt;Description&lt;/th&gt;&lt;/tr&gt;&lt;/thead&gt;&lt;tbody&gt;&lt;tr&gt;&lt;td class&#x3D;&#39;code&#39;&gt;EditAccounts&lt;/td&gt;&lt;td&gt;Viewing and updating user account info (including name, business name, address and phone number/account number)&lt;/td&gt;&lt;/tr&gt;&lt;/tbody&gt;&lt;/table&gt;&lt;h4&gt;API Group&lt;/h4&gt;&lt;p&gt;Light&lt;/p&gt;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param accountId Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session
 @param extensionId Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session
 @param body JSON body
 @return ReportingSettings*/
func (a *ReportingApiService) UpdateReportingSettings(ctx context.Context, accountId string, extensionId string, body SetReportingSettingsRequest) (ReportingSettings,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  ReportingSettings
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/restapi/v1.0/account/{accountId}/extension/{extensionId}/reporting/settings"
	localVarPath = strings.Replace(localVarPath, "{"+"accountId"+"}", fmt.Sprintf("%v", accountId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"extensionId"+"}", fmt.Sprintf("%v", extensionId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

