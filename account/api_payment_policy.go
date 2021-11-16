/*
Account API

The <b>Account API</b> gives sellers the ability to configure their eBay seller accounts, including the seller's policies (the Fulfillment Policy, Payment Policy, and Return Policy), opt in and out of eBay seller programs, configure sales tax tables, and get account information.  <br><br>For details on the availability of the methods in this API, see <a href=\"/api-docs/sell/account/overview.html#requirements\">Account API requirements and restrictions</a>.

API version: v1.6.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package account

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

// PaymentPolicyApiService PaymentPolicyApi service
type PaymentPolicyApiService service

type ApiCreatePaymentPolicyRequest struct {
	ctx _context.Context
	ApiService *PaymentPolicyApiService
	paymentPolicyRequest *PaymentPolicyRequest
}

// Payment policy request
func (r ApiCreatePaymentPolicyRequest) PaymentPolicyRequest(paymentPolicyRequest PaymentPolicyRequest) ApiCreatePaymentPolicyRequest {
	r.paymentPolicyRequest = &paymentPolicyRequest
	return r
}

func (r ApiCreatePaymentPolicyRequest) Execute() (SetPaymentPolicyResponse, *_nethttp.Response, error) {
	return r.ApiService.CreatePaymentPolicyExecute(r)
}

/*
CreatePaymentPolicy Method for CreatePaymentPolicy

This method creates a new payment policy where the policy encapsulates seller's terms for purchase payments.  <br><br>Each policy targets a <b>marketplaceId</b> and <code>categoryTypes.</code><b>name</b> combination and you can create multiple policies for each combination. Be aware that some marketplaces require a specific payment policy for vehicle listings.  <br><br>A successful request returns the URI to the new policy in the <b>Location</b> response header and the ID for the new policy is returned in the response payload.  <p class="tablenote"><b>Tip:</b> For details on creating and using the business policies supported by the Account API, see <a href="/api-docs/sell/static/seller-accounts/business-policies.html">eBay business policies</a>.</p>  <p><b>Marketplaces and locales</b></p>  <p>Policy instructions can be localized by providing a locale in the <code>Accept-Language</code> HTTP request header. For example, the following setting displays field values from the request body in German: <code>Accept-Language: de-DE</code>.</p>  <p>Target the specific locale of a marketplace that supports multiple locales using the <code>Content-Language</code> request header. For example, target the French locale of the Canadian marketplace by specifying the <code>fr-CA</code> locale for <code>Content-Language</code>. Likewise, target the Dutch locale of the Belgium marketplace by setting <code>Content-Language: nl-BE</code>.</p> <p class="tablenote"><b>Tip:</b> For details on headers, see <a href="/api-docs/static/rest-request-components.html#HTTP">HTTP request headers</a>.</p>

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreatePaymentPolicyRequest
*/
func (a *PaymentPolicyApiService) CreatePaymentPolicy(ctx _context.Context) ApiCreatePaymentPolicyRequest {
	return ApiCreatePaymentPolicyRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return SetPaymentPolicyResponse
func (a *PaymentPolicyApiService) CreatePaymentPolicyExecute(r ApiCreatePaymentPolicyRequest) (SetPaymentPolicyResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  SetPaymentPolicyResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PaymentPolicyApiService.CreatePaymentPolicy")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/payment_policy"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.paymentPolicyRequest == nil {
		return localVarReturnValue, nil, reportError("paymentPolicyRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.paymentPolicyRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiDeletePaymentPolicyRequest struct {
	ctx _context.Context
	ApiService *PaymentPolicyApiService
	paymentPolicyId string
}


func (r ApiDeletePaymentPolicyRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.DeletePaymentPolicyExecute(r)
}

/*
DeletePaymentPolicy Method for DeletePaymentPolicy

This method deletes a payment policy. Supply the ID of the policy you want to delete in the <b>paymentPolicyId</b> path parameter. Note that you cannot delete the default payment policy.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param paymentPolicyId This path parameter specifies the ID of the payment policy you want to delete.
 @return ApiDeletePaymentPolicyRequest
*/
func (a *PaymentPolicyApiService) DeletePaymentPolicy(ctx _context.Context, paymentPolicyId string) ApiDeletePaymentPolicyRequest {
	return ApiDeletePaymentPolicyRequest{
		ApiService: a,
		ctx: ctx,
		paymentPolicyId: paymentPolicyId,
	}
}

// Execute executes the request
func (a *PaymentPolicyApiService) DeletePaymentPolicyExecute(r ApiDeletePaymentPolicyRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PaymentPolicyApiService.DeletePaymentPolicy")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/payment_policy/{payment_policy_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"payment_policy_id"+"}", _neturl.PathEscape(parameterToString(r.paymentPolicyId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiGetPaymentPoliciesRequest struct {
	ctx _context.Context
	ApiService *PaymentPolicyApiService
	marketplaceId *string
}

// This query parameter specifies the eBay marketplace of the policies you want to retrieve. For implementation help, refer to eBay API documentation at https://developer.ebay.com/api-docs/sell/account/types/ba:MarketplaceIdEnum
func (r ApiGetPaymentPoliciesRequest) MarketplaceId(marketplaceId string) ApiGetPaymentPoliciesRequest {
	r.marketplaceId = &marketplaceId
	return r
}

func (r ApiGetPaymentPoliciesRequest) Execute() (PaymentPolicyResponse, *_nethttp.Response, error) {
	return r.ApiService.GetPaymentPoliciesExecute(r)
}

/*
GetPaymentPolicies Method for GetPaymentPolicies

This method retrieves all the payment policies configured for the marketplace you specify using the <code>marketplace_id</code> query parameter.  <br><br><b>Marketplaces and locales</b>  <br><br>Get the correct policies for a marketplace that supports multiple locales using the <code>Content-Language</code> request header. For example, get the policies for the French locale of the Canadian marketplace by specifying <code>fr-CA</code> for the <code>Content-Language</code> header. Likewise, target the Dutch locale of the Belgium marketplace by setting <code>Content-Language: nl-BE</code>. For details on header values, see <a href="/api-docs/static/rest-request-components.html#HTTP" target="_blank">HTTP request headers</a>.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetPaymentPoliciesRequest
*/
func (a *PaymentPolicyApiService) GetPaymentPolicies(ctx _context.Context) ApiGetPaymentPoliciesRequest {
	return ApiGetPaymentPoliciesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PaymentPolicyResponse
func (a *PaymentPolicyApiService) GetPaymentPoliciesExecute(r ApiGetPaymentPoliciesRequest) (PaymentPolicyResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  PaymentPolicyResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PaymentPolicyApiService.GetPaymentPolicies")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/payment_policy"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.marketplaceId == nil {
		return localVarReturnValue, nil, reportError("marketplaceId is required and must be specified")
	}

	localVarQueryParams.Add("marketplace_id", parameterToString(*r.marketplaceId, ""))
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetPaymentPolicyRequest struct {
	ctx _context.Context
	ApiService *PaymentPolicyApiService
	paymentPolicyId string
}


func (r ApiGetPaymentPolicyRequest) Execute() (PaymentPolicy, *_nethttp.Response, error) {
	return r.ApiService.GetPaymentPolicyExecute(r)
}

/*
GetPaymentPolicy Method for GetPaymentPolicy

This method retrieves the complete details of a payment policy. Supply the ID of the policy you want to retrieve using the <b>paymentPolicyId</b> path parameter.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param paymentPolicyId This path parameter specifies the ID of the payment policy you want to retrieve.
 @return ApiGetPaymentPolicyRequest
*/
func (a *PaymentPolicyApiService) GetPaymentPolicy(ctx _context.Context, paymentPolicyId string) ApiGetPaymentPolicyRequest {
	return ApiGetPaymentPolicyRequest{
		ApiService: a,
		ctx: ctx,
		paymentPolicyId: paymentPolicyId,
	}
}

// Execute executes the request
//  @return PaymentPolicy
func (a *PaymentPolicyApiService) GetPaymentPolicyExecute(r ApiGetPaymentPolicyRequest) (PaymentPolicy, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  PaymentPolicy
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PaymentPolicyApiService.GetPaymentPolicy")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/payment_policy/{payment_policy_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"payment_policy_id"+"}", _neturl.PathEscape(parameterToString(r.paymentPolicyId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetPaymentPolicyByNameRequest struct {
	ctx _context.Context
	ApiService *PaymentPolicyApiService
	marketplaceId *string
	name *string
}

// This query parameter specifies the eBay marketplace of the policy you want to retrieve. For implementation help, refer to eBay API documentation at https://developer.ebay.com/api-docs/sell/account/types/ba:MarketplaceIdEnum
func (r ApiGetPaymentPolicyByNameRequest) MarketplaceId(marketplaceId string) ApiGetPaymentPolicyByNameRequest {
	r.marketplaceId = &marketplaceId
	return r
}
// This query parameter specifies the user-defined name of the payment policy you want to retrieve.
func (r ApiGetPaymentPolicyByNameRequest) Name(name string) ApiGetPaymentPolicyByNameRequest {
	r.name = &name
	return r
}

func (r ApiGetPaymentPolicyByNameRequest) Execute() (PaymentPolicy, *_nethttp.Response, error) {
	return r.ApiService.GetPaymentPolicyByNameExecute(r)
}

/*
GetPaymentPolicyByName Method for GetPaymentPolicyByName

This method retrieves the complete details of a single payment policy. Supply both the policy <code>name</code> and its associated <code>marketplace_id</code> in the request query parameters.   <br><br><b>Marketplaces and locales</b>  <br><br>Get the correct policy for a marketplace that supports multiple locales using the <code>Content-Language</code> request header. For example, get a policy for the French locale of the Canadian marketplace by specifying <code>fr-CA</code> for the <code>Content-Language</code> header. Likewise, target the Dutch locale of the Belgium marketplace by setting <code>Content-Language: nl-BE</code>. For details on header values, see <a href="/api-docs/static/rest-request-components.html#HTTP">HTTP request headers</a>.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetPaymentPolicyByNameRequest
*/
func (a *PaymentPolicyApiService) GetPaymentPolicyByName(ctx _context.Context) ApiGetPaymentPolicyByNameRequest {
	return ApiGetPaymentPolicyByNameRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PaymentPolicy
func (a *PaymentPolicyApiService) GetPaymentPolicyByNameExecute(r ApiGetPaymentPolicyByNameRequest) (PaymentPolicy, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  PaymentPolicy
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PaymentPolicyApiService.GetPaymentPolicyByName")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/payment_policy/get_by_policy_name"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.marketplaceId == nil {
		return localVarReturnValue, nil, reportError("marketplaceId is required and must be specified")
	}
	if r.name == nil {
		return localVarReturnValue, nil, reportError("name is required and must be specified")
	}

	localVarQueryParams.Add("marketplace_id", parameterToString(*r.marketplaceId, ""))
	localVarQueryParams.Add("name", parameterToString(*r.name, ""))
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiUpdatePaymentPolicyRequest struct {
	ctx _context.Context
	ApiService *PaymentPolicyApiService
	paymentPolicyId string
	paymentPolicyRequest *PaymentPolicyRequest
}

// Payment policy request
func (r ApiUpdatePaymentPolicyRequest) PaymentPolicyRequest(paymentPolicyRequest PaymentPolicyRequest) ApiUpdatePaymentPolicyRequest {
	r.paymentPolicyRequest = &paymentPolicyRequest
	return r
}

func (r ApiUpdatePaymentPolicyRequest) Execute() (SetPaymentPolicyResponse, *_nethttp.Response, error) {
	return r.ApiService.UpdatePaymentPolicyExecute(r)
}

/*
UpdatePaymentPolicy Method for UpdatePaymentPolicy

This method updates an existing payment policy. Specify the policy you want to update using the <b>payment_policy_id</b> path parameter. Supply a complete policy payload with the updates you want to make; this call overwrites the existing policy with the new details specified in the payload.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param paymentPolicyId This path parameter specifies the ID of the payment policy you want to update.
 @return ApiUpdatePaymentPolicyRequest
*/
func (a *PaymentPolicyApiService) UpdatePaymentPolicy(ctx _context.Context, paymentPolicyId string) ApiUpdatePaymentPolicyRequest {
	return ApiUpdatePaymentPolicyRequest{
		ApiService: a,
		ctx: ctx,
		paymentPolicyId: paymentPolicyId,
	}
}

// Execute executes the request
//  @return SetPaymentPolicyResponse
func (a *PaymentPolicyApiService) UpdatePaymentPolicyExecute(r ApiUpdatePaymentPolicyRequest) (SetPaymentPolicyResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  SetPaymentPolicyResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PaymentPolicyApiService.UpdatePaymentPolicy")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/payment_policy/{payment_policy_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"payment_policy_id"+"}", _neturl.PathEscape(parameterToString(r.paymentPolicyId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.paymentPolicyRequest == nil {
		return localVarReturnValue, nil, reportError("paymentPolicyRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.paymentPolicyRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
