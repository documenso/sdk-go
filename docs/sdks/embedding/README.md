# Embedding
(*Embedding*)

## Overview

### Available Operations

* [EmbeddingPresignCreateEmbeddingPresignToken](#embeddingpresigncreateembeddingpresigntoken) - Create embedding presign token
* [EmbeddingPresignVerifyEmbeddingPresignToken](#embeddingpresignverifyembeddingpresigntoken) - Verify embedding presign token

## EmbeddingPresignCreateEmbeddingPresignToken

Creates a presign token for embedding operations with configurable expiration time

### Example Usage

<!-- UsageSnippet language="go" operationID="embeddingPresign-createEmbeddingPresignToken" method="post" path="/embedding/create-presign-token" -->
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

    res, err := s.Embedding.EmbeddingPresignCreateEmbeddingPresignToken(ctx, operations.EmbeddingPresignCreateEmbeddingPresignTokenRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                      | Type                                                                                                                                           | Required                                                                                                                                       | Description                                                                                                                                    |
| ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                                          | :heavy_check_mark:                                                                                                                             | The context to use for the request.                                                                                                            |
| `request`                                                                                                                                      | [operations.EmbeddingPresignCreateEmbeddingPresignTokenRequest](../../models/operations/embeddingpresigncreateembeddingpresigntokenrequest.md) | :heavy_check_mark:                                                                                                                             | The request object to use for the request.                                                                                                     |
| `opts`                                                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                                                       | :heavy_minus_sign:                                                                                                                             | The options for this request.                                                                                                                  |

### Response

**[*operations.EmbeddingPresignCreateEmbeddingPresignTokenResponse](../../models/operations/embeddingpresigncreateembeddingpresigntokenresponse.md), error**

### Errors

| Error Type                                                               | Status Code                                                              | Content Type                                                             |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| apierrors.EmbeddingPresignCreateEmbeddingPresignTokenBadRequestError     | 400                                                                      | application/json                                                         |
| apierrors.EmbeddingPresignCreateEmbeddingPresignTokenInternalServerError | 500                                                                      | application/json                                                         |
| apierrors.APIError                                                       | 4XX, 5XX                                                                 | \*/\*                                                                    |

## EmbeddingPresignVerifyEmbeddingPresignToken

Verifies a presign token for embedding operations and returns the associated API token

### Example Usage

<!-- UsageSnippet language="go" operationID="embeddingPresign-verifyEmbeddingPresignToken" method="post" path="/embedding/verify-presign-token" -->
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

    res, err := s.Embedding.EmbeddingPresignVerifyEmbeddingPresignToken(ctx, operations.EmbeddingPresignVerifyEmbeddingPresignTokenRequest{
        Token: "<value>",
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

| Parameter                                                                                                                                      | Type                                                                                                                                           | Required                                                                                                                                       | Description                                                                                                                                    |
| ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                                          | :heavy_check_mark:                                                                                                                             | The context to use for the request.                                                                                                            |
| `request`                                                                                                                                      | [operations.EmbeddingPresignVerifyEmbeddingPresignTokenRequest](../../models/operations/embeddingpresignverifyembeddingpresigntokenrequest.md) | :heavy_check_mark:                                                                                                                             | The request object to use for the request.                                                                                                     |
| `opts`                                                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                                                       | :heavy_minus_sign:                                                                                                                             | The options for this request.                                                                                                                  |

### Response

**[*operations.EmbeddingPresignVerifyEmbeddingPresignTokenResponse](../../models/operations/embeddingpresignverifyembeddingpresigntokenresponse.md), error**

### Errors

| Error Type                                                               | Status Code                                                              | Content Type                                                             |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| apierrors.EmbeddingPresignVerifyEmbeddingPresignTokenBadRequestError     | 400                                                                      | application/json                                                         |
| apierrors.EmbeddingPresignVerifyEmbeddingPresignTokenInternalServerError | 500                                                                      | application/json                                                         |
| apierrors.APIError                                                       | 4XX, 5XX                                                                 | \*/\*                                                                    |