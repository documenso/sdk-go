# ~~Templates.Fields~~

> [!WARNING]
> This SDK is **DEPRECATED**

## Overview

### Available Operations

* [~~Create~~](#create) - Create template field :warning: **Deprecated**
* [~~Get~~](#get) - Get template field :warning: **Deprecated**
* [~~CreateMany~~](#createmany) - Create template fields :warning: **Deprecated**
* [~~Update~~](#update) - Update template field :warning: **Deprecated**
* [~~UpdateMany~~](#updatemany) - Update template fields :warning: **Deprecated**
* [~~Delete~~](#delete) - Delete template field :warning: **Deprecated**

## ~~Create~~

Deprecated: this endpoint is being replaced by the Envelope API. See https://docs.documenso.com/docs/developers/api/migrate-to-envelopes for the migration guide. Create a single field for a template.

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage

<!-- UsageSnippet language="go" operationID="field-createTemplateField" method="post" path="/template/field/create" -->
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

    res, err := s.Templates.Fields.Create(ctx, operations.FieldCreateTemplateFieldRequest{
        TemplateID: 1203.71,
        Field: operations.CreateFieldCreateTemplateFieldFieldUnionFieldCreateTemplateFieldFieldDate(
            operations.FieldCreateTemplateFieldFieldDate{
                Type: operations.FieldCreateTemplateFieldTypeDateRequest1Date,
                RecipientID: 2738.54,
                PageNumber: 5735.12,
                PageX: 2936.28,
                PageY: 8594.41,
                Width: 7589.39,
                Height: 3122.23,
            },
        ),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        switch res.Object.FieldMeta.Type {
            case operations.FieldCreateTemplateFieldFieldMetaUnionTypeFieldCreateTemplateFieldFieldMetaSignatureResponse:
                // res.Object.FieldMeta.FieldCreateTemplateFieldFieldMetaSignatureResponse is populated
            case operations.FieldCreateTemplateFieldFieldMetaUnionTypeFieldCreateTemplateFieldFieldMetaInitialsResponse:
                // res.Object.FieldMeta.FieldCreateTemplateFieldFieldMetaInitialsResponse is populated
            case operations.FieldCreateTemplateFieldFieldMetaUnionTypeFieldCreateTemplateFieldFieldMetaNameResponse:
                // res.Object.FieldMeta.FieldCreateTemplateFieldFieldMetaNameResponse is populated
            case operations.FieldCreateTemplateFieldFieldMetaUnionTypeFieldCreateTemplateFieldFieldMetaEmailResponse:
                // res.Object.FieldMeta.FieldCreateTemplateFieldFieldMetaEmailResponse is populated
            case operations.FieldCreateTemplateFieldFieldMetaUnionTypeFieldCreateTemplateFieldFieldMetaDateResponse:
                // res.Object.FieldMeta.FieldCreateTemplateFieldFieldMetaDateResponse is populated
            case operations.FieldCreateTemplateFieldFieldMetaUnionTypeFieldCreateTemplateFieldFieldMetaTextResponse:
                // res.Object.FieldMeta.FieldCreateTemplateFieldFieldMetaTextResponse is populated
            case operations.FieldCreateTemplateFieldFieldMetaUnionTypeFieldCreateTemplateFieldFieldMetaNumberResponse:
                // res.Object.FieldMeta.FieldCreateTemplateFieldFieldMetaNumberResponse is populated
            case operations.FieldCreateTemplateFieldFieldMetaUnionTypeFieldCreateTemplateFieldFieldMetaRadioResponse:
                // res.Object.FieldMeta.FieldCreateTemplateFieldFieldMetaRadioResponse is populated
            case operations.FieldCreateTemplateFieldFieldMetaUnionTypeFieldCreateTemplateFieldFieldMetaCheckboxResponse:
                // res.Object.FieldMeta.FieldCreateTemplateFieldFieldMetaCheckboxResponse is populated
            case operations.FieldCreateTemplateFieldFieldMetaUnionTypeFieldCreateTemplateFieldFieldMetaDropdownResponse:
                // res.Object.FieldMeta.FieldCreateTemplateFieldFieldMetaDropdownResponse is populated
        }

    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.FieldCreateTemplateFieldRequest](../../models/operations/fieldcreatetemplatefieldrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.FieldCreateTemplateFieldResponse](../../models/operations/fieldcreatetemplatefieldresponse.md), error**

### Errors

| Error Type                                            | Status Code                                           | Content Type                                          |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| apierrors.FieldCreateTemplateFieldBadRequestError     | 400                                                   | application/json                                      |
| apierrors.FieldCreateTemplateFieldUnauthorizedError   | 401                                                   | application/json                                      |
| apierrors.FieldCreateTemplateFieldForbiddenError      | 403                                                   | application/json                                      |
| apierrors.FieldCreateTemplateFieldInternalServerError | 500                                                   | application/json                                      |
| apierrors.APIError                                    | 4XX, 5XX                                              | \*/\*                                                 |

## ~~Get~~

Deprecated: this endpoint is being replaced by the Envelope API. See https://docs.documenso.com/docs/developers/api/migrate-to-envelopes for the migration guide. Returns a single field. If you want to retrieve all the fields for a template, use the "Get Template" endpoint.

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage

<!-- UsageSnippet language="go" operationID="field-getTemplateField" method="get" path="/template/field/{fieldId}" -->
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

    res, err := s.Templates.Fields.Get(ctx, 1152.82)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        switch res.Object.FieldMeta.Type {
            case operations.FieldGetTemplateFieldFieldMetaUnionTypeFieldGetTemplateFieldFieldMetaSignature:
                // res.Object.FieldMeta.FieldGetTemplateFieldFieldMetaSignature is populated
            case operations.FieldGetTemplateFieldFieldMetaUnionTypeFieldGetTemplateFieldFieldMetaInitials:
                // res.Object.FieldMeta.FieldGetTemplateFieldFieldMetaInitials is populated
            case operations.FieldGetTemplateFieldFieldMetaUnionTypeFieldGetTemplateFieldFieldMetaName:
                // res.Object.FieldMeta.FieldGetTemplateFieldFieldMetaName is populated
            case operations.FieldGetTemplateFieldFieldMetaUnionTypeFieldGetTemplateFieldFieldMetaEmail:
                // res.Object.FieldMeta.FieldGetTemplateFieldFieldMetaEmail is populated
            case operations.FieldGetTemplateFieldFieldMetaUnionTypeFieldGetTemplateFieldFieldMetaDate:
                // res.Object.FieldMeta.FieldGetTemplateFieldFieldMetaDate is populated
            case operations.FieldGetTemplateFieldFieldMetaUnionTypeFieldGetTemplateFieldFieldMetaText:
                // res.Object.FieldMeta.FieldGetTemplateFieldFieldMetaText is populated
            case operations.FieldGetTemplateFieldFieldMetaUnionTypeFieldGetTemplateFieldFieldMetaNumber:
                // res.Object.FieldMeta.FieldGetTemplateFieldFieldMetaNumber is populated
            case operations.FieldGetTemplateFieldFieldMetaUnionTypeFieldGetTemplateFieldFieldMetaRadio:
                // res.Object.FieldMeta.FieldGetTemplateFieldFieldMetaRadio is populated
            case operations.FieldGetTemplateFieldFieldMetaUnionTypeFieldGetTemplateFieldFieldMetaCheckbox:
                // res.Object.FieldMeta.FieldGetTemplateFieldFieldMetaCheckbox is populated
            case operations.FieldGetTemplateFieldFieldMetaUnionTypeFieldGetTemplateFieldFieldMetaDropdown:
                // res.Object.FieldMeta.FieldGetTemplateFieldFieldMetaDropdown is populated
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

**[*operations.FieldGetTemplateFieldResponse](../../models/operations/fieldgettemplatefieldresponse.md), error**

### Errors

| Error Type                                         | Status Code                                        | Content Type                                       |
| -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- |
| apierrors.FieldGetTemplateFieldBadRequestError     | 400                                                | application/json                                   |
| apierrors.FieldGetTemplateFieldUnauthorizedError   | 401                                                | application/json                                   |
| apierrors.FieldGetTemplateFieldForbiddenError      | 403                                                | application/json                                   |
| apierrors.FieldGetTemplateFieldNotFoundError       | 404                                                | application/json                                   |
| apierrors.FieldGetTemplateFieldInternalServerError | 500                                                | application/json                                   |
| apierrors.APIError                                 | 4XX, 5XX                                           | \*/\*                                              |

## ~~CreateMany~~

Deprecated: this endpoint is being replaced by the Envelope API. See https://docs.documenso.com/docs/developers/api/migrate-to-envelopes for the migration guide. Create multiple fields for a template.

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage

<!-- UsageSnippet language="go" operationID="field-createTemplateFields" method="post" path="/template/field/create-many" -->
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

    res, err := s.Templates.Fields.CreateMany(ctx, operations.FieldCreateTemplateFieldsRequest{
        TemplateID: 586.2,
        Fields: []operations.FieldCreateTemplateFieldsFieldUnion{
            operations.CreateFieldCreateTemplateFieldsFieldUnionFieldCreateTemplateFieldsFieldSignature(
                operations.FieldCreateTemplateFieldsFieldSignature{
                    Type: operations.FieldCreateTemplateFieldsTypeSignatureRequest1Signature,
                    RecipientID: 6990.12,
                    PageNumber: 3472.45,
                    PageX: 4747.87,
                    PageY: 1673.94,
                    Width: 7215.37,
                    Height: 9417.43,
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
| `request`                                                                                                  | [operations.FieldCreateTemplateFieldsRequest](../../models/operations/fieldcreatetemplatefieldsrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.FieldCreateTemplateFieldsResponse](../../models/operations/fieldcreatetemplatefieldsresponse.md), error**

### Errors

| Error Type                                             | Status Code                                            | Content Type                                           |
| ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ |
| apierrors.FieldCreateTemplateFieldsBadRequestError     | 400                                                    | application/json                                       |
| apierrors.FieldCreateTemplateFieldsUnauthorizedError   | 401                                                    | application/json                                       |
| apierrors.FieldCreateTemplateFieldsForbiddenError      | 403                                                    | application/json                                       |
| apierrors.FieldCreateTemplateFieldsInternalServerError | 500                                                    | application/json                                       |
| apierrors.APIError                                     | 4XX, 5XX                                               | \*/\*                                                  |

## ~~Update~~

Deprecated: this endpoint is being replaced by the Envelope API. See https://docs.documenso.com/docs/developers/api/migrate-to-envelopes for the migration guide. Update a single field for a template.

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage

<!-- UsageSnippet language="go" operationID="field-updateTemplateField" method="post" path="/template/field/update" -->
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

    res, err := s.Templates.Fields.Update(ctx, operations.FieldUpdateTemplateFieldRequest{
        TemplateID: 5083.07,
        Field: operations.CreateFieldUpdateTemplateFieldFieldUnionFieldUpdateTemplateFieldFieldText(
            operations.FieldUpdateTemplateFieldFieldText{
                Type: operations.FieldUpdateTemplateFieldTypeTextRequest1Text,
                ID: 1792.29,
            },
        ),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        switch res.Object.FieldMeta.Type {
            case operations.FieldUpdateTemplateFieldFieldMetaUnionTypeFieldUpdateTemplateFieldFieldMetaSignatureResponse:
                // res.Object.FieldMeta.FieldUpdateTemplateFieldFieldMetaSignatureResponse is populated
            case operations.FieldUpdateTemplateFieldFieldMetaUnionTypeFieldUpdateTemplateFieldFieldMetaInitialsResponse:
                // res.Object.FieldMeta.FieldUpdateTemplateFieldFieldMetaInitialsResponse is populated
            case operations.FieldUpdateTemplateFieldFieldMetaUnionTypeFieldUpdateTemplateFieldFieldMetaNameResponse:
                // res.Object.FieldMeta.FieldUpdateTemplateFieldFieldMetaNameResponse is populated
            case operations.FieldUpdateTemplateFieldFieldMetaUnionTypeFieldUpdateTemplateFieldFieldMetaEmailResponse:
                // res.Object.FieldMeta.FieldUpdateTemplateFieldFieldMetaEmailResponse is populated
            case operations.FieldUpdateTemplateFieldFieldMetaUnionTypeFieldUpdateTemplateFieldFieldMetaDateResponse:
                // res.Object.FieldMeta.FieldUpdateTemplateFieldFieldMetaDateResponse is populated
            case operations.FieldUpdateTemplateFieldFieldMetaUnionTypeFieldUpdateTemplateFieldFieldMetaTextResponse:
                // res.Object.FieldMeta.FieldUpdateTemplateFieldFieldMetaTextResponse is populated
            case operations.FieldUpdateTemplateFieldFieldMetaUnionTypeFieldUpdateTemplateFieldFieldMetaNumberResponse:
                // res.Object.FieldMeta.FieldUpdateTemplateFieldFieldMetaNumberResponse is populated
            case operations.FieldUpdateTemplateFieldFieldMetaUnionTypeFieldUpdateTemplateFieldFieldMetaRadioResponse:
                // res.Object.FieldMeta.FieldUpdateTemplateFieldFieldMetaRadioResponse is populated
            case operations.FieldUpdateTemplateFieldFieldMetaUnionTypeFieldUpdateTemplateFieldFieldMetaCheckboxResponse:
                // res.Object.FieldMeta.FieldUpdateTemplateFieldFieldMetaCheckboxResponse is populated
            case operations.FieldUpdateTemplateFieldFieldMetaUnionTypeFieldUpdateTemplateFieldFieldMetaDropdownResponse:
                // res.Object.FieldMeta.FieldUpdateTemplateFieldFieldMetaDropdownResponse is populated
        }

    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.FieldUpdateTemplateFieldRequest](../../models/operations/fieldupdatetemplatefieldrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.FieldUpdateTemplateFieldResponse](../../models/operations/fieldupdatetemplatefieldresponse.md), error**

### Errors

| Error Type                                            | Status Code                                           | Content Type                                          |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| apierrors.FieldUpdateTemplateFieldBadRequestError     | 400                                                   | application/json                                      |
| apierrors.FieldUpdateTemplateFieldUnauthorizedError   | 401                                                   | application/json                                      |
| apierrors.FieldUpdateTemplateFieldForbiddenError      | 403                                                   | application/json                                      |
| apierrors.FieldUpdateTemplateFieldInternalServerError | 500                                                   | application/json                                      |
| apierrors.APIError                                    | 4XX, 5XX                                              | \*/\*                                                 |

## ~~UpdateMany~~

Deprecated: this endpoint is being replaced by the Envelope API. See https://docs.documenso.com/docs/developers/api/migrate-to-envelopes for the migration guide. Update multiple fields for a template.

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage

<!-- UsageSnippet language="go" operationID="field-updateTemplateFields" method="post" path="/template/field/update-many" -->
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

    res, err := s.Templates.Fields.UpdateMany(ctx, operations.FieldUpdateTemplateFieldsRequest{
        TemplateID: 3969.1,
        Fields: []operations.FieldUpdateTemplateFieldsFieldUnion{
            operations.CreateFieldUpdateTemplateFieldsFieldUnionFieldUpdateTemplateFieldsFieldDropdown(
                operations.FieldUpdateTemplateFieldsFieldDropdown{
                    Type: operations.FieldUpdateTemplateFieldsTypeDropdownRequest1Dropdown,
                    ID: 2460.72,
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
| `request`                                                                                                  | [operations.FieldUpdateTemplateFieldsRequest](../../models/operations/fieldupdatetemplatefieldsrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.FieldUpdateTemplateFieldsResponse](../../models/operations/fieldupdatetemplatefieldsresponse.md), error**

### Errors

| Error Type                                             | Status Code                                            | Content Type                                           |
| ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ |
| apierrors.FieldUpdateTemplateFieldsBadRequestError     | 400                                                    | application/json                                       |
| apierrors.FieldUpdateTemplateFieldsUnauthorizedError   | 401                                                    | application/json                                       |
| apierrors.FieldUpdateTemplateFieldsForbiddenError      | 403                                                    | application/json                                       |
| apierrors.FieldUpdateTemplateFieldsInternalServerError | 500                                                    | application/json                                       |
| apierrors.APIError                                     | 4XX, 5XX                                               | \*/\*                                                  |

## ~~Delete~~

Deprecated: this endpoint is being replaced by the Envelope API. See https://docs.documenso.com/docs/developers/api/migrate-to-envelopes for the migration guide.

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage

<!-- UsageSnippet language="go" operationID="field-deleteTemplateField" method="post" path="/template/field/delete" -->
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

    res, err := s.Templates.Fields.Delete(ctx, operations.FieldDeleteTemplateFieldRequest{
        FieldID: 7996.49,
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
| `request`                                                                                                | [operations.FieldDeleteTemplateFieldRequest](../../models/operations/fielddeletetemplatefieldrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.FieldDeleteTemplateFieldResponse](../../models/operations/fielddeletetemplatefieldresponse.md), error**

### Errors

| Error Type                                            | Status Code                                           | Content Type                                          |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| apierrors.FieldDeleteTemplateFieldBadRequestError     | 400                                                   | application/json                                      |
| apierrors.FieldDeleteTemplateFieldUnauthorizedError   | 401                                                   | application/json                                      |
| apierrors.FieldDeleteTemplateFieldForbiddenError      | 403                                                   | application/json                                      |
| apierrors.FieldDeleteTemplateFieldInternalServerError | 500                                                   | application/json                                      |
| apierrors.APIError                                    | 4XX, 5XX                                              | \*/\*                                                 |