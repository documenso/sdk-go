# EnvelopesAttachments
(*Envelopes.Attachments*)

## Overview

### Available Operations

* [Find](#find) - Find attachments
* [Create](#create) - Create attachment
* [Update](#update) - Update attachment
* [Delete](#delete) - Delete attachment

## Find

Find all attachments for an envelope

### Example Usage

<!-- UsageSnippet language="go" operationID="envelope-attachment-find" method="get" path="/envelope/attachment" -->
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

    res, err := s.Envelopes.Attachments.Find(ctx, "<id>", nil)
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
| `token`                                                  | **string*                                                | :heavy_minus_sign:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.EnvelopeAttachmentFindResponse](../../models/operations/envelopeattachmentfindresponse.md), error**

### Errors

| Error Type                                          | Status Code                                         | Content Type                                        |
| --------------------------------------------------- | --------------------------------------------------- | --------------------------------------------------- |
| apierrors.EnvelopeAttachmentFindBadRequestError     | 400                                                 | application/json                                    |
| apierrors.EnvelopeAttachmentFindUnauthorizedError   | 401                                                 | application/json                                    |
| apierrors.EnvelopeAttachmentFindForbiddenError      | 403                                                 | application/json                                    |
| apierrors.EnvelopeAttachmentFindNotFoundError       | 404                                                 | application/json                                    |
| apierrors.EnvelopeAttachmentFindInternalServerError | 500                                                 | application/json                                    |
| apierrors.APIError                                  | 4XX, 5XX                                            | \*/\*                                               |

## Create

Create a new attachment for an envelope

### Example Usage

<!-- UsageSnippet language="go" operationID="envelope-attachment-create" method="post" path="/envelope/attachment/create" -->
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

    res, err := s.Envelopes.Attachments.Create(ctx, operations.EnvelopeAttachmentCreateRequest{
        EnvelopeID: "<id>",
        Data: operations.EnvelopeAttachmentCreateData{
            Label: "<value>",
            Data: "https://lustrous-skeleton.info",
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

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.EnvelopeAttachmentCreateRequest](../../models/operations/envelopeattachmentcreaterequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.EnvelopeAttachmentCreateResponse](../../models/operations/envelopeattachmentcreateresponse.md), error**

### Errors

| Error Type                                            | Status Code                                           | Content Type                                          |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| apierrors.EnvelopeAttachmentCreateBadRequestError     | 400                                                   | application/json                                      |
| apierrors.EnvelopeAttachmentCreateUnauthorizedError   | 401                                                   | application/json                                      |
| apierrors.EnvelopeAttachmentCreateForbiddenError      | 403                                                   | application/json                                      |
| apierrors.EnvelopeAttachmentCreateInternalServerError | 500                                                   | application/json                                      |
| apierrors.APIError                                    | 4XX, 5XX                                              | \*/\*                                                 |

## Update

Update an existing attachment

### Example Usage

<!-- UsageSnippet language="go" operationID="envelope-attachment-update" method="post" path="/envelope/attachment/update" -->
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

    res, err := s.Envelopes.Attachments.Update(ctx, operations.EnvelopeAttachmentUpdateRequest{
        ID: "<id>",
        Data: operations.EnvelopeAttachmentUpdateData{
            Label: "<value>",
            Data: "https://tough-premier.biz",
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

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.EnvelopeAttachmentUpdateRequest](../../models/operations/envelopeattachmentupdaterequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.EnvelopeAttachmentUpdateResponse](../../models/operations/envelopeattachmentupdateresponse.md), error**

### Errors

| Error Type                                            | Status Code                                           | Content Type                                          |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| apierrors.EnvelopeAttachmentUpdateBadRequestError     | 400                                                   | application/json                                      |
| apierrors.EnvelopeAttachmentUpdateUnauthorizedError   | 401                                                   | application/json                                      |
| apierrors.EnvelopeAttachmentUpdateForbiddenError      | 403                                                   | application/json                                      |
| apierrors.EnvelopeAttachmentUpdateInternalServerError | 500                                                   | application/json                                      |
| apierrors.APIError                                    | 4XX, 5XX                                              | \*/\*                                                 |

## Delete

Delete an attachment from an envelope

### Example Usage

<!-- UsageSnippet language="go" operationID="envelope-attachment-delete" method="post" path="/envelope/attachment/delete" -->
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

    res, err := s.Envelopes.Attachments.Delete(ctx, operations.EnvelopeAttachmentDeleteRequest{
        ID: "<id>",
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

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.EnvelopeAttachmentDeleteRequest](../../models/operations/envelopeattachmentdeleterequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.EnvelopeAttachmentDeleteResponse](../../models/operations/envelopeattachmentdeleteresponse.md), error**

### Errors

| Error Type                                            | Status Code                                           | Content Type                                          |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| apierrors.EnvelopeAttachmentDeleteBadRequestError     | 400                                                   | application/json                                      |
| apierrors.EnvelopeAttachmentDeleteUnauthorizedError   | 401                                                   | application/json                                      |
| apierrors.EnvelopeAttachmentDeleteForbiddenError      | 403                                                   | application/json                                      |
| apierrors.EnvelopeAttachmentDeleteInternalServerError | 500                                                   | application/json                                      |
| apierrors.APIError                                    | 4XX, 5XX                                              | \*/\*                                                 |