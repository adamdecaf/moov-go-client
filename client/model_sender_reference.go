/*
 * Moov API
 *
 * _Note_: The Moov API and services are under development and could introduce breaking changes while reaching a stable stus. We are looking for community feedback so please try out our code, [join the slack organization](https://slack.moov.io/) and give us some feedback! We announce releases on the [mailing list](https://groups.google.com/forum/#!forum/moov-users).  The Moov API is organized around [REST](http://en.wikipedia.org/wiki/Representational_State_Transfer). Our API has predictable, resource-oriented URLs, and uses HTTP response codes to indicate API errors. We use built-in HTTP features, like HTTP authentication and HTTP verbs, which are understood by off-the-shelf HTTP clients. We support [cross-origin resource sharing](http://en.wikipedia.org/wiki/Cross-origin_resource_sharing), allowing you to interact securely with our API from client-side web applications (never expose your secret API key in any public website's client-side code). [JSON](http://www.json.org/) is returned by all API responses, including errors, although you can generate client code via [OpenAPI code generation](https://github.com/OpenAPITools/openapi-generator) or the [OpenAPI editor](https://editor.swagger.io/) to convert responses to appropriate language-specific objects.  The Moov API offers two methods of authentication, Cookie and OAuth2 access tokens. The cookie auth is designed for web browsers while the OAuth2 authentication is designed for automated access of our API.  When an API requires a token generated using OAuth (2-legged), no end user is involved. You generate the token by passing your client credentials (Client ID and Client Secret) in a simple call to Create access token (`/oauth2/token`). The operation returns a token that is valid for a few hours and can be renewed; when it expires, you just repeat the call and get a new token. Making additional token requests will keep generating tokens. There are no hard or soft limits.  Cookie auth is setup by provided (`/users/login`) a valid email and password combination. A `Set-Cookie` header is returned on success, which can be used in later calls. Cookie auth is required to generate OAuth2 client credentials.  The following order of API operations is suggested to start developing against the Moov API:  1. [Create a Moov API user](#operation/createUser) with a unique email address 1. [Login with user/password credentials](#operation/userLogin) 1. [Create an OAuth2 client](#operation/createOAuth2Client) and [Generate an OAuth access token](#operation/createOAuth2Token) 1. Using the OAuth credentials create:    - [Originator](#operation/addOriginator) and [Originator Depository](#operation/addDepository) (requires micro deposit setup)    - [Receiver](#operation/addReceivers) and [Receiver Depository](#operation/addDepository) (requires micro deposit setup) 1. [Submit the Transfer](#operation/addTransfer)  After signup clients can [submit ACH files](#operation/addFile) (either in JSON or plaintext) for [validation](#operation/validateFile) and [tabulation](#operation/getFileContents).  The Moov API offers many services: - Automated Clearing House (ACH) origination and file management - Transfers and ACH Receiver management. - X9 / Image Cash Ledger (ICL) specification support (image uplaod)  ACH is implemented a RESTful API enabling ACH transactions to be submitted and received without a deep understanding of a full NACHA file specification.  An Originator can initiate a Transfer as either a push (credit) or pull (debit) to a Receiver. Originators and Receivers must have a valid Depository account for a Transfer. A Transfer is initiated by an Originator to a Receiver with an amount and flow of funds. ``` Originator                 ->   Gateway   ->   Receiver  - OriginatorDepository                         - ReceiverDepository  - Type   (Push or Pull)  - Amount (USD 12.43)  - Status (Pending)  ```  If you find a security related problem please contact us at [`security@moov.io`](mailto:security@moov.io).
 *
 * API version: v1
 * Contact: security@moov.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// SenderReference SenderReference is sender reference
type SenderReference struct {
	// SenderReference
	SenderReference string `json:"senderReference,omitempty"`
}
