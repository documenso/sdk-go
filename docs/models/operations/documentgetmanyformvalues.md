# DocumentGetManyFormValues


## Supported Types

### 

```go
documentGetManyFormValues := operations.CreateDocumentGetManyFormValuesStr(string{/* values here */})
```

### 

```go
documentGetManyFormValues := operations.CreateDocumentGetManyFormValuesBoolean(bool{/* values here */})
```

### 

```go
documentGetManyFormValues := operations.CreateDocumentGetManyFormValuesNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch documentGetManyFormValues.Type {
	case operations.DocumentGetManyFormValuesTypeStr:
		// documentGetManyFormValues.Str is populated
	case operations.DocumentGetManyFormValuesTypeBoolean:
		// documentGetManyFormValues.Boolean is populated
	case operations.DocumentGetManyFormValuesTypeNumber:
		// documentGetManyFormValues.Number is populated
}
```
