# CheckEmailOutputMisc

Additional information about the email account.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**is_disposable** | **boolean** | Indicates if the email address is from a known disposable email provider. | [default to undefined]
**is_role_account** | **boolean** | Indicates if the email address is a role-based account. | [default to undefined]
**gravatar_url** | **string** | URL to the Gravatar profile picture associated with the email, if available and requested. | [optional] [default to undefined]
**is_b2c** | **boolean** | Is this a B2C email address? | [default to undefined]
**type** | **string** | The type of error. | [default to undefined]
**message** | **string** | A human-readable description of the error. | [default to undefined]

## Example

```typescript
import { CheckEmailOutputMisc } from '@oppulence-engineering/reacher-sdk';

const instance: CheckEmailOutputMisc = {
    is_disposable,
    is_role_account,
    gravatar_url,
    is_b2c,
    type,
    message,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
