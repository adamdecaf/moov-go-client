/*
 * Moov API
 *
 * The Moov API is an HTTP API served by Moov Financial, Inc for initiating money movements across the ACH payment rail. We follow [RESTful](http://en.wikipedia.org/wiki/Representational_State_Transfer) operations and naming conventions with predictable and standard HTTP status codes. We are available to help with onboarding or issues related to our services on [the Moov slack organization](https://slack.moov.io/) or via [support email](mailto:support@moov.io).  ## Tenants and Organizations  The Moov API offers two groups for organizating `Customer` records. A `Tenant` is the largest grouping which covers an entire business entity such as an LCC or corporation. Login credentials are tied to a Tenant and is extracted from the credentials provided on each request. An `Organization` is a grouping within a Tenant designed to represent a department (sales, marketing) and can be used for the entire LLC. On signup a Tenant is created with an Organization within through the web UI.  Organizations allow for custom underwriting, additional risk tolerances, and advanced access controls for Customer and Account objects. They can be used to keep departments of your company separate or restrict specific underwriting conditions.  <a href=\"https://raw.githubusercontent.com/moov-io/paygate/master/docs/images/tenant-in-paygate.png\" target=\"_blank\">   <img src=\"https://raw.githubusercontent.com/moov-io/paygate/master/docs/images/tenant-in-paygate.png\" /> </a>  ## Errors  The API will respond with various standard HTTP status codes for errors which indicate how to resolve the request's problem. All errors will be in the `application/json` Content-Type with the below structure.  ``` {   \"error\": \"Descriptive message\" } ```  | Status Code | Summary           | Description                                                                                                       | |:-----------:|-------------------|-------------------------------------------------------------------------------------------------------------------| | 200         | OK                | The request was successful.                                                                                       | | 400         | Bad Request       | The request could not be understood by the server. The Incoming parameters might not be valid.                    | | 404         | Not Found         | The requested resource is not found or the credentials are not authorized to access it.                           | | 429         | Too Many Requests | Too many requests have been made in a short period of time. Please make requests at a slower rate or contact us.  | | 500         | Server Error      | The server could not return the representation due to an internal server error.                                   | | 501         | Not Implemented   | The requested operation is not supported (e.g. supports PUT but not POST etc.)                                    |  ## Content-Type  All requests and responses will be in the `application/json` MIME Content-Type unless otherwise specified.  ## Cross-Origin Request Sharing  We support [cross-origin resource sharing](http://en.wikipedia.org/wiki/Cross-origin_resource_sharing), allowing you to interact securely with our API from client-side web applications (never expose your secret API key in any public website's client-side code).  ## Versioning  The Moov API is currently using `/v1/` as the versioning prefix for all endpoints. This results in a base URI of `https://api.moov.io/v1/`.  ## Clients  Currently Moov offers a generated [Go client](https://github.com/moov-io/go-client) for usage with our API. The [OpenAPI specification](https://github.com/moov-io/api/blob/master/openapi.yaml) can be used to generate clients in other languages and we are open to supporting additonal languages. Please [contact us](mailto:support@moov.io) with feedback or suggestions.  ## Authorization  The Moov API offers one authorization method via a configured OIDC provider for your Tenant. This provider can be Google, Github, LDAP, or another vendor. We leverage OIDC becasue it allows immediate credential revocation, two-factor verification with that provider and a faster signup flow for users.  ## Security  Moov continiously monitors and scans our API services for security and privacy issues, but if you find a security related problem please contact us at [`security@moov.io`](mailto:security@moov.io).
 *
 * API version: v1
 * Contact: support@moov.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client

import (
	_context "context"
	"github.com/antihax/optional"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

// OrganizationsApiService OrganizationsApi service
type OrganizationsApiService service

// CreateOrganizationOpts Optional parameters for the method 'CreateOrganization'
type CreateOrganizationOpts struct {
	XRequestID optional.String
}

/*
CreateOrganization Create Organization
Create a new Organization under PayGate
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param xUserID Unique userID set by an auth proxy or client to identify and isolate objects.
 * @param createOrganization
 * @param optional nil or *CreateOrganizationOpts - Optional Parameters:
 * @param "XRequestID" (optional.String) -  Optional requestID allows application developer to trace requests through the systems logs
@return Organization
*/
func (a *OrganizationsApiService) CreateOrganization(ctx _context.Context, xUserID string, createOrganization CreateOrganization, localVarOptionals *CreateOrganizationOpts) (Organization, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Organization
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v1/organizations"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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
	if localVarOptionals != nil && localVarOptionals.XRequestID.IsSet() {
		localVarHeaderParams["X-Request-ID"] = parameterToString(localVarOptionals.XRequestID.Value(), "")
	}
	localVarHeaderParams["X-User-ID"] = parameterToString(xUserID, "")
	// body params
	localVarPostBody = &createOrganization
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
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

// GetOrganizationsOpts Optional parameters for the method 'GetOrganizations'
type GetOrganizationsOpts struct {
	XRequestID optional.String
}

/*
GetOrganizations Get Organizations
Retrieve all Organizations for the given userID
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param xUserID Unique userID set by an auth proxy or client to identify and isolate objects.
 * @param optional nil or *GetOrganizationsOpts - Optional Parameters:
 * @param "XRequestID" (optional.String) -  Optional requestID allows application developer to trace requests through the systems logs
@return []Organization
*/
func (a *OrganizationsApiService) GetOrganizations(ctx _context.Context, xUserID string, localVarOptionals *GetOrganizationsOpts) ([]Organization, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []Organization
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v1/organizations"
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
	if localVarOptionals != nil && localVarOptionals.XRequestID.IsSet() {
		localVarHeaderParams["X-Request-ID"] = parameterToString(localVarOptionals.XRequestID.Value(), "")
	}
	localVarHeaderParams["X-User-ID"] = parameterToString(xUserID, "")
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
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

// UpdateOrganizationOpts Optional parameters for the method 'UpdateOrganization'
type UpdateOrganizationOpts struct {
	XRequestID optional.String
}

/*
UpdateOrganization Update Organization
Update metadata for an Organization
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param organizationID organizationID for the Organization to update
 * @param xUserID Unique userID set by an auth proxy or client to identify and isolate objects.
 * @param createOrganization
 * @param optional nil or *UpdateOrganizationOpts - Optional Parameters:
 * @param "XRequestID" (optional.String) -  Optional requestID allows application developer to trace requests through the systems logs
@return Organization
*/
func (a *OrganizationsApiService) UpdateOrganization(ctx _context.Context, organizationID string, xUserID string, createOrganization CreateOrganization, localVarOptionals *UpdateOrganizationOpts) (Organization, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Organization
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v1/organizations/{organizationID}"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationID"+"}", _neturl.QueryEscape(parameterToString(organizationID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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
	if localVarOptionals != nil && localVarOptionals.XRequestID.IsSet() {
		localVarHeaderParams["X-Request-ID"] = parameterToString(localVarOptionals.XRequestID.Value(), "")
	}
	localVarHeaderParams["X-User-ID"] = parameterToString(xUserID, "")
	// body params
	localVarPostBody = &createOrganization
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
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