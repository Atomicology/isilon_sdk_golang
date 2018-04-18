/*
 * Isilon SDK
 *
 * Isilon SDK - Language bindings for the OneFS API
 *
 * API version: 5
 * Contact: sdk@isilon.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package isi_sdk_8_1_0

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

type SyncTargetApiService service


/* SyncTargetApiService 
 Cancel the most recent SyncIQ job for this policy from the target side.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param policiesPolicyCancelItem 
 @param policy 
 @return CreateResponse*/
func (a *SyncTargetApiService) CreatePoliciesPolicyCancelItem(ctx context.Context, policiesPolicyCancelItem Empty, policy string) (CreateResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  CreateResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/platform/1/sync/target/policies/{Policy}/cancel"
	localVarPath = strings.Replace(localVarPath, "{"+"Policy"+"}", fmt.Sprintf("%v", policy), -1)

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
	localVarPostBody = &policiesPolicyCancelItem
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

/* SyncTargetApiService 
 View a single SyncIQ target subreport.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param reportsReportSubreportId View a single SyncIQ target subreport.
 @param rid 
 @return ReportsReportSubreports*/
func (a *SyncTargetApiService) GetReportsReportSubreport(ctx context.Context, reportsReportSubreportId string, rid string) (ReportsReportSubreports,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  ReportsReportSubreports
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/platform/4/sync/target/reports/{Rid}/subreports/{ReportsReportSubreportId}"
	localVarPath = strings.Replace(localVarPath, "{"+"ReportsReportSubreportId"+"}", fmt.Sprintf("%v", reportsReportSubreportId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"Rid"+"}", fmt.Sprintf("%v", rid), -1)

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

/* SyncTargetApiService 
 Get a list of SyncIQ target subreports for a report.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param rid 
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "sort" (string) The field that will be used for sorting.
     @param "resume" (string) Continue returning results from previous call using this token (token should come from the previous call, resume cannot be used with other options).
     @param "newerThan" (int32) Filter the returned reports to include only those whose jobs started more recently than the specified number of days ago.
     @param "state" (string) Filter the returned reports to include only those whose jobs are in this state.
     @param "limit" (int32) Return no more than this many results at once (see resume).
     @param "dir" (string) The direction of the sort.
 @return ReportsReportSubreportsExtended*/
func (a *SyncTargetApiService) GetReportsReportSubreports(ctx context.Context, rid string, localVarOptionals map[string]interface{}) (ReportsReportSubreportsExtended,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  ReportsReportSubreportsExtended
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/platform/4/sync/target/reports/{Rid}/subreports"
	localVarPath = strings.Replace(localVarPath, "{"+"Rid"+"}", fmt.Sprintf("%v", rid), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["sort"], "string", "sort"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["resume"], "string", "resume"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["newerThan"], "int32", "newerThan"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["state"], "string", "state"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["limit"], "int32", "limit"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["dir"], "string", "dir"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["sort"].(string); localVarOk {
		localVarQueryParams.Add("sort", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["resume"].(string); localVarOk {
		localVarQueryParams.Add("resume", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["newerThan"].(int32); localVarOk {
		localVarQueryParams.Add("newer_than", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["state"].(string); localVarOk {
		localVarQueryParams.Add("state", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["limit"].(int32); localVarOk {
		localVarQueryParams.Add("limit", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["dir"].(string); localVarOk {
		localVarQueryParams.Add("dir", parameterToString(localVarTempParam, ""))
	}
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
