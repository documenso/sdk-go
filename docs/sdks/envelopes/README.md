# Envelopes
(*Envelopes*)

## Overview

### Available Operations

* [Get](#get) - Get envelope
* [Create](#create) - Create envelope
* [Use](#use) - Use envelope
* [Update](#update) - Update envelope
* [Delete](#delete) - Delete envelope
* [Duplicate](#duplicate) - Duplicate envelope
* [Distribute](#distribute) - Distribute envelope
* [Redistribute](#redistribute) - Redistribute envelope

## Get

Returns an envelope given an ID

### Example Usage

<!-- UsageSnippet language="go" operationID="envelope-get" method="get" path="/envelope/{envelopeId}" -->
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

    res, err := s.Envelopes.Get(ctx, "<id>")
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
| `envelopeID`                                             | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.EnvelopeGetResponse](../../models/operations/envelopegetresponse.md), error**

### Errors

| Error Type                               | Status Code                              | Content Type                             |
| ---------------------------------------- | ---------------------------------------- | ---------------------------------------- |
| apierrors.EnvelopeGetBadRequestError     | 400                                      | application/json                         |
| apierrors.EnvelopeGetUnauthorizedError   | 401                                      | application/json                         |
| apierrors.EnvelopeGetForbiddenError      | 403                                      | application/json                         |
| apierrors.EnvelopeGetNotFoundError       | 404                                      | application/json                         |
| apierrors.EnvelopeGetInternalServerError | 500                                      | application/json                         |
| apierrors.APIError                       | 4XX, 5XX                                 | \*/\*                                    |

## Create

Create an envelope using form data.

### Example Usage

<!-- UsageSnippet language="go" operationID="envelope-create" method="post" path="/envelope/create" -->
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

    res, err := s.Envelopes.Create(ctx, operations.EnvelopeCreateRequest{
        Payload: operations.EnvelopeCreatePayload{
            Title: "<value>",
            Type: operations.EnvelopeCreateTypeTemplate,
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
| `request`                                                                            | [operations.EnvelopeCreateRequest](../../models/operations/envelopecreaterequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.EnvelopeCreateResponse](../../models/operations/envelopecreateresponse.md), error**

### Errors

| Error Type                                  | Status Code                                 | Content Type                                |
| ------------------------------------------- | ------------------------------------------- | ------------------------------------------- |
| apierrors.EnvelopeCreateBadRequestError     | 400                                         | application/json                            |
| apierrors.EnvelopeCreateUnauthorizedError   | 401                                         | application/json                            |
| apierrors.EnvelopeCreateForbiddenError      | 403                                         | application/json                            |
| apierrors.EnvelopeCreateInternalServerError | 500                                         | application/json                            |
| apierrors.APIError                          | 4XX, 5XX                                    | \*/\*                                       |

## Use

Create a document envelope from a template envelope. Upload custom files to replace the template PDFs and map them to specific envelope items using identifiers.

### Example Usage

<!-- UsageSnippet language="go" operationID="envelope-use" method="post" path="/envelope/use" -->
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

    res, err := s.Envelopes.Use(ctx, operations.EnvelopeUseRequest{
        Payload: operations.EnvelopeUsePayload{
            EnvelopeID: "<id>",
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

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.EnvelopeUseRequest](../../models/operations/envelopeuserequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.EnvelopeUseResponse](../../models/operations/envelopeuseresponse.md), error**

### Errors

| Error Type                               | Status Code                              | Content Type                             |
| ---------------------------------------- | ---------------------------------------- | ---------------------------------------- |
| apierrors.EnvelopeUseBadRequestError     | 400                                      | application/json                         |
| apierrors.EnvelopeUseUnauthorizedError   | 401                                      | application/json                         |
| apierrors.EnvelopeUseForbiddenError      | 403                                      | application/json                         |
| apierrors.EnvelopeUseInternalServerError | 500                                      | application/json                         |
| apierrors.APIError                       | 4XX, 5XX                                 | \*/\*                                    |

## Update

Update envelope

### Example Usage

<!-- UsageSnippet language="go" operationID="envelope-update" method="post" path="/envelope/update" -->
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

    res, err := s.Envelopes.Update(ctx, operations.EnvelopeUpdateRequest{
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
| `request`                                                                            | [operations.EnvelopeUpdateRequest](../../models/operations/envelopeupdaterequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.EnvelopeUpdateResponse](../../models/operations/envelopeupdateresponse.md), error**

### Errors

| Error Type                                  | Status Code                                 | Content Type                                |
| ------------------------------------------- | ------------------------------------------- | ------------------------------------------- |
| apierrors.EnvelopeUpdateBadRequestError     | 400                                         | application/json                            |
| apierrors.EnvelopeUpdateUnauthorizedError   | 401                                         | application/json                            |
| apierrors.EnvelopeUpdateForbiddenError      | 403                                         | application/json                            |
| apierrors.EnvelopeUpdateInternalServerError | 500                                         | application/json                            |
| apierrors.APIError                          | 4XX, 5XX                                    | \*/\*                                       |

## Delete

Delete envelope

### Example Usage

<!-- UsageSnippet language="go" operationID="envelope-delete" method="post" path="/envelope/delete" -->
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

    res, err := s.Envelopes.Delete(ctx, operations.EnvelopeDeleteRequest{
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
| `request`                                                                            | [operations.EnvelopeDeleteRequest](../../models/operations/envelopedeleterequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.EnvelopeDeleteResponse](../../models/operations/envelopedeleteresponse.md), error**

### Errors

| Error Type                                  | Status Code                                 | Content Type                                |
| ------------------------------------------- | ------------------------------------------- | ------------------------------------------- |
| apierrors.EnvelopeDeleteBadRequestError     | 400                                         | application/json                            |
| apierrors.EnvelopeDeleteUnauthorizedError   | 401                                         | application/json                            |
| apierrors.EnvelopeDeleteForbiddenError      | 403                                         | application/json                            |
| apierrors.EnvelopeDeleteInternalServerError | 500                                         | application/json                            |
| apierrors.APIError                          | 4XX, 5XX                                    | \*/\*                                       |

## Duplicate

Duplicate an envelope with all its settings

### Example Usage

<!-- UsageSnippet language="go" operationID="envelope-duplicate" method="post" path="/envelope/duplicate" -->
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

    res, err := s.Envelopes.Duplicate(ctx, operations.EnvelopeDuplicateRequest{
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

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.EnvelopeDuplicateRequest](../../models/operations/envelopeduplicaterequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.EnvelopeDuplicateResponse](../../models/operations/envelopeduplicateresponse.md), error**

### Errors

| Error Type                                     | Status Code                                    | Content Type                                   |
| ---------------------------------------------- | ---------------------------------------------- | ---------------------------------------------- |
| apierrors.EnvelopeDuplicateBadRequestError     | 400                                            | application/json                               |
| apierrors.EnvelopeDuplicateUnauthorizedError   | 401                                            | application/json                               |
| apierrors.EnvelopeDuplicateForbiddenError      | 403                                            | application/json                               |
| apierrors.EnvelopeDuplicateInternalServerError | 500                                            | application/json                               |
| apierrors.APIError                             | 4XX, 5XX                                       | \*/\*                                          |

## Distribute

Send the envelope to recipients based on your distribution method

### Example Usage

<!-- UsageSnippet language="go" operationID="envelope-distribute" method="post" path="/envelope/distribute" -->
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

    res, err := s.Envelopes.Distribute(ctx, operations.EnvelopeDistributeRequest{
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

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.EnvelopeDistributeRequest](../../models/operations/envelopedistributerequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.EnvelopeDistributeResponse](../../models/operations/envelopedistributeresponse.md), error**

### Errors

| Error Type                                      | Status Code                                     | Content Type                                    |
| ----------------------------------------------- | ----------------------------------------------- | ----------------------------------------------- |
| apierrors.EnvelopeDistributeBadRequestError     | 400                                             | application/json                                |
| apierrors.EnvelopeDistributeUnauthorizedError   | 401                                             | application/json                                |
| apierrors.EnvelopeDistributeForbiddenError      | 403                                             | application/json                                |
| apierrors.EnvelopeDistributeInternalServerError | 500                                             | application/json                                |
| apierrors.APIError                              | 4XX, 5XX                                        | \*/\*                                           |

## Redistribute

Redistribute the envelope to the provided recipients who have not actioned the envelope. Will use the distribution method set in the envelope

### Example Usage

<!-- UsageSnippet language="go" operationID="envelope-redistribute" method="post" path="/envelope/redistribute" -->
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

    res, err := s.Envelopes.Redistribute(ctx, operations.EnvelopeRedistributeRequest{
        EnvelopeID: "<id>",
        Recipients: []float64{},
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
| `request`                                                                                        | [operations.EnvelopeRedistributeRequest](../../models/operations/enveloperedistributerequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.EnvelopeRedistributeResponse](../../models/operations/enveloperedistributeresponse.md), error**

### Errors

| Error Type                                        | Status Code                                       | Content Type                                      |
| ------------------------------------------------- | ------------------------------------------------- | ------------------------------------------------- |
| apierrors.EnvelopeRedistributeBadRequestError     | 400                                               | application/json                                  |
| apierrors.EnvelopeRedistributeUnauthorizedError   | 401                                               | application/json                                  |
| apierrors.EnvelopeRedistributeForbiddenError      | 403                                               | application/json                                  |
| apierrors.EnvelopeRedistributeInternalServerError | 500                                               | application/json                                  |
| apierrors.APIError                                | 4XX, 5XX                                          | \*/\*                                             |