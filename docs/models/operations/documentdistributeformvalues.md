# DocumentDistributeFormValues


## Supported Types

### 

```go
documentDistributeFormValues := operations.CreateDocumentDistributeFormValuesStr(string{/* values here */})
```

### 

```go
documentDistributeFormValues := operations.CreateDocumentDistributeFormValuesBoolean(bool{/* values here */})
```

### 

```go
documentDistributeFormValues := operations.CreateDocumentDistributeFormValuesNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch documentDistributeFormValues.Type {
	case operations.DocumentDistributeFormValuesTypeStr:
		// documentDistributeFormValues.Str is populated
	case operations.DocumentDistributeFormValuesTypeBoolean:
		// documentDistributeFormValues.Boolean is populated
	case operations.DocumentDistributeFormValuesTypeNumber:
		// documentDistributeFormValues.Number is populated
}
```
