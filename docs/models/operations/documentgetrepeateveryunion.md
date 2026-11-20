# DocumentGetRepeatEveryUnion


## Supported Types

### DocumentGetRepeatEvery1

```go
documentGetRepeatEveryUnion := operations.CreateDocumentGetRepeatEveryUnionDocumentGetRepeatEvery1(operations.DocumentGetRepeatEvery1{/* values here */})
```

### DocumentGetRepeatEvery2

```go
documentGetRepeatEveryUnion := operations.CreateDocumentGetRepeatEveryUnionDocumentGetRepeatEvery2(operations.DocumentGetRepeatEvery2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch documentGetRepeatEveryUnion.Type {
	case operations.DocumentGetRepeatEveryUnionTypeDocumentGetRepeatEvery1:
		// documentGetRepeatEveryUnion.DocumentGetRepeatEvery1 is populated
	case operations.DocumentGetRepeatEveryUnionTypeDocumentGetRepeatEvery2:
		// documentGetRepeatEveryUnion.DocumentGetRepeatEvery2 is populated
}
```
