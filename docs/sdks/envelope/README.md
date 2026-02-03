# Envelope

## Overview

### Available Operations

* [EnvelopeFind](#envelopefind) - Find envelopes
* [EnvelopeAuditLogFind](#envelopeauditlogfind) - Get envelope audit logs
* [EnvelopeGetMany](#envelopegetmany) - Get multiple envelopes

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