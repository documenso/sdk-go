# Document
(*Document*)

## Overview

### Available Operations

* [DocumentDownload](#documentdownload) - Download document (beta)

## DocumentDownload

Get a pre-signed download URL for the original or signed version of a document

### Example Usage

<!-- UsageSnippet language="go" operationID="document-downloadBeta" method="get" path="/document/{documentId}/download-beta" -->
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

    res, err := s.Document.DocumentDownload(ctx, 9550.11, operations.DocumentDownloadBetaVersionSigned.ToPointer())
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                            | Type                                                                                                                                                 | Required                                                                                                                                             | Description                                                                                                                                          |
| ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                                | :heavy_check_mark:                                                                                                                                   | The context to use for the request.                                                                                                                  |
| `documentID`                                                                                                                                         | *float64*                                                                                                                                            | :heavy_check_mark:                                                                                                                                   | The ID of the document to download.                                                                                                                  |
| `version`                                                                                                                                            | [*operations.DocumentDownloadBetaVersion](../../models/operations/documentdownloadbetaversion.md)                                                    | :heavy_minus_sign:                                                                                                                                   | The version of the document to download. "signed" returns the completed document with signatures, "original" returns the original uploaded document. |
| `opts`                                                                                                                                               | [][operations.Option](../../models/operations/option.md)                                                                                             | :heavy_minus_sign:                                                                                                                                   | The options for this request.                                                                                                                        |

### Response

**[*operations.DocumentDownloadBetaResponse](../../models/operations/documentdownloadbetaresponse.md), error**

### Errors

| Error Type                                        | Status Code                                       | Content Type                                      |
| ------------------------------------------------- | ------------------------------------------------- | ------------------------------------------------- |
| apierrors.DocumentDownloadBetaBadRequestError     | 400                                               | application/json                                  |
| apierrors.DocumentDownloadBetaUnauthorizedError   | 401                                               | application/json                                  |
| apierrors.DocumentDownloadBetaForbiddenError      | 403                                               | application/json                                  |
| apierrors.DocumentDownloadBetaNotFoundError       | 404                                               | application/json                                  |
| apierrors.DocumentDownloadBetaInternalServerError | 500                                               | application/json                                  |
| apierrors.APIError                                | 4XX, 5XX                                          | \*/\*                                             |