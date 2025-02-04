# Recipients
(*Documents.Recipients*)

## Overview

### Available Operations

* [Get](#get) - Get document recipient
* [Create](#create) - Create document recipient
* [CreateMany](#createmany) - Create document recipients
* [Update](#update) - Update document recipient
* [UpdateMany](#updatemany) - Update document recipients
* [Delete](#delete) - Delete document recipient

## Get

Returns a single recipient. If you want to retrieve all the recipients for a document, use the "Get Document" endpoint.

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

    res, err := s.Documents.Recipients.Get(ctx, 7003.47)
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

**[*operations.RecipientGetDocumentRecipientResponse](../../models/operations/recipientgetdocumentrecipientresponse.md), error**

### Errors

| Error Type                                                                     | Status Code                                                                    | Content Type                                                                   |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| apierrors.RecipientGetDocumentRecipientResponseBody                            | 400                                                                            | application/json                                                               |
| apierrors.RecipientGetDocumentRecipientDocumentsRecipientsResponseBody         | 404                                                                            | application/json                                                               |
| apierrors.RecipientGetDocumentRecipientDocumentsRecipientsResponseResponseBody | 500                                                                            | application/json                                                               |
| apierrors.APIError                                                             | 4XX, 5XX                                                                       | \*/\*                                                                          |

## Create

Create a single recipient for a document.

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

    res, err := s.Documents.Recipients.Create(ctx, operations.RecipientCreateDocumentRecipientRequestBody{
        DocumentID: 4865.89,
        Recipient: operations.Recipient{
            Email: "Haylie_Bernhard95@yahoo.com",
            Name: "<value>",
            Role: operations.RecipientCreateDocumentRecipientRoleCc,
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

| Parameter                                                                                                                        | Type                                                                                                                             | Required                                                                                                                         | Description                                                                                                                      |
| -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                            | :heavy_check_mark:                                                                                                               | The context to use for the request.                                                                                              |
| `request`                                                                                                                        | [operations.RecipientCreateDocumentRecipientRequestBody](../../models/operations/recipientcreatedocumentrecipientrequestbody.md) | :heavy_check_mark:                                                                                                               | The request object to use for the request.                                                                                       |
| `opts`                                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                                         | :heavy_minus_sign:                                                                                                               | The options for this request.                                                                                                    |

### Response

**[*operations.RecipientCreateDocumentRecipientResponse](../../models/operations/recipientcreatedocumentrecipientresponse.md), error**

### Errors

| Error Type                                                                | Status Code                                                               | Content Type                                                              |
| ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- |
| apierrors.RecipientCreateDocumentRecipientResponseBody                    | 400                                                                       | application/json                                                          |
| apierrors.RecipientCreateDocumentRecipientDocumentsRecipientsResponseBody | 500                                                                       | application/json                                                          |
| apierrors.APIError                                                        | 4XX, 5XX                                                                  | \*/\*                                                                     |

## CreateMany

Create multiple recipients for a document.

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

    res, err := s.Documents.Recipients.CreateMany(ctx, operations.RecipientCreateDocumentRecipientsRequestBody{
        DocumentID: 5158.41,
        Recipients: []operations.RecipientCreateDocumentRecipientsRecipients{
            operations.RecipientCreateDocumentRecipientsRecipients{
                Email: "Demetrius.Sanford35@hotmail.com",
                Name: "<value>",
                Role: operations.RecipientCreateDocumentRecipientsRoleViewer,
            },
            operations.RecipientCreateDocumentRecipientsRecipients{
                Email: "Lyla50@yahoo.com",
                Name: "<value>",
                Role: operations.RecipientCreateDocumentRecipientsRoleApprover,
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

| Parameter                                                                                                                          | Type                                                                                                                               | Required                                                                                                                           | Description                                                                                                                        |
| ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                              | :heavy_check_mark:                                                                                                                 | The context to use for the request.                                                                                                |
| `request`                                                                                                                          | [operations.RecipientCreateDocumentRecipientsRequestBody](../../models/operations/recipientcreatedocumentrecipientsrequestbody.md) | :heavy_check_mark:                                                                                                                 | The request object to use for the request.                                                                                         |
| `opts`                                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                                           | :heavy_minus_sign:                                                                                                                 | The options for this request.                                                                                                      |

### Response

**[*operations.RecipientCreateDocumentRecipientsResponse](../../models/operations/recipientcreatedocumentrecipientsresponse.md), error**

### Errors

| Error Type                                                                 | Status Code                                                                | Content Type                                                               |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| apierrors.RecipientCreateDocumentRecipientsResponseBody                    | 400                                                                        | application/json                                                           |
| apierrors.RecipientCreateDocumentRecipientsDocumentsRecipientsResponseBody | 500                                                                        | application/json                                                           |
| apierrors.APIError                                                         | 4XX, 5XX                                                                   | \*/\*                                                                      |

## Update

Update a single recipient for a document.

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

    res, err := s.Documents.Recipients.Update(ctx, operations.RecipientUpdateDocumentRecipientRequestBody{
        DocumentID: 8574.78,
        Recipient: operations.RecipientUpdateDocumentRecipientRecipient{
            ID: 5971.29,
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

| Parameter                                                                                                                        | Type                                                                                                                             | Required                                                                                                                         | Description                                                                                                                      |
| -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                            | :heavy_check_mark:                                                                                                               | The context to use for the request.                                                                                              |
| `request`                                                                                                                        | [operations.RecipientUpdateDocumentRecipientRequestBody](../../models/operations/recipientupdatedocumentrecipientrequestbody.md) | :heavy_check_mark:                                                                                                               | The request object to use for the request.                                                                                       |
| `opts`                                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                                         | :heavy_minus_sign:                                                                                                               | The options for this request.                                                                                                    |

### Response

**[*operations.RecipientUpdateDocumentRecipientResponse](../../models/operations/recipientupdatedocumentrecipientresponse.md), error**

### Errors

| Error Type                                                                | Status Code                                                               | Content Type                                                              |
| ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- |
| apierrors.RecipientUpdateDocumentRecipientResponseBody                    | 400                                                                       | application/json                                                          |
| apierrors.RecipientUpdateDocumentRecipientDocumentsRecipientsResponseBody | 500                                                                       | application/json                                                          |
| apierrors.APIError                                                        | 4XX, 5XX                                                                  | \*/\*                                                                     |

## UpdateMany

Update multiple recipients for a document.

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

    res, err := s.Documents.Recipients.UpdateMany(ctx, operations.RecipientUpdateDocumentRecipientsRequestBody{
        DocumentID: 4057.69,
        Recipients: []operations.RecipientUpdateDocumentRecipientsRecipients{
            operations.RecipientUpdateDocumentRecipientsRecipients{
                ID: 5359.16,
            },
            operations.RecipientUpdateDocumentRecipientsRecipients{
                ID: 8982.15,
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

| Parameter                                                                                                                          | Type                                                                                                                               | Required                                                                                                                           | Description                                                                                                                        |
| ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                              | :heavy_check_mark:                                                                                                                 | The context to use for the request.                                                                                                |
| `request`                                                                                                                          | [operations.RecipientUpdateDocumentRecipientsRequestBody](../../models/operations/recipientupdatedocumentrecipientsrequestbody.md) | :heavy_check_mark:                                                                                                                 | The request object to use for the request.                                                                                         |
| `opts`                                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                                           | :heavy_minus_sign:                                                                                                                 | The options for this request.                                                                                                      |

### Response

**[*operations.RecipientUpdateDocumentRecipientsResponse](../../models/operations/recipientupdatedocumentrecipientsresponse.md), error**

### Errors

| Error Type                                                                 | Status Code                                                                | Content Type                                                               |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| apierrors.RecipientUpdateDocumentRecipientsResponseBody                    | 400                                                                        | application/json                                                           |
| apierrors.RecipientUpdateDocumentRecipientsDocumentsRecipientsResponseBody | 500                                                                        | application/json                                                           |
| apierrors.APIError                                                         | 4XX, 5XX                                                                   | \*/\*                                                                      |

## Delete

Delete document recipient

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

    res, err := s.Documents.Recipients.Delete(ctx, operations.RecipientDeleteDocumentRecipientRequestBody{
        RecipientID: 5459.07,
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

| Parameter                                                                                                                        | Type                                                                                                                             | Required                                                                                                                         | Description                                                                                                                      |
| -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                            | :heavy_check_mark:                                                                                                               | The context to use for the request.                                                                                              |
| `request`                                                                                                                        | [operations.RecipientDeleteDocumentRecipientRequestBody](../../models/operations/recipientdeletedocumentrecipientrequestbody.md) | :heavy_check_mark:                                                                                                               | The request object to use for the request.                                                                                       |
| `opts`                                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                                         | :heavy_minus_sign:                                                                                                               | The options for this request.                                                                                                    |

### Response

**[*operations.RecipientDeleteDocumentRecipientResponse](../../models/operations/recipientdeletedocumentrecipientresponse.md), error**

### Errors

| Error Type                                                                | Status Code                                                               | Content Type                                                              |
| ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- |
| apierrors.RecipientDeleteDocumentRecipientResponseBody                    | 400                                                                       | application/json                                                          |
| apierrors.RecipientDeleteDocumentRecipientDocumentsRecipientsResponseBody | 500                                                                       | application/json                                                          |
| apierrors.APIError                                                        | 4XX, 5XX                                                                  | \*/\*                                                                     |