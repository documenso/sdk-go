# Folders

## Overview

### Available Operations

* [Find](#find) - Find folders
* [Create](#create) - Create new folder
* [Update](#update) - Update folder
* [Delete](#delete) - Delete folder

## Find

Find folders based on a search criteria

### Example Usage

<!-- UsageSnippet language="go" operationID="folder-findFolders" method="get" path="/folder" -->
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

    res, err := s.Folders.Find(ctx, operations.FolderFindFoldersRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.FolderFindFoldersRequest](../../models/operations/folderfindfoldersrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.FolderFindFoldersResponse](../../models/operations/folderfindfoldersresponse.md), error**

### Errors

| Error Type                                     | Status Code                                    | Content Type                                   |
| ---------------------------------------------- | ---------------------------------------------- | ---------------------------------------------- |
| apierrors.FolderFindFoldersBadRequestError     | 400                                            | application/json                               |
| apierrors.FolderFindFoldersUnauthorizedError   | 401                                            | application/json                               |
| apierrors.FolderFindFoldersForbiddenError      | 403                                            | application/json                               |
| apierrors.FolderFindFoldersNotFoundError       | 404                                            | application/json                               |
| apierrors.FolderFindFoldersInternalServerError | 500                                            | application/json                               |
| apierrors.APIError                             | 4XX, 5XX                                       | \*/\*                                          |

## Create

Creates a new folder in your team

### Example Usage

<!-- UsageSnippet language="go" operationID="folder-createFolder" method="post" path="/folder/create" -->
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

    res, err := s.Folders.Create(ctx, operations.FolderCreateFolderRequest{
        Name: "<value>",
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

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.FolderCreateFolderRequest](../../models/operations/foldercreatefolderrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.FolderCreateFolderResponse](../../models/operations/foldercreatefolderresponse.md), error**

### Errors

| Error Type                                      | Status Code                                     | Content Type                                    |
| ----------------------------------------------- | ----------------------------------------------- | ----------------------------------------------- |
| apierrors.FolderCreateFolderBadRequestError     | 400                                             | application/json                                |
| apierrors.FolderCreateFolderUnauthorizedError   | 401                                             | application/json                                |
| apierrors.FolderCreateFolderForbiddenError      | 403                                             | application/json                                |
| apierrors.FolderCreateFolderInternalServerError | 500                                             | application/json                                |
| apierrors.APIError                              | 4XX, 5XX                                        | \*/\*                                           |

## Update

Updates an existing folder

### Example Usage

<!-- UsageSnippet language="go" operationID="folder-updateFolder" method="post" path="/folder/update" -->
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

    res, err := s.Folders.Update(ctx, operations.FolderUpdateFolderRequest{
        FolderID: "<id>",
        Data: operations.FolderUpdateFolderData{},
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

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.FolderUpdateFolderRequest](../../models/operations/folderupdatefolderrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.FolderUpdateFolderResponse](../../models/operations/folderupdatefolderresponse.md), error**

### Errors

| Error Type                                      | Status Code                                     | Content Type                                    |
| ----------------------------------------------- | ----------------------------------------------- | ----------------------------------------------- |
| apierrors.FolderUpdateFolderBadRequestError     | 400                                             | application/json                                |
| apierrors.FolderUpdateFolderUnauthorizedError   | 401                                             | application/json                                |
| apierrors.FolderUpdateFolderForbiddenError      | 403                                             | application/json                                |
| apierrors.FolderUpdateFolderInternalServerError | 500                                             | application/json                                |
| apierrors.APIError                              | 4XX, 5XX                                        | \*/\*                                           |

## Delete

Deletes an existing folder

### Example Usage

<!-- UsageSnippet language="go" operationID="folder-deleteFolder" method="post" path="/folder/delete" -->
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

    res, err := s.Folders.Delete(ctx, operations.FolderDeleteFolderRequest{
        FolderID: "<id>",
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

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.FolderDeleteFolderRequest](../../models/operations/folderdeletefolderrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.FolderDeleteFolderResponse](../../models/operations/folderdeletefolderresponse.md), error**

### Errors

| Error Type                                      | Status Code                                     | Content Type                                    |
| ----------------------------------------------- | ----------------------------------------------- | ----------------------------------------------- |
| apierrors.FolderDeleteFolderBadRequestError     | 400                                             | application/json                                |
| apierrors.FolderDeleteFolderUnauthorizedError   | 401                                             | application/json                                |
| apierrors.FolderDeleteFolderForbiddenError      | 403                                             | application/json                                |
| apierrors.FolderDeleteFolderInternalServerError | 500                                             | application/json                                |
| apierrors.APIError                              | 4XX, 5XX                                        | \*/\*                                           |