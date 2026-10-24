# EnvelopeCreateIdentifier2


## Supported Types

### 

```go
envelopeCreateIdentifier2 := operations.CreateEnvelopeCreateIdentifier2Str(string{/* values here */})
```

### 

```go
envelopeCreateIdentifier2 := operations.CreateEnvelopeCreateIdentifier2Number(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch envelopeCreateIdentifier2.Type {
	case operations.EnvelopeCreateIdentifier2TypeStr:
		// envelopeCreateIdentifier2.Str is populated
	case operations.EnvelopeCreateIdentifier2TypeNumber:
		// envelopeCreateIdentifier2.Number is populated
}
```
