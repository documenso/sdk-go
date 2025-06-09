# TemplatesFields
(*Templates.Fields*)

## Overview

### Available Operations

* [Create](#create) - Create template field
* [Get](#get) - Get template field
* [CreateMany](#createmany) - Create template fields
* [Update](#update) - Update template field
* [UpdateMany](#updatemany) - Update template fields
* [Delete](#delete) - Delete template field

## Create

Create a single field for a template.

### Example Usage

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
                Type: operations.FieldCreateTemplateFieldTypeDateRequestBody1Date,
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
        // handle response
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
| apierrors.FieldCreateTemplateFieldInternalServerError | 500                                                   | application/json                                      |
| apierrors.APIError                                    | 4XX, 5XX                                              | \*/\*                                                 |

## Get

Returns a single field. If you want to retrieve all the fields for a template, use the "Get Template" endpoint.

### Example Usage

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

    res, err := s.Templates.Fields.Get(ctx, 1152.82)
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

**[*operations.FieldGetTemplateFieldResponse](../../models/operations/fieldgettemplatefieldresponse.md), error**

### Errors

| Error Type                                         | Status Code                                        | Content Type                                       |
| -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- |
| apierrors.FieldGetTemplateFieldBadRequestError     | 400                                                | application/json                                   |
| apierrors.FieldGetTemplateFieldNotFoundError       | 404                                                | application/json                                   |
| apierrors.FieldGetTemplateFieldInternalServerError | 500                                                | application/json                                   |
| apierrors.APIError                                 | 4XX, 5XX                                           | \*/\*                                              |

## CreateMany

Create multiple fields for a template.

### Example Usage

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
                    Type: operations.FieldCreateTemplateFieldsTypeSignatureSignature,
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
| apierrors.FieldCreateTemplateFieldsInternalServerError | 500                                                    | application/json                                       |
| apierrors.APIError                                     | 4XX, 5XX                                               | \*/\*                                                  |

## Update

Update a single field for a template.

### Example Usage

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
                Type: operations.FieldUpdateTemplateFieldTypeTextRequestBody1Text,
                ID: 1792.29,
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
| `request`                                                                                                | [operations.FieldUpdateTemplateFieldRequest](../../models/operations/fieldupdatetemplatefieldrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.FieldUpdateTemplateFieldResponse](../../models/operations/fieldupdatetemplatefieldresponse.md), error**

### Errors

| Error Type                                            | Status Code                                           | Content Type                                          |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| apierrors.FieldUpdateTemplateFieldBadRequestError     | 400                                                   | application/json                                      |
| apierrors.FieldUpdateTemplateFieldInternalServerError | 500                                                   | application/json                                      |
| apierrors.APIError                                    | 4XX, 5XX                                              | \*/\*                                                 |

## UpdateMany

Update multiple fields for a template.

### Example Usage

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
                    Type: operations.FieldUpdateTemplateFieldsTypeDropdownRequestBody1Dropdown,
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
| apierrors.FieldUpdateTemplateFieldsInternalServerError | 500                                                    | application/json                                       |
| apierrors.APIError                                     | 4XX, 5XX                                               | \*/\*                                                  |

## Delete

Delete template field

### Example Usage

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
| apierrors.FieldDeleteTemplateFieldInternalServerError | 500                                                   | application/json                                      |
| apierrors.APIError                                    | 4XX, 5XX                                              | \*/\*                                                 |