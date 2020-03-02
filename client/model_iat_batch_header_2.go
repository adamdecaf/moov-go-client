/*
 * Moov API
 *
 * _Note_: The Moov API and services are under development and could introduce breaking changes while reaching a stable status. We are looking for community feedback so please try out our code, [join the slack organization](https://slack.moov.io/) and give us some feedback! We announce releases on the [mailing list](https://groups.google.com/forum/#!forum/moov-users).  The Moov API is organized around [REST](http://en.wikipedia.org/wiki/Representational_State_Transfer). Our API has predictable, resource-oriented URLs, and uses HTTP response codes to indicate API errors. We use built-in HTTP features, like HTTP authentication and HTTP verbs, which are understood by off-the-shelf HTTP clients. We support [cross-origin resource sharing](http://en.wikipedia.org/wiki/Cross-origin_resource_sharing), allowing you to interact securely with our API from client-side web applications (never expose your secret API key in any public website's client-side code). [JSON](http://www.json.org/) is returned by all API responses, including errors, although you can generate client code via [OpenAPI code generation](https://github.com/OpenAPITools/openapi-generator) or the [OpenAPI editor](https://editor.swagger.io/) to convert responses to appropriate language-specific objects.  The Moov API offers two methods of authentication, Cookie and OAuth2 access tokens. The cookie auth is designed for web browsers while the OAuth2 authentication is designed for automated access of our API.  When an API requires a token generated using OAuth (2-legged), no end user is involved. You generate the token by passing your client credentials (Client ID and Client Secret) in a simple call to Create access token (`/oauth2/token`). The operation returns a token that is valid for a few hours and can be renewed; when it expires, you just repeat the call and get a new token. Making additional token requests will keep generating tokens. There are no hard or soft limits.  Cookie auth is setup by provided (`/users/login`) a valid email and password combination. A `Set-Cookie` header is returned on success, which can be used in later calls. Cookie auth is required to generate OAuth2 client credentials.  The following order of API operations is suggested to start developing against the Moov API:  1. [Create a Moov API user](#operation/createUser) with a unique email address 1. [Login with user/password credentials](#operation/userLogin) 1. [Create an OAuth2 client](#operation/createOAuth2Client) and [Generate an OAuth access token](#operation/createOAuth2Token) 1. Using the OAuth credentials create:    - [Originator](#operation/addOriginator) and [Originator Depository](#operation/addDepository) (requires micro deposit setup)    - [Receiver](#operation/addReceivers) and [Receiver Depository](#operation/addDepository) (requires micro deposit setup) 1. [Submit the Transfer](#operation/addTransfer)  After signup clients can [submit ACH files](#operation/addFile) (either in JSON or plaintext) for [validation](#operation/validateFile) and [tabulation](#operation/getFileContents).  The Moov API offers many services: - Automated Clearing House (ACH) origination and file management - Transfers and ACH Receiver management - Image Cash Ledger (ICL) file creation and modification API - Fed WIRE file creation and modification API  ACH is implemented a RESTful API enabling ACH transactions to be submitted and received without a deep understanding of a full NACHA file specification.  An `Originator` can initiate a `Transfer` as either a push (credit) or pull (debit) to a `Receiver`. Originators and Receivers must have a valid `Depository` account for a `Transfer`. A `Transfer` is initiated by an `Originator` to a `Receiver` with an amount and flow of funds.  If you find a security related problem please contact us at [`security@moov.io`](mailto:security@moov.io).
 *
 * API version: v1
 * Contact: security@moov.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client

// IatBatchHeader2 struct for IatBatchHeader2
type IatBatchHeader2 struct {
	// ID is a client defined string used as a reference to this record.
	ID string `json:"ID,omitempty"`
	// ServiceClassCode ACH Mixed Debits and Credits '200' ACH Credits Only '220' ACH Debits Only '225'
	ServiceClassCode int32 `json:"serviceClassCode,omitempty"`
	// Leave Blank. Only used for corrected IAT entries
	IATIndicator string `json:"IATIndicator,omitempty"`
	// Code indicating currency conversion. FV Fixed-to-Variable – Entry is originated in a fixed-value amount and is to be received in a variable amount resulting from the execution of the foreign exchange conversion. VF Variable-to-Fixed – Entry is originated in a variable-value amount based on a specific foreign exchange rate for conversion to a fixed-value amount in which the entry is to be received. FF Fixed-to-Fixed – Entry is originated in a fixed-value amount and is to be received in the same fixed-value amount in the same currency denomination. There is no foreign exchange conversion for entries transmitted using this code. For entries originated in a fixed value amount, the foreign Exchange Reference Field will be space filled.
	ForeignExchangeIndicator string `json:"foreignExchangeIndicator,omitempty"`
	// Code used to indicate the content of the Foreign Exchange Reference Field and is filled by the gateway operator. Valid entries are 1 - Foreign Exchange Rate; 2 - Foreign Exchange Reference Number; or 3 - Space Filled
	ForeignExchangeReferenceIndicator int32 `json:"foreignExchangeReferenceIndicator,omitempty"`
	// Contains either the foreign exchange rate used to execute the foreign exchange conversion of a cross-border entry or another reference to the foreign exchange transaction.
	ForeignExchangeReference string `json:"foreignExchangeReference,omitempty"`
	// Two-character code, as approved by the International Organization for Standardization (ISO), to identify the country in which the entry is to be received. For United States use US.
	ISODestinationCountryCode string `json:"ISODestinationCountryCode,omitempty"`
	// For U.S. entities: the number assigned will be your tax ID (often Social Security Number) For non-U.S. entities: the number assigned will be your DDA number, or the last 9 characters of your account number if it exceeds 9 characters
	OriginatorIdentification string `json:"originatorIdentification,omitempty"`
	// StandardEntryClassCode for consumer and non consumer international payments is IAT. Identifies the payment type (product) found within an ACH batch-using a 3-character code. The SEC Code pertains to all items within batch. Determines format of the detail records. Determines addenda records (required or optional PLUS one or up to 9,999 records). Determines rules to follow (return time frames). Some SEC codes require specific data in predetermined fields within the ACH record
	StandardEntryClassCode string `json:"standardEntryClassCode,omitempty"`
	// A description of the entries contained in the batch The Originator establishes the value of this field to provide a description of the purpose of the entry to be displayed back to the receive For example, \"GAS BILL,\" \"REG. SALARY,\" \"INS. PREM,\" \"SOC. SEC.,\" \"DTC,\" \"TRADE PAY,\" \"PURCHASE,\" etc. This field must contain the word \"REVERSAL\" (left justified) when the batch contains reversing entries. This field must contain the word \"RECLAIM\" (left justified) when the batch contains reclamation entries. This field must contain the word \"NONSETTLED\" (left justified) when the batch contains entries which could not settle.
	CompanyEntryDescription string `json:"companyEntryDescription,omitempty"`
	// Three-character code, as approved by the International Organization for Standardization (ISO), to identify the currency denomination in which the entry was first originated. If the source of funds is within the territorial jurisdiction of the U.S., enter 'USD', otherwise refer to International Organization for Standardization website for value: www.iso.org
	ISOOriginatingCurrencyCode string `json:"ISOOriginatingCurrencyCode,omitempty"`
	// ISODestinationCurrencyCode is the three-character code, as approved by the International Organization for Standardization (ISO), to identify the currency denomination in which the entry will ultimately be settled. If the final destination of funds is within the territorial jurisdiction of the U.S., enter \"USD\", otherwise refer to International Organization for Standardization website for value: www.iso.org
	ISODestinationCurrencyCode string `json:"ISODestinationCurrencyCode,omitempty"`
	// EffectiveEntryDate the date on which the entries are to settle format YYMMDD (Y=Year, M=Month, D=Day)
	EffectiveEntryDate string `json:"effectiveEntryDate,omitempty"`
	// SettlementDate Leave blank, this field is inserted by the ACH operator settlementDate string OriginatorStatusCode refers to the ODFI initiating the Entry. 0 ADV File prepared by an ACH Operator. 1 This code identifies the Originator as a depository financial institution. 2 This code identifies the Originator as a Federal Government entity or agency.
	OriginatorStatusCode int32 `json:"originatorStatusCode,omitempty"`
	// ODFIIdentification First 8 digits of the originating DFI transit routing number for Inbound IAT Entries, this field contains the routing number of the U.S. Gateway Operator.  For Outbound IAT Entries, this field contains the standard routing number, as assigned by Accuity, that identifies the U.S. ODFI initiating the Entry. Format - TTTTAAAA
	ODFIIdentification string `json:"ODFIIdentification,omitempty"`
	// BatchNumber is assigned in ascending sequence to each batch by the ODFI or its Sending Point in a given file of entries. Since the batch number in the Batch Header Record and the Batch Control Record is the same, the ascending sequence number should be assigned by batch and not by record.
	BatchNumber int32 `json:"batchNumber,omitempty"`
}
