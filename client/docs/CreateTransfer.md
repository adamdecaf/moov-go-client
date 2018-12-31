# CreateTransfer

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransferType** | **string** | Type of transaction being actioned against the receiving institution. Expected values are pull (debits) or push (credits). | [optional] 
**Amount** | **string** | Amount of money. USD - United States. | 
**Originator** | **string** | ID of the Originator account initiating the transfer. | 
**OriginatorDepository** | **string** | ID of the Originator Depository to be be used to override the default depository. | [optional] 
**Customer** | **string** | ID of the Customer account the transfer was sent to. | 
**CustomerDepository** | **string** | ID of the Customer Depository to be used to override the default depository | [optional] 
**Description** | **string** | Brief description of the transaction, that may appear on the receiving entity’s financial statement | 
**StandardEntryClassCode** | **string** | Standard Entry Class code will be generated based on Customer type for CCD and PPD | [optional] 
**SameDay** | **bool** | When set to true this indicates the transfer should be processed the same day if possible. | [optional] [default to false]
**WEBDetail** | [**WebDetail**](WEBDetail.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

