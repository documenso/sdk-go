# DocumentCreateSendAfterUnion


## Supported Types

### DocumentCreateSendAfter1

```go
documentCreateSendAfterUnion := operations.CreateDocumentCreateSendAfterUnionDocumentCreateSendAfter1(operations.DocumentCreateSendAfter1{/* values here */})
```

### DocumentCreateSendAfter2

```go
documentCreateSendAfterUnion := operations.CreateDocumentCreateSendAfterUnionDocumentCreateSendAfter2(operations.DocumentCreateSendAfter2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch documentCreateSendAfterUnion.Type {
	case operations.DocumentCreateSendAfterUnionTypeDocumentCreateSendAfter1:
		// documentCreateSendAfterUnion.DocumentCreateSendAfter1 is populated
	case operations.DocumentCreateSendAfterUnionTypeDocumentCreateSendAfter2:
		// documentCreateSendAfterUnion.DocumentCreateSendAfter2 is populated
}
```
