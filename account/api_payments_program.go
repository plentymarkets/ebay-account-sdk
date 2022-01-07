/*
Account API

The <b>Account API</b> gives sellers the ability to configure their eBay seller accounts, including the seller's policies (seller-defined custom policies and eBay business policies), opt in and out of eBay seller programs, configure sales tax tables, and get account information.  <br><br>For details on the availability of the methods in this API, see <a href=\"/api-docs/sell/account/overview.html#requirements\">Account API requirements and restrictions</a>.

API version: v1.7.0
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

// PaymentsProgramApiService PaymentsProgramApi service
type PaymentsProgramApiService service

type ApiGetPaymentsProgramRequest struct {
	ctx _context.Context
	ApiService *PaymentsProgramApiService
	marketplaceId string
	paymentsProgramType string
}


func (r ApiGetPaymentsProgramRequest) Execute() (PaymentsProgramResponse, *_nethttp.Response, error) {
	return r.ApiService.GetPaymentsProgramExecute(r)
}

/*
GetPaymentsProgram Method for GetPaymentsProgram

This method returns whether or not the user is opted-in to the specified payments program. Sellers opt-in to payments programs by marketplace and you use the <b>marketplace_id</b> path parameter to specify the marketplace of the status flag you want returned.  <br><br><span class="tablenote"><b>Note:</b> Currently, the only supported payments program is <b>eBay Payments</b>. For details, see: <ul><li><a href="https://pages.ebay.com/seller-center/service-and-payments/managed-payments-on-ebay.html" target="_blank">Managed Payments on eBay</a></li> <li><a href="https://pages.ebay.com/payment/2.0/terms.html" target="_blank">Payments Terms of Use</a></li></ul></span>

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param marketplaceId This path parameter specifies the eBay marketplace of the payments program for which you want to retrieve the seller's status.
 @param paymentsProgramType This path parameter specifies the payments program whose status is returned by the call.  <br><br>Currently the only supported payments program is <code>EBAY_PAYMENTS</code>. For details on the program, see <a href=\"https://pages.ebay.com/payment/2.0/terms.html\" target=\"_blank\">Payments Terms of Use</a>.
 @return ApiGetPaymentsProgramRequest
*/
func (a *PaymentsProgramApiService) GetPaymentsProgram(ctx _context.Context, marketplaceId string, paymentsProgramType string) ApiGetPaymentsProgramRequest {
	return ApiGetPaymentsProgramRequest{
		ApiService: a,
		ctx: ctx,
		marketplaceId: marketplaceId,
		paymentsProgramType: paymentsProgramType,
	}
}

// Execute executes the request
//  @return PaymentsProgramResponse
func (a *PaymentsProgramApiService) GetPaymentsProgramExecute(r ApiGetPaymentsProgramRequest) (PaymentsProgramResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  PaymentsProgramResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PaymentsProgramApiService.GetPaymentsProgram")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/payments_program/{marketplace_id}/{payments_program_type}"
	localVarPath = strings.Replace(localVarPath, "{"+"marketplace_id"+"}", _neturl.PathEscape(parameterToString(r.marketplaceId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"payments_program_type"+"}", _neturl.PathEscape(parameterToString(r.paymentsProgramType, "")), -1)

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
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
