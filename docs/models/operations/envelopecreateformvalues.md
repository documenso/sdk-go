# EnvelopeCreateFormValues


## Supported Types

### 

```go
envelopeCreateFormValues := operations.CreateEnvelopeCreateFormValuesStr(string{/* values here */})
```

### 

```go
envelopeCreateFormValues := operations.CreateEnvelopeCreateFormValuesBoolean(bool{/* values here */})
```

### 

```go
envelopeCreateFormValues := operations.CreateEnvelopeCreateFormValuesNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch envelopeCreateFormValues.Type {
	case operations.EnvelopeCreateFormValuesTypeStr:
		// envelopeCreateFormValues.Str is populated
	case operations.EnvelopeCreateFormValuesTypeBoolean:
		// envelopeCreateFormValues.Boolean is populated
	case operations.EnvelopeCreateFormValuesTypeNumber:
		// envelopeCreateFormValues.Number is populated
}
```
