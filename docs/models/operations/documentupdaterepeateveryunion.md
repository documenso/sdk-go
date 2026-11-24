# DocumentUpdateRepeatEveryUnion


## Supported Types

### DocumentUpdateRepeatEvery1

```go
documentUpdateRepeatEveryUnion := operations.CreateDocumentUpdateRepeatEveryUnionDocumentUpdateRepeatEvery1(operations.DocumentUpdateRepeatEvery1{/* values here */})
```

### DocumentUpdateRepeatEvery2

```go
documentUpdateRepeatEveryUnion := operations.CreateDocumentUpdateRepeatEveryUnionDocumentUpdateRepeatEvery2(operations.DocumentUpdateRepeatEvery2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch documentUpdateRepeatEveryUnion.Type {
	case operations.DocumentUpdateRepeatEveryUnionTypeDocumentUpdateRepeatEvery1:
		// documentUpdateRepeatEveryUnion.DocumentUpdateRepeatEvery1 is populated
	case operations.DocumentUpdateRepeatEveryUnionTypeDocumentUpdateRepeatEvery2:
		// documentUpdateRepeatEveryUnion.DocumentUpdateRepeatEvery2 is populated
}
```
