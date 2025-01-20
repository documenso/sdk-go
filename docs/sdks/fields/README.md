# Fields
(*Documents.Fields*)

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

    res, err := s.Documents.Fields.Get(ctx, 7003.46)
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

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| apierrors.ErrorBADREQUEST          | 400                                | application/json                   |
| apierrors.ErrorNOTFOUND            | 404                                | application/json                   |
| apierrors.ERRORINTERNALSERVERERROR | 500                                | application/json                   |
| apierrors.APIError                 | 4XX, 5XX                           | \*/\*                              |

## Create

Create a single field for a document.

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

    res, err := s.Documents.Fields.Create(ctx, operations.FieldCreateDocumentFieldRequestBody{
        DocumentID: 4865.89,
        Field: operations.CreateFieldField1(
            operations.Field1{
                Type: operations.FieldTypeSignature,
                RecipientID: 1697.27,
                PageNumber: 899.64,
                PageX: 7926.20,
                PageY: 8165.87,
                Width: 5862.20,
                Height: 7524.37,
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
| `request`                                                                                                        | [operations.FieldCreateDocumentFieldRequestBody](../../models/operations/fieldcreatedocumentfieldrequestbody.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `opts`                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |

### Response

**[*operations.FieldCreateDocumentFieldResponse](../../models/operations/fieldcreatedocumentfieldresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| apierrors.ErrorBADREQUEST          | 400                                | application/json                   |
| apierrors.ERRORINTERNALSERVERERROR | 500                                | application/json                   |
| apierrors.APIError                 | 4XX, 5XX                           | \*/\*                              |

## CreateMany

Create multiple fields for a document.

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

    res, err := s.Documents.Fields.CreateMany(ctx, operations.FieldCreateDocumentFieldsRequestBody{
        DocumentID: 5158.41,
        Fields: []operations.FieldCreateDocumentFieldsFields{
            operations.CreateFieldCreateDocumentFieldsFieldsFields8(
                operations.Fields8{
                    Type: operations.FieldCreateDocumentFieldsFieldsDocumentsFieldsRequestRequestBody8TypeNumber,
                    RecipientID: 5043.41,
                    PageNumber: 9910.51,
                    PageX: 7633.87,
                    PageY: 2229.58,
                    Width: 4749.13,
                    Height: 5816.06,
                },
            ),
            operations.CreateFieldCreateDocumentFieldsFieldsFields9(
                operations.Fields9{
                    Type: operations.FieldCreateDocumentFieldsFieldsDocumentsFieldsRequestRequestBody9TypeRadio,
                    RecipientID: 6129.81,
                    PageNumber: 8630.00,
                    PageX: 6459.94,
                    PageY: 6070.00,
                    Width: 2530.57,
                    Height: 4918.06,
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
| `request`                                                                                                          | [operations.FieldCreateDocumentFieldsRequestBody](../../models/operations/fieldcreatedocumentfieldsrequestbody.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |
| `opts`                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                           | :heavy_minus_sign:                                                                                                 | The options for this request.                                                                                      |

### Response

**[*operations.FieldCreateDocumentFieldsResponse](../../models/operations/fieldcreatedocumentfieldsresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| apierrors.ErrorBADREQUEST          | 400                                | application/json                   |
| apierrors.ERRORINTERNALSERVERERROR | 500                                | application/json                   |
| apierrors.APIError                 | 4XX, 5XX                           | \*/\*                              |

## Update

Update a single field for a document.

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

    res, err := s.Documents.Fields.Update(ctx, operations.FieldUpdateDocumentFieldRequestBody{
        DocumentID: 8574.77,
        Field: operations.CreateFieldUpdateDocumentFieldFieldFieldUpdateDocumentFieldField4(
            operations.FieldUpdateDocumentFieldField4{
                Type: operations.FieldUpdateDocumentFieldFieldDocumentsFieldsRequestRequestBody4TypeName,
                ID: 9914.64,
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
| `request`                                                                                                        | [operations.FieldUpdateDocumentFieldRequestBody](../../models/operations/fieldupdatedocumentfieldrequestbody.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `opts`                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |

### Response

**[*operations.FieldUpdateDocumentFieldResponse](../../models/operations/fieldupdatedocumentfieldresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| apierrors.ErrorBADREQUEST          | 400                                | application/json                   |
| apierrors.ERRORINTERNALSERVERERROR | 500                                | application/json                   |
| apierrors.APIError                 | 4XX, 5XX                           | \*/\*                              |

## UpdateMany

Update multiple fields for a document.

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

    res, err := s.Documents.Fields.UpdateMany(ctx, operations.FieldUpdateDocumentFieldsRequestBody{
        DocumentID: 4057.69,
        Fields: []operations.FieldUpdateDocumentFieldsFields{
            operations.CreateFieldUpdateDocumentFieldsFieldsFieldUpdateDocumentFieldsFields4(
                operations.FieldUpdateDocumentFieldsFields4{
                    Type: operations.FieldUpdateDocumentFieldsFieldsDocumentsFieldsRequestRequestBody4TypeName,
                    ID: 310.20,
                },
            ),
            operations.CreateFieldUpdateDocumentFieldsFieldsFieldUpdateDocumentFieldsFields2(
                operations.FieldUpdateDocumentFieldsFields2{
                    Type: operations.FieldUpdateDocumentFieldsFieldsDocumentsFieldsTypeFreeSignature,
                    ID: 8948.10,
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
| `request`                                                                                                          | [operations.FieldUpdateDocumentFieldsRequestBody](../../models/operations/fieldupdatedocumentfieldsrequestbody.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |
| `opts`                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                           | :heavy_minus_sign:                                                                                                 | The options for this request.                                                                                      |

### Response

**[*operations.FieldUpdateDocumentFieldsResponse](../../models/operations/fieldupdatedocumentfieldsresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| apierrors.ErrorBADREQUEST          | 400                                | application/json                   |
| apierrors.ERRORINTERNALSERVERERROR | 500                                | application/json                   |
| apierrors.APIError                 | 4XX, 5XX                           | \*/\*                              |

## Delete

Delete document field

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

    res, err := s.Documents.Fields.Delete(ctx, operations.FieldDeleteDocumentFieldRequestBody{
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
| `request`                                                                                                        | [operations.FieldDeleteDocumentFieldRequestBody](../../models/operations/fielddeletedocumentfieldrequestbody.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `opts`                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |

### Response

**[*operations.FieldDeleteDocumentFieldResponse](../../models/operations/fielddeletedocumentfieldresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| apierrors.ErrorBADREQUEST          | 400                                | application/json                   |
| apierrors.ERRORINTERNALSERVERERROR | 500                                | application/json                   |
| apierrors.APIError                 | 4XX, 5XX                           | \*/\*                              |