# GetV1Bulk200Response


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**job_id** | **number** |  | [default to undefined]
**created_at** | **string** | The date and time when the bulk verification job was created. | [default to undefined]
**finished_at** | **string** | If the bulk verification job is completed, the date and time when it was finished. | [optional] [default to undefined]
**total_records** | **number** | The number of emails to verify in the bulk verification job. | [default to undefined]
**total_processed** | **number** | The number of emails that have been verified at the time of the query. | [default to undefined]
**summary** | [**GetV1Bulk200ResponseSummary**](GetV1Bulk200ResponseSummary.md) |  | [default to undefined]
**job_status** | **string** | The status of the job, either \&quot;Running\&quot; or \&quot;Completed\&quot;. | [default to undefined]

## Example

```typescript
import { GetV1Bulk200Response } from '@oppulence/reacher-sdk';

const instance: GetV1Bulk200Response = {
    job_id,
    created_at,
    finished_at,
    total_records,
    total_processed,
    summary,
    job_status,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
