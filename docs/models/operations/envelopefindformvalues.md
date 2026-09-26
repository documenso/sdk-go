# EnvelopeFindFormValues


## Supported Types

### 

```go
envelopeFindFormValues := operations.CreateEnvelopeFindFormValuesStr(string{/* values here */})
```

### 

```go
envelopeFindFormValues := operations.CreateEnvelopeFindFormValuesBoolean(bool{/* values here */})
```

### 

```go
envelopeFindFormValues := operations.CreateEnvelopeFindFormValuesNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch envelopeFindFormValues.Type {
	case operations.EnvelopeFindFormValuesTypeStr:
		// envelopeFindFormValues.Str is populated
	case operations.EnvelopeFindFormValuesTypeBoolean:
		// envelopeFindFormValues.Boolean is populated
	case operations.EnvelopeFindFormValuesTypeNumber:
		// envelopeFindFormValues.Number is populated
}
```
