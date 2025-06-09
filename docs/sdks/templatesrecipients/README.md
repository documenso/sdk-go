# TemplatesRecipients
(*Templates.Recipients*)

## Overview

### Available Operations

* [Get](#get) - Get template recipient
* [Create](#create) - Create template recipient
* [CreateMany](#createmany) - Create template recipients
* [Update](#update) - Update template recipient
* [UpdateMany](#updatemany) - Update template recipients
* [Delete](#delete) - Delete template recipient

## Get

Returns a single recipient. If you want to retrieve all the recipients for a template, use the "Get Template" endpoint.

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

    res, err := s.Templates.Recipients.Get(ctx, 9436.42)
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

**[*operations.RecipientGetTemplateRecipientResponse](../../models/operations/recipientgettemplaterecipientresponse.md), error**

### Errors

| Error Type                                                 | Status Code                                                | Content Type                                               |
| ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- |
| apierrors.RecipientGetTemplateRecipientBadRequestError     | 400                                                        | application/json                                           |
| apierrors.RecipientGetTemplateRecipientNotFoundError       | 404                                                        | application/json                                           |
| apierrors.RecipientGetTemplateRecipientInternalServerError | 500                                                        | application/json                                           |
| apierrors.APIError                                         | 4XX, 5XX                                                   | \*/\*                                                      |

## Create

Create a single recipient for a template.

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

    res, err := s.Templates.Recipients.Create(ctx, operations.RecipientCreateTemplateRecipientRequest{
        TemplateID: 5712.95,
        Recipient: operations.RecipientCreateTemplateRecipientRecipient{
            Email: "Gerhard88@yahoo.com",
            Name: "<value>",
            Role: operations.RecipientCreateTemplateRecipientRoleRequestBodySigner,
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
| `request`                                                                                                                | [operations.RecipientCreateTemplateRecipientRequest](../../models/operations/recipientcreatetemplaterecipientrequest.md) | :heavy_check_mark:                                                                                                       | The request object to use for the request.                                                                               |
| `opts`                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                 | :heavy_minus_sign:                                                                                                       | The options for this request.                                                                                            |

### Response

**[*operations.RecipientCreateTemplateRecipientResponse](../../models/operations/recipientcreatetemplaterecipientresponse.md), error**

### Errors

| Error Type                                                    | Status Code                                                   | Content Type                                                  |
| ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- |
| apierrors.RecipientCreateTemplateRecipientBadRequestError     | 400                                                           | application/json                                              |
| apierrors.RecipientCreateTemplateRecipientInternalServerError | 500                                                           | application/json                                              |
| apierrors.APIError                                            | 4XX, 5XX                                                      | \*/\*                                                         |

## CreateMany

Create multiple recipients for a template.

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

    res, err := s.Templates.Recipients.CreateMany(ctx, operations.RecipientCreateTemplateRecipientsRequest{
        TemplateID: 5642.48,
        Recipients: []operations.RecipientCreateTemplateRecipientsRecipientRequestBody{},
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
| `request`                                                                                                                  | [operations.RecipientCreateTemplateRecipientsRequest](../../models/operations/recipientcreatetemplaterecipientsrequest.md) | :heavy_check_mark:                                                                                                         | The request object to use for the request.                                                                                 |
| `opts`                                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                                   | :heavy_minus_sign:                                                                                                         | The options for this request.                                                                                              |

### Response

**[*operations.RecipientCreateTemplateRecipientsResponse](../../models/operations/recipientcreatetemplaterecipientsresponse.md), error**

### Errors

| Error Type                                                     | Status Code                                                    | Content Type                                                   |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| apierrors.RecipientCreateTemplateRecipientsBadRequestError     | 400                                                            | application/json                                               |
| apierrors.RecipientCreateTemplateRecipientsInternalServerError | 500                                                            | application/json                                               |
| apierrors.APIError                                             | 4XX, 5XX                                                       | \*/\*                                                          |

## Update

Update a single recipient for a template.

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

    res, err := s.Templates.Recipients.Update(ctx, operations.RecipientUpdateTemplateRecipientRequest{
        TemplateID: 2984.61,
        Recipient: operations.RecipientUpdateTemplateRecipientRecipient{
            ID: 8617.99,
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
| `request`                                                                                                                | [operations.RecipientUpdateTemplateRecipientRequest](../../models/operations/recipientupdatetemplaterecipientrequest.md) | :heavy_check_mark:                                                                                                       | The request object to use for the request.                                                                               |
| `opts`                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                 | :heavy_minus_sign:                                                                                                       | The options for this request.                                                                                            |

### Response

**[*operations.RecipientUpdateTemplateRecipientResponse](../../models/operations/recipientupdatetemplaterecipientresponse.md), error**

### Errors

| Error Type                                                    | Status Code                                                   | Content Type                                                  |
| ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- |
| apierrors.RecipientUpdateTemplateRecipientBadRequestError     | 400                                                           | application/json                                              |
| apierrors.RecipientUpdateTemplateRecipientInternalServerError | 500                                                           | application/json                                              |
| apierrors.APIError                                            | 4XX, 5XX                                                      | \*/\*                                                         |

## UpdateMany

Update multiple recipients for a template.

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

    res, err := s.Templates.Recipients.UpdateMany(ctx, operations.RecipientUpdateTemplateRecipientsRequest{
        TemplateID: 5597.58,
        Recipients: []operations.RecipientUpdateTemplateRecipientsRecipientRequestBody{
            operations.RecipientUpdateTemplateRecipientsRecipientRequestBody{
                ID: 1630.42,
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
| `request`                                                                                                                  | [operations.RecipientUpdateTemplateRecipientsRequest](../../models/operations/recipientupdatetemplaterecipientsrequest.md) | :heavy_check_mark:                                                                                                         | The request object to use for the request.                                                                                 |
| `opts`                                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                                   | :heavy_minus_sign:                                                                                                         | The options for this request.                                                                                              |

### Response

**[*operations.RecipientUpdateTemplateRecipientsResponse](../../models/operations/recipientupdatetemplaterecipientsresponse.md), error**

### Errors

| Error Type                                                     | Status Code                                                    | Content Type                                                   |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| apierrors.RecipientUpdateTemplateRecipientsBadRequestError     | 400                                                            | application/json                                               |
| apierrors.RecipientUpdateTemplateRecipientsInternalServerError | 500                                                            | application/json                                               |
| apierrors.APIError                                             | 4XX, 5XX                                                       | \*/\*                                                          |

## Delete

Delete template recipient

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

    res, err := s.Templates.Recipients.Delete(ctx, operations.RecipientDeleteTemplateRecipientRequest{
        RecipientID: 312.69,
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
| `request`                                                                                                                | [operations.RecipientDeleteTemplateRecipientRequest](../../models/operations/recipientdeletetemplaterecipientrequest.md) | :heavy_check_mark:                                                                                                       | The request object to use for the request.                                                                               |
| `opts`                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                 | :heavy_minus_sign:                                                                                                       | The options for this request.                                                                                            |

### Response

**[*operations.RecipientDeleteTemplateRecipientResponse](../../models/operations/recipientdeletetemplaterecipientresponse.md), error**

### Errors

| Error Type                                                    | Status Code                                                   | Content Type                                                  |
| ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- |
| apierrors.RecipientDeleteTemplateRecipientBadRequestError     | 400                                                           | application/json                                              |
| apierrors.RecipientDeleteTemplateRecipientInternalServerError | 500                                                           | application/json                                              |
| apierrors.APIError                                            | 4XX, 5XX                                                      | \*/\*                                                         |