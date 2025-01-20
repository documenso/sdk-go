# DocumensoFields
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

    res, err := s.Templates.Fields.Create(ctx, operations.FieldCreateTemplateFieldRequestBody{
        TemplateID: 4865.89,
        Field: operations.CreateFieldCreateTemplateFieldFieldFieldCreateTemplateFieldField8(
            operations.FieldCreateTemplateFieldField8{
                Type: operations.FieldCreateTemplateFieldFieldTemplatesFieldsRequestRequestBody8TypeNumber,
                RecipientID: 4174.58,
                PageNumber: 1343.65,
                PageX: 690.25,
                PageY: 7964.74,
                Width: 9510.62,
                Height: 0.86,
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

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.FieldCreateTemplateFieldRequestBody](../../models/operations/fieldcreatetemplatefieldrequestbody.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `opts`                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |

### Response

**[*operations.FieldCreateTemplateFieldResponse](../../models/operations/fieldcreatetemplatefieldresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| apierrors.ErrorBADREQUEST          | 400                                | application/json                   |
| apierrors.ERRORINTERNALSERVERERROR | 500                                | application/json                   |
| apierrors.APIError                 | 4XX, 5XX                           | \*/\*                              |

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

    res, err := s.Templates.Fields.Get(ctx, 7003.47)
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

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| apierrors.ErrorBADREQUEST          | 400                                | application/json                   |
| apierrors.ErrorNOTFOUND            | 404                                | application/json                   |
| apierrors.ERRORINTERNALSERVERERROR | 500                                | application/json                   |
| apierrors.APIError                 | 4XX, 5XX                           | \*/\*                              |

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

    res, err := s.Templates.Fields.CreateMany(ctx, operations.FieldCreateTemplateFieldsRequestBody{
        TemplateID: 5158.41,
        Fields: []operations.FieldCreateTemplateFieldsFields{
            operations.CreateFieldCreateTemplateFieldsFieldsFieldCreateTemplateFieldsFields10(
                operations.FieldCreateTemplateFieldsFields10{
                    Type: operations.FieldCreateTemplateFieldsFieldsTemplatesFieldsRequestRequestBody10TypeCheckbox,
                    RecipientID: 2516.72,
                    PageNumber: 2304.17,
                    PageX: 7760.32,
                    PageY: 3376.66,
                    Width: 3566.94,
                    Height: 2768.94,
                },
            ),
            operations.CreateFieldCreateTemplateFieldsFieldsFieldCreateTemplateFieldsFields8(
                operations.FieldCreateTemplateFieldsFields8{
                    Type: operations.FieldCreateTemplateFieldsFieldsTemplatesFieldsRequestRequestBody8TypeNumber,
                    RecipientID: 5689.64,
                    PageNumber: 6483.69,
                    PageX: 7271.79,
                    PageY: 1891.56,
                    Width: 7263.21,
                    Height: 5043.41,
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

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.FieldCreateTemplateFieldsRequestBody](../../models/operations/fieldcreatetemplatefieldsrequestbody.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |
| `opts`                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                           | :heavy_minus_sign:                                                                                                 | The options for this request.                                                                                      |

### Response

**[*operations.FieldCreateTemplateFieldsResponse](../../models/operations/fieldcreatetemplatefieldsresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| apierrors.ErrorBADREQUEST          | 400                                | application/json                   |
| apierrors.ERRORINTERNALSERVERERROR | 500                                | application/json                   |
| apierrors.APIError                 | 4XX, 5XX                           | \*/\*                              |

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

    res, err := s.Templates.Fields.Update(ctx, operations.FieldUpdateTemplateFieldRequestBody{
        TemplateID: 8574.78,
        Field: operations.CreateFieldUpdateTemplateFieldFieldFieldUpdateTemplateFieldField7(
            operations.FieldUpdateTemplateFieldField7{
                Type: operations.FieldUpdateTemplateFieldFieldTemplatesFieldsRequestRequestBody7TypeText,
                ID: 3446.2,
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

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.FieldUpdateTemplateFieldRequestBody](../../models/operations/fieldupdatetemplatefieldrequestbody.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `opts`                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |

### Response

**[*operations.FieldUpdateTemplateFieldResponse](../../models/operations/fieldupdatetemplatefieldresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| apierrors.ErrorBADREQUEST          | 400                                | application/json                   |
| apierrors.ERRORINTERNALSERVERERROR | 500                                | application/json                   |
| apierrors.APIError                 | 4XX, 5XX                           | \*/\*                              |

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

    res, err := s.Templates.Fields.UpdateMany(ctx, operations.FieldUpdateTemplateFieldsRequestBody{
        TemplateID: 4057.69,
        Fields: []operations.FieldUpdateTemplateFieldsFields{
            operations.CreateFieldUpdateTemplateFieldsFieldsFieldUpdateTemplateFieldsFields6(
                operations.FieldUpdateTemplateFieldsFields6{
                    Type: operations.FieldUpdateTemplateFieldsFieldsTemplatesFieldsRequestRequestBody6TypeDate,
                    ID: 8982.15,
                },
            ),
            operations.CreateFieldUpdateTemplateFieldsFieldsFieldUpdateTemplateFieldsFields4(
                operations.FieldUpdateTemplateFieldsFields4{
                    Type: operations.FieldUpdateTemplateFieldsFieldsTemplatesFieldsRequestRequestBody4TypeName,
                    ID: 310.19,
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

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.FieldUpdateTemplateFieldsRequestBody](../../models/operations/fieldupdatetemplatefieldsrequestbody.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |
| `opts`                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                           | :heavy_minus_sign:                                                                                                 | The options for this request.                                                                                      |

### Response

**[*operations.FieldUpdateTemplateFieldsResponse](../../models/operations/fieldupdatetemplatefieldsresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| apierrors.ErrorBADREQUEST          | 400                                | application/json                   |
| apierrors.ERRORINTERNALSERVERERROR | 500                                | application/json                   |
| apierrors.APIError                 | 4XX, 5XX                           | \*/\*                              |

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

    res, err := s.Templates.Fields.Delete(ctx, operations.FieldDeleteTemplateFieldRequestBody{
        FieldID: 5459.07,
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

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.FieldDeleteTemplateFieldRequestBody](../../models/operations/fielddeletetemplatefieldrequestbody.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `opts`                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |

### Response

**[*operations.FieldDeleteTemplateFieldResponse](../../models/operations/fielddeletetemplatefieldresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| apierrors.ErrorBADREQUEST          | 400                                | application/json                   |
| apierrors.ERRORINTERNALSERVERERROR | 500                                | application/json                   |
| apierrors.APIError                 | 4XX, 5XX                           | \*/\*                              |