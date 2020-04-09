# AdvBatchControl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ID** | **string** | Batch ID | [optional] 
**ServiceClassCode** | **int32** | Same as ServiceClassCode in BatchHeader record | [optional] 
**EntryAddendaCount** | **int32** | EntryAddendaCount is a tally of each Entry Detail Record and each Addenda Record processed, within either the batch or file as appropriate. | [optional] 
**EntryHash** | **int32** | Validate the Receiving DFI Identification in each Entry Detail Record is hashed to provide a check against inadvertent alteration of data contents due to hardware failure or program error. In this context the Entry Hash is the sum of the corresponding fields in the Entry Detail Records on the file.  | [optional] 
**TotalDebit** | **int32** | Contains accumulated Entry debit totals within the batch. | [optional] 
**TotalCredit** | **int32** | Contains accumulated Entry credit totals within the batch. | [optional] 
**AchOperatorData** | **string** | Alphanumeric code used to identify an ACH Operator | [optional] 
**ODFIIdentification** | **string** | The routing number is used to identify the DFI originating entries within a given branch. | [optional] 
**BatchNumber** | **string** | BatchNumber is assigned in ascending sequence to each batch by the ODFI or its Sending Point in a given file of entries. Since the batch number in the Batch Header Record and the Batch Control Record is the same, the ascending sequence number should be assigned by batch and not by record. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


