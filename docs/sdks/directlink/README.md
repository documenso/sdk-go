# DirectLink
(*Templates.DirectLink*)

## Overview

### Available Operations

* [Create](#create) - Create direct link
* [Delete](#delete) - Delete direct link
* [Toggle](#toggle) - Toggle direct link

## Create

Create a direct link for a template

### Example Usage

<!-- UsageSnippet language="go" operationID="template-createTemplateDirectLink" method="post" path="/template/direct/create" -->
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

    res, err := s.Templates.DirectLink.Create(ctx, operations.TemplateCreateTemplateDirectLinkRequest{
        TemplateID: 5094.31,
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
| `request`                                                                                                                | [operations.TemplateCreateTemplateDirectLinkRequest](../../models/operations/templatecreatetemplatedirectlinkrequest.md) | :heavy_check_mark:                                                                                                       | The request object to use for the request.                                                                               |
| `opts`                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                 | :heavy_minus_sign:                                                                                                       | The options for this request.                                                                                            |

### Response

**[*operations.TemplateCreateTemplateDirectLinkResponse](../../models/operations/templatecreatetemplatedirectlinkresponse.md), error**

### Errors

| Error Type                                                    | Status Code                                                   | Content Type                                                  |
| ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- |
| apierrors.TemplateCreateTemplateDirectLinkBadRequestError     | 400                                                           | application/json                                              |
| apierrors.TemplateCreateTemplateDirectLinkInternalServerError | 500                                                           | application/json                                              |
| apierrors.APIError                                            | 4XX, 5XX                                                      | \*/\*                                                         |

## Delete

Delete a direct link for a template

### Example Usage

<!-- UsageSnippet language="go" operationID="template-deleteTemplateDirectLink" method="post" path="/template/direct/delete" -->
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

    res, err := s.Templates.DirectLink.Delete(ctx, operations.TemplateDeleteTemplateDirectLinkRequest{
        TemplateID: 9950.03,
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
| `request`                                                                                                                | [operations.TemplateDeleteTemplateDirectLinkRequest](../../models/operations/templatedeletetemplatedirectlinkrequest.md) | :heavy_check_mark:                                                                                                       | The request object to use for the request.                                                                               |
| `opts`                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                 | :heavy_minus_sign:                                                                                                       | The options for this request.                                                                                            |

### Response

**[*operations.TemplateDeleteTemplateDirectLinkResponse](../../models/operations/templatedeletetemplatedirectlinkresponse.md), error**

### Errors

| Error Type                                                    | Status Code                                                   | Content Type                                                  |
| ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- |
| apierrors.TemplateDeleteTemplateDirectLinkBadRequestError     | 400                                                           | application/json                                              |
| apierrors.TemplateDeleteTemplateDirectLinkInternalServerError | 500                                                           | application/json                                              |
| apierrors.APIError                                            | 4XX, 5XX                                                      | \*/\*                                                         |

## Toggle

Enable or disable a direct link for a template

### Example Usage

<!-- UsageSnippet language="go" operationID="template-toggleTemplateDirectLink" method="post" path="/template/direct/toggle" -->
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

    res, err := s.Templates.DirectLink.Toggle(ctx, operations.TemplateToggleTemplateDirectLinkRequest{
        TemplateID: 6583.54,
        Enabled: false,
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
| `request`                                                                                                                | [operations.TemplateToggleTemplateDirectLinkRequest](../../models/operations/templatetoggletemplatedirectlinkrequest.md) | :heavy_check_mark:                                                                                                       | The request object to use for the request.                                                                               |
| `opts`                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                 | :heavy_minus_sign:                                                                                                       | The options for this request.                                                                                            |

### Response

**[*operations.TemplateToggleTemplateDirectLinkResponse](../../models/operations/templatetoggletemplatedirectlinkresponse.md), error**

### Errors

| Error Type                                                    | Status Code                                                   | Content Type                                                  |
| ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- |
| apierrors.TemplateToggleTemplateDirectLinkBadRequestError     | 400                                                           | application/json                                              |
| apierrors.TemplateToggleTemplateDirectLinkInternalServerError | 500                                                           | application/json                                              |
| apierrors.APIError                                            | 4XX, 5XX                                                      | \*/\*                                                         |