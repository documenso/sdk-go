# DocumentGetSendAfterUnion


## Supported Types

### DocumentGetSendAfter1

```go
documentGetSendAfterUnion := operations.CreateDocumentGetSendAfterUnionDocumentGetSendAfter1(operations.DocumentGetSendAfter1{/* values here */})
```

### DocumentGetSendAfter2

```go
documentGetSendAfterUnion := operations.CreateDocumentGetSendAfterUnionDocumentGetSendAfter2(operations.DocumentGetSendAfter2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch documentGetSendAfterUnion.Type {
	case operations.DocumentGetSendAfterUnionTypeDocumentGetSendAfter1:
		// documentGetSendAfterUnion.DocumentGetSendAfter1 is populated
	case operations.DocumentGetSendAfterUnionTypeDocumentGetSendAfter2:
		// documentGetSendAfterUnion.DocumentGetSendAfter2 is populated
}
```
