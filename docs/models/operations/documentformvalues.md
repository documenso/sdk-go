# DocumentFormValues


## Supported Types

### 

```go
documentFormValues := operations.CreateDocumentFormValuesStr(string{/* values here */})
```

### 

```go
documentFormValues := operations.CreateDocumentFormValuesBoolean(bool{/* values here */})
```

### 

```go
documentFormValues := operations.CreateDocumentFormValuesNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch documentFormValues.Type {
	case operations.DocumentFormValuesTypeStr:
		// documentFormValues.Str is populated
	case operations.DocumentFormValuesTypeBoolean:
		// documentFormValues.Boolean is populated
	case operations.DocumentFormValuesTypeNumber:
		// documentFormValues.Number is populated
}
```
