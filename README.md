# github.com/documenso/sdk-go

Developer-friendly & type-safe Go SDK specifically catered to leverage *github.com/documenso/sdk-go* API.

<div align="left">
    <a href="https://www.speakeasy.com/?utm_source=github-com/documenso/sdk-go&utm_campaign=go"><img src="https://custom-icon-badges.demolab.com/badge/-Built%20By%20Speakeasy-212015?style=for-the-badge&logoColor=FBE331&logo=speakeasy&labelColor=545454" /></a>
    <a href="https://opensource.org/licenses/MIT">
        <img src="https://img.shields.io/badge/License-MIT-blue.svg" style="width: 100px; height: 28px;" />
    </a>
</div>


<br /><br />
> [!IMPORTANT]
> This SDK is not yet ready for production use. To complete setup please follow the steps outlined in your [workspace](https://app.speakeasy.com/org/documenso-inc/api). Delete this section before > publishing to a package manager.

<!-- Start Summary [summary] -->
## Summary

Documenso v2 beta API: Subject to breaking changes until v2 is fully released.
<!-- End Summary [summary] -->

<!-- Start Table of Contents [toc] -->
## Table of Contents
<!-- $toc-max-depth=2 -->
* [github.com/documenso/sdk-go](#githubcomdocumensosdk-go)
  * [SDK Installation](#sdk-installation)
  * [SDK Example Usage](#sdk-example-usage)
  * [Authentication](#authentication)
  * [Available Resources and Operations](#available-resources-and-operations)
  * [Retries](#retries)
  * [Error Handling](#error-handling)
  * [Server Selection](#server-selection)
  * [Custom HTTP Client](#custom-http-client)
* [Development](#development)
  * [Maturity](#maturity)
  * [Contributions](#contributions)

<!-- End Table of Contents [toc] -->

<!-- Start SDK Installation [installation] -->
## SDK Installation

To add the SDK as a dependency to your project:
```bash
go get github.com/documenso/sdk-go
```
<!-- End SDK Installation [installation] -->

<!-- Start SDK Example Usage [usage] -->
## SDK Example Usage

### Example

```go
package main

import (
	"context"
	sdkgo "github.com/documenso/sdk-go"
	"github.com/documenso/sdk-go/models/operations"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := sdkgo.New(
		sdkgo.WithSecurity(os.Getenv("DOCUMENSO_API_KEY")),
	)

	res, err := s.Documents.Find(ctx, operations.DocumentFindDocumentsRequest{
		OrderByDirection: operations.OrderByDirectionDesc.ToPointer(),
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.Object != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->

<!-- Start Authentication [security] -->
## Authentication

### Per-Client Security Schemes

This SDK supports the following security scheme globally:

| Name     | Type   | Scheme  | Environment Variable |
| -------- | ------ | ------- | -------------------- |
| `APIKey` | apiKey | API key | `DOCUMENSO_API_KEY`  |

You can configure it using the `WithSecurity` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	sdkgo "github.com/documenso/sdk-go"
	"github.com/documenso/sdk-go/models/operations"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := sdkgo.New(
		sdkgo.WithSecurity(os.Getenv("DOCUMENSO_API_KEY")),
	)

	res, err := s.Documents.Find(ctx, operations.DocumentFindDocumentsRequest{
		OrderByDirection: operations.OrderByDirectionDesc.ToPointer(),
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.Object != nil {
		// handle response
	}
}

```
<!-- End Authentication [security] -->

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

<details open>
<summary>Available methods</summary>


### [Documents](docs/sdks/documents/README.md)

* [Find](docs/sdks/documents/README.md#find) - Find documents
* [Get](docs/sdks/documents/README.md#get) - Get document
* [CreateV0](docs/sdks/documents/README.md#createv0) - Create document
* [Update](docs/sdks/documents/README.md#update) - Update document
* [Delete](docs/sdks/documents/README.md#delete) - Delete document
* [MoveToTeam](docs/sdks/documents/README.md#movetoteam) - Move document
* [Distribute](docs/sdks/documents/README.md#distribute) - Distribute document
* [Redistribute](docs/sdks/documents/README.md#redistribute) - Redistribute document
* [Duplicate](docs/sdks/documents/README.md#duplicate) - Duplicate document

#### [Documents.Fields](docs/sdks/fields/README.md)

* [Get](docs/sdks/fields/README.md#get) - Get document field
* [Create](docs/sdks/fields/README.md#create) - Create document field
* [CreateMany](docs/sdks/fields/README.md#createmany) - Create document fields
* [Update](docs/sdks/fields/README.md#update) - Update document field
* [UpdateMany](docs/sdks/fields/README.md#updatemany) - Update document fields
* [Delete](docs/sdks/fields/README.md#delete) - Delete document field

#### [Documents.Recipients](docs/sdks/recipients/README.md)

* [Get](docs/sdks/recipients/README.md#get) - Get document recipient
* [Create](docs/sdks/recipients/README.md#create) - Create document recipient
* [CreateMany](docs/sdks/recipients/README.md#createmany) - Create document recipients
* [Update](docs/sdks/recipients/README.md#update) - Update document recipient
* [UpdateMany](docs/sdks/recipients/README.md#updatemany) - Update document recipients
* [Delete](docs/sdks/recipients/README.md#delete) - Delete document recipient

### [Templates](docs/sdks/templates/README.md)

* [Find](docs/sdks/templates/README.md#find) - Find templates
* [Get](docs/sdks/templates/README.md#get) - Get template
* [Update](docs/sdks/templates/README.md#update) - Update template
* [Duplicate](docs/sdks/templates/README.md#duplicate) - Duplicate template
* [Delete](docs/sdks/templates/README.md#delete) - Delete template
* [Use](docs/sdks/templates/README.md#use) - Use template
* [MoveToTeam](docs/sdks/templates/README.md#movetoteam) - Move template

#### [Templates.DirectLink](docs/sdks/directlink/README.md)

* [Create](docs/sdks/directlink/README.md#create) - Create direct link
* [Delete](docs/sdks/directlink/README.md#delete) - Delete direct link
* [Toggle](docs/sdks/directlink/README.md#toggle) - Toggle direct link

#### [Templates.Fields](docs/sdks/documensofields/README.md)

* [Create](docs/sdks/documensofields/README.md#create) - Create template field
* [Get](docs/sdks/documensofields/README.md#get) - Get template field
* [CreateMany](docs/sdks/documensofields/README.md#createmany) - Create template fields
* [Update](docs/sdks/documensofields/README.md#update) - Update template field
* [UpdateMany](docs/sdks/documensofields/README.md#updatemany) - Update template fields
* [Delete](docs/sdks/documensofields/README.md#delete) - Delete template field

#### [Templates.Recipients](docs/sdks/documensorecipients/README.md)

* [Get](docs/sdks/documensorecipients/README.md#get) - Get template recipient
* [Create](docs/sdks/documensorecipients/README.md#create) - Create template recipient
* [CreateMany](docs/sdks/documensorecipients/README.md#createmany) - Create template recipients
* [Update](docs/sdks/documensorecipients/README.md#update) - Update template recipient
* [UpdateMany](docs/sdks/documensorecipients/README.md#updatemany) - Update template recipients
* [Delete](docs/sdks/documensorecipients/README.md#delete) - Delete template recipient

</details>
<!-- End Available Resources and Operations [operations] -->

<!-- Start Retries [retries] -->
## Retries

Some of the endpoints in this SDK support retries. If you use the SDK without any configuration, it will fall back to the default retry strategy provided by the API. However, the default retry strategy can be overridden on a per-operation basis, or across the entire SDK.

To change the default retry strategy for a single API call, simply provide a `retry.Config` object to the call by using the `WithRetries` option:
```go
package main

import (
	"context"
	sdkgo "github.com/documenso/sdk-go"
	"github.com/documenso/sdk-go/models/operations"
	"github.com/documenso/sdk-go/retry"
	"log"
	"models/operations"
	"os"
)

func main() {
	ctx := context.Background()

	s := sdkgo.New(
		sdkgo.WithSecurity(os.Getenv("DOCUMENSO_API_KEY")),
	)

	res, err := s.Documents.Find(ctx, operations.DocumentFindDocumentsRequest{
		OrderByDirection: operations.OrderByDirectionDesc.ToPointer(),
	}, operations.WithRetries(
		retry.Config{
			Strategy: "backoff",
			Backoff: &retry.BackoffStrategy{
				InitialInterval: 1,
				MaxInterval:     50,
				Exponent:        1.1,
				MaxElapsedTime:  100,
			},
			RetryConnectionErrors: false,
		}))
	if err != nil {
		log.Fatal(err)
	}
	if res.Object != nil {
		// handle response
	}
}

```

If you'd like to override the default retry strategy for all operations that support retries, you can use the `WithRetryConfig` option at SDK initialization:
```go
package main

import (
	"context"
	sdkgo "github.com/documenso/sdk-go"
	"github.com/documenso/sdk-go/models/operations"
	"github.com/documenso/sdk-go/retry"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := sdkgo.New(
		sdkgo.WithRetryConfig(
			retry.Config{
				Strategy: "backoff",
				Backoff: &retry.BackoffStrategy{
					InitialInterval: 1,
					MaxInterval:     50,
					Exponent:        1.1,
					MaxElapsedTime:  100,
				},
				RetryConnectionErrors: false,
			}),
		sdkgo.WithSecurity(os.Getenv("DOCUMENSO_API_KEY")),
	)

	res, err := s.Documents.Find(ctx, operations.DocumentFindDocumentsRequest{
		OrderByDirection: operations.OrderByDirectionDesc.ToPointer(),
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.Object != nil {
		// handle response
	}
}

```
<!-- End Retries [retries] -->

<!-- Start Error Handling [errors] -->
## Error Handling

Handling errors in this SDK should largely match your expectations. All operations return a response object or an error, they will never return both.

By Default, an API error will return `apierrors.APIError`. When custom error responses are specified for an operation, the SDK may also return their associated error. You can refer to respective *Errors* tables in SDK docs for more details on possible error types for each operation.

For example, the `Find` function may return the following errors:

| Error Type                         | Status Code | Content Type     |
| ---------------------------------- | ----------- | ---------------- |
| apierrors.ErrorBADREQUEST          | 400         | application/json |
| apierrors.ErrorNOTFOUND            | 404         | application/json |
| apierrors.ERRORINTERNALSERVERERROR | 500         | application/json |
| apierrors.APIError                 | 4XX, 5XX    | \*/\*            |

### Example

```go
package main

import (
	"context"
	"errors"
	sdkgo "github.com/documenso/sdk-go"
	"github.com/documenso/sdk-go/models/apierrors"
	"github.com/documenso/sdk-go/models/operations"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := sdkgo.New(
		sdkgo.WithSecurity(os.Getenv("DOCUMENSO_API_KEY")),
	)

	res, err := s.Documents.Find(ctx, operations.DocumentFindDocumentsRequest{
		OrderByDirection: operations.OrderByDirectionDesc.ToPointer(),
	})
	if err != nil {

		var e *apierrors.ErrorBADREQUEST
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *apierrors.ErrorNOTFOUND
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *apierrors.ERRORINTERNALSERVERERROR
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *apierrors.APIError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}
	}
}

```
<!-- End Error Handling [errors] -->

<!-- Start Server Selection [server] -->
## Server Selection

### Override Server URL Per-Client

The default server can also be overridden globally using the `WithServerURL(serverURL string)` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	sdkgo "github.com/documenso/sdk-go"
	"github.com/documenso/sdk-go/models/operations"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := sdkgo.New(
		sdkgo.WithServerURL("https://app.documenso.com/api/v2-beta"),
		sdkgo.WithSecurity(os.Getenv("DOCUMENSO_API_KEY")),
	)

	res, err := s.Documents.Find(ctx, operations.DocumentFindDocumentsRequest{
		OrderByDirection: operations.OrderByDirectionDesc.ToPointer(),
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.Object != nil {
		// handle response
	}
}

```
<!-- End Server Selection [server] -->

<!-- Start Custom HTTP Client [http-client] -->
## Custom HTTP Client

The Go SDK makes API calls that wrap an internal HTTP client. The requirements for the HTTP client are very simple. It must match this interface:

```go
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
```

The built-in `net/http` client satisfies this interface and a default client based on the built-in is provided by default. To replace this default with a client of your own, you can implement this interface yourself or provide your own client configured as desired. Here's a simple example, which adds a client with a 30 second timeout.

```go
import (
	"net/http"
	"time"
	"github.com/myorg/your-go-sdk"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = sdk.New(sdk.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client [http-client] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->

# Development

## Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

## Contributions

While we value open-source contributions to this SDK, this library is generated programmatically. Any manual changes added to internal files will be overwritten on the next generation. 
We look forward to hearing your feedback. Feel free to open a PR or an issue with a proof of concept and we'll do our best to include it in a future release. 

### SDK Created by [Speakeasy](https://www.speakeasy.com/?utm_source=github-com/documenso/sdk-go&utm_campaign=go)
