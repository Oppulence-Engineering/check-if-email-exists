## @oppulence-engineering/reacher-sdk@0.11.0

This generator creates TypeScript/JavaScript client that utilizes [axios](https://github.com/axios/axios). The generated Node module can be used in the following environments:

Environment
* Node.js
* Webpack
* Browserify

Language level
* ES5 - you must have a Promises/A+ library installed
* ES6

Module system
* CommonJS
* ES6 module system

It can be used in both TypeScript and JavaScript. In TypeScript, the definition will be automatically resolved via `package.json`. ([Reference](https://www.typescriptlang.org/docs/handbook/declaration-files/consumption.html))

### Building

To build and compile the typescript sources to javascript use:
```
npm install
npm run build
```

### Publishing

First build the package then run `npm publish`

### Consuming

navigate to the folder of your consuming project and run one of the following commands.

_published:_

```
npm install @oppulence-engineering/reacher-sdk@0.11.0 --save
```

_unPublished (not recommended):_

```
npm install PATH_TO_GENERATED_PACKAGE --save
```

### Documentation for API Endpoints

All URIs are relative to *https://api.reacher.email*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DefaultApi* | [**getV1Bulk**](docs/DefaultApi.md#getv1bulk) | **GET** /v1/bulk/{job_id} | /v1/bulk/{job_id}
*DefaultApi* | [**getV1BulkResults**](docs/DefaultApi.md#getv1bulkresults) | **GET** /v1/bulk/{job_id}/results | Retrieve bulk verification results
*DefaultApi* | [**postV0CheckEmail**](docs/DefaultApi.md#postv0checkemail) | **POST** /v0/check_email | /v0/check_email
*DefaultApi* | [**postV1Bulk**](docs/DefaultApi.md#postv1bulk) | **POST** /v1/bulk | /v1/bulk
*DefaultApi* | [**postV1CheckEmail**](docs/DefaultApi.md#postv1checkemail) | **POST** /v1/check_email | /v1/check_email


### Documentation For Models

 - [CheckEmailInputProxy](docs/CheckEmailInputProxy.md)
 - [CheckEmailOutput](docs/CheckEmailOutput.md)
 - [CheckEmailOutputMisc](docs/CheckEmailOutputMisc.md)
 - [CheckEmailOutputMx](docs/CheckEmailOutputMx.md)
 - [CheckEmailOutputSmtp](docs/CheckEmailOutputSmtp.md)
 - [CheckEmailRequest](docs/CheckEmailRequest.md)
 - [CoreError](docs/CoreError.md)
 - [DebugDetails](docs/DebugDetails.md)
 - [DebugDetailsSmtp](docs/DebugDetailsSmtp.md)
 - [Duration](docs/Duration.md)
 - [GetV1Bulk200Response](docs/GetV1Bulk200Response.md)
 - [GetV1Bulk200ResponseSummary](docs/GetV1Bulk200ResponseSummary.md)
 - [GetV1BulkResults200Response](docs/GetV1BulkResults200Response.md)
 - [GmailVerifMethod](docs/GmailVerifMethod.md)
 - [HotmailB2BVerifMethod](docs/HotmailB2BVerifMethod.md)
 - [HotmailB2CVerifMethod](docs/HotmailB2CVerifMethod.md)
 - [MiscDetails](docs/MiscDetails.md)
 - [MxDetails](docs/MxDetails.md)
 - [PostV1Bulk200Response](docs/PostV1Bulk200Response.md)
 - [PostV1BulkRequest](docs/PostV1BulkRequest.md)
 - [Reachable](docs/Reachable.md)
 - [SmtpDetails](docs/SmtpDetails.md)
 - [SyntaxDetails](docs/SyntaxDetails.md)
 - [TaskWebhook](docs/TaskWebhook.md)
 - [VerifMethod](docs/VerifMethod.md)
 - [Webhook](docs/Webhook.md)
 - [YahooVerifMethod](docs/YahooVerifMethod.md)


<a id="documentation-for-authorization"></a>
## Documentation For Authorization


Authentication schemes defined for the API:
<a id="Authorization"></a>
### Authorization

- **Type**: API key
- **API key parameter name**: Authorization
- **Location**: HTTP header

