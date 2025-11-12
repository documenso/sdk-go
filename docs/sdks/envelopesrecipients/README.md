# EnvelopesRecipients
(*Envelopes.Recipients*)

## Overview

### Available Operations

* [Get](#get) - Get envelope recipient
* [CreateMany](#createmany) - Create envelope recipients
* [UpdateMany](#updatemany) - Update envelope recipients
* [Delete](#delete) - Delete envelope recipient

## Get

Returns an envelope recipient given an ID

### Example Usage

<!-- UsageSnippet language="go" operationID="envelope-recipient-get" method="get" path="/envelope/recipient/{recipientId}" -->
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

    res, err := s.Envelopes.Recipients.Get(ctx, 8771.72)
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
| `recipientID`                                            | *float64*                                                | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.EnvelopeRecipientGetResponse](../../models/operations/enveloperecipientgetresponse.md), error**

### Errors

| Error Type                                        | Status Code                                       | Content Type                                      |
| ------------------------------------------------- | ------------------------------------------------- | ------------------------------------------------- |
| apierrors.EnvelopeRecipientGetBadRequestError     | 400                                               | application/json                                  |
| apierrors.EnvelopeRecipientGetUnauthorizedError   | 401                                               | application/json                                  |
| apierrors.EnvelopeRecipientGetForbiddenError      | 403                                               | application/json                                  |
| apierrors.EnvelopeRecipientGetNotFoundError       | 404                                               | application/json                                  |
| apierrors.EnvelopeRecipientGetInternalServerError | 500                                               | application/json                                  |
| apierrors.APIError                                | 4XX, 5XX                                          | \*/\*                                             |

## CreateMany

Create multiple recipients for an envelope

### Example Usage

<!-- UsageSnippet language="go" operationID="envelope-recipient-createMany" method="post" path="/envelope/recipient/create-many" -->
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

    res, err := s.Envelopes.Recipients.CreateMany(ctx, operations.EnvelopeRecipientCreateManyRequest{
        EnvelopeID: "<id>",
        Data: []operations.EnvelopeRecipientCreateManyDataRequest{
            operations.EnvelopeRecipientCreateManyDataRequest{
                Email: "Ed16@yahoo.com",
                Name: "<value>",
                Role: operations.EnvelopeRecipientCreateManyRoleRequestSigner,
            },
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

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.EnvelopeRecipientCreateManyRequest](../../models/operations/enveloperecipientcreatemanyrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `opts`                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.EnvelopeRecipientCreateManyResponse](../../models/operations/enveloperecipientcreatemanyresponse.md), error**

### Errors

| Error Type                                               | Status Code                                              | Content Type                                             |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| apierrors.EnvelopeRecipientCreateManyBadRequestError     | 400                                                      | application/json                                         |
| apierrors.EnvelopeRecipientCreateManyUnauthorizedError   | 401                                                      | application/json                                         |
| apierrors.EnvelopeRecipientCreateManyForbiddenError      | 403                                                      | application/json                                         |
| apierrors.EnvelopeRecipientCreateManyInternalServerError | 500                                                      | application/json                                         |
| apierrors.APIError                                       | 4XX, 5XX                                                 | \*/\*                                                    |

## UpdateMany

Update multiple recipients for an envelope

### Example Usage

<!-- UsageSnippet language="go" operationID="envelope-recipient-updateMany" method="post" path="/envelope/recipient/update-many" -->
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

    res, err := s.Envelopes.Recipients.UpdateMany(ctx, operations.EnvelopeRecipientUpdateManyRequest{
        EnvelopeID: "<id>",
        Data: []operations.EnvelopeRecipientUpdateManyDataRequest{
            operations.EnvelopeRecipientUpdateManyDataRequest{
                ID: 8894.57,
            },
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

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.EnvelopeRecipientUpdateManyRequest](../../models/operations/enveloperecipientupdatemanyrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `opts`                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.EnvelopeRecipientUpdateManyResponse](../../models/operations/enveloperecipientupdatemanyresponse.md), error**

### Errors

| Error Type                                               | Status Code                                              | Content Type                                             |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| apierrors.EnvelopeRecipientUpdateManyBadRequestError     | 400                                                      | application/json                                         |
| apierrors.EnvelopeRecipientUpdateManyUnauthorizedError   | 401                                                      | application/json                                         |
| apierrors.EnvelopeRecipientUpdateManyForbiddenError      | 403                                                      | application/json                                         |
| apierrors.EnvelopeRecipientUpdateManyInternalServerError | 500                                                      | application/json                                         |
| apierrors.APIError                                       | 4XX, 5XX                                                 | \*/\*                                                    |

## Delete

Delete an envelope recipient

### Example Usage

<!-- UsageSnippet language="go" operationID="envelope-recipient-delete" method="post" path="/envelope/recipient/delete" -->
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

    res, err := s.Envelopes.Recipients.Delete(ctx, operations.EnvelopeRecipientDeleteRequest{
        RecipientID: 4834.93,
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
| `request`                                                                                              | [operations.EnvelopeRecipientDeleteRequest](../../models/operations/enveloperecipientdeleterequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../models/operations/option.md)                                               | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.EnvelopeRecipientDeleteResponse](../../models/operations/enveloperecipientdeleteresponse.md), error**

### Errors

| Error Type                                           | Status Code                                          | Content Type                                         |
| ---------------------------------------------------- | ---------------------------------------------------- | ---------------------------------------------------- |
| apierrors.EnvelopeRecipientDeleteBadRequestError     | 400                                                  | application/json                                     |
| apierrors.EnvelopeRecipientDeleteUnauthorizedError   | 401                                                  | application/json                                     |
| apierrors.EnvelopeRecipientDeleteForbiddenError      | 403                                                  | application/json                                     |
| apierrors.EnvelopeRecipientDeleteInternalServerError | 500                                                  | application/json                                     |
| apierrors.APIError                                   | 4XX, 5XX                                             | \*/\*                                                |