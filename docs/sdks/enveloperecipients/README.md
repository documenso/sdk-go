# EnvelopeRecipients

## Overview

### Available Operations

* [EnvelopeRecipientRejectOnBehalfOf](#enveloperecipientrejectonbehalfof) - Reject envelope recipient on behalf of

## EnvelopeRecipientRejectOnBehalfOf

Records a rejection on behalf of a recipient. Use this when a recipient has declined to sign outside of the platform. The rejection is flagged as external in the document audit log. By default the action is attributed to the API user; supply `actAsEmail` to attribute it to a specific team member.

### Example Usage

<!-- UsageSnippet language="go" operationID="envelope-recipient-rejectOnBehalfOf" method="post" path="/envelope/recipient/{recipientId}/reject" -->
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

    res, err := s.EnvelopeRecipients.EnvelopeRecipientRejectOnBehalfOf(ctx, 51.94, operations.EnvelopeRecipientRejectOnBehalfOfRequestBody{
        EnvelopeID: "<id>",
        Reason: "<value>",
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
| `recipientID`                                                                                                                      | `float64`                                                                                                                          | :heavy_check_mark:                                                                                                                 | The ID of the recipient to reject the document on behalf of.                                                                       |
| `requestBody`                                                                                                                      | [operations.EnvelopeRecipientRejectOnBehalfOfRequestBody](../../models/operations/enveloperecipientrejectonbehalfofrequestbody.md) | :heavy_check_mark:                                                                                                                 | N/A                                                                                                                                |
| `opts`                                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                                           | :heavy_minus_sign:                                                                                                                 | The options for this request.                                                                                                      |

### Response

**[*operations.EnvelopeRecipientRejectOnBehalfOfResponse](../../models/operations/enveloperecipientrejectonbehalfofresponse.md), error**

### Errors

| Error Type                                                     | Status Code                                                    | Content Type                                                   |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| apierrors.EnvelopeRecipientRejectOnBehalfOfBadRequestError     | 400                                                            | application/json                                               |
| apierrors.EnvelopeRecipientRejectOnBehalfOfUnauthorizedError   | 401                                                            | application/json                                               |
| apierrors.EnvelopeRecipientRejectOnBehalfOfForbiddenError      | 403                                                            | application/json                                               |
| apierrors.EnvelopeRecipientRejectOnBehalfOfInternalServerError | 500                                                            | application/json                                               |
| apierrors.APIError                                             | 4XX, 5XX                                                       | \*/\*                                                          |