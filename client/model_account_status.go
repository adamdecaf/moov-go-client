/*
 * Moov API
 *
 * _Note_: The Moov API and services are under development and could introduce breaking changes while reaching a stable status. We are looking for community feedback so please try out our code, [join the slack organization](https://slack.moov.io/) and give us some feedback! We announce releases on the [mailing list](https://groups.google.com/forum/#!forum/moov-users).  The Moov API is organized around [REST](http://en.wikipedia.org/wiki/Representational_State_Transfer). Our API has predictable, resource-oriented URLs, and uses HTTP response codes to indicate API errors. We use built-in HTTP features, like HTTP authentication and HTTP verbs, which are understood by off-the-shelf HTTP clients. We support [cross-origin resource sharing](http://en.wikipedia.org/wiki/Cross-origin_resource_sharing), allowing you to interact securely with our API from client-side web applications (never expose your secret API key in any public website's client-side code). [JSON](http://www.json.org/) is returned by all API responses, including errors, although you can generate client code via [OpenAPI code generation](https://github.com/OpenAPITools/openapi-generator) or the [OpenAPI editor](https://editor.swagger.io/) to convert responses to appropriate language-specific objects.  The Moov API offers one method of authentication -- OAuth2 access tokens. The OAuth2 authentication is designed for automated access of our API. When an API requires a token generated using OAuth (2-legged), no end user is involved. You generate the token by passing your client credentials (Client ID and Client Secret) in a simple call to Create access token (`/oauth2/token`). The operation returns a token that is valid for a few hours and can be renewed; when it expires, you just repeat the call and get a new token. Making additional token requests will keep generating tokens. There are no hard or soft limits.  The Moov API offers many services:   - Automated Clearing House (ACH) origination and file management   - Transfers management   <!-- - Image Cash Ledger (ICL) file creation and modification API -->   <!-- - Fed WIRE file creation and modification API -->  The following order of API operations is suggested to start developing against the Moov API:    1. [Create a Moov API user](#operation/...) with a unique email address   1. Create a source Customer      1. Approve      1. Create an Account         1. Verify or Approve   1. Create a destination Customer      1. Approve      1. Create an Account         1. Verify or Approve   1. Initiate a Transfer  ACH is implemented a RESTful API enabling ACH transactions to be submitted and received without a deep understanding of a full NACHA file specification. A `Customer` can initiate a `Transfer` as either a push (credit) or pull (debit) to another `Customer`. Customers must have a valid `Account` account for a `Transfer`.  If you find a security related problem please contact us at [`security@moov.io`](mailto:security@moov.io).
 *
 * API version: v1
 * Contact: security@moov.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client

// AccountStatus Enumeration if this account has been validated through micro-deposits or another available strategy.
type AccountStatus string

// List of AccountStatus
const (
	NONE      AccountStatus = "none"
	VALIDATED AccountStatus = "validated"
)
