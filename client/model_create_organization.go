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

// CreateOrganization struct for CreateOrganization
type CreateOrganization struct {
	// Legal name for this Organization
	Name string `json:"name"`
	// tenantID to create this Organization under
	TenantID string `json:"tenantID"`
	// A customerID from the Customers service to use in Transfers with this Organization. When transferring to or from the Organization this Customer and Account(s) are used. The Customer assigned here should represent the legal entity that manages the Organization.
	PrimaryCustomer string `json:"primaryCustomer"`
}
