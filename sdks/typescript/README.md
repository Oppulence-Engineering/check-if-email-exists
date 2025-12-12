# @oppulence-engineering/reacher-sdk

Official TypeScript SDK for the Reacher Email Verification API. Auto-generated from the OpenAPI specification.

## Installation

```bash
npm install @oppulence-engineering/reacher-sdk
# or
yarn add @oppulence-engineering/reacher-sdk
# or
pnpm add @oppulence-engineering/reacher-sdk
```

## Quick Start

```typescript
import { Configuration, DefaultApi } from '@oppulence-engineering/reacher-sdk';

// Initialize the client
const config = new Configuration({
  basePath: 'https://api.reacher.email',
  baseOptions: {
    headers: {
      'Authorization': 'YOUR_API_KEY'
    }
  }
});

const api = new DefaultApi(config);

// Verify a single email
async function verifyEmail() {
  try {
    const result = await api.postV1CheckEmail({
      checkEmailRequest: {
        toEmail: 'test@example.com'
      }
    });

    console.log('Verification result:', result.data);
    console.log('Is reachable:', result.data.isReachable);
  } catch (error) {
    console.error('Verification failed:', error);
  }
}

// Bulk email verification
async function bulkVerify() {
  // Start bulk job
  const bulkResult = await api.postV1Bulk({
    postV1BulkRequest: {
      input: ['email1@example.com', 'email2@example.com']
    }
  });

  const jobId = bulkResult.data.jobId;
  console.log('Bulk job started:', jobId);

  // Check progress
  const progress = await api.getV1Bulk({ jobId });
  console.log('Progress:', progress.data.totalProcessed, '/', progress.data.totalRecords);

  // Get results when completed
  if (progress.data.jobStatus === 'Completed') {
    const results = await api.getV1BulkResults({ jobId: String(jobId) });
    console.log('Results:', results.data);
  }
}
```

## API Reference

### Single Email Verification

```typescript
const result = await api.postV1CheckEmail({
  checkEmailRequest: {
    toEmail: 'user@domain.com',
    fromEmail: 'verify@yourdomain.com',      // Optional
    helloName: 'yourdomain.com',              // Optional
    smtpPort: 25,                              // Optional
    checkGravatar: true,                       // Optional
    proxy: {                                   // Optional
      host: 'proxy.example.com',
      port: 1080,
      username: 'user',
      password: 'pass'
    }
  }
});
```

### Response Types

```typescript
interface CheckEmailOutput {
  input: string;
  isReachable: 'safe' | 'risky' | 'invalid' | 'unknown';
  misc: MiscDetails | CoreError;
  mx: MxDetails | CoreError;
  smtp: SmtpDetails | CoreError;
  syntax: SyntaxDetails;
  debug?: DebugDetails;
}
```

## Configuration Options

```typescript
const config = new Configuration({
  basePath: 'https://api.reacher.email',  // API base URL
  baseOptions: {
    headers: {
      'Authorization': 'YOUR_API_KEY'       // Required for all requests
    },
    timeout: 30000                          // Optional timeout
  }
});
```

## Error Handling

```typescript
import { AxiosError } from 'axios';

try {
  const result = await api.postV1CheckEmail({ ... });
} catch (error) {
  if (error instanceof AxiosError) {
    console.error('HTTP Error:', error.response?.status);
    console.error('Message:', error.response?.data);
  }
}
```

## License

MIT
