# EnvelopeGetManyFormValues


## Supported Types

### 

```go
envelopeGetManyFormValues := operations.CreateEnvelopeGetManyFormValuesStr(string{/* values here */})
```

### 

```go
envelopeGetManyFormValues := operations.CreateEnvelopeGetManyFormValuesBoolean(bool{/* values here */})
```

### 

```go
envelopeGetManyFormValues := operations.CreateEnvelopeGetManyFormValuesNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch envelopeGetManyFormValues.Type {
	case operations.EnvelopeGetManyFormValuesTypeStr:
		// envelopeGetManyFormValues.Str is populated
	case operations.EnvelopeGetManyFormValuesTypeBoolean:
		// envelopeGetManyFormValues.Boolean is populated
	case operations.EnvelopeGetManyFormValuesTypeNumber:
		// envelopeGetManyFormValues.Number is populated
}
```
