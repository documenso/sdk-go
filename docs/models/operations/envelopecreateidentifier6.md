# EnvelopeCreateIdentifier6


## Supported Types

### 

```go
envelopeCreateIdentifier6 := operations.CreateEnvelopeCreateIdentifier6Str(string{/* values here */})
```

### 

```go
envelopeCreateIdentifier6 := operations.CreateEnvelopeCreateIdentifier6Number(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch envelopeCreateIdentifier6.Type {
	case operations.EnvelopeCreateIdentifier6TypeStr:
		// envelopeCreateIdentifier6.Str is populated
	case operations.EnvelopeCreateIdentifier6TypeNumber:
		// envelopeCreateIdentifier6.Number is populated
}
```
