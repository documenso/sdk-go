# Envelopes.Fields

## Overview

### Available Operations

* [Get](#get) - Get envelope field
* [CreateMany](#createmany) - Create envelope fields
* [UpdateMany](#updatemany) - Update envelope fields
* [Delete](#delete) - Delete envelope field

## Get

Returns an envelope field given an ID

### Example Usage

<!-- UsageSnippet language="go" operationID="envelope-field-get" method="get" path="/envelope/field/{fieldId}" -->
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

    res, err := s.Envelopes.Fields.Get(ctx, 6981.76)
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
| `fieldID`                                                | *float64*                                                | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.EnvelopeFieldGetResponse](../../models/operations/envelopefieldgetresponse.md), error**

### Errors

| Error Type                                    | Status Code                                   | Content Type                                  |
| --------------------------------------------- | --------------------------------------------- | --------------------------------------------- |
| apierrors.EnvelopeFieldGetBadRequestError     | 400                                           | application/json                              |
| apierrors.EnvelopeFieldGetUnauthorizedError   | 401                                           | application/json                              |
| apierrors.EnvelopeFieldGetForbiddenError      | 403                                           | application/json                              |
| apierrors.EnvelopeFieldGetNotFoundError       | 404                                           | application/json                              |
| apierrors.EnvelopeFieldGetInternalServerError | 500                                           | application/json                              |
| apierrors.APIError                            | 4XX, 5XX                                      | \*/\*                                         |

## CreateMany

Create multiple fields for an envelope

### Example Usage

<!-- UsageSnippet language="go" operationID="envelope-field-createMany" method="post" path="/envelope/field/create-many" -->
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

    res, err := s.Envelopes.Fields.CreateMany(ctx, operations.EnvelopeFieldCreateManyRequest{
        EnvelopeID: "<id>",
        Data: []operations.EnvelopeFieldCreateManyDataUnion{},
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

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.EnvelopeFieldCreateManyRequest](../../models/operations/envelopefieldcreatemanyrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../models/operations/option.md)                                               | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.EnvelopeFieldCreateManyResponse](../../models/operations/envelopefieldcreatemanyresponse.md), error**

### Errors

| Error Type                                           | Status Code                                          | Content Type                                         |
| ---------------------------------------------------- | ---------------------------------------------------- | ---------------------------------------------------- |
| apierrors.EnvelopeFieldCreateManyBadRequestError     | 400                                                  | application/json                                     |
| apierrors.EnvelopeFieldCreateManyUnauthorizedError   | 401                                                  | application/json                                     |
| apierrors.EnvelopeFieldCreateManyForbiddenError      | 403                                                  | application/json                                     |
| apierrors.EnvelopeFieldCreateManyInternalServerError | 500                                                  | application/json                                     |
| apierrors.APIError                                   | 4XX, 5XX                                             | \*/\*                                                |

## UpdateMany

Update multiple envelope fields for an envelope

### Example Usage

<!-- UsageSnippet language="go" operationID="envelope-field-updateMany" method="post" path="/envelope/field/update-many" -->
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

    res, err := s.Envelopes.Fields.UpdateMany(ctx, operations.EnvelopeFieldUpdateManyRequest{
        EnvelopeID: "<id>",
        Data: []operations.EnvelopeFieldUpdateManyDataUnion{},
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

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.EnvelopeFieldUpdateManyRequest](../../models/operations/envelopefieldupdatemanyrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../models/operations/option.md)                                               | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.EnvelopeFieldUpdateManyResponse](../../models/operations/envelopefieldupdatemanyresponse.md), error**

### Errors

| Error Type                                           | Status Code                                          | Content Type                                         |
| ---------------------------------------------------- | ---------------------------------------------------- | ---------------------------------------------------- |
| apierrors.EnvelopeFieldUpdateManyBadRequestError     | 400                                                  | application/json                                     |
| apierrors.EnvelopeFieldUpdateManyUnauthorizedError   | 401                                                  | application/json                                     |
| apierrors.EnvelopeFieldUpdateManyForbiddenError      | 403                                                  | application/json                                     |
| apierrors.EnvelopeFieldUpdateManyInternalServerError | 500                                                  | application/json                                     |
| apierrors.APIError                                   | 4XX, 5XX                                             | \*/\*                                                |

## Delete

Delete an envelope field

### Example Usage

<!-- UsageSnippet language="go" operationID="envelope-field-delete" method="post" path="/envelope/field/delete" -->
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

    res, err := s.Envelopes.Fields.Delete(ctx, operations.EnvelopeFieldDeleteRequest{
        FieldID: 2481.37,
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

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.EnvelopeFieldDeleteRequest](../../models/operations/envelopefielddeleterequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.EnvelopeFieldDeleteResponse](../../models/operations/envelopefielddeleteresponse.md), error**

### Errors

| Error Type                                       | Status Code                                      | Content Type                                     |
| ------------------------------------------------ | ------------------------------------------------ | ------------------------------------------------ |
| apierrors.EnvelopeFieldDeleteBadRequestError     | 400                                              | application/json                                 |
| apierrors.EnvelopeFieldDeleteUnauthorizedError   | 401                                              | application/json                                 |
| apierrors.EnvelopeFieldDeleteForbiddenError      | 403                                              | application/json                                 |
| apierrors.EnvelopeFieldDeleteInternalServerError | 500                                              | application/json                                 |
| apierrors.APIError                               | 4XX, 5XX                                         | \*/\*                                            |