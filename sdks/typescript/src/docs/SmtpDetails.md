# SmtpDetails

Results from SMTP connection attempts to the mail server.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**can_connect_smtp** | **boolean** | Indicates if the mail exchanger can be contacted successfully. | [default to undefined]
**has_full_inbox** | **boolean** | Indicates if the mailbox is full. | [default to undefined]
**is_catch_all** | **boolean** | Indicates if the email address is a catch-all address. | [default to undefined]
**is_deliverable** | **boolean** | Indicates if an email sent to this address is deliverable. | [default to undefined]
**is_disabled** | **boolean** | Indicates if the email address has been disabled by the provider. | [default to undefined]

## Example

```typescript
import { SmtpDetails } from '@oppulence-engineering/reacher-sdk';

const instance: SmtpDetails = {
    can_connect_smtp,
    has_full_inbox,
    is_catch_all,
    is_deliverable,
    is_disabled,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
