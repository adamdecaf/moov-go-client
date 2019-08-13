/*
 * Moov API
 *
 * _Note_: We're currently in pre-release of our API. We expect breaking changes before launching v1 so please join our [slack organization](http://moov-io.slack.com/) ([request an invite](https://join.slack.com/t/moov-io/shared_invite/enQtNDE5NzIwNTYxODEwLTRkYTcyZDI5ZTlkZWRjMzlhMWVhMGZlOTZiOTk4MmM3MmRhZDY4OTJiMDVjOTE2MGEyNWYzYzY1MGMyMThiZjg)) or [mailing list](https://groups.google.com/forum/#!forum/moov-users) for more updates and notices.  The Moov API is organized around [REST](http://en.wikipedia.org/wiki/Representational_State_Transfer). Our API has predictable, resource-oriented URLs, and uses HTTP response codes to indicate API errors. We use built-in HTTP features, like HTTP authentication and HTTP verbs, which are understood by off-the-shelf HTTP clients. We support [cross-origin resource sharing](http://en.wikipedia.org/wiki/Cross-origin_resource_sharing), allowing you to interact securely with our API from client-side web applications (never expose your secret API key in any public website's client-side code). [JSON](http://www.json.org/) is returned by all API responses, including errors, although you can generate client code via [OpenAPI code generation](https://github.com/OpenAPITools/openapi-generator) or the [OpenAPI editor](https://editor.swagger.io/) to convert responses to appropriate language-specific objects.  The Moov API offers two methods of authentication, Cookie and OAuth2 access tokens. The cookie auth is designed for web browsers while the OAuth2 authentication is designed for automated access of our API.  When an API requires a token generated using OAuth (2-legged), no end user is involved. You generate the token by passing your client credentials (Client ID and Client Secret) in a simple call to Create access token (`/oauth2/token`). The operation returns a token that is valid for a few hours and can be renewed; when it expires, you just repeat the call and get a new token. Making additional token requests will keep generating tokens. There are no hard or soft limits.  Cookie auth is setup by provided (`/users/login`) a valid email and password combination. A `Set-Cookie` header is returned on success, which can be used in later calls. Cookie auth is required to generate OAuth2 client credentials.  The following order of API operations is suggested to start developing against the Moov API:  1. [Create a Moov API user](#operation/createUser) with a unique email address 1. [Login with user/password credentials](#operation/userLogin) 1. [Create an OAuth2 client](#operation/createOAuth2Client) and [Generate an OAuth access token](#operation/createOAuth2Token) 1. Using the OAuth credentials create:    - [Originator](#operation/addOriginator) and [Originator Depository](#operation/addDepository) (requires micro deposit setup)    - [Receiver](#operation/addReceivers) and [Receiver Depository](#operation/addDepository) (requires micro deposit setup) 1. [Submit the Transfer](#operation/addTransfer)  After signup clients can [submit ACH files](#operation/addFile) (either in JSON or plaintext) for [validation](#operation/validateFile) and [tabulation](#operation/getFileContents).  The Moov API offers many services: - Automated Clearing House (ACH) origination and file management - Transfers and ACH Receiver management. - X9 / Image Cash Ledger (ICL) specification support (image uplaod)  ACH is implemented a RESTful API enabling ACH transactions to be submitted and received without a deep understanding of a full NACHA file specification.  An Originator can initiate a Transfer as either a push (credit) or pull (debit) to a Receiver. Originators and Receivers must have a valid Depository account for a Transfer. A Transfer is initiated by an Originator to a Receiver with an amount and flow of funds. ``` Originator                 ->   Gateway   ->   Receiver  - OriginatorDepository                         - ReceiverDepository  - Type   (Push or Pull)  - Amount (USD 12.43)  - Status (Pending)  ```  If you find a security related problem please contact us at [`security@moov.io`](mailto:security@moov.io).
 *
 * API version: v1
 * Contact: security@moov.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type BatchHeader struct {
	// Batch Header ID
	ID string `json:"ID,omitempty"`
	// Service Class Code - ACH Credits Only 220 and ACH Debits Only 225
	ServiceClassCode int32 `json:"serviceClassCode"`
	// Company originating the entries in the batch
	CompanyName string `json:"companyName"`
	// The 9 digit FEIN number (proceeded by a predetermined alpha or numeric character) of the entity in the company name field
	CompanyDiscretionaryData string `json:"companyDiscretionaryData,omitempty"`
	// Identifies the payment type (product) found within an ACH batch-using a 3-character code.
	StandardEntryClassCode string `json:"standardEntryClassCode,omitempty"`
	// A description of the entries contained in the batch. The Originator establishes the value of this field to provide a description of the purpose of the entry to be displayed back to the receive For example, \"GAS BILL,\" \"REG. SALARY,\" \"INS. PREM,\", \"SOC. SEC.,\" \"DTC,\" \"TRADE PAY,\" \"PURCHASE,\" etc. This field must contain the word \"REVERSAL\" (left justified) when the batch contains reversing entries. This field must contain the word \"RECLAIM\" (left justified) when the batch contains reclamation entries. This field must contain the word \"NONSETTLED\" (left justified) when the batch contains entries which could not settle.
	CompanyEntryDescription string `json:"companyEntryDescription,omitempty"`
	// The Originator establishes this field as the date it would like to see displayed to the receiver for descriptive purposes. This field is never used to control timing of any computer or manual operation. It is solely for descriptive purposes. The RDFI should not assume any specific format.
	CompanyDescriptiveDate string `json:"companyDescriptiveDate,omitempty"`
	// Date on which the entries are to settle. Format YYMMDD (Y=Year, M=Month, D=Day)
	EffectiveEntryDate string `json:"effectiveEntryDate,omitempty"`
	// ODFI initiating the Entry. 0 ADV File prepared by an ACH Operator. 1 This code identifies the Originator as a depository financial institution. 2 This code identifies the Originator as a Federal Government entity or agency.
	OriginatorStatusCode int32 `json:"originatorStatusCode,omitempty"`
	// First 8 digits of the originating DFI transit routing number
	ODFIIdentification string `json:"ODFIIdentification"`
	// BatchNumber is assigned in ascending sequence to each batch by the ODFI or its Sending Point in a given file of entries. Since the batch number in the Batch Header Record and the Batch Control Record is the same, the ascending sequence number should be assigned by batch and not by record.
	BatchNumber string `json:"batchNumber,omitempty"`
}
