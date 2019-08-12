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

import (
	"time"
)

type ImageViewData struct {
	// ImageViewData ID
	Id string `json:"id,omitempty"`
	// ECEInstitutionRoutingNumber contains the routing and transit number of the institution that that creates the bundle header.  This number is imported from the Bundle Header Record (Clause 9.4) associated with the image view conveyed in this Image View Data Property.
	ECEInstitutionRoutingNumber string `json:"eCEInstitutionRoutingNumber,omitempty"`
	// BundleBusinessDate is the business date of the bundle.
	BundleBusinessDate time.Time `json:"bundleBusinessDate,omitempty"`
	// CycleNumber is a code assigned by the institution that creates the bundle.  Denotes the cycle under which the bundle is created.
	CycleNumber string `json:"cycleNumber,omitempty"`
	// ECEInstitutionItemSequenceNumber is a number assigned by the institution that creates the CheckDetail or Return.  This number is imported from the CheckDetail.ECEInstitutionItemSequenceNumber or Return.ECEInstitutionItemSequenceNumber associated with the image view conveyed in this Image View Data Record. The ECE institution must construct the sequence number to guarantee uniqueness for a given routing number, business day, and cycle number. Must contain a numeric value.
	IvDataECEInstitutionItemSequenceNumber string `json:"ivData.ECEInstitutionItemSequenceNumber,omitempty"`
	// SecurityOriginatorName is a unique name that creates the Digital Signature for data to be exchanged. Shall be present only under clearing arrangements and when ImageViewDetail.DigitalSignatureIndicator is 1 Shall not be present when ImageViewDetail.ImageIndicator is 0.
	IvDataSecurityOriginatorName string `json:"ivData.SecurityOriginatorName,omitempty"`
	// SecurityAuthenticatorName is the unique name that performs authentication on received data. Shall be present only under clearing arrangements and when ImageViewDetail.DigitalSignatureIndicator is 1 Shall not be present when ImageViewDetail.ImageIndicator is 0.
	IvDataSecurityAuthenticatorName string `json:"ivData.SecurityAuthenticatorName,omitempty"`
	// SecurityKeyName is a name or character sequence used by the signer (originator) to communicate a key identifierto the recipient (authenticator) so the recipient can obtain the key needed to validate the signature. The name is typically used as an identifier related to the key pair used to sign the image. The name is mutually known to the security originator and the security authenticator and is unique to this relationship. Shall be present only under clearing arrangements and when ImageViewDetail.DigitalSignatureIndicator is 1 Shall not be present when ImageViewDetail.ImageIndicator is 0.
	IvDataSecurityKeyName string `json:"ivData.SecurityKeyName,omitempty"`
	// ClippingOrigin is a code that defines the corner of the conveyed image view that is taken as the reference point for the clipping coordinates. Top, bottom, left, and right references apply to a view that presents a visually correct orientation. When clipping information is present, the nature of the Area of Interest defined by the clipping rectangle is determined by the value of the ImageViewDetail.ViewDescriptor. Primary front and rear views shall only have a Defined Value of 0.  Can be blank.  * `0` - Clipping information is not present–full view present * `1` - Clipping origin is top left corner of image view * `2` - Clipping origin is top right corner of image view * `3` - Clipping origin is bottom right corner of image view * `4` - Clipping origin is bottom left corner of image view
	IvDataClippingOrigin int32 `json:"ivData.ClippingOrigin,omitempty"`
	// ClippingCoordinateH1 is a number that represents the horizontal offset in pixels from the clipping origin to the nearest vertical side of the clipping rectangle. The clipping coordinates (h1, h2, v1, v2) convey the clipping rectangle’s offsets in both horizontal (h) and vertical (v) directions. The offset values collectively establish the boundary sides of the clipping rectangle. Pixels on the boundary of the clipping rectangle are included in the selected array of pixels. That is, the first pixel of the selected array is at offset (h1, v1) and the last pixel of the selected array is at offset (h2, v2). The corner pixel at the origin of the image view is assumed to have the offset value (0, 0). Shall be present if Image View Data.ClippingOrigin is present and non-zero. Shall not be present when ImageViewDetail.ImageIndicator is 0. Valid values - 0000–9999
	IvDataClippingCoordinateH1 string `json:"ivData.ClippingCoordinateH1,omitempty"`
	// ClippingCoordinateH2 is a number that represents the horizontal offset in pixels from the clipping origin to the furthermost vertical side of the clipping rectangle. Shall be present if Image View Data.ClippingOrigin is present and non-zero. Shall not be present when ImageViewDetail.ImageIndicator is 0. Valid values - 0000–9999
	IvDataClippingCoordinateH2 string `json:"ivData.ClippingCoordinateH2,omitempty"`
	// ClippingCoordinateV1 is a number that represents the vertical offset in pixels from the clipping origin to the nearest horizontal side of the clipping rectangle. Shall be present if Image View Data.ClippingOrigin is present and non-zero. Shall not be present when ImageViewDetail.ImageIndicator is 0. Valid values - 0000–9999
	IvDataClippingCoordinateV1 string `json:"ivData.ClippingCoordinateV1,omitempty"`
	// ClippingCoordinateV2 is number that represents the vertical offset in pixels from the clipping origin to the furthermost horizontal side of the clipping rectangle. Shall be present if Image View Data.ClippingOrigin is present and non-zero. Shall not be present when ImageViewDetail.ImageIndicator is 0. Valid values - 0000–9999
	IvDataClippingCoordinateV2 string `json:"ivData.ClippingCoordinateV2,omitempty"`
	// LengthImageReferenceKey is the number of characters in the ImageViewData.ImageReferenceKey. Shall not be present when ImageViewDetail.ImageIndicator is 0. Valid values - 0000ImageReferenceKey is not present 0001–9999  Valid when ImageReferenceKey is present
	IvDataLengthImageReferenceKey string `json:"ivData.LengthImageReferenceKey,omitempty"`
	// ImageReferenceKey is assigned by the ECE institution that creates the CheckDetail or Return, and the related Image View Records. This designator, when used, shall uniquely identify the item image to the ECE institution. This designator is a special key with significance to the creating institution. It is intended to be used to locate within an archive the unique image associated with the item. The designator could be a full access path and name that would allow direct look up and access to the image, for example a URL. This shall match CheckDetailAddendumB.ImageReferenceKey, or ReturnAddendumCImageReferenceKey Record, if used. Valid size - 0 – 9999
	IvDataImageReferenceKey string `json:"ivData.ImageReferenceKey,omitempty"`
	// LengthDigitalSignature is the number of bytes in the Image View Data.DigitalSignature. Shall not be present when ImageViewDetail.ImageIndicator is 0.
	IvDataLengthDigitalSignature string `json:"ivData.LengthDigitalSignature,omitempty"`
	// DigitalSignature is created by applying the cryptographic algorithm and private/secret key against the data to be protected. The Digital Signature provides user authentication and data integrity. Shall be present only under clearing arrangements and when ImageViewDetail.DigitalSignatureIndicator is 1 Shall not be present when ImageViewDetail.ImageIndicator is 0. Valid size - 0-99999
	IvDataDigitalSignature string `json:"ivData.DigitalSignature,omitempty"`
	// LengthImageData is the number of bytes in the ImageViewData.ImageData. Shall be present when ImageViewDetail.ImageIndicator is NOT 0 Valid values - 0000001–99999999
	IvDataLengthImageData string `json:"ivData.LengthImageData,omitempty"`
	// ImageData contains the image view. The Image Data generally consists of an image header and the image raster data. The image header provides information that is required to interpret the image raster data. The image raster data contains the scanned image of the physical item in raster (line by line) format. Each scan line comprises a set of concatenated pixels. The image comprises a set of scan lines. The image raster data is typically compressed to reduce the number of bytes needed to transmit and store the image. The header/image format type is defined by the ImageViewDetail.ImageViewFormatIndicator . The syntax and semantics of the image header/image format are understood by referring to the appropriate image format specification. The compression scheme used to compress the image raster data is specified in the ImageViewCompressionAlgorithmIdentifier and in the image header portion of the Image Data or by association with the selected image format. Shall be present when ImageViewDetail.ImageIndicator Record is NOT 0. Valid size - 0-9999999
	IvDataImageData string `json:"ivData.ImageData,omitempty"`
}
