# EnvelopeCreateIdentifier7


## Supported Types

### 

```go
envelopeCreateIdentifier7 := operations.CreateEnvelopeCreateIdentifier7Str(string{/* values here */})
```

### 

```go
envelopeCreateIdentifier7 := operations.CreateEnvelopeCreateIdentifier7Number(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch envelopeCreateIdentifier7.Type {
	case operations.EnvelopeCreateIdentifier7TypeStr:
		// envelopeCreateIdentifier7.Str is populated
	case operations.EnvelopeCreateIdentifier7TypeNumber:
		// envelopeCreateIdentifier7.Number is populated
}
```
