# CheckEmailOutput

The result of the email verification process.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**input** | **string** | The email address that was verified. | [default to undefined]
**is_reachable** | [**Reachable**](Reachable.md) |  | [default to undefined]
**misc** | [**CheckEmailOutputMisc**](CheckEmailOutputMisc.md) |  | [default to undefined]
**mx** | [**CheckEmailOutputMx**](CheckEmailOutputMx.md) |  | [default to undefined]
**smtp** | [**CheckEmailOutputSmtp**](CheckEmailOutputSmtp.md) |  | [default to undefined]
**syntax** | [**SyntaxDetails**](SyntaxDetails.md) |  | [default to undefined]
**debug** | [**DebugDetails**](DebugDetails.md) |  | [optional] [default to undefined]

## Example

```typescript
import { CheckEmailOutput } from '@oppulence/reacher-sdk';

const instance: CheckEmailOutput = {
    input,
    is_reachable,
    misc,
    mx,
    smtp,
    syntax,
    debug,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
