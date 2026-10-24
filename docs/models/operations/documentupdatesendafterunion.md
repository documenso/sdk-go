# DocumentUpdateSendAfterUnion


## Supported Types

### DocumentUpdateSendAfter1

```go
documentUpdateSendAfterUnion := operations.CreateDocumentUpdateSendAfterUnionDocumentUpdateSendAfter1(operations.DocumentUpdateSendAfter1{/* values here */})
```

### DocumentUpdateSendAfter2

```go
documentUpdateSendAfterUnion := operations.CreateDocumentUpdateSendAfterUnionDocumentUpdateSendAfter2(operations.DocumentUpdateSendAfter2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch documentUpdateSendAfterUnion.Type {
	case operations.DocumentUpdateSendAfterUnionTypeDocumentUpdateSendAfter1:
		// documentUpdateSendAfterUnion.DocumentUpdateSendAfter1 is populated
	case operations.DocumentUpdateSendAfterUnionTypeDocumentUpdateSendAfter2:
		// documentUpdateSendAfterUnion.DocumentUpdateSendAfter2 is populated
}
```
