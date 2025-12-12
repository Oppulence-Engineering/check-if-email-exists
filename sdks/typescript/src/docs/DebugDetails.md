# DebugDetails


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**start_time** | **string** | The timestamp when the email verification started. | [default to undefined]
**end_time** | **string** | The timestamp when the email verification ended. | [default to undefined]
**duration** | [**Duration**](Duration.md) |  | [default to undefined]
**server_name** | **string** | The name of the server that performed the verification. | [default to undefined]
**smtp** | [**DebugDetailsSmtp**](DebugDetailsSmtp.md) |  | [default to undefined]

## Example

```typescript
import { DebugDetails } from '@oppulence-engineering/reacher-sdk';

const instance: DebugDetails = {
    start_time,
    end_time,
    duration,
    server_name,
    smtp,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
