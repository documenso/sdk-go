# EnvelopeCreateIdentifier1


## Supported Types

### 

```go
envelopeCreateIdentifier1 := operations.CreateEnvelopeCreateIdentifier1Str(string{/* values here */})
```

### 

```go
envelopeCreateIdentifier1 := operations.CreateEnvelopeCreateIdentifier1Number(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch envelopeCreateIdentifier1.Type {
	case operations.EnvelopeCreateIdentifier1TypeStr:
		// envelopeCreateIdentifier1.Str is populated
	case operations.EnvelopeCreateIdentifier1TypeNumber:
		// envelopeCreateIdentifier1.Number is populated
}
```
