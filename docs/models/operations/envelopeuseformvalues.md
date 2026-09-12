# EnvelopeUseFormValues


## Supported Types

### 

```go
envelopeUseFormValues := operations.CreateEnvelopeUseFormValuesStr(string{/* values here */})
```

### 

```go
envelopeUseFormValues := operations.CreateEnvelopeUseFormValuesBoolean(bool{/* values here */})
```

### 

```go
envelopeUseFormValues := operations.CreateEnvelopeUseFormValuesNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch envelopeUseFormValues.Type {
	case operations.EnvelopeUseFormValuesTypeStr:
		// envelopeUseFormValues.Str is populated
	case operations.EnvelopeUseFormValuesTypeBoolean:
		// envelopeUseFormValues.Boolean is populated
	case operations.EnvelopeUseFormValuesTypeNumber:
		// envelopeUseFormValues.Number is populated
}
```
