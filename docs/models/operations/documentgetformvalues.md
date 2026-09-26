# DocumentGetFormValues


## Supported Types

### 

```go
documentGetFormValues := operations.CreateDocumentGetFormValuesStr(string{/* values here */})
```

### 

```go
documentGetFormValues := operations.CreateDocumentGetFormValuesBoolean(bool{/* values here */})
```

### 

```go
documentGetFormValues := operations.CreateDocumentGetFormValuesNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch documentGetFormValues.Type {
	case operations.DocumentGetFormValuesTypeStr:
		// documentGetFormValues.Str is populated
	case operations.DocumentGetFormValuesTypeBoolean:
		// documentGetFormValues.Boolean is populated
	case operations.DocumentGetFormValuesTypeNumber:
		// documentGetFormValues.Number is populated
}
```
