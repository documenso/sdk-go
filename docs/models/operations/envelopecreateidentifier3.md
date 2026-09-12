# EnvelopeCreateIdentifier3


## Supported Types

### 

```go
envelopeCreateIdentifier3 := operations.CreateEnvelopeCreateIdentifier3Str(string{/* values here */})
```

### 

```go
envelopeCreateIdentifier3 := operations.CreateEnvelopeCreateIdentifier3Number(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch envelopeCreateIdentifier3.Type {
	case operations.EnvelopeCreateIdentifier3TypeStr:
		// envelopeCreateIdentifier3.Str is populated
	case operations.EnvelopeCreateIdentifier3TypeNumber:
		// envelopeCreateIdentifier3.Number is populated
}
```
