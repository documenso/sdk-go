# DocumentUpdateFormValues


## Supported Types

### 

```go
documentUpdateFormValues := operations.CreateDocumentUpdateFormValuesStr(string{/* values here */})
```

### 

```go
documentUpdateFormValues := operations.CreateDocumentUpdateFormValuesBoolean(bool{/* values here */})
```

### 

```go
documentUpdateFormValues := operations.CreateDocumentUpdateFormValuesNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch documentUpdateFormValues.Type {
	case operations.DocumentUpdateFormValuesTypeStr:
		// documentUpdateFormValues.Str is populated
	case operations.DocumentUpdateFormValuesTypeBoolean:
		// documentUpdateFormValues.Boolean is populated
	case operations.DocumentUpdateFormValuesTypeNumber:
		// documentUpdateFormValues.Number is populated
}
```
