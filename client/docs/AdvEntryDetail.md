# AdvEntryDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ID** | **string** | Entry Detail ID | [optional] 
**TransactionCode** | **int32** | TransactionCode representing Accounting Entries Credit for ACH debits originated - 81 Debit for ACH credits originated - 82 Credit for ACH credits received 83 Debit for ACH debits received 84 Credit for ACH credits in rejected batches 85 Debit for ACH debits in rejected batches - 86 Summary credit for respondent ACH activity - 87 Summary debit for respondent ACH activity - 88  | 
**RDFIIdentification** | **string** | RDFI&#39;s routing number without the last digit. | 
**CheckDigit** | **string** | Last digit in RDFI routing number. | 
**DFIAccountNumber** | **string** | The receiver&#39;s bank account number you are crediting/debiting. It important to note that this is an alphanumeric field, so its space padded, no zero padded  | 
**Amount** | **int32** | Number of cents you are debiting/crediting this account | 
**AdviceRoutingNumber** | **string** | Suggested routing number to use | [optional] 
**FileIdentification** | **string** | Unique identifier for the File | [optional] 
**AchOperatorData** | **string** | Information related to the ACH opreator | [optional] 
**IndividualName** | **string** | The name of the receiver, usually the name on the bank account | 
**DiscretionaryData** | **string** | DiscretionaryData allows ODFIs to include codes, of significance only to them, to enable specialized handling of the entry. There will be no standardized interpretation for the value of this field. It can either be a single two-character code, or two distinct one-character codes, according to the needs of the ODFI and/or Originator involved. This field must be returned intact for any returned entry. WEB uses the Discretionary Data Field as the Payment Type Code  | [optional] 
**AddendaRecordIndicator** | **int32** | AddendaRecordIndicator indicates the existence of an Addenda Record. A value of \&quot;1\&quot; indicates that one ore more addenda records follow, and \&quot;0\&quot; means no such record is present.  | [optional] 
**AchOperatorRoutingNumber** | **string** | Routing number for ACH operator | [optional] 
**JulianDay** | **float32** | Julian Day of the year | [optional] 
**SequenceNumber** | **float32** | SequenceNumber is consecutively assigned to each Addenda05 Record following an Entry Detail Record. The first addenda05 sequence number must always be a 1. | [optional] 
**Addenda99** | [**[]Addenda99**](Addenda99.md) | Addenda99 record for the Entry Detail | [optional] 
**Category** | **string** | Category defines if the entry is a Forward, Return, or NOC | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


