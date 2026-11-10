# EnvelopeGetFormValues


## Supported Types

### 

```go
envelopeGetFormValues := operations.CreateEnvelopeGetFormValuesStr(string{/* values here */})
```

### 

```go
envelopeGetFormValues := operations.CreateEnvelopeGetFormValuesBoolean(bool{/* values here */})
```

### 

```go
envelopeGetFormValues := operations.CreateEnvelopeGetFormValuesNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch envelopeGetFormValues.Type {
	case operations.EnvelopeGetFormValuesTypeStr:
		// envelopeGetFormValues.Str is populated
	case operations.EnvelopeGetFormValuesTypeBoolean:
		// envelopeGetFormValues.Boolean is populated
	case operations.EnvelopeGetFormValuesTypeNumber:
		// envelopeGetFormValues.Number is populated
}
```
