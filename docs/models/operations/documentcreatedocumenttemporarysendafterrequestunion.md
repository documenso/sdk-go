# DocumentCreateDocumentTemporarySendAfterRequestUnion


## Supported Types

### DocumentCreateDocumentTemporarySendAfterRequest1

```go
documentCreateDocumentTemporarySendAfterRequestUnion := operations.CreateDocumentCreateDocumentTemporarySendAfterRequestUnionDocumentCreateDocumentTemporarySendAfterRequest1(operations.DocumentCreateDocumentTemporarySendAfterRequest1{/* values here */})
```

### DocumentCreateDocumentTemporarySendAfterRequest2

```go
documentCreateDocumentTemporarySendAfterRequestUnion := operations.CreateDocumentCreateDocumentTemporarySendAfterRequestUnionDocumentCreateDocumentTemporarySendAfterRequest2(operations.DocumentCreateDocumentTemporarySendAfterRequest2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch documentCreateDocumentTemporarySendAfterRequestUnion.Type {
	case operations.DocumentCreateDocumentTemporarySendAfterRequestUnionTypeDocumentCreateDocumentTemporarySendAfterRequest1:
		// documentCreateDocumentTemporarySendAfterRequestUnion.DocumentCreateDocumentTemporarySendAfterRequest1 is populated
	case operations.DocumentCreateDocumentTemporarySendAfterRequestUnionTypeDocumentCreateDocumentTemporarySendAfterRequest2:
		// documentCreateDocumentTemporarySendAfterRequestUnion.DocumentCreateDocumentTemporarySendAfterRequest2 is populated
}
```
