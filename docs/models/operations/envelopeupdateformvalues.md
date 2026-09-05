# EnvelopeUpdateFormValues


## Supported Types

### 

```go
envelopeUpdateFormValues := operations.CreateEnvelopeUpdateFormValuesStr(string{/* values here */})
```

### 

```go
envelopeUpdateFormValues := operations.CreateEnvelopeUpdateFormValuesBoolean(bool{/* values here */})
```

### 

```go
envelopeUpdateFormValues := operations.CreateEnvelopeUpdateFormValuesNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch envelopeUpdateFormValues.Type {
	case operations.EnvelopeUpdateFormValuesTypeStr:
		// envelopeUpdateFormValues.Str is populated
	case operations.EnvelopeUpdateFormValuesTypeBoolean:
		// envelopeUpdateFormValues.Boolean is populated
	case operations.EnvelopeUpdateFormValuesTypeNumber:
		// envelopeUpdateFormValues.Number is populated
}
```
