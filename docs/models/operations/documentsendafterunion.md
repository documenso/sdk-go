# DocumentSendAfterUnion


## Supported Types

### SendAfterDocument1

```go
documentSendAfterUnion := operations.CreateDocumentSendAfterUnionSendAfterDocument1(operations.SendAfterDocument1{/* values here */})
```

### SendAfterDocument2

```go
documentSendAfterUnion := operations.CreateDocumentSendAfterUnionSendAfterDocument2(operations.SendAfterDocument2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch documentSendAfterUnion.Type {
	case operations.DocumentSendAfterUnionTypeSendAfterDocument1:
		// documentSendAfterUnion.SendAfterDocument1 is populated
	case operations.DocumentSendAfterUnionTypeSendAfterDocument2:
		// documentSendAfterUnion.SendAfterDocument2 is populated
}
```
