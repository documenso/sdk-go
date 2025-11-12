# DocumentsRecipients
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

<!-- UsageSnippet language="go" operationID="recipient-getDocumentRecipient" method="get" path="/document/recipient/{recipientId}" -->
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

    res, err := s.Documents.Recipients.Get(ctx, 874.3)
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

| Error Type                                                 | Status Code                                                | Content Type                                               |
| ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- |
| apierrors.RecipientGetDocumentRecipientBadRequestError     | 400                                                        | application/json                                           |
| apierrors.RecipientGetDocumentRecipientUnauthorizedError   | 401                                                        | application/json                                           |
| apierrors.RecipientGetDocumentRecipientForbiddenError      | 403                                                        | application/json                                           |
| apierrors.RecipientGetDocumentRecipientNotFoundError       | 404                                                        | application/json                                           |
| apierrors.RecipientGetDocumentRecipientInternalServerError | 500                                                        | application/json                                           |
| apierrors.APIError                                         | 4XX, 5XX                                                   | \*/\*                                                      |

## Create

Create a single recipient for a document.

### Example Usage

<!-- UsageSnippet language="go" operationID="recipient-createDocumentRecipient" method="post" path="/document/recipient/create" -->
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

    res, err := s.Documents.Recipients.Create(ctx, operations.RecipientCreateDocumentRecipientRequest{
        DocumentID: 3058.31,
        Recipient: operations.RecipientCreateDocumentRecipientRecipient{
            Email: "Ila.Steuber@yahoo.com",
            Name: "<value>",
            Role: operations.RecipientCreateDocumentRecipientRoleRequestAssistant,
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

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `request`                                                                                                                | [operations.RecipientCreateDocumentRecipientRequest](../../models/operations/recipientcreatedocumentrecipientrequest.md) | :heavy_check_mark:                                                                                                       | The request object to use for the request.                                                                               |
| `opts`                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                 | :heavy_minus_sign:                                                                                                       | The options for this request.                                                                                            |

### Response

**[*operations.RecipientCreateDocumentRecipientResponse](../../models/operations/recipientcreatedocumentrecipientresponse.md), error**

### Errors

| Error Type                                                    | Status Code                                                   | Content Type                                                  |
| ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- |
| apierrors.RecipientCreateDocumentRecipientBadRequestError     | 400                                                           | application/json                                              |
| apierrors.RecipientCreateDocumentRecipientUnauthorizedError   | 401                                                           | application/json                                              |
| apierrors.RecipientCreateDocumentRecipientForbiddenError      | 403                                                           | application/json                                              |
| apierrors.RecipientCreateDocumentRecipientInternalServerError | 500                                                           | application/json                                              |
| apierrors.APIError                                            | 4XX, 5XX                                                      | \*/\*                                                         |

## CreateMany

Create multiple recipients for a document.

### Example Usage

<!-- UsageSnippet language="go" operationID="recipient-createDocumentRecipients" method="post" path="/document/recipient/create-many" -->
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

    res, err := s.Documents.Recipients.CreateMany(ctx, operations.RecipientCreateDocumentRecipientsRequest{
        DocumentID: 9983.95,
        Recipients: []operations.RecipientCreateDocumentRecipientsRecipientRequest{
            operations.RecipientCreateDocumentRecipientsRecipientRequest{
                Email: "Roosevelt_Baumbach@yahoo.com",
                Name: "<value>",
                Role: operations.RecipientCreateDocumentRecipientsRoleRequestCc,
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

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |
| `request`                                                                                                                  | [operations.RecipientCreateDocumentRecipientsRequest](../../models/operations/recipientcreatedocumentrecipientsrequest.md) | :heavy_check_mark:                                                                                                         | The request object to use for the request.                                                                                 |
| `opts`                                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                                   | :heavy_minus_sign:                                                                                                         | The options for this request.                                                                                              |

### Response

**[*operations.RecipientCreateDocumentRecipientsResponse](../../models/operations/recipientcreatedocumentrecipientsresponse.md), error**

### Errors

| Error Type                                                     | Status Code                                                    | Content Type                                                   |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| apierrors.RecipientCreateDocumentRecipientsBadRequestError     | 400                                                            | application/json                                               |
| apierrors.RecipientCreateDocumentRecipientsUnauthorizedError   | 401                                                            | application/json                                               |
| apierrors.RecipientCreateDocumentRecipientsForbiddenError      | 403                                                            | application/json                                               |
| apierrors.RecipientCreateDocumentRecipientsInternalServerError | 500                                                            | application/json                                               |
| apierrors.APIError                                             | 4XX, 5XX                                                       | \*/\*                                                          |

## Update

Update a single recipient for a document.

### Example Usage

<!-- UsageSnippet language="go" operationID="recipient-updateDocumentRecipient" method="post" path="/document/recipient/update" -->
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

    res, err := s.Documents.Recipients.Update(ctx, operations.RecipientUpdateDocumentRecipientRequest{
        DocumentID: 7045.62,
        Recipient: operations.RecipientUpdateDocumentRecipientRecipient{
            ID: 2224.05,
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

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `request`                                                                                                                | [operations.RecipientUpdateDocumentRecipientRequest](../../models/operations/recipientupdatedocumentrecipientrequest.md) | :heavy_check_mark:                                                                                                       | The request object to use for the request.                                                                               |
| `opts`                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                 | :heavy_minus_sign:                                                                                                       | The options for this request.                                                                                            |

### Response

**[*operations.RecipientUpdateDocumentRecipientResponse](../../models/operations/recipientupdatedocumentrecipientresponse.md), error**

### Errors

| Error Type                                                    | Status Code                                                   | Content Type                                                  |
| ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- |
| apierrors.RecipientUpdateDocumentRecipientBadRequestError     | 400                                                           | application/json                                              |
| apierrors.RecipientUpdateDocumentRecipientUnauthorizedError   | 401                                                           | application/json                                              |
| apierrors.RecipientUpdateDocumentRecipientForbiddenError      | 403                                                           | application/json                                              |
| apierrors.RecipientUpdateDocumentRecipientInternalServerError | 500                                                           | application/json                                              |
| apierrors.APIError                                            | 4XX, 5XX                                                      | \*/\*                                                         |

## UpdateMany

Update multiple recipients for a document.

### Example Usage

<!-- UsageSnippet language="go" operationID="recipient-updateDocumentRecipients" method="post" path="/document/recipient/update-many" -->
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

    res, err := s.Documents.Recipients.UpdateMany(ctx, operations.RecipientUpdateDocumentRecipientsRequest{
        DocumentID: 3189.76,
        Recipients: []operations.RecipientUpdateDocumentRecipientsRecipientRequest{},
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

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |
| `request`                                                                                                                  | [operations.RecipientUpdateDocumentRecipientsRequest](../../models/operations/recipientupdatedocumentrecipientsrequest.md) | :heavy_check_mark:                                                                                                         | The request object to use for the request.                                                                                 |
| `opts`                                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                                   | :heavy_minus_sign:                                                                                                         | The options for this request.                                                                                              |

### Response

**[*operations.RecipientUpdateDocumentRecipientsResponse](../../models/operations/recipientupdatedocumentrecipientsresponse.md), error**

### Errors

| Error Type                                                     | Status Code                                                    | Content Type                                                   |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| apierrors.RecipientUpdateDocumentRecipientsBadRequestError     | 400                                                            | application/json                                               |
| apierrors.RecipientUpdateDocumentRecipientsUnauthorizedError   | 401                                                            | application/json                                               |
| apierrors.RecipientUpdateDocumentRecipientsForbiddenError      | 403                                                            | application/json                                               |
| apierrors.RecipientUpdateDocumentRecipientsInternalServerError | 500                                                            | application/json                                               |
| apierrors.APIError                                             | 4XX, 5XX                                                       | \*/\*                                                          |

## Delete

Delete document recipient

### Example Usage

<!-- UsageSnippet language="go" operationID="recipient-deleteDocumentRecipient" method="post" path="/document/recipient/delete" -->
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

    res, err := s.Documents.Recipients.Delete(ctx, operations.RecipientDeleteDocumentRecipientRequest{
        RecipientID: 5490.43,
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

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `request`                                                                                                                | [operations.RecipientDeleteDocumentRecipientRequest](../../models/operations/recipientdeletedocumentrecipientrequest.md) | :heavy_check_mark:                                                                                                       | The request object to use for the request.                                                                               |
| `opts`                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                 | :heavy_minus_sign:                                                                                                       | The options for this request.                                                                                            |

### Response

**[*operations.RecipientDeleteDocumentRecipientResponse](../../models/operations/recipientdeletedocumentrecipientresponse.md), error**

### Errors

| Error Type                                                    | Status Code                                                   | Content Type                                                  |
| ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- |
| apierrors.RecipientDeleteDocumentRecipientBadRequestError     | 400                                                           | application/json                                              |
| apierrors.RecipientDeleteDocumentRecipientUnauthorizedError   | 401                                                           | application/json                                              |
| apierrors.RecipientDeleteDocumentRecipientForbiddenError      | 403                                                           | application/json                                              |
| apierrors.RecipientDeleteDocumentRecipientInternalServerError | 500                                                           | application/json                                              |
| apierrors.APIError                                            | 4XX, 5XX                                                      | \*/\*                                                         |