# Templates
(*Templates*)

## Overview

### Available Operations

* [Find](#find) - Find templates
* [Get](#get) - Get template
* [Create](#create) - Create template
* [Update](#update) - Update template
* [Duplicate](#duplicate) - Duplicate template
* [Delete](#delete) - Delete template
* [Use](#use) - Use template

## Find

Find templates based on a search criteria

### Example Usage

<!-- UsageSnippet language="go" operationID="template-findTemplates" method="get" path="/template" -->
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

    res, err := s.Templates.Find(ctx, operations.TemplateFindTemplatesRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.TemplateFindTemplatesRequest](../../models/operations/templatefindtemplatesrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.TemplateFindTemplatesResponse](../../models/operations/templatefindtemplatesresponse.md), error**

### Errors

| Error Type                                         | Status Code                                        | Content Type                                       |
| -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- |
| apierrors.TemplateFindTemplatesBadRequestError     | 400                                                | application/json                                   |
| apierrors.TemplateFindTemplatesUnauthorizedError   | 401                                                | application/json                                   |
| apierrors.TemplateFindTemplatesForbiddenError      | 403                                                | application/json                                   |
| apierrors.TemplateFindTemplatesNotFoundError       | 404                                                | application/json                                   |
| apierrors.TemplateFindTemplatesInternalServerError | 500                                                | application/json                                   |
| apierrors.APIError                                 | 4XX, 5XX                                           | \*/\*                                              |

## Get

Get template

### Example Usage

<!-- UsageSnippet language="go" operationID="template-getTemplateById" method="get" path="/template/{templateId}" -->
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

    res, err := s.Templates.Get(ctx, 2128.54)
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
| `templateID`                                             | *float64*                                                | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.TemplateGetTemplateByIDResponse](../../models/operations/templategettemplatebyidresponse.md), error**

### Errors

| Error Type                                           | Status Code                                          | Content Type                                         |
| ---------------------------------------------------- | ---------------------------------------------------- | ---------------------------------------------------- |
| apierrors.TemplateGetTemplateByIDBadRequestError     | 400                                                  | application/json                                     |
| apierrors.TemplateGetTemplateByIDUnauthorizedError   | 401                                                  | application/json                                     |
| apierrors.TemplateGetTemplateByIDForbiddenError      | 403                                                  | application/json                                     |
| apierrors.TemplateGetTemplateByIDNotFoundError       | 404                                                  | application/json                                     |
| apierrors.TemplateGetTemplateByIDInternalServerError | 500                                                  | application/json                                     |
| apierrors.APIError                                   | 4XX, 5XX                                             | \*/\*                                                |

## Create

Create a new template

### Example Usage

<!-- UsageSnippet language="go" operationID="template-createTemplate" method="post" path="/template/create" -->
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

    example, fileErr := os.Open("example.file")
    if fileErr != nil {
        panic(fileErr)
    }

    res, err := s.Templates.Create(ctx, operations.TemplateCreateTemplateRequest{
        Payload: operations.TemplateCreateTemplatePayload{
            Title: "<value>",
        },
        File: operations.TemplateCreateTemplateFile{
            FileName: "example.file",
            Content: example,
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

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.TemplateCreateTemplateRequest](../../models/operations/templatecreatetemplaterequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.TemplateCreateTemplateResponse](../../models/operations/templatecreatetemplateresponse.md), error**

### Errors

| Error Type                                          | Status Code                                         | Content Type                                        |
| --------------------------------------------------- | --------------------------------------------------- | --------------------------------------------------- |
| apierrors.TemplateCreateTemplateBadRequestError     | 400                                                 | application/json                                    |
| apierrors.TemplateCreateTemplateUnauthorizedError   | 401                                                 | application/json                                    |
| apierrors.TemplateCreateTemplateForbiddenError      | 403                                                 | application/json                                    |
| apierrors.TemplateCreateTemplateInternalServerError | 500                                                 | application/json                                    |
| apierrors.APIError                                  | 4XX, 5XX                                            | \*/\*                                               |

## Update

Update template

### Example Usage

<!-- UsageSnippet language="go" operationID="template-updateTemplate" method="post" path="/template/update" -->
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

    res, err := s.Templates.Update(ctx, operations.TemplateUpdateTemplateRequest{
        TemplateID: 9404.77,
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

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.TemplateUpdateTemplateRequest](../../models/operations/templateupdatetemplaterequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.TemplateUpdateTemplateResponse](../../models/operations/templateupdatetemplateresponse.md), error**

### Errors

| Error Type                                          | Status Code                                         | Content Type                                        |
| --------------------------------------------------- | --------------------------------------------------- | --------------------------------------------------- |
| apierrors.TemplateUpdateTemplateBadRequestError     | 400                                                 | application/json                                    |
| apierrors.TemplateUpdateTemplateUnauthorizedError   | 401                                                 | application/json                                    |
| apierrors.TemplateUpdateTemplateForbiddenError      | 403                                                 | application/json                                    |
| apierrors.TemplateUpdateTemplateInternalServerError | 500                                                 | application/json                                    |
| apierrors.APIError                                  | 4XX, 5XX                                            | \*/\*                                               |

## Duplicate

Duplicate template

### Example Usage

<!-- UsageSnippet language="go" operationID="template-duplicateTemplate" method="post" path="/template/duplicate" -->
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

    res, err := s.Templates.Duplicate(ctx, operations.TemplateDuplicateTemplateRequest{
        TemplateID: 2490.16,
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
| `request`                                                                                                  | [operations.TemplateDuplicateTemplateRequest](../../models/operations/templateduplicatetemplaterequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.TemplateDuplicateTemplateResponse](../../models/operations/templateduplicatetemplateresponse.md), error**

### Errors

| Error Type                                             | Status Code                                            | Content Type                                           |
| ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ |
| apierrors.TemplateDuplicateTemplateBadRequestError     | 400                                                    | application/json                                       |
| apierrors.TemplateDuplicateTemplateUnauthorizedError   | 401                                                    | application/json                                       |
| apierrors.TemplateDuplicateTemplateForbiddenError      | 403                                                    | application/json                                       |
| apierrors.TemplateDuplicateTemplateInternalServerError | 500                                                    | application/json                                       |
| apierrors.APIError                                     | 4XX, 5XX                                               | \*/\*                                                  |

## Delete

Delete template

### Example Usage

<!-- UsageSnippet language="go" operationID="template-deleteTemplate" method="post" path="/template/delete" -->
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

    res, err := s.Templates.Delete(ctx, operations.TemplateDeleteTemplateRequest{
        TemplateID: 536.89,
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

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.TemplateDeleteTemplateRequest](../../models/operations/templatedeletetemplaterequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.TemplateDeleteTemplateResponse](../../models/operations/templatedeletetemplateresponse.md), error**

### Errors

| Error Type                                          | Status Code                                         | Content Type                                        |
| --------------------------------------------------- | --------------------------------------------------- | --------------------------------------------------- |
| apierrors.TemplateDeleteTemplateBadRequestError     | 400                                                 | application/json                                    |
| apierrors.TemplateDeleteTemplateUnauthorizedError   | 401                                                 | application/json                                    |
| apierrors.TemplateDeleteTemplateForbiddenError      | 403                                                 | application/json                                    |
| apierrors.TemplateDeleteTemplateInternalServerError | 500                                                 | application/json                                    |
| apierrors.APIError                                  | 4XX, 5XX                                            | \*/\*                                               |

## Use

Use the template to create a document

### Example Usage

<!-- UsageSnippet language="go" operationID="template-createDocumentFromTemplate" method="post" path="/template/use" -->
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

    res, err := s.Templates.Use(ctx, operations.TemplateCreateDocumentFromTemplateRequest{
        TemplateID: 7392.96,
        Recipients: []operations.TemplateCreateDocumentFromTemplateRecipientRequest{},
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

| Parameter                                                                                                                    | Type                                                                                                                         | Required                                                                                                                     | Description                                                                                                                  |
| ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                        | :heavy_check_mark:                                                                                                           | The context to use for the request.                                                                                          |
| `request`                                                                                                                    | [operations.TemplateCreateDocumentFromTemplateRequest](../../models/operations/templatecreatedocumentfromtemplaterequest.md) | :heavy_check_mark:                                                                                                           | The request object to use for the request.                                                                                   |
| `opts`                                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                                     | :heavy_minus_sign:                                                                                                           | The options for this request.                                                                                                |

### Response

**[*operations.TemplateCreateDocumentFromTemplateResponse](../../models/operations/templatecreatedocumentfromtemplateresponse.md), error**

### Errors

| Error Type                                                      | Status Code                                                     | Content Type                                                    |
| --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- |
| apierrors.TemplateCreateDocumentFromTemplateBadRequestError     | 400                                                             | application/json                                                |
| apierrors.TemplateCreateDocumentFromTemplateUnauthorizedError   | 401                                                             | application/json                                                |
| apierrors.TemplateCreateDocumentFromTemplateForbiddenError      | 403                                                             | application/json                                                |
| apierrors.TemplateCreateDocumentFromTemplateInternalServerError | 500                                                             | application/json                                                |
| apierrors.APIError                                              | 4XX, 5XX                                                        | \*/\*                                                           |