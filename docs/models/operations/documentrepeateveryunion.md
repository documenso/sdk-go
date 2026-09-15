# DocumentRepeatEveryUnion


## Supported Types

### RepeatEveryDocument1

```go
documentRepeatEveryUnion := operations.CreateDocumentRepeatEveryUnionRepeatEveryDocument1(operations.RepeatEveryDocument1{/* values here */})
```

### RepeatEveryDocument2

```go
documentRepeatEveryUnion := operations.CreateDocumentRepeatEveryUnionRepeatEveryDocument2(operations.RepeatEveryDocument2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch documentRepeatEveryUnion.Type {
	case operations.DocumentRepeatEveryUnionTypeRepeatEveryDocument1:
		// documentRepeatEveryUnion.RepeatEveryDocument1 is populated
	case operations.DocumentRepeatEveryUnionTypeRepeatEveryDocument2:
		// documentRepeatEveryUnion.RepeatEveryDocument2 is populated
}
```
