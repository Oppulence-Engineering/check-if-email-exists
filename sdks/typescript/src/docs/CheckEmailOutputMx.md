# CheckEmailOutputMx

Details obtained from querying the mail server\'s MX records.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**accepts_mail** | **boolean** | Indicates if the mail server accepts emails. | [default to undefined]
**records** | **Array&lt;string&gt;** | List of Fully Qualified Domain Names (FQDN) of the mail server. | [default to undefined]
**type** | **string** | The type of error. | [default to undefined]
**message** | **string** | A human-readable description of the error. | [default to undefined]

## Example

```typescript
import { CheckEmailOutputMx } from '@oppulence-engineering/reacher-sdk';

const instance: CheckEmailOutputMx = {
    accepts_mail,
    records,
    type,
    message,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
