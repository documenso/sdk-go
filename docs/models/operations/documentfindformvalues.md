# DocumentFindFormValues


## Supported Types

### 

```go
documentFindFormValues := operations.CreateDocumentFindFormValuesStr(string{/* values here */})
```

### 

```go
documentFindFormValues := operations.CreateDocumentFindFormValuesBoolean(bool{/* values here */})
```

### 

```go
documentFindFormValues := operations.CreateDocumentFindFormValuesNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch documentFindFormValues.Type {
	case operations.DocumentFindFormValuesTypeStr:
		// documentFindFormValues.Str is populated
	case operations.DocumentFindFormValuesTypeBoolean:
		// documentFindFormValues.Boolean is populated
	case operations.DocumentFindFormValuesTypeNumber:
		// documentFindFormValues.Number is populated
}
```
