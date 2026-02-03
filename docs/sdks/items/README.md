# Envelopes.Items

## Overview

### Available Operations

* [CreateMany](#createmany) - Create envelope items
* [UpdateMany](#updatemany) - Update envelope items
* [Delete](#delete) - Delete envelope item
* [Download](#download) - Download an envelope item

## CreateMany

Create multiple envelope items for an envelope

### Example Usage

<!-- UsageSnippet language="go" operationID="envelope-item-createMany" method="post" path="/envelope/item/create-many" -->
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

    res, err := s.Envelopes.Items.CreateMany(ctx, operations.EnvelopeItemCreateManyRequest{
        Payload: operations.EnvelopeItemCreateManyPayload{
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

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.EnvelopeItemCreateManyRequest](../../models/operations/envelopeitemcreatemanyrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.EnvelopeItemCreateManyResponse](../../models/operations/envelopeitemcreatemanyresponse.md), error**

### Errors

| Error Type                                          | Status Code                                         | Content Type                                        |
| --------------------------------------------------- | --------------------------------------------------- | --------------------------------------------------- |
| apierrors.EnvelopeItemCreateManyBadRequestError     | 400                                                 | application/json                                    |
| apierrors.EnvelopeItemCreateManyUnauthorizedError   | 401                                                 | application/json                                    |
| apierrors.EnvelopeItemCreateManyForbiddenError      | 403                                                 | application/json                                    |
| apierrors.EnvelopeItemCreateManyInternalServerError | 500                                                 | application/json                                    |
| apierrors.APIError                                  | 4XX, 5XX                                            | \*/\*                                               |

## UpdateMany

Update multiple envelope items for an envelope

### Example Usage

<!-- UsageSnippet language="go" operationID="envelope-item-updateMany" method="post" path="/envelope/item/update-many" -->
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

    res, err := s.Envelopes.Items.UpdateMany(ctx, operations.EnvelopeItemUpdateManyRequest{
        EnvelopeID: "<id>",
        Data: []operations.EnvelopeItemUpdateManyDataRequest{},
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

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.EnvelopeItemUpdateManyRequest](../../models/operations/envelopeitemupdatemanyrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.EnvelopeItemUpdateManyResponse](../../models/operations/envelopeitemupdatemanyresponse.md), error**

### Errors

| Error Type                                          | Status Code                                         | Content Type                                        |
| --------------------------------------------------- | --------------------------------------------------- | --------------------------------------------------- |
| apierrors.EnvelopeItemUpdateManyBadRequestError     | 400                                                 | application/json                                    |
| apierrors.EnvelopeItemUpdateManyUnauthorizedError   | 401                                                 | application/json                                    |
| apierrors.EnvelopeItemUpdateManyForbiddenError      | 403                                                 | application/json                                    |
| apierrors.EnvelopeItemUpdateManyInternalServerError | 500                                                 | application/json                                    |
| apierrors.APIError                                  | 4XX, 5XX                                            | \*/\*                                               |

## Delete

Delete an envelope item from an envelope

### Example Usage

<!-- UsageSnippet language="go" operationID="envelope-item-delete" method="post" path="/envelope/item/delete" -->
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

    res, err := s.Envelopes.Items.Delete(ctx, operations.EnvelopeItemDeleteRequest{
        EnvelopeID: "<id>",
        EnvelopeItemID: "<id>",
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
| `request`                                                                                    | [operations.EnvelopeItemDeleteRequest](../../models/operations/envelopeitemdeleterequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.EnvelopeItemDeleteResponse](../../models/operations/envelopeitemdeleteresponse.md), error**

### Errors

| Error Type                                      | Status Code                                     | Content Type                                    |
| ----------------------------------------------- | ----------------------------------------------- | ----------------------------------------------- |
| apierrors.EnvelopeItemDeleteBadRequestError     | 400                                             | application/json                                |
| apierrors.EnvelopeItemDeleteUnauthorizedError   | 401                                             | application/json                                |
| apierrors.EnvelopeItemDeleteForbiddenError      | 403                                             | application/json                                |
| apierrors.EnvelopeItemDeleteInternalServerError | 500                                             | application/json                                |
| apierrors.APIError                              | 4XX, 5XX                                        | \*/\*                                           |

## Download

Download an envelope item by its ID

### Example Usage

<!-- UsageSnippet language="go" operationID="envelope-item-download" method="get" path="/envelope/item/{envelopeItemId}/download" -->
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

    res, err := s.Envelopes.Items.Download(ctx, "<id>", operations.EnvelopeItemDownloadVersionSigned.ToPointer())
    if err != nil {
        log.Fatal(err)
    }
    if res.Any != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                 | Type                                                                                                                                                      | Required                                                                                                                                                  | Description                                                                                                                                               |
| --------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                     | [context.Context](https://pkg.go.dev/context#Context)                                                                                                     | :heavy_check_mark:                                                                                                                                        | The context to use for the request.                                                                                                                       |
| `envelopeItemID`                                                                                                                                          | *string*                                                                                                                                                  | :heavy_check_mark:                                                                                                                                        | The ID of the envelope item to download.                                                                                                                  |
| `version`                                                                                                                                                 | [*operations.EnvelopeItemDownloadVersion](../../models/operations/envelopeitemdownloadversion.md)                                                         | :heavy_minus_sign:                                                                                                                                        | The version of the envelope item to download. "signed" returns the completed document with signatures, "original" returns the original uploaded document. |
| `opts`                                                                                                                                                    | [][operations.Option](../../models/operations/option.md)                                                                                                  | :heavy_minus_sign:                                                                                                                                        | The options for this request.                                                                                                                             |

### Response

**[*operations.EnvelopeItemDownloadResponse](../../models/operations/envelopeitemdownloadresponse.md), error**

### Errors

| Error Type                                        | Status Code                                       | Content Type                                      |
| ------------------------------------------------- | ------------------------------------------------- | ------------------------------------------------- |
| apierrors.EnvelopeItemDownloadBadRequestError     | 400                                               | application/json                                  |
| apierrors.EnvelopeItemDownloadUnauthorizedError   | 401                                               | application/json                                  |
| apierrors.EnvelopeItemDownloadForbiddenError      | 403                                               | application/json                                  |
| apierrors.EnvelopeItemDownloadNotFoundError       | 404                                               | application/json                                  |
| apierrors.EnvelopeItemDownloadInternalServerError | 500                                               | application/json                                  |
| apierrors.APIError                                | 4XX, 5XX                                          | \*/\*                                             |