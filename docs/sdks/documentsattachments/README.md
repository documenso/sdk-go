# Documents.Attachments

## Overview

### Available Operations

* [Create](#create) - Create attachment
* [Update](#update) - Update attachment
* [Delete](#delete) - Delete attachment
* [Find](#find) - Find attachments

## Create

Create a new attachment for a document

### Example Usage

<!-- UsageSnippet language="go" operationID="document-attachment-create" method="post" path="/document/attachment/create" -->
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

    res, err := s.Documents.Attachments.Create(ctx, operations.DocumentAttachmentCreateRequest{
        DocumentID: 7014.36,
        Data: operations.DocumentAttachmentCreateData{
            Label: "<value>",
            Data: "https://cheerful-bourgeoisie.org/",
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
| `request`                                                                                                | [operations.DocumentAttachmentCreateRequest](../../models/operations/documentattachmentcreaterequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.DocumentAttachmentCreateResponse](../../models/operations/documentattachmentcreateresponse.md), error**

### Errors

| Error Type                                            | Status Code                                           | Content Type                                          |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| apierrors.DocumentAttachmentCreateBadRequestError     | 400                                                   | application/json                                      |
| apierrors.DocumentAttachmentCreateUnauthorizedError   | 401                                                   | application/json                                      |
| apierrors.DocumentAttachmentCreateForbiddenError      | 403                                                   | application/json                                      |
| apierrors.DocumentAttachmentCreateInternalServerError | 500                                                   | application/json                                      |
| apierrors.APIError                                    | 4XX, 5XX                                              | \*/\*                                                 |

## Update

Update an existing attachment

### Example Usage

<!-- UsageSnippet language="go" operationID="document-attachment-update" method="post" path="/document/attachment/update" -->
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

    res, err := s.Documents.Attachments.Update(ctx, operations.DocumentAttachmentUpdateRequest{
        ID: "<id>",
        Data: operations.DocumentAttachmentUpdateData{
            Label: "<value>",
            Data: "https://tinted-ceramics.biz",
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
| `request`                                                                                                | [operations.DocumentAttachmentUpdateRequest](../../models/operations/documentattachmentupdaterequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.DocumentAttachmentUpdateResponse](../../models/operations/documentattachmentupdateresponse.md), error**

### Errors

| Error Type                                            | Status Code                                           | Content Type                                          |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| apierrors.DocumentAttachmentUpdateBadRequestError     | 400                                                   | application/json                                      |
| apierrors.DocumentAttachmentUpdateUnauthorizedError   | 401                                                   | application/json                                      |
| apierrors.DocumentAttachmentUpdateForbiddenError      | 403                                                   | application/json                                      |
| apierrors.DocumentAttachmentUpdateInternalServerError | 500                                                   | application/json                                      |
| apierrors.APIError                                    | 4XX, 5XX                                              | \*/\*                                                 |

## Delete

Delete an attachment from a document

### Example Usage

<!-- UsageSnippet language="go" operationID="document-attachment-delete" method="post" path="/document/attachment/delete" -->
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

    res, err := s.Documents.Attachments.Delete(ctx, operations.DocumentAttachmentDeleteRequest{
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
| `request`                                                                                                | [operations.DocumentAttachmentDeleteRequest](../../models/operations/documentattachmentdeleterequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.DocumentAttachmentDeleteResponse](../../models/operations/documentattachmentdeleteresponse.md), error**

### Errors

| Error Type                                            | Status Code                                           | Content Type                                          |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| apierrors.DocumentAttachmentDeleteBadRequestError     | 400                                                   | application/json                                      |
| apierrors.DocumentAttachmentDeleteUnauthorizedError   | 401                                                   | application/json                                      |
| apierrors.DocumentAttachmentDeleteForbiddenError      | 403                                                   | application/json                                      |
| apierrors.DocumentAttachmentDeleteInternalServerError | 500                                                   | application/json                                      |
| apierrors.APIError                                    | 4XX, 5XX                                              | \*/\*                                                 |

## Find

Find all attachments for a document

### Example Usage

<!-- UsageSnippet language="go" operationID="document-attachment-find" method="get" path="/document/attachment" -->
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

    res, err := s.Documents.Attachments.Find(ctx, 965.17)
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
| `documentID`                                             | *float64*                                                | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DocumentAttachmentFindResponse](../../models/operations/documentattachmentfindresponse.md), error**

### Errors

| Error Type                                          | Status Code                                         | Content Type                                        |
| --------------------------------------------------- | --------------------------------------------------- | --------------------------------------------------- |
| apierrors.DocumentAttachmentFindBadRequestError     | 400                                                 | application/json                                    |
| apierrors.DocumentAttachmentFindUnauthorizedError   | 401                                                 | application/json                                    |
| apierrors.DocumentAttachmentFindForbiddenError      | 403                                                 | application/json                                    |
| apierrors.DocumentAttachmentFindNotFoundError       | 404                                                 | application/json                                    |
| apierrors.DocumentAttachmentFindInternalServerError | 500                                                 | application/json                                    |
| apierrors.APIError                                  | 4XX, 5XX                                            | \*/\*                                               |