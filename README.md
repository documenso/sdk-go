<img src="https://github.com/documenso/documenso/assets/13398220/a643571f-0239-46a6-a73e-6bef38d1228b" alt="Documenso Logo">

&nbsp;

<div align="center">
    <a href="https://www.speakeasy.com/?utm_source=github-com/documenso/sdk-go&utm_campaign=go"><img src="https://custom-icon-badges.demolab.com/badge/-Built%20By%20Speakeasy-212015?style=for-the-badge&logoColor=FBE331&logo=speakeasy&labelColor=545454" /></a>
    <a href="https://opensource.org/licenses/MIT">
        <img src="https://img.shields.io/badge/License-MIT-blue.svg" style="width: 100px; height: 28px;" />
    </a>
</div>

## Documenso Go SDK

A SDK for seamless integration with Documenso v2 API.

The full Documenso API can be viewed here (**todo**), which includes examples.

## ⚠️ Warning

Documenso v2 API and SDKs are currently in beta. There may be to breaking changes.

To keep updated, please follow the discussions here:
- [Feedback](https://github.com/documenso/documenso/discussions/1611)
- [Breaking change alerts](https://github.com/documenso/documenso/discussions/1612)

<!-- No Summary [summary] -->

## Table of Contents

<!-- $toc-max-depth=2 -->
* [Overview](#documenso-go-sdk)
  * [SDK Installation](#sdk-installation)
  * [Document creation example](#document-creation-example)
  * [Authentication](#authentication)
  * [Available Resources and Operations](#available-resources-and-operations)
  * [Retries](#retries)
  * [Error Handling](#error-handling)
  * [Server Selection](#server-selection)
* [Development](#development)
  * [Maturity](#maturity)
  * [Contributions](#contributions)

<!-- No Table of Contents [toc] -->

<!-- Start SDK Installation [installation] -->
## SDK Installation

To add the SDK as a dependency to your project:
```bash
go get github.com/documenso/sdk-go
```
<!-- End SDK Installation [installation] -->

## Authentication

To use the SDK, you will need a Documenso API key which can be created [here](https://docs.documenso.com/developers/public-api/authentication#creating-an-api-key
).

```go
documenso := documensoSdk.New(
  documensoSdk.WithSecurity(os.Getenv("DOCUMENSO_API_KEY")),
)
```
<!-- No Authentication [security] -->

## Document creation example

Currently creating a document involves two steps:

1. Create the document
2. Upload the PDF

This is a temporary measure, in the near future prior to the full release we will merge these two tasks into one request. 

Here is a full example of the document creation process which you can copy and run.

Note that the function is temporarily called `createV0`, which will be replaced by `create` once we resolve the 2 step workaround.

```go
package main

import (
	"bytes"
	"context"
	"errors"
	"io"
	"log"
	"net/http"
	"os"

	documensoSdk "github.com/documenso/sdk-go"
	"github.com/documenso/sdk-go/models/operations"
)

func uploadFileToPresignedUrl(filePath string, uploadUrl string) error {
	// Read file
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	// Create request
	req, err := http.NewRequest("PUT", uploadUrl, bytes.NewReader(fileContent))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/octet-stream")

	// Make request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check response
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return errors.New("upload failed with status: " + resp.Status + " body: " + string(body))
	}

	return nil
}

func main() {
	ctx := context.Background()

	documenso := documensoSdk.New(
		documensoSdk.WithSecurity("<API_KEY>"),
	)

	res, err := documenso.Documents.CreateV0(ctx, operations.DocumentCreateDocumentTemporaryRequestBody{
		Title: "Document title",
		Meta: &operations.Meta{
			Subject:              documensoSdk.String("Email subject"),
			Message:              documensoSdk.String("Email message"),
			TypedSignatureEnabled: documensoSdk.Bool(false),
		},
	})


	if err != nil {
		log.Fatal(err)
	}


	// Upload file
	err = uploadFileToPresignedUrl("./demo.pdf", res.Object.UploadURL)

	if err != nil {
		log.Fatal(err)
	}
}
```
<!-- No SDK Example Usage [usage] -->

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

<details open>
<summary>Available methods</summary>

### [Document](docs/sdks/document/README.md)

* [DocumentGetMany](docs/sdks/document/README.md#documentgetmany) - Get multiple documents
* [DocumentDownload](docs/sdks/document/README.md#documentdownload) - Download document (beta)

### [Documents](docs/sdks/documents/README.md)

* [Get](docs/sdks/documents/README.md#get) - Get document
* [Find](docs/sdks/documents/README.md#find) - Find documents
* [Create](docs/sdks/documents/README.md#create) - Create document
* [Update](docs/sdks/documents/README.md#update) - Update document
* [Delete](docs/sdks/documents/README.md#delete) - Delete document
* [Duplicate](docs/sdks/documents/README.md#duplicate) - Duplicate document
* [Distribute](docs/sdks/documents/README.md#distribute) - Distribute document
* [Redistribute](docs/sdks/documents/README.md#redistribute) - Redistribute document
* [Download](docs/sdks/documents/README.md#download) - Download document
* [CreateV0](docs/sdks/documents/README.md#createv0) - Create document

#### [Documents.Attachments](docs/sdks/documentsattachments/README.md)

* [Create](docs/sdks/documentsattachments/README.md#create) - Create attachment
* [Update](docs/sdks/documentsattachments/README.md#update) - Update attachment
* [Delete](docs/sdks/documentsattachments/README.md#delete) - Delete attachment
* [Find](docs/sdks/documentsattachments/README.md#find) - Find attachments

#### [Documents.Fields](docs/sdks/documentsfields/README.md)

* [Get](docs/sdks/documentsfields/README.md#get) - Get document field
* [Create](docs/sdks/documentsfields/README.md#create) - Create document field
* [CreateMany](docs/sdks/documentsfields/README.md#createmany) - Create document fields
* [Update](docs/sdks/documentsfields/README.md#update) - Update document field
* [UpdateMany](docs/sdks/documentsfields/README.md#updatemany) - Update document fields
* [Delete](docs/sdks/documentsfields/README.md#delete) - Delete document field

#### [Documents.Recipients](docs/sdks/documentsrecipients/README.md)

* [Get](docs/sdks/documentsrecipients/README.md#get) - Get document recipient
* [Create](docs/sdks/documentsrecipients/README.md#create) - Create document recipient
* [CreateMany](docs/sdks/documentsrecipients/README.md#createmany) - Create document recipients
* [Update](docs/sdks/documentsrecipients/README.md#update) - Update document recipient
* [UpdateMany](docs/sdks/documentsrecipients/README.md#updatemany) - Update document recipients
* [Delete](docs/sdks/documentsrecipients/README.md#delete) - Delete document recipient

### [Embedding](docs/sdks/embedding/README.md)

* [EmbeddingPresignCreateEmbeddingPresignToken](docs/sdks/embedding/README.md#embeddingpresigncreateembeddingpresigntoken) - Create embedding presign token
* [EmbeddingPresignVerifyEmbeddingPresignToken](docs/sdks/embedding/README.md#embeddingpresignverifyembeddingpresigntoken) - Verify embedding presign token

### [Envelope](docs/sdks/envelope/README.md)

* [EnvelopeFind](docs/sdks/envelope/README.md#envelopefind) - Find envelopes
* [EnvelopeAuditLogFind](docs/sdks/envelope/README.md#envelopeauditlogfind) - Get envelope audit logs
* [EnvelopeGetMany](docs/sdks/envelope/README.md#envelopegetmany) - Get multiple envelopes

### [Envelopes](docs/sdks/envelopes/README.md)

* [Get](docs/sdks/envelopes/README.md#get) - Get envelope
* [Create](docs/sdks/envelopes/README.md#create) - Create envelope
* [Use](docs/sdks/envelopes/README.md#use) - Use envelope
* [Update](docs/sdks/envelopes/README.md#update) - Update envelope
* [Delete](docs/sdks/envelopes/README.md#delete) - Delete envelope
* [Duplicate](docs/sdks/envelopes/README.md#duplicate) - Duplicate envelope
* [Distribute](docs/sdks/envelopes/README.md#distribute) - Distribute envelope
* [Redistribute](docs/sdks/envelopes/README.md#redistribute) - Redistribute envelope

#### [Envelopes.Attachments](docs/sdks/envelopesattachments/README.md)

* [Find](docs/sdks/envelopesattachments/README.md#find) - Find attachments
* [Create](docs/sdks/envelopesattachments/README.md#create) - Create attachment
* [Update](docs/sdks/envelopesattachments/README.md#update) - Update attachment
* [Delete](docs/sdks/envelopesattachments/README.md#delete) - Delete attachment

#### [Envelopes.Fields](docs/sdks/envelopesfields/README.md)

* [Get](docs/sdks/envelopesfields/README.md#get) - Get envelope field
* [CreateMany](docs/sdks/envelopesfields/README.md#createmany) - Create envelope fields
* [UpdateMany](docs/sdks/envelopesfields/README.md#updatemany) - Update envelope fields
* [Delete](docs/sdks/envelopesfields/README.md#delete) - Delete envelope field

#### [Envelopes.Items](docs/sdks/items/README.md)

* [CreateMany](docs/sdks/items/README.md#createmany) - Create envelope items
* [UpdateMany](docs/sdks/items/README.md#updatemany) - Update envelope items
* [Delete](docs/sdks/items/README.md#delete) - Delete envelope item
* [Download](docs/sdks/items/README.md#download) - Download an envelope item

#### [Envelopes.Recipients](docs/sdks/envelopesrecipients/README.md)

* [Get](docs/sdks/envelopesrecipients/README.md#get) - Get envelope recipient
* [CreateMany](docs/sdks/envelopesrecipients/README.md#createmany) - Create envelope recipients
* [UpdateMany](docs/sdks/envelopesrecipients/README.md#updatemany) - Update envelope recipients
* [Delete](docs/sdks/envelopesrecipients/README.md#delete) - Delete envelope recipient

### [Folders](docs/sdks/folders/README.md)

* [Find](docs/sdks/folders/README.md#find) - Find folders
* [Create](docs/sdks/folders/README.md#create) - Create new folder
* [Update](docs/sdks/folders/README.md#update) - Update folder
* [Delete](docs/sdks/folders/README.md#delete) - Delete folder

### [Template](docs/sdks/template/README.md)

* [TemplateGetMany](docs/sdks/template/README.md#templategetmany) - Get multiple templates
* [TemplateCreateTemplateTemporary](docs/sdks/template/README.md#templatecreatetemplatetemporary) - Create template

### [Templates](docs/sdks/templates/README.md)

* [Find](docs/sdks/templates/README.md#find) - Find templates
* [Get](docs/sdks/templates/README.md#get) - Get template
* [Create](docs/sdks/templates/README.md#create) - Create template
* [Update](docs/sdks/templates/README.md#update) - Update template
* [Duplicate](docs/sdks/templates/README.md#duplicate) - Duplicate template
* [Delete](docs/sdks/templates/README.md#delete) - Delete template
* [Use](docs/sdks/templates/README.md#use) - Use template

#### [Templates.DirectLink](docs/sdks/directlink/README.md)

* [Create](docs/sdks/directlink/README.md#create) - Create direct link
* [Delete](docs/sdks/directlink/README.md#delete) - Delete direct link
* [Toggle](docs/sdks/directlink/README.md#toggle) - Toggle direct link

#### [Templates.Fields](docs/sdks/templatesfields/README.md)

* [Create](docs/sdks/templatesfields/README.md#create) - Create template field
* [Get](docs/sdks/templatesfields/README.md#get) - Get template field
* [CreateMany](docs/sdks/templatesfields/README.md#createmany) - Create template fields
* [Update](docs/sdks/templatesfields/README.md#update) - Update template field
* [UpdateMany](docs/sdks/templatesfields/README.md#updatemany) - Update template fields
* [Delete](docs/sdks/templatesfields/README.md#delete) - Delete template field

#### [Templates.Recipients](docs/sdks/templatesrecipients/README.md)

* [Get](docs/sdks/templatesrecipients/README.md#get) - Get template recipient
* [Create](docs/sdks/templatesrecipients/README.md#create) - Create template recipient
* [CreateMany](docs/sdks/templatesrecipients/README.md#createmany) - Create template recipients
* [Update](docs/sdks/templatesrecipients/README.md#update) - Update template recipient
* [UpdateMany](docs/sdks/templatesrecipients/README.md#updatemany) - Update template recipients
* [Delete](docs/sdks/templatesrecipients/README.md#delete) - Delete template recipient

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

	res, err := s.Envelopes.Get(ctx, "<id>", operations.WithRetries(
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

	res, err := s.Envelopes.Get(ctx, "<id>")
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

For example, the `Get` function may return the following errors:

| Error Type                               | Status Code | Content Type     |
| ---------------------------------------- | ----------- | ---------------- |
| apierrors.EnvelopeGetBadRequestError     | 400         | application/json |
| apierrors.EnvelopeGetUnauthorizedError   | 401         | application/json |
| apierrors.EnvelopeGetForbiddenError      | 403         | application/json |
| apierrors.EnvelopeGetNotFoundError       | 404         | application/json |
| apierrors.EnvelopeGetInternalServerError | 500         | application/json |
| apierrors.APIError                       | 4XX, 5XX    | \*/\*            |

### Example

```go
package main

import (
	"context"
	"errors"
	sdkgo "github.com/documenso/sdk-go"
	"github.com/documenso/sdk-go/models/apierrors"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := sdkgo.New(
		sdkgo.WithSecurity(os.Getenv("DOCUMENSO_API_KEY")),
	)

	res, err := s.Envelopes.Get(ctx, "<id>")
	if err != nil {

		var e *apierrors.EnvelopeGetBadRequestError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *apierrors.EnvelopeGetUnauthorizedError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *apierrors.EnvelopeGetForbiddenError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *apierrors.EnvelopeGetNotFoundError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *apierrors.EnvelopeGetInternalServerError
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

The default server can be overridden globally using the `WithServerURL(serverURL string)` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	sdkgo "github.com/documenso/sdk-go"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := sdkgo.New(
		sdkgo.WithServerURL("https://app.documenso.com/api/v2"),
		sdkgo.WithSecurity(os.Getenv("DOCUMENSO_API_KEY")),
	)

	res, err := s.Envelopes.Get(ctx, "<id>")
	if err != nil {
		log.Fatal(err)
	}
	if res.Object != nil {
		// handle response
	}
}

```
<!-- End Server Selection [server] -->

<!-- No Custom HTTP Client [http-client] -->

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
