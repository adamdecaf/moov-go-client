/*
 * Moov API
 *
 * _Note_: We're currently in pre-release of our API. We expect breaking changes before launching v1 so please join our [slack organization](http://moov-io.slack.com/) ([request an invite](https://join.slack.com/t/moov-io/shared_invite/enQtNDE5NzIwNTYxODEwLTRkYTcyZDI5ZTlkZWRjMzlhMWVhMGZlOTZiOTk4MmM3MmRhZDY4OTJiMDVjOTE2MGEyNWYzYzY1MGMyMThiZjg)) or [mailing list](https://groups.google.com/forum/#!forum/moov-users) for more updates and notices.  The Moov API is organized around [REST](http://en.wikipedia.org/wiki/Representational_State_Transfer). Our API has predictable, resource-oriented URLs, and uses HTTP response codes to indicate API errors. We use built-in HTTP features, like HTTP authentication and HTTP verbs, which are understood by off-the-shelf HTTP clients. We support [cross-origin resource sharing](http://en.wikipedia.org/wiki/Cross-origin_resource_sharing), allowing you to interact securely with our API from client-side web applications (never expose your secret API key in any public website's client-side code). [JSON](http://www.json.org/) is returned by all API responses, including errors, although you can generate client code via [OpenAPI](https://swagger.io/) code generation to convert responses to appropriate language-specific objects.  The Moov API offers two methods of authentication, Cookie and OAuth2 access tokens. The cookie auth is designed for web browsers while the OAuth2 authentication is designed for automated access of our API.  When an API requires a token generated using OAuth (2-legged), no end user is involved. You generate the token by passing your client credentials (Client Id and Client Secret) in a simple call to Create access token (`/oauth2/token`). The operation returns a token that is valid for a few hours and can be renewed; when it expires, you just repeat the call and get a new token. Making additional token requests will keep generating tokens. There are no hard or soft limits.  Cookie auth is setup by provided (`/users/login`) a valid email and password combination. A `Set-Cookie` header is returned on success, which can be used in later calls. Cookie auth is required to generate OAuth2 client credentials.  The following order of API operations is suggested to start developing against the Moov API:  1. [Create a Moov API user](#operation/createUser) with a unique email address 1. [Login with user/password credentials](#operation/userLogin) 1. [Create an OAuth2 client](#operation/createOAuth2Client) and [Generate an OAuth access token](#operation/createOAuth2Token) 1. Using the OAuth credentials create:    - [Originator](#operation/addOriginator) and [Originator Depository](#operation/addDepository) (requires micro deposit setup)    - [Customer](#operation/addCustomers) and [Customer Depository](#operation/addDepository) (requires micro deposit setup) 1. [Submit the Transfer](#operation/addTransfer)  After signup clients can [submit ACH files](#operation/addFile) (either in JSON or plaintext) for [validation](#operation/validateFile) and [tabulation](#operation/getFileContents).  The Moov API offers many services: - Automated Clearing House (ACH) origination and file management - Transfers and ACH Customer management. - X9 / Image Cash Ledger (ICL) specification support (image uplaod)  ACH is implemented a RESTful API enabling ACH transactions to be submitted and received without a deep understanding of a full NACHA file specification.  An Originator can initiate a Transfer as either a push (credit) or pull (debit) to a Customer. Originators and Customers must have a valid Depository account for a Transfer. A Transfer is initiated by an Originator to a Customer with an amount and flow of funds. ``` Originator                 ->   Gateway   ->   Customer  - OriginatorDepository                         - CustomerDepository  - Type   (Push or Pull)  - Amount (USD 12.43)  - Status (Pending)  ```  If you find a security related problem please contact us at [`security@moov.io`](mailto:security@moov.io).
 *
 * API version: v1
 * Contact: security@moov.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"context"
	"fmt"
	"github.com/antihax/optional"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Linger please
var (
	_ context.Context
)

type CustomersApiService service

/*
CustomersApiService Create a new Customer object
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param createCustomer
 * @param optional nil or *AddCustomersOpts - Optional Parameters:
 * @param "XIdempotencyKey" (optional.String) -  Idempotent key in the header which expires after 24 hours. These strings should contain enough entropy for to not collide with each other in your requests.
 * @param "XRequestId" (optional.String) -  Optional Request ID allows application developer to trace requests through the systems logs
@return Customer
*/

type AddCustomersOpts struct {
	XIdempotencyKey optional.String
	XRequestId      optional.String
}

func (a *CustomersApiService) AddCustomers(ctx context.Context, createCustomer CreateCustomer, localVarOptionals *AddCustomersOpts) (Customer, *http.Response, error) {
	var (
		localVarHttpMethod   = strings.ToUpper("Post")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Customer
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v1/ach/customers"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

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
	if localVarOptionals != nil && localVarOptionals.XIdempotencyKey.IsSet() {
		localVarHeaderParams["X-Idempotency-Key"] = parameterToString(localVarOptionals.XIdempotencyKey.Value(), "")
	}
	if localVarOptionals != nil && localVarOptionals.XRequestId.IsSet() {
		localVarHeaderParams["X-Request-Id"] = parameterToString(localVarOptionals.XRequestId.Value(), "")
	}
	// body params
	localVarPostBody = &createCustomer
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Cookie"] = key
		}
	}

	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
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

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 201 {
			var v Customer
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
CustomersApiService Permanently deletes a customer and associated depositories and transfers. It cannot be undone. Immediately cancels any active Transfers for the customer.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param customerId Customer ID
 * @param optional nil or *DeleteCustomerOpts - Optional Parameters:
 * @param "Authorization" (optional.String) -  OAuth2 Bearer token
 * @param "XRequestId" (optional.String) -  Optional Request ID allows application developer to trace requests through the systems logs
*/

type DeleteCustomerOpts struct {
	Authorization optional.String
	XRequestId    optional.String
}

func (a *CustomersApiService) DeleteCustomer(ctx context.Context, customerId string, localVarOptionals *DeleteCustomerOpts) (*http.Response, error) {
	var (
		localVarHttpMethod   = strings.ToUpper("Delete")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v1/ach/customers/{customerId}"
	localVarPath = strings.Replace(localVarPath, "{"+"customerId"+"}", fmt.Sprintf("%v", customerId), -1)

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
	localVarHttpHeaderAccepts := []string{}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarOptionals != nil && localVarOptionals.Authorization.IsSet() {
		localVarHeaderParams["Authorization"] = parameterToString(localVarOptionals.Authorization.Value(), "")
	}
	if localVarOptionals != nil && localVarOptionals.XRequestId.IsSet() {
		localVarHeaderParams["X-Request-Id"] = parameterToString(localVarOptionals.XRequestId.Value(), "")
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Cookie"] = key
		}
	}

	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
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
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}

/*
CustomersApiService Get a Customer by ID
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param customerId Customer ID
 * @param optional nil or *GetCustomerByIDOpts - Optional Parameters:
 * @param "Offset" (optional.Int32) -  The number of items to skip before starting to collect the result set
 * @param "Limit" (optional.Int32) -  The number of items to return
 * @param "XRequestId" (optional.String) -  Optional Request ID allows application developer to trace requests through the systems logs
@return Customer
*/

type GetCustomerByIDOpts struct {
	Offset     optional.Int32
	Limit      optional.Int32
	XRequestId optional.String
}

func (a *CustomersApiService) GetCustomerByID(ctx context.Context, customerId string, localVarOptionals *GetCustomerByIDOpts) (Customer, *http.Response, error) {
	var (
		localVarHttpMethod   = strings.ToUpper("Get")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Customer
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v1/ach/customers/{customerId}"
	localVarPath = strings.Replace(localVarPath, "{"+"customerId"+"}", fmt.Sprintf("%v", customerId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Offset.IsSet() {
		localVarQueryParams.Add("offset", parameterToString(localVarOptionals.Offset.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Limit.IsSet() {
		localVarQueryParams.Add("limit", parameterToString(localVarOptionals.Limit.Value(), ""))
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
	if localVarOptionals != nil && localVarOptionals.XRequestId.IsSet() {
		localVarHeaderParams["X-Request-Id"] = parameterToString(localVarOptionals.XRequestId.Value(), "")
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Cookie"] = key
		}
	}

	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
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

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v Customer
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
CustomersApiService Gets a list of Customers
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *GetCustomersOpts - Optional Parameters:
 * @param "Offset" (optional.Int32) -  The number of items to skip before starting to collect the result set
 * @param "Limit" (optional.Int32) -  The number of items to return
 * @param "XRequestId" (optional.String) -  Optional Request ID allows application developer to trace requests through the systems logs
@return []Customer
*/

type GetCustomersOpts struct {
	Offset     optional.Int32
	Limit      optional.Int32
	XRequestId optional.String
}

func (a *CustomersApiService) GetCustomers(ctx context.Context, localVarOptionals *GetCustomersOpts) ([]Customer, *http.Response, error) {
	var (
		localVarHttpMethod   = strings.ToUpper("Get")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []Customer
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v1/ach/customers"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Offset.IsSet() {
		localVarQueryParams.Add("offset", parameterToString(localVarOptionals.Offset.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Limit.IsSet() {
		localVarQueryParams.Add("limit", parameterToString(localVarOptionals.Limit.Value(), ""))
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
	if localVarOptionals != nil && localVarOptionals.XRequestId.IsSet() {
		localVarHeaderParams["X-Request-Id"] = parameterToString(localVarOptionals.XRequestId.Value(), "")
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Cookie"] = key
		}
	}

	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
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

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v []Customer
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
CustomersApiService Get a list of Depository accounts for a Customer
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param customerId Customer ID
 * @param optional nil or *GetDepositoriesByCustomerIDOpts - Optional Parameters:
 * @param "Offset" (optional.Int32) -  The number of items to skip before starting to collect the result set
 * @param "Limit" (optional.Int32) -  The number of items to return
 * @param "XRequestId" (optional.String) -  Optional Request ID allows application developer to trace requests through the systems logs
@return []Depository
*/

type GetDepositoriesByCustomerIDOpts struct {
	Offset     optional.Int32
	Limit      optional.Int32
	XRequestId optional.String
}

func (a *CustomersApiService) GetDepositoriesByCustomerID(ctx context.Context, customerId string, localVarOptionals *GetDepositoriesByCustomerIDOpts) ([]Depository, *http.Response, error) {
	var (
		localVarHttpMethod   = strings.ToUpper("Get")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []Depository
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v1/ach/customers/{customerId}/depositories"
	localVarPath = strings.Replace(localVarPath, "{"+"customerId"+"}", fmt.Sprintf("%v", customerId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Offset.IsSet() {
		localVarQueryParams.Add("offset", parameterToString(localVarOptionals.Offset.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Limit.IsSet() {
		localVarQueryParams.Add("limit", parameterToString(localVarOptionals.Limit.Value(), ""))
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
	if localVarOptionals != nil && localVarOptionals.XRequestId.IsSet() {
		localVarHeaderParams["X-Request-Id"] = parameterToString(localVarOptionals.XRequestId.Value(), "")
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Cookie"] = key
		}
	}

	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
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

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v []Depository
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
CustomersApiService Get a Depository accounts for a Customer based on it's ID
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param customerId Customer ID
 * @param depositoryId Depository ID
 * @param optional nil or *GetDepositoriesByIDOpts - Optional Parameters:
 * @param "Offset" (optional.Int32) -  The number of items to skip before starting to collect the result set
 * @param "Limit" (optional.Int32) -  The number of items to return
 * @param "XRequestId" (optional.String) -  Optional Request ID allows application developer to trace requests through the systems logs
@return Depository
*/

type GetDepositoriesByIDOpts struct {
	Offset     optional.Int32
	Limit      optional.Int32
	XRequestId optional.String
}

func (a *CustomersApiService) GetDepositoriesByID(ctx context.Context, customerId string, depositoryId string, localVarOptionals *GetDepositoriesByIDOpts) (Depository, *http.Response, error) {
	var (
		localVarHttpMethod   = strings.ToUpper("Get")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Depository
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v1/ach/customers/{customerId}/depositories/{depositoryId}"
	localVarPath = strings.Replace(localVarPath, "{"+"customerId"+"}", fmt.Sprintf("%v", customerId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"depositoryId"+"}", fmt.Sprintf("%v", depositoryId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Offset.IsSet() {
		localVarQueryParams.Add("offset", parameterToString(localVarOptionals.Offset.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Limit.IsSet() {
		localVarQueryParams.Add("limit", parameterToString(localVarOptionals.Limit.Value(), ""))
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
	if localVarOptionals != nil && localVarOptionals.XRequestId.IsSet() {
		localVarHeaderParams["X-Request-Id"] = parameterToString(localVarOptionals.XRequestId.Value(), "")
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Cookie"] = key
		}
	}

	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
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

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v Depository
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
CustomersApiService Updates the specified Customer by setting the values of the parameters passed. Any parameters not provided will be left unchanged.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param customerId Customer ID
 * @param createCustomer
 * @param optional nil or *UpdateCustomerOpts - Optional Parameters:
 * @param "XIdempotencyKey" (optional.String) -  Idempotent key in the header which expires after 24 hours. These strings should contain enough entropy for to not collide with each other in your requests.
 * @param "XRequestId" (optional.String) -  Optional Request ID allows application developer to trace requests through the systems logs
@return Customer
*/

type UpdateCustomerOpts struct {
	XIdempotencyKey optional.String
	XRequestId      optional.String
}

func (a *CustomersApiService) UpdateCustomer(ctx context.Context, customerId string, createCustomer CreateCustomer, localVarOptionals *UpdateCustomerOpts) (Customer, *http.Response, error) {
	var (
		localVarHttpMethod   = strings.ToUpper("Patch")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Customer
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v1/ach/customers/{customerId}"
	localVarPath = strings.Replace(localVarPath, "{"+"customerId"+"}", fmt.Sprintf("%v", customerId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

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
	if localVarOptionals != nil && localVarOptionals.XIdempotencyKey.IsSet() {
		localVarHeaderParams["X-Idempotency-Key"] = parameterToString(localVarOptionals.XIdempotencyKey.Value(), "")
	}
	if localVarOptionals != nil && localVarOptionals.XRequestId.IsSet() {
		localVarHeaderParams["X-Request-Id"] = parameterToString(localVarOptionals.XRequestId.Value(), "")
	}
	// body params
	localVarPostBody = &createCustomer
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Cookie"] = key
		}
	}

	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
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

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 201 {
			var v Customer
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
