# SyntaxDetails

Validation of the email address syntax.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**domain** | **string** | The domain part of the email address. | [default to undefined]
**is_valid_syntax** | **boolean** | Indicates if the email address syntax is valid. | [default to undefined]
**username** | **string** | The username part of the email address. | [default to undefined]

## Example

```typescript
import { SyntaxDetails } from '@oppulence-engineering/reacher-sdk';

const instance: SyntaxDetails = {
    domain,
    is_valid_syntax,
    username,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
