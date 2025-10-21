<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	sdkgo "github.com/documenso/sdk-go"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := sdkgo.New(
		sdkgo.WithSecurity(os.Getenv("DOCUMENSO_API_KEY")),
	)

	res, err := s.Documents.Get(ctx, 6150.61)
	if err != nil {
		log.Fatal(err)
	}
	if res.Object != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->