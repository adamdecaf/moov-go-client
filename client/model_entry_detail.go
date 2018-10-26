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

type EntryDetail struct {
	// Entry Detail ID
	Id string `json:"id"`
	// transactionCode if the receivers account is: Credit (deposit) to checking account 22 Prenote for credit to checking account 23 Debit (withdrawal) to checking account 27 Prenote for debit to checking account 28 Credit to savings account 32 Prenote for credit to savings account 33 Debit to savings account 37 Prenote for debit to savings account 38 
	TransactionCode int32 `json:"transactionCode"`
	// RDFI's routing number without the last digit.
	RDFIIdentification string `json:"RDFIIdentification"`
	// Last digit in RDFI routing number.
	CheckDigit string `json:"checkDigit"`
	// The receiver's bank account number you are crediting/debiting. It important to note that this is an alphanumeric field, so its space padded, no zero padded 
	DFIAccountNumber string `json:"DFIAccountNumber"`
	// Number of cents you are debiting/crediting this account
	Amount int32 `json:"amount"`
	// Internal identification (alphanumeric) that you use to uniquely identify this Entry Detail Record
	IdentificationNumber string `json:"identificationNumber,omitempty"`
	// The name of the receiver, usually the name on the bank account
	IndividualName string `json:"individualName"`
	// DiscretionaryData allows ODFIs to include codes, of significance only to them, to enable specialized handling of the entry. There will be no standardized interpretation for the value of this field. It can either be a single two-character code, or two distinct one-character codes, according to the needs of the ODFI and/or Originator involved. This field must be returned intact for any returned entry. WEB uses the Discretionary Data Field as the Payment Type Code 
	DiscretionaryData string `json:"discretionaryData,omitempty"`
	// AddendaRecordIndicator indicates the existence of an Addenda Record. A value of \"1\" indicates that one ore more addenda records follow, and \"0\" means no such record is present. 
	AddendaRecordIndicator int32 `json:"addendaRecordIndicator,omitempty"`
	// TraceNumber assigned by the ODFI in ascending sequence, is included in each Entry Detail Record, Corporate Entry Detail Record, and addenda Record. Trace Numbers uniquely identify each entry within a batch in an ACH input file. In association with the Batch Number, transmission (File Creation) Date, and File ID Modifier, the Trace Number uniquely identifies an entry within a given file. For addenda Records, the Trace Number will be identical to the Trace Number in the associated Entry Detail Record, since the Trace Number is associated with an entry or item rather than a physical record. 
	TraceNumber int32 `json:"traceNumber,omitempty"`
	// List of Addenda for the Entry Detail
	Addendum []Addendum `json:"addendum,omitempty"`
	// Category defines if the entry is a Forward, Return, or NOC
	Category string `json:"category,omitempty"`
}
