# MxDetails

Details about the mail server\'s MX records.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**accepts_mail** | **boolean** | Indicates if the mail server accepts emails. | [default to undefined]
**records** | **Array&lt;string&gt;** | List of Fully Qualified Domain Names (FQDN) of the mail server. | [default to undefined]

## Example

```typescript
import { MxDetails } from '@oppulence-engineering/reacher-sdk';

const instance: MxDetails = {
    accepts_mail,
    records,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
