# DocumentDownloadVersion

The version of the document to download. "signed" returns the completed document with signatures, "original" returns the original uploaded document.

## Example Usage

```go
import (
	"github.com/documenso/sdk-go/models/operations"
)

value := operations.DocumentDownloadVersionOriginal
```


## Values

| Name                              | Value                             |
| --------------------------------- | --------------------------------- |
| `DocumentDownloadVersionOriginal` | original                          |
| `DocumentDownloadVersionSigned`   | signed                            |