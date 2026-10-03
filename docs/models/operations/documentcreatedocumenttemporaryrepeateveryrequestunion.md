# DocumentCreateDocumentTemporaryRepeatEveryRequestUnion


## Supported Types

### DocumentCreateDocumentTemporaryRepeatEveryRequest1

```go
documentCreateDocumentTemporaryRepeatEveryRequestUnion := operations.CreateDocumentCreateDocumentTemporaryRepeatEveryRequestUnionDocumentCreateDocumentTemporaryRepeatEveryRequest1(operations.DocumentCreateDocumentTemporaryRepeatEveryRequest1{/* values here */})
```

### DocumentCreateDocumentTemporaryRepeatEveryRequest2

```go
documentCreateDocumentTemporaryRepeatEveryRequestUnion := operations.CreateDocumentCreateDocumentTemporaryRepeatEveryRequestUnionDocumentCreateDocumentTemporaryRepeatEveryRequest2(operations.DocumentCreateDocumentTemporaryRepeatEveryRequest2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch documentCreateDocumentTemporaryRepeatEveryRequestUnion.Type {
	case operations.DocumentCreateDocumentTemporaryRepeatEveryRequestUnionTypeDocumentCreateDocumentTemporaryRepeatEveryRequest1:
		// documentCreateDocumentTemporaryRepeatEveryRequestUnion.DocumentCreateDocumentTemporaryRepeatEveryRequest1 is populated
	case operations.DocumentCreateDocumentTemporaryRepeatEveryRequestUnionTypeDocumentCreateDocumentTemporaryRepeatEveryRequest2:
		// documentCreateDocumentTemporaryRepeatEveryRequestUnion.DocumentCreateDocumentTemporaryRepeatEveryRequest2 is populated
}
```
