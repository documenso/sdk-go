# ~~Documents.Fields~~

> [!WARNING]
> This SDK is **DEPRECATED**

## Overview

### Available Operations

* [~~Get~~](#get) - Get document field :warning: **Deprecated**
* [~~Create~~](#create) - Create document field :warning: **Deprecated**
* [~~CreateMany~~](#createmany) - Create document fields :warning: **Deprecated**
* [~~Update~~](#update) - Update document field :warning: **Deprecated**
* [~~UpdateMany~~](#updatemany) - Update document fields :warning: **Deprecated**
* [~~Delete~~](#delete) - Delete document field :warning: **Deprecated**

## ~~Get~~

Deprecated: this endpoint is being replaced by the Envelope API. See https://docs.documenso.com/docs/developers/api/migrate-to-envelopes for the migration guide. Returns a single field. If you want to retrieve all the fields for a document, use the "Get Document" endpoint.

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage

<!-- UsageSnippet language="go" operationID="field-getDocumentField" method="get" path="/document/field/{fieldId}" -->
```go
package main

import(
	"context"
	"os"
	sdkgo "github.com/documenso/sdk-go"
	"log"
	"github.com/documenso/sdk-go/models/operations"
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
        switch res.Object.FieldMeta.Type {
            case operations.FieldGetDocumentFieldFieldMetaUnionTypeFieldGetDocumentFieldFieldMetaSignature:
                // res.Object.FieldMeta.FieldGetDocumentFieldFieldMetaSignature is populated
            case operations.FieldGetDocumentFieldFieldMetaUnionTypeFieldGetDocumentFieldFieldMetaInitials:
                // res.Object.FieldMeta.FieldGetDocumentFieldFieldMetaInitials is populated
            case operations.FieldGetDocumentFieldFieldMetaUnionTypeFieldGetDocumentFieldFieldMetaName:
                // res.Object.FieldMeta.FieldGetDocumentFieldFieldMetaName is populated
            case operations.FieldGetDocumentFieldFieldMetaUnionTypeFieldGetDocumentFieldFieldMetaEmail:
                // res.Object.FieldMeta.FieldGetDocumentFieldFieldMetaEmail is populated
            case operations.FieldGetDocumentFieldFieldMetaUnionTypeFieldGetDocumentFieldFieldMetaDate:
                // res.Object.FieldMeta.FieldGetDocumentFieldFieldMetaDate is populated
            case operations.FieldGetDocumentFieldFieldMetaUnionTypeFieldGetDocumentFieldFieldMetaText:
                // res.Object.FieldMeta.FieldGetDocumentFieldFieldMetaText is populated
            case operations.FieldGetDocumentFieldFieldMetaUnionTypeFieldGetDocumentFieldFieldMetaNumber:
                // res.Object.FieldMeta.FieldGetDocumentFieldFieldMetaNumber is populated
            case operations.FieldGetDocumentFieldFieldMetaUnionTypeFieldGetDocumentFieldFieldMetaRadio:
                // res.Object.FieldMeta.FieldGetDocumentFieldFieldMetaRadio is populated
            case operations.FieldGetDocumentFieldFieldMetaUnionTypeFieldGetDocumentFieldFieldMetaCheckbox:
                // res.Object.FieldMeta.FieldGetDocumentFieldFieldMetaCheckbox is populated
            case operations.FieldGetDocumentFieldFieldMetaUnionTypeFieldGetDocumentFieldFieldMetaDropdown:
                // res.Object.FieldMeta.FieldGetDocumentFieldFieldMetaDropdown is populated
        }

    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `fieldID`                                                | `float64`                                                | :heavy_check_mark:                                       | N/A                                                      |
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

## ~~Create~~

Deprecated: this endpoint is being replaced by the Envelope API. See https://docs.documenso.com/docs/developers/api/migrate-to-envelopes for the migration guide. Create a single field for a document.

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

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
        switch res.Object.FieldMeta.Type {
            case operations.FieldCreateDocumentFieldFieldMetaUnionTypeFieldCreateDocumentFieldFieldMetaSignatureResponse:
                // res.Object.FieldMeta.FieldCreateDocumentFieldFieldMetaSignatureResponse is populated
            case operations.FieldCreateDocumentFieldFieldMetaUnionTypeFieldCreateDocumentFieldFieldMetaInitialsResponse:
                // res.Object.FieldMeta.FieldCreateDocumentFieldFieldMetaInitialsResponse is populated
            case operations.FieldCreateDocumentFieldFieldMetaUnionTypeFieldCreateDocumentFieldFieldMetaNameResponse:
                // res.Object.FieldMeta.FieldCreateDocumentFieldFieldMetaNameResponse is populated
            case operations.FieldCreateDocumentFieldFieldMetaUnionTypeFieldCreateDocumentFieldFieldMetaEmailResponse:
                // res.Object.FieldMeta.FieldCreateDocumentFieldFieldMetaEmailResponse is populated
            case operations.FieldCreateDocumentFieldFieldMetaUnionTypeFieldCreateDocumentFieldFieldMetaDateResponse:
                // res.Object.FieldMeta.FieldCreateDocumentFieldFieldMetaDateResponse is populated
            case operations.FieldCreateDocumentFieldFieldMetaUnionTypeFieldCreateDocumentFieldFieldMetaTextResponse:
                // res.Object.FieldMeta.FieldCreateDocumentFieldFieldMetaTextResponse is populated
            case operations.FieldCreateDocumentFieldFieldMetaUnionTypeFieldCreateDocumentFieldFieldMetaNumberResponse:
                // res.Object.FieldMeta.FieldCreateDocumentFieldFieldMetaNumberResponse is populated
            case operations.FieldCreateDocumentFieldFieldMetaUnionTypeFieldCreateDocumentFieldFieldMetaRadioResponse:
                // res.Object.FieldMeta.FieldCreateDocumentFieldFieldMetaRadioResponse is populated
            case operations.FieldCreateDocumentFieldFieldMetaUnionTypeFieldCreateDocumentFieldFieldMetaCheckboxResponse:
                // res.Object.FieldMeta.FieldCreateDocumentFieldFieldMetaCheckboxResponse is populated
            case operations.FieldCreateDocumentFieldFieldMetaUnionTypeFieldCreateDocumentFieldFieldMetaDropdownResponse:
                // res.Object.FieldMeta.FieldCreateDocumentFieldFieldMetaDropdownResponse is populated
        }

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

## ~~CreateMany~~

Deprecated: this endpoint is being replaced by the Envelope API. See https://docs.documenso.com/docs/developers/api/migrate-to-envelopes for the migration guide. Create multiple fields for a document.

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

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

## ~~Update~~

Deprecated: this endpoint is being replaced by the Envelope API. See https://docs.documenso.com/docs/developers/api/migrate-to-envelopes for the migration guide. Update a single field for a document.

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

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
        switch res.Object.FieldMeta.Type {
            case operations.FieldUpdateDocumentFieldFieldMetaUnionTypeFieldUpdateDocumentFieldFieldMetaSignatureResponse:
                // res.Object.FieldMeta.FieldUpdateDocumentFieldFieldMetaSignatureResponse is populated
            case operations.FieldUpdateDocumentFieldFieldMetaUnionTypeFieldUpdateDocumentFieldFieldMetaInitialsResponse:
                // res.Object.FieldMeta.FieldUpdateDocumentFieldFieldMetaInitialsResponse is populated
            case operations.FieldUpdateDocumentFieldFieldMetaUnionTypeFieldUpdateDocumentFieldFieldMetaNameResponse:
                // res.Object.FieldMeta.FieldUpdateDocumentFieldFieldMetaNameResponse is populated
            case operations.FieldUpdateDocumentFieldFieldMetaUnionTypeFieldUpdateDocumentFieldFieldMetaEmailResponse:
                // res.Object.FieldMeta.FieldUpdateDocumentFieldFieldMetaEmailResponse is populated
            case operations.FieldUpdateDocumentFieldFieldMetaUnionTypeFieldUpdateDocumentFieldFieldMetaDateResponse:
                // res.Object.FieldMeta.FieldUpdateDocumentFieldFieldMetaDateResponse is populated
            case operations.FieldUpdateDocumentFieldFieldMetaUnionTypeFieldUpdateDocumentFieldFieldMetaTextResponse:
                // res.Object.FieldMeta.FieldUpdateDocumentFieldFieldMetaTextResponse is populated
            case operations.FieldUpdateDocumentFieldFieldMetaUnionTypeFieldUpdateDocumentFieldFieldMetaNumberResponse:
                // res.Object.FieldMeta.FieldUpdateDocumentFieldFieldMetaNumberResponse is populated
            case operations.FieldUpdateDocumentFieldFieldMetaUnionTypeFieldUpdateDocumentFieldFieldMetaRadioResponse:
                // res.Object.FieldMeta.FieldUpdateDocumentFieldFieldMetaRadioResponse is populated
            case operations.FieldUpdateDocumentFieldFieldMetaUnionTypeFieldUpdateDocumentFieldFieldMetaCheckboxResponse:
                // res.Object.FieldMeta.FieldUpdateDocumentFieldFieldMetaCheckboxResponse is populated
            case operations.FieldUpdateDocumentFieldFieldMetaUnionTypeFieldUpdateDocumentFieldFieldMetaDropdownResponse:
                // res.Object.FieldMeta.FieldUpdateDocumentFieldFieldMetaDropdownResponse is populated
        }

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

## ~~UpdateMany~~

Deprecated: this endpoint is being replaced by the Envelope API. See https://docs.documenso.com/docs/developers/api/migrate-to-envelopes for the migration guide. Update multiple fields for a document.

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

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

## ~~Delete~~

Deprecated: this endpoint is being replaced by the Envelope API. See https://docs.documenso.com/docs/developers/api/migrate-to-envelopes for the migration guide.

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

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