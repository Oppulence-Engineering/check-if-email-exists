# CheckEmailRequest

A request object to perform an email verification. The `to_email` field is required, all other fields are optional.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**from_email** | **string** | In the SMTP connection, the FROM email address. | [optional] [default to undefined]
**to_email** | **string** | The email address to check. | [default to undefined]
**hello_name** | **string** | In the SMTP connection, the EHLO hostname. | [optional] [default to undefined]
**proxy** | [**CheckEmailInputProxy**](CheckEmailInputProxy.md) |  | [optional] [default to undefined]
**smtp_port** | **number** | SMTP port to use for email validation. Defaults to 25, but 465, 587, and 2525 are sometimes also used. | [optional] [default to undefined]
**gmail_verif_method** | [**GmailVerifMethod**](GmailVerifMethod.md) |  | [optional] [default to undefined]
**hotmailb2b_verif_method** | [**HotmailB2BVerifMethod**](HotmailB2BVerifMethod.md) |  | [optional] [default to undefined]
**hotmailb2c_verif_method** | [**HotmailB2CVerifMethod**](HotmailB2CVerifMethod.md) |  | [optional] [default to undefined]
**yahoo_verif_method** | [**YahooVerifMethod**](YahooVerifMethod.md) |  | [optional] [default to undefined]
**check_gravatar** | **boolean** | Whether to check if a Gravatar image exists for the given email. | [optional] [default to undefined]

## Example

```typescript
import { CheckEmailRequest } from '@oppulence/reacher-sdk';

const instance: CheckEmailRequest = {
    from_email,
    to_email,
    hello_name,
    proxy,
    smtp_port,
    gmail_verif_method,
    hotmailb2b_verif_method,
    hotmailb2c_verif_method,
    yahoo_verif_method,
    check_gravatar,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
