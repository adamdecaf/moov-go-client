/*
 * Moov API
 *
 * _Note_: We're currently in pre-release of our API. We expect breaking changes before launching v1 so please join our [mailing list](https://groups.google.com/forum/#!forum/moov-users) for more updates and notices.  The Moov API is organized around [REST](http://en.wikipedia.org/wiki/Representational_State_Transfer). Our API has predictable, resource-oriented URLs, and uses HTTP response codes to indicate API errors. We use built-in HTTP features, like HTTP authentication and HTTP verbs, which are understood by off-the-shelf HTTP clients. We support [cross-origin resource sharing](http://en.wikipedia.org/wiki/Cross-origin_resource_sharing), allowing you to interact securely with our API from client-side web applications (never expose your secret API key in any public website's client-side code). [JSON](http://www.json.org/) is returned by all API responses, including errors, although you can generate client code via [OpenAPI](https://swagger.io/) code generation to convert responses to appropriate language-specific objects.  The Moov API offers two methods of authentication, Cookie and OAuth2 access tokens. The cookie auth is designed for web browsers while the OAuth2 authentication is designed for automated access of our API.  When an API requires a token generated using OAuth (2-legged), no end user is involved. You generate the token by passing your client credentials (Client Id and Client Secret) in a simple call to Create access token (`/oauth2/token`). The operation returns a token that is valid for a few hours and can be renewed; when it expires, you just repeat the call and get a new token. Making additional token requests will keep generating tokens. There are no hard or soft limits.  Cookie auth is setup by provided (`/users/login`) a valid email and password combination. A `Set-Cookie` header is returned on success, which can be used in later calls. Cookie auth is required to generate OAuth2 client credentials.  The Moov API offers many services: - Automated Clearing House (ACH) origination and file management - Transfers and ACH Customer management. - X9 / Image Cash Ledger (ICL) specification support (image uplaod)  ACH is implemented a RESTful API enabling ACH transactions to be submitted and received without a deep understanding of a full NACHA file specification.  An *Originator* can initiate a *Transfer* as either a push (credit) or pull (debit) to a *Customer*. Originators and Customers must have a valid *Depository* account for a Transfer. A *Transfer* is initiated by an Originator to a Customer with an amount and flow of funds. ``` Originator                 ->   Gateway   ->   Customer  - OriginatorDepository                         - CustomerDepository  - Type   (Push or Pull)  - Amount (USD 12.43)  - Status (Pending)  ```  If you find a security related problem please contact us at [`security@moov.io`](mailto:security@moov.io). 
 *
 * API version: v1
 * Contact: security@moov.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type BatchControl struct {
	// Batch ID
	Id string `json:"id,omitempty"`
	// Same as ServiceClassCode in BatchHeaderRecord
	ServiceClassCode int32 `json:"serviceClassCode,omitempty"`
	// EntryAddendaCount is a tally of each Entry Detail Record and each Addenda Record processed, within either the batch or file as appropriate.
	EntryAddendaount int32 `json:"entryAddendaÇount,omitempty"`
	// Validate the Receiving DFI Identification in each Entry Detail Record is hashed to provide a check against inadvertent alteration of data contents due to hardware failure or program error. In this context the Entry Hash is the sum of the corresponding fields in the Entry Detail Records on the file. 
	EntryHash int32 `json:"entryHash,omitempty"`
	// Contains accumulated Entry debit totals within the batch.
	TotalDebit int32 `json:"totalDebit,omitempty"`
	// Contains accumulated Entry credit totals within the batch.
	TotalCredit int32 `json:"totalCredit,omitempty"`
	// Alphanumeric code used to identify an Originator The Company Identification Field must be included on all prenotification records and on each entry initiated pursuant to such prenotification. The Company ID may begin with the ANSI one-digit Identification Code Designator (ICD), followed by the identification number The ANSI Identification Numbers and related Identification Code IRS Employer Identification Number (EIN) \"1\" Data Universal Numbering Systems (DUNS) \"3\" User Assigned Number \"9\" 
	CompanyIdentification string `json:"companyIdentification,omitempty"`
	// MAC is an eight character code derived from a special key used in conjunction with the DES algorithm. The purpose of the MAC is to validate the authenticity of ACH entries. The DES algorithm and key message standards must be in accordance with standards adopted by the American National Standards Institute. The remaining eleven characters of this field are blank.
	MessageAuthentication string `json:"messageAuthentication,omitempty"`
	// The routing number is used to identify the DFI originating entries within a given branch.
	ODFIIdentification string `json:"ODFIIdentification,omitempty"`
	// BatchNumber is assigned in ascending sequence to each batch by the ODFI or its Sending Point in a given file of entries. Since the batch number in the Batch Header Record and the Batch Control Record is the same, the ascending sequence number should be assigned by batch and not by record.
	BatchNumber string `json:"batchNumber,omitempty"`
}
