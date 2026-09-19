# ~~Documents~~

> [!WARNING]
> This SDK is **DEPRECATED**

## Overview

### Available Operations

* [~~Get~~](#get) - Get document :warning: **Deprecated**
* [~~Find~~](#find) - Find documents :warning: **Deprecated**
* [~~Create~~](#create) - Create document :warning: **Deprecated**
* [~~Update~~](#update) - Update document :warning: **Deprecated**
* [~~Delete~~](#delete) - Delete document :warning: **Deprecated**
* [~~Duplicate~~](#duplicate) - Duplicate document :warning: **Deprecated**
* [~~Distribute~~](#distribute) - Distribute document :warning: **Deprecated**
* [~~Redistribute~~](#redistribute) - Redistribute document :warning: **Deprecated**
* [~~Download~~](#download) - Download document :warning: **Deprecated**
* [~~CreateV0~~](#createv0) - Create document :warning: **Deprecated**

## ~~Get~~

Deprecated: this endpoint is being replaced by the Envelope API. See https://docs.documenso.com/docs/developers/api/migrate-to-envelopes for the migration guide. Returns a document given an ID

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage

<!-- UsageSnippet language="go" operationID="document-get" method="get" path="/document/{documentId}" -->
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

    res, err := s.Documents.Get(ctx, 6150.61)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `documentID`                                             | `float64`                                                | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DocumentGetResponse](../../models/operations/documentgetresponse.md), error**

### Errors

| Error Type                               | Status Code                              | Content Type                             |
| ---------------------------------------- | ---------------------------------------- | ---------------------------------------- |
| apierrors.DocumentGetBadRequestError     | 400                                      | application/json                         |
| apierrors.DocumentGetUnauthorizedError   | 401                                      | application/json                         |
| apierrors.DocumentGetForbiddenError      | 403                                      | application/json                         |
| apierrors.DocumentGetNotFoundError       | 404                                      | application/json                         |
| apierrors.DocumentGetInternalServerError | 500                                      | application/json                         |
| apierrors.APIError                       | 4XX, 5XX                                 | \*/\*                                    |

## ~~Find~~

Deprecated: this endpoint is being replaced by the Envelope API. See https://docs.documenso.com/docs/developers/api/migrate-to-envelopes for the migration guide. Find documents based on a search criteria

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage

<!-- UsageSnippet language="go" operationID="document-find" method="get" path="/document" -->
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

    res, err := s.Documents.Find(ctx, operations.DocumentFindRequest{})
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
| `request`                                                                        | [operations.DocumentFindRequest](../../models/operations/documentfindrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.DocumentFindResponse](../../models/operations/documentfindresponse.md), error**

### Errors

| Error Type                                | Status Code                               | Content Type                              |
| ----------------------------------------- | ----------------------------------------- | ----------------------------------------- |
| apierrors.DocumentFindBadRequestError     | 400                                       | application/json                          |
| apierrors.DocumentFindUnauthorizedError   | 401                                       | application/json                          |
| apierrors.DocumentFindForbiddenError      | 403                                       | application/json                          |
| apierrors.DocumentFindNotFoundError       | 404                                       | application/json                          |
| apierrors.DocumentFindInternalServerError | 500                                       | application/json                          |
| apierrors.APIError                        | 4XX, 5XX                                  | \*/\*                                     |

## ~~Create~~

Deprecated: this endpoint is being replaced by the Envelope API. See https://docs.documenso.com/docs/developers/api/migrate-to-envelopes for the migration guide.

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage

<!-- UsageSnippet language="go" operationID="document-create" method="post" path="/document/create" -->
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

    example, fileErr := os.Open("example.file")
    if fileErr != nil {
        panic(fileErr)
    }

    res, err := s.Documents.Create(ctx, operations.DocumentCreateRequest{
        Payload: operations.DocumentCreatePayload{
            Title: "<value>",
        },
        File: operations.DocumentCreateFile{
            FileName: "example.file",
            Content: example,
        },
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
| `request`                                                                            | [operations.DocumentCreateRequest](../../models/operations/documentcreaterequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.DocumentCreateResponse](../../models/operations/documentcreateresponse.md), error**

### Errors

| Error Type                                  | Status Code                                 | Content Type                                |
| ------------------------------------------- | ------------------------------------------- | ------------------------------------------- |
| apierrors.DocumentCreateBadRequestError     | 400                                         | application/json                            |
| apierrors.DocumentCreateUnauthorizedError   | 401                                         | application/json                            |
| apierrors.DocumentCreateForbiddenError      | 403                                         | application/json                            |
| apierrors.DocumentCreateInternalServerError | 500                                         | application/json                            |
| apierrors.APIError                          | 4XX, 5XX                                    | \*/\*                                       |

## ~~Update~~

Deprecated: this endpoint is being replaced by the Envelope API. See https://docs.documenso.com/docs/developers/api/migrate-to-envelopes for the migration guide.

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage

<!-- UsageSnippet language="go" operationID="document-update" method="post" path="/document/update" -->
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

    res, err := s.Documents.Update(ctx, operations.DocumentUpdateRequest{
        DocumentID: 3428.95,
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
| `request`                                                                            | [operations.DocumentUpdateRequest](../../models/operations/documentupdaterequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.DocumentUpdateResponse](../../models/operations/documentupdateresponse.md), error**

### Errors

| Error Type                                  | Status Code                                 | Content Type                                |
| ------------------------------------------- | ------------------------------------------- | ------------------------------------------- |
| apierrors.DocumentUpdateBadRequestError     | 400                                         | application/json                            |
| apierrors.DocumentUpdateUnauthorizedError   | 401                                         | application/json                            |
| apierrors.DocumentUpdateForbiddenError      | 403                                         | application/json                            |
| apierrors.DocumentUpdateInternalServerError | 500                                         | application/json                            |
| apierrors.APIError                          | 4XX, 5XX                                    | \*/\*                                       |

## ~~Delete~~

Deprecated: this endpoint is being replaced by the Envelope API. See https://docs.documenso.com/docs/developers/api/migrate-to-envelopes for the migration guide.

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage

<!-- UsageSnippet language="go" operationID="document-delete" method="post" path="/document/delete" -->
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

    res, err := s.Documents.Delete(ctx, operations.DocumentDeleteRequest{
        DocumentID: 3963.4,
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
| `request`                                                                            | [operations.DocumentDeleteRequest](../../models/operations/documentdeleterequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.DocumentDeleteResponse](../../models/operations/documentdeleteresponse.md), error**

### Errors

| Error Type                                  | Status Code                                 | Content Type                                |
| ------------------------------------------- | ------------------------------------------- | ------------------------------------------- |
| apierrors.DocumentDeleteBadRequestError     | 400                                         | application/json                            |
| apierrors.DocumentDeleteUnauthorizedError   | 401                                         | application/json                            |
| apierrors.DocumentDeleteForbiddenError      | 403                                         | application/json                            |
| apierrors.DocumentDeleteInternalServerError | 500                                         | application/json                            |
| apierrors.APIError                          | 4XX, 5XX                                    | \*/\*                                       |

## ~~Duplicate~~

Deprecated: this endpoint is being replaced by the Envelope API. See https://docs.documenso.com/docs/developers/api/migrate-to-envelopes for the migration guide.

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage

<!-- UsageSnippet language="go" operationID="document-duplicate" method="post" path="/document/duplicate" -->
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

    res, err := s.Documents.Duplicate(ctx, operations.DocumentDuplicateRequest{
        DocumentID: 5285.3,
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

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.DocumentDuplicateRequest](../../models/operations/documentduplicaterequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.DocumentDuplicateResponse](../../models/operations/documentduplicateresponse.md), error**

### Errors

| Error Type                                     | Status Code                                    | Content Type                                   |
| ---------------------------------------------- | ---------------------------------------------- | ---------------------------------------------- |
| apierrors.DocumentDuplicateBadRequestError     | 400                                            | application/json                               |
| apierrors.DocumentDuplicateUnauthorizedError   | 401                                            | application/json                               |
| apierrors.DocumentDuplicateForbiddenError      | 403                                            | application/json                               |
| apierrors.DocumentDuplicateInternalServerError | 500                                            | application/json                               |
| apierrors.APIError                             | 4XX, 5XX                                       | \*/\*                                          |

## ~~Distribute~~

Deprecated: this endpoint is being replaced by the Envelope API. See https://docs.documenso.com/docs/developers/api/migrate-to-envelopes for the migration guide. Send the document out to recipients based on your distribution method

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage

<!-- UsageSnippet language="go" operationID="document-distribute" method="post" path="/document/distribute" -->
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

    res, err := s.Documents.Distribute(ctx, operations.DocumentDistributeRequest{
        DocumentID: 7548.74,
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

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.DocumentDistributeRequest](../../models/operations/documentdistributerequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.DocumentDistributeResponse](../../models/operations/documentdistributeresponse.md), error**

### Errors

| Error Type                                      | Status Code                                     | Content Type                                    |
| ----------------------------------------------- | ----------------------------------------------- | ----------------------------------------------- |
| apierrors.DocumentDistributeBadRequestError     | 400                                             | application/json                                |
| apierrors.DocumentDistributeUnauthorizedError   | 401                                             | application/json                                |
| apierrors.DocumentDistributeForbiddenError      | 403                                             | application/json                                |
| apierrors.DocumentDistributeInternalServerError | 500                                             | application/json                                |
| apierrors.APIError                              | 4XX, 5XX                                        | \*/\*                                           |

## ~~Redistribute~~

Deprecated: this endpoint is being replaced by the Envelope API. See https://docs.documenso.com/docs/developers/api/migrate-to-envelopes for the migration guide. Redistribute the document to the provided recipients who have not actioned the document. Will use the distribution method set in the document. This also refreshes the signing-link expiration for the targeted unsigned recipients, renewing any expired links.

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage

<!-- UsageSnippet language="go" operationID="document-redistribute" method="post" path="/document/redistribute" -->
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

    res, err := s.Documents.Redistribute(ctx, operations.DocumentRedistributeRequest{
        DocumentID: 9084.69,
        Recipients: []float64{
            6011.8,
            4441.56,
            4251.15,
        },
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
| `request`                                                                                        | [operations.DocumentRedistributeRequest](../../models/operations/documentredistributerequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.DocumentRedistributeResponse](../../models/operations/documentredistributeresponse.md), error**

### Errors

| Error Type                                        | Status Code                                       | Content Type                                      |
| ------------------------------------------------- | ------------------------------------------------- | ------------------------------------------------- |
| apierrors.DocumentRedistributeBadRequestError     | 400                                               | application/json                                  |
| apierrors.DocumentRedistributeUnauthorizedError   | 401                                               | application/json                                  |
| apierrors.DocumentRedistributeForbiddenError      | 403                                               | application/json                                  |
| apierrors.DocumentRedistributeInternalServerError | 500                                               | application/json                                  |
| apierrors.APIError                                | 4XX, 5XX                                          | \*/\*                                             |

## ~~Download~~

Deprecated: this endpoint is being replaced by the Envelope API. See https://docs.documenso.com/docs/developers/api/migrate-to-envelopes for the migration guide.

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage

<!-- UsageSnippet language="go" operationID="document-download" method="get" path="/document/{documentId}/download" -->
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

    res, err := s.Documents.Download(ctx, 5396.97, operations.DocumentDownloadVersionSigned.ToPointer())
    if err != nil {
        log.Fatal(err)
    }
    if res.Any != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                            | Type                                                                                                                                                 | Required                                                                                                                                             | Description                                                                                                                                          |
| ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                                | :heavy_check_mark:                                                                                                                                   | The context to use for the request.                                                                                                                  |
| `documentID`                                                                                                                                         | `float64`                                                                                                                                            | :heavy_check_mark:                                                                                                                                   | The ID of the document to download.                                                                                                                  |
| `version`                                                                                                                                            | [*operations.DocumentDownloadVersion](../../models/operations/documentdownloadversion.md)                                                            | :heavy_minus_sign:                                                                                                                                   | The version of the document to download. "signed" returns the completed document with signatures, "original" returns the original uploaded document. |
| `opts`                                                                                                                                               | [][operations.Option](../../models/operations/option.md)                                                                                             | :heavy_minus_sign:                                                                                                                                   | The options for this request.                                                                                                                        |

### Response

**[*operations.DocumentDownloadResponse](../../models/operations/documentdownloadresponse.md), error**

### Errors

| Error Type                                    | Status Code                                   | Content Type                                  |
| --------------------------------------------- | --------------------------------------------- | --------------------------------------------- |
| apierrors.DocumentDownloadBadRequestError     | 400                                           | application/json                              |
| apierrors.DocumentDownloadUnauthorizedError   | 401                                           | application/json                              |
| apierrors.DocumentDownloadForbiddenError      | 403                                           | application/json                              |
| apierrors.DocumentDownloadNotFoundError       | 404                                           | application/json                              |
| apierrors.DocumentDownloadInternalServerError | 500                                           | application/json                              |
| apierrors.APIError                            | 4XX, 5XX                                      | \*/\*                                         |

## ~~CreateV0~~

Deprecated: this endpoint is being replaced by the Envelope API. See https://docs.documenso.com/docs/developers/api/migrate-to-envelopes for the migration guide. You will need to upload the PDF to the provided URL returned. This endpoint will be removed since we will allow direct uploads, instead of using an upload URL.

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage

<!-- UsageSnippet language="go" operationID="document-createDocumentTemporary" method="post" path="/document/create/beta" -->
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

    res, err := s.Documents.CreateV0(ctx, operations.DocumentCreateDocumentTemporaryRequest{
        Title: "<value>",
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

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                  | :heavy_check_mark:                                                                                                     | The context to use for the request.                                                                                    |
| `request`                                                                                                              | [operations.DocumentCreateDocumentTemporaryRequest](../../models/operations/documentcreatedocumenttemporaryrequest.md) | :heavy_check_mark:                                                                                                     | The request object to use for the request.                                                                             |
| `opts`                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                               | :heavy_minus_sign:                                                                                                     | The options for this request.                                                                                          |

### Response

**[*operations.DocumentCreateDocumentTemporaryResponse](../../models/operations/documentcreatedocumenttemporaryresponse.md), error**

### Errors

| Error Type                                                   | Status Code                                                  | Content Type                                                 |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| apierrors.DocumentCreateDocumentTemporaryBadRequestError     | 400                                                          | application/json                                             |
| apierrors.DocumentCreateDocumentTemporaryUnauthorizedError   | 401                                                          | application/json                                             |
| apierrors.DocumentCreateDocumentTemporaryForbiddenError      | 403                                                          | application/json                                             |
| apierrors.DocumentCreateDocumentTemporaryInternalServerError | 500                                                          | application/json                                             |
| apierrors.APIError                                           | 4XX, 5XX                                                     | \*/\*                                                        |