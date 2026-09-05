# DocumentCreateDocumentTemporaryFormValuesRequest


## Supported Types

### 

```go
documentCreateDocumentTemporaryFormValuesRequest := operations.CreateDocumentCreateDocumentTemporaryFormValuesRequestStr(string{/* values here */})
```

### 

```go
documentCreateDocumentTemporaryFormValuesRequest := operations.CreateDocumentCreateDocumentTemporaryFormValuesRequestBoolean(bool{/* values here */})
```

### 

```go
documentCreateDocumentTemporaryFormValuesRequest := operations.CreateDocumentCreateDocumentTemporaryFormValuesRequestNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch documentCreateDocumentTemporaryFormValuesRequest.Type {
	case operations.DocumentCreateDocumentTemporaryFormValuesRequestTypeStr:
		// documentCreateDocumentTemporaryFormValuesRequest.Str is populated
	case operations.DocumentCreateDocumentTemporaryFormValuesRequestTypeBoolean:
		// documentCreateDocumentTemporaryFormValuesRequest.Boolean is populated
	case operations.DocumentCreateDocumentTemporaryFormValuesRequestTypeNumber:
		// documentCreateDocumentTemporaryFormValuesRequest.Number is populated
}
```
