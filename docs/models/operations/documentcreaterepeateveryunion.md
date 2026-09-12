# DocumentCreateRepeatEveryUnion


## Supported Types

### DocumentCreateRepeatEvery1

```go
documentCreateRepeatEveryUnion := operations.CreateDocumentCreateRepeatEveryUnionDocumentCreateRepeatEvery1(operations.DocumentCreateRepeatEvery1{/* values here */})
```

### DocumentCreateRepeatEvery2

```go
documentCreateRepeatEveryUnion := operations.CreateDocumentCreateRepeatEveryUnionDocumentCreateRepeatEvery2(operations.DocumentCreateRepeatEvery2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch documentCreateRepeatEveryUnion.Type {
	case operations.DocumentCreateRepeatEveryUnionTypeDocumentCreateRepeatEvery1:
		// documentCreateRepeatEveryUnion.DocumentCreateRepeatEvery1 is populated
	case operations.DocumentCreateRepeatEveryUnionTypeDocumentCreateRepeatEvery2:
		// documentCreateRepeatEveryUnion.DocumentCreateRepeatEvery2 is populated
}
```
