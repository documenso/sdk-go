# Documents.Fields

## Overview

### Available Operations

* [Get](#get) - Get document field
* [Create](#create) - Create document field
* [CreateMany](#createmany) - Create document fields
* [Update](#update) - Update document field
* [UpdateMany](#updatemany) - Update document fields
* [Delete](#delete) - Delete document field

## Get

Returns a single field. If you want to retrieve all the fields for a document, use the "Get Document" endpoint.

### Example Usage

<!-- UsageSnippet language="go" operationID="field-getDocumentField" method="get" path="/document/field/{fieldId}" -->
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

    res, err := s.Documents.Fields.Get(ctx, 6077.81)
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

**[*operations.FieldGetDocumentFieldResponse](../../models/operations/fieldgetdocumentfieldresponse.md), error**

### Errors

| Error Type                                         | Status Code                                        | Content Type                                       |
| -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- |
| apierrors.FieldGetDocumentFieldBadRequestError     | 400                                                | application/json                                   |
| apierrors.FieldGetDocumentFieldUnauthorizedError   | 401                                                | application/json                                   |
| apierrors.FieldGetDocumentFieldForbiddenError      | 403                                                | application/json                                   |
| apierrors.FieldGetDocumentFieldNotFoundError       | 404                                                | application/json                                   |
| apierrors.FieldGetDocumentFieldInternalServerError | 500                                                | application/json                                   |
| apierrors.APIError                                 | 4XX, 5XX                                           | \*/\*                                              |

## Create

Create a single field for a document.

### Example Usage

<!-- UsageSnippet language="go" operationID="field-createDocumentField" method="post" path="/document/field/create" -->
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

    res, err := s.Documents.Fields.Create(ctx, operations.FieldCreateDocumentFieldRequest{
        DocumentID: 8001.93,
        Field: operations.CreateFieldCreateDocumentFieldFieldUnionFieldCreateDocumentFieldFieldName(
            operations.FieldCreateDocumentFieldFieldName{
                Type: operations.FieldCreateDocumentFieldTypeNameRequest1Name,
                RecipientID: 2564.68,
                PageNumber: 791.77,
                PageX: 7845.22,
                PageY: 6843.16,
                Width: 3932.15,
                Height: 8879.89,
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

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.FieldCreateDocumentFieldRequest](../../models/operations/fieldcreatedocumentfieldrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.FieldCreateDocumentFieldResponse](../../models/operations/fieldcreatedocumentfieldresponse.md), error**

### Errors

| Error Type                                            | Status Code                                           | Content Type                                          |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| apierrors.FieldCreateDocumentFieldBadRequestError     | 400                                                   | application/json                                      |
| apierrors.FieldCreateDocumentFieldUnauthorizedError   | 401                                                   | application/json                                      |
| apierrors.FieldCreateDocumentFieldForbiddenError      | 403                                                   | application/json                                      |
| apierrors.FieldCreateDocumentFieldInternalServerError | 500                                                   | application/json                                      |
| apierrors.APIError                                    | 4XX, 5XX                                              | \*/\*                                                 |

## CreateMany

Create multiple fields for a document.

### Example Usage

<!-- UsageSnippet language="go" operationID="field-createDocumentFields" method="post" path="/document/field/create-many" -->
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

    res, err := s.Documents.Fields.CreateMany(ctx, operations.FieldCreateDocumentFieldsRequest{
        DocumentID: 6257.51,
        Fields: []operations.FieldCreateDocumentFieldsFieldUnion{
            operations.CreateFieldCreateDocumentFieldsFieldUnionFieldCreateDocumentFieldsFieldFreeSignature(
                operations.FieldCreateDocumentFieldsFieldFreeSignature{
                    Type: operations.FieldCreateDocumentFieldsTypeFreeSignatureFreeSignature,
                    RecipientID: 679.35,
                    PageNumber: 5914.59,
                    PageX: 7253.11,
                    PageY: 8426.91,
                    Width: 8995.55,
                    Height: 9808.97,
                },
            ),
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

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.FieldCreateDocumentFieldsRequest](../../models/operations/fieldcreatedocumentfieldsrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.FieldCreateDocumentFieldsResponse](../../models/operations/fieldcreatedocumentfieldsresponse.md), error**

### Errors

| Error Type                                             | Status Code                                            | Content Type                                           |
| ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ |
| apierrors.FieldCreateDocumentFieldsBadRequestError     | 400                                                    | application/json                                       |
| apierrors.FieldCreateDocumentFieldsUnauthorizedError   | 401                                                    | application/json                                       |
| apierrors.FieldCreateDocumentFieldsForbiddenError      | 403                                                    | application/json                                       |
| apierrors.FieldCreateDocumentFieldsInternalServerError | 500                                                    | application/json                                       |
| apierrors.APIError                                     | 4XX, 5XX                                               | \*/\*                                                  |

## Update

Update a single field for a document.

### Example Usage

<!-- UsageSnippet language="go" operationID="field-updateDocumentField" method="post" path="/document/field/update" -->
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

    res, err := s.Documents.Fields.Update(ctx, operations.FieldUpdateDocumentFieldRequest{
        DocumentID: 5956.26,
        Field: operations.CreateFieldUpdateDocumentFieldFieldUnionFieldUpdateDocumentFieldFieldFreeSignature(
            operations.FieldUpdateDocumentFieldFieldFreeSignature{
                Type: operations.FieldUpdateDocumentFieldTypeFreeSignatureFreeSignature,
                ID: 6955.16,
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

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.FieldUpdateDocumentFieldRequest](../../models/operations/fieldupdatedocumentfieldrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.FieldUpdateDocumentFieldResponse](../../models/operations/fieldupdatedocumentfieldresponse.md), error**

### Errors

| Error Type                                            | Status Code                                           | Content Type                                          |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| apierrors.FieldUpdateDocumentFieldBadRequestError     | 400                                                   | application/json                                      |
| apierrors.FieldUpdateDocumentFieldUnauthorizedError   | 401                                                   | application/json                                      |
| apierrors.FieldUpdateDocumentFieldForbiddenError      | 403                                                   | application/json                                      |
| apierrors.FieldUpdateDocumentFieldInternalServerError | 500                                                   | application/json                                      |
| apierrors.APIError                                    | 4XX, 5XX                                              | \*/\*                                                 |

## UpdateMany

Update multiple fields for a document.

### Example Usage

<!-- UsageSnippet language="go" operationID="field-updateDocumentFields" method="post" path="/document/field/update-many" -->
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

    res, err := s.Documents.Fields.UpdateMany(ctx, operations.FieldUpdateDocumentFieldsRequest{
        DocumentID: 9317.43,
        Fields: []operations.FieldUpdateDocumentFieldsFieldUnion{},
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

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.FieldUpdateDocumentFieldsRequest](../../models/operations/fieldupdatedocumentfieldsrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.FieldUpdateDocumentFieldsResponse](../../models/operations/fieldupdatedocumentfieldsresponse.md), error**

### Errors

| Error Type                                             | Status Code                                            | Content Type                                           |
| ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ |
| apierrors.FieldUpdateDocumentFieldsBadRequestError     | 400                                                    | application/json                                       |
| apierrors.FieldUpdateDocumentFieldsUnauthorizedError   | 401                                                    | application/json                                       |
| apierrors.FieldUpdateDocumentFieldsForbiddenError      | 403                                                    | application/json                                       |
| apierrors.FieldUpdateDocumentFieldsInternalServerError | 500                                                    | application/json                                       |
| apierrors.APIError                                     | 4XX, 5XX                                               | \*/\*                                                  |

## Delete

Delete document field

### Example Usage

<!-- UsageSnippet language="go" operationID="field-deleteDocumentField" method="post" path="/document/field/delete" -->
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

    res, err := s.Documents.Fields.Delete(ctx, operations.FieldDeleteDocumentFieldRequest{
        FieldID: 4748.27,
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
| `request`                                                                                                | [operations.FieldDeleteDocumentFieldRequest](../../models/operations/fielddeletedocumentfieldrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.FieldDeleteDocumentFieldResponse](../../models/operations/fielddeletedocumentfieldresponse.md), error**

### Errors

| Error Type                                            | Status Code                                           | Content Type                                          |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| apierrors.FieldDeleteDocumentFieldBadRequestError     | 400                                                   | application/json                                      |
| apierrors.FieldDeleteDocumentFieldUnauthorizedError   | 401                                                   | application/json                                      |
| apierrors.FieldDeleteDocumentFieldForbiddenError      | 403                                                   | application/json                                      |
| apierrors.FieldDeleteDocumentFieldInternalServerError | 500                                                   | application/json                                      |
| apierrors.APIError                                    | 4XX, 5XX                                              | \*/\*                                                 |