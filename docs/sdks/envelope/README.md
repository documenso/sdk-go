# Envelope

## Overview

### Available Operations

* [EnvelopeFind](#envelopefind) - Find envelopes
* [EnvelopeAuditLogFind](#envelopeauditlogfind) - Get envelope audit logs
* [EnvelopeAuditLogDownloadPdf](#envelopeauditlogdownloadpdf) - Download envelope audit log PDF
* [EnvelopeCertificateDownloadPdf](#envelopecertificatedownloadpdf) - Download envelope certificate PDF
* [EnvelopeGetMany](#envelopegetmany) - Get multiple envelopes
* [EnvelopeCancel](#envelopecancel) - Cancel envelope

## EnvelopeFind

Find envelopes based on search criteria

### Example Usage

<!-- UsageSnippet language="go" operationID="envelope-find" method="get" path="/envelope" -->
```go
package main

import(
	"context"
	"os"
	sdkgo "github.com/documenso/sdk-go"
	"github.com/documenso/sdk-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdkgo.New(
        sdkgo.WithSecurity(os.Getenv("DOCUMENSO_API_KEY")),
    )

    res, err := s.Envelope.EnvelopeFind(ctx, operations.EnvelopeFindRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.EnvelopeFindRequest](../../models/operations/envelopefindrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.EnvelopeFindResponse](../../models/operations/envelopefindresponse.md), error**

### Errors

| Error Type                                | Status Code                               | Content Type                              |
| ----------------------------------------- | ----------------------------------------- | ----------------------------------------- |
| apierrors.EnvelopeFindBadRequestError     | 400                                       | application/json                          |
| apierrors.EnvelopeFindUnauthorizedError   | 401                                       | application/json                          |
| apierrors.EnvelopeFindForbiddenError      | 403                                       | application/json                          |
| apierrors.EnvelopeFindNotFoundError       | 404                                       | application/json                          |
| apierrors.EnvelopeFindInternalServerError | 500                                       | application/json                          |
| apierrors.APIError                        | 4XX, 5XX                                  | \*/\*                                     |

## EnvelopeAuditLogFind

Find audit logs based on a search criteria

### Example Usage

<!-- UsageSnippet language="go" operationID="envelope-auditLog-find" method="get" path="/envelope/{envelopeId}/audit-log" -->
```go
package main

import(
	"context"
	"os"
	sdkgo "github.com/documenso/sdk-go"
	"github.com/documenso/sdk-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdkgo.New(
        sdkgo.WithSecurity(os.Getenv("DOCUMENSO_API_KEY")),
    )

    res, err := s.Envelope.EnvelopeAuditLogFind(ctx, operations.EnvelopeAuditLogFindRequest{
        EnvelopeID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.EnvelopeAuditLogFindRequest](../../models/operations/envelopeauditlogfindrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.EnvelopeAuditLogFindResponse](../../models/operations/envelopeauditlogfindresponse.md), error**

### Errors

| Error Type                                        | Status Code                                       | Content Type                                      |
| ------------------------------------------------- | ------------------------------------------------- | ------------------------------------------------- |
| apierrors.EnvelopeAuditLogFindBadRequestError     | 400                                               | application/json                                  |
| apierrors.EnvelopeAuditLogFindUnauthorizedError   | 401                                               | application/json                                  |
| apierrors.EnvelopeAuditLogFindForbiddenError      | 403                                               | application/json                                  |
| apierrors.EnvelopeAuditLogFindNotFoundError       | 404                                               | application/json                                  |
| apierrors.EnvelopeAuditLogFindInternalServerError | 500                                               | application/json                                  |
| apierrors.APIError                                | 4XX, 5XX                                          | \*/\*                                             |

## EnvelopeAuditLogDownloadPdf

Download the audit log for a document as a PDF.

### Example Usage

<!-- UsageSnippet language="go" operationID="envelope-auditLog-downloadPdf" method="get" path="/envelope/{envelopeId}/audit-log/download" -->
```go
package main

import(
	"context"
	"os"
	sdkgo "github.com/documenso/sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdkgo.New(
        sdkgo.WithSecurity(os.Getenv("DOCUMENSO_API_KEY")),
    )

    res, err := s.Envelope.EnvelopeAuditLogDownloadPdf(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.Any != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `envelopeID`                                             | `string`                                                 | :heavy_check_mark:                                       | The ID of the envelope to download the audit log for.    |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.EnvelopeAuditLogDownloadPdfResponse](../../models/operations/envelopeauditlogdownloadpdfresponse.md), error**

### Errors

| Error Type                                               | Status Code                                              | Content Type                                             |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| apierrors.EnvelopeAuditLogDownloadPdfBadRequestError     | 400                                                      | application/json                                         |
| apierrors.EnvelopeAuditLogDownloadPdfUnauthorizedError   | 401                                                      | application/json                                         |
| apierrors.EnvelopeAuditLogDownloadPdfForbiddenError      | 403                                                      | application/json                                         |
| apierrors.EnvelopeAuditLogDownloadPdfNotFoundError       | 404                                                      | application/json                                         |
| apierrors.EnvelopeAuditLogDownloadPdfInternalServerError | 500                                                      | application/json                                         |
| apierrors.APIError                                       | 4XX, 5XX                                                 | \*/\*                                                    |

## EnvelopeCertificateDownloadPdf

Download the signing certificate for a completed document as a PDF.

### Example Usage

<!-- UsageSnippet language="go" operationID="envelope-certificate-downloadPdf" method="get" path="/envelope/{envelopeId}/certificate/download" -->
```go
package main

import(
	"context"
	"os"
	sdkgo "github.com/documenso/sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdkgo.New(
        sdkgo.WithSecurity(os.Getenv("DOCUMENSO_API_KEY")),
    )

    res, err := s.Envelope.EnvelopeCertificateDownloadPdf(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.Any != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `envelopeID`                                             | `string`                                                 | :heavy_check_mark:                                       | The ID of the envelope to download the certificate for.  |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.EnvelopeCertificateDownloadPdfResponse](../../models/operations/envelopecertificatedownloadpdfresponse.md), error**

### Errors

| Error Type                                                  | Status Code                                                 | Content Type                                                |
| ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- |
| apierrors.EnvelopeCertificateDownloadPdfBadRequestError     | 400                                                         | application/json                                            |
| apierrors.EnvelopeCertificateDownloadPdfUnauthorizedError   | 401                                                         | application/json                                            |
| apierrors.EnvelopeCertificateDownloadPdfForbiddenError      | 403                                                         | application/json                                            |
| apierrors.EnvelopeCertificateDownloadPdfNotFoundError       | 404                                                         | application/json                                            |
| apierrors.EnvelopeCertificateDownloadPdfInternalServerError | 500                                                         | application/json                                            |
| apierrors.APIError                                          | 4XX, 5XX                                                    | \*/\*                                                       |

## EnvelopeGetMany

Retrieve multiple envelopes by their IDs

### Example Usage

<!-- UsageSnippet language="go" operationID="envelope-getMany" method="post" path="/envelope/get-many" -->
```go
package main

import(
	"context"
	"os"
	sdkgo "github.com/documenso/sdk-go"
	"github.com/documenso/sdk-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdkgo.New(
        sdkgo.WithSecurity(os.Getenv("DOCUMENSO_API_KEY")),
    )

    res, err := s.Envelope.EnvelopeGetMany(ctx, operations.EnvelopeGetManyRequest{
        Ids: operations.CreateIdsIdsEnvelopeID(
            operations.IdsEnvelopeID{
                Type: operations.TypeEnvelopeIDEnvelopeID,
                Ids: []string{},
            },
        ),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.EnvelopeGetManyRequest](../../models/operations/envelopegetmanyrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.EnvelopeGetManyResponse](../../models/operations/envelopegetmanyresponse.md), error**

### Errors

| Error Type                                   | Status Code                                  | Content Type                                 |
| -------------------------------------------- | -------------------------------------------- | -------------------------------------------- |
| apierrors.EnvelopeGetManyBadRequestError     | 400                                          | application/json                             |
| apierrors.EnvelopeGetManyUnauthorizedError   | 401                                          | application/json                             |
| apierrors.EnvelopeGetManyForbiddenError      | 403                                          | application/json                             |
| apierrors.EnvelopeGetManyInternalServerError | 500                                          | application/json                             |
| apierrors.APIError                           | 4XX, 5XX                                     | \*/\*                                        |

## EnvelopeCancel

Cancel envelope

### Example Usage

<!-- UsageSnippet language="go" operationID="envelope-cancel" method="post" path="/envelope/cancel" -->
```go
package main

import(
	"context"
	"os"
	sdkgo "github.com/documenso/sdk-go"
	"github.com/documenso/sdk-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdkgo.New(
        sdkgo.WithSecurity(os.Getenv("DOCUMENSO_API_KEY")),
    )

    res, err := s.Envelope.EnvelopeCancel(ctx, operations.EnvelopeCancelRequest{
        EnvelopeID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.EnvelopeCancelRequest](../../models/operations/envelopecancelrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.EnvelopeCancelResponse](../../models/operations/envelopecancelresponse.md), error**

### Errors

| Error Type                                  | Status Code                                 | Content Type                                |
| ------------------------------------------- | ------------------------------------------- | ------------------------------------------- |
| apierrors.EnvelopeCancelBadRequestError     | 400                                         | application/json                            |
| apierrors.EnvelopeCancelUnauthorizedError   | 401                                         | application/json                            |
| apierrors.EnvelopeCancelForbiddenError      | 403                                         | application/json                            |
| apierrors.EnvelopeCancelInternalServerError | 500                                         | application/json                            |
| apierrors.APIError                          | 4XX, 5XX                                    | \*/\*                                       |