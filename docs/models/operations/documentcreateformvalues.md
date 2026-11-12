# DocumentCreateFormValues


## Supported Types

### 

```go
documentCreateFormValues := operations.CreateDocumentCreateFormValuesStr(string{/* values here */})
```

### 

```go
documentCreateFormValues := operations.CreateDocumentCreateFormValuesBoolean(bool{/* values here */})
```

### 

```go
documentCreateFormValues := operations.CreateDocumentCreateFormValuesNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch documentCreateFormValues.Type {
	case operations.DocumentCreateFormValuesTypeStr:
		// documentCreateFormValues.Str is populated
	case operations.DocumentCreateFormValuesTypeBoolean:
		// documentCreateFormValues.Boolean is populated
	case operations.DocumentCreateFormValuesTypeNumber:
		// documentCreateFormValues.Number is populated
}
```
