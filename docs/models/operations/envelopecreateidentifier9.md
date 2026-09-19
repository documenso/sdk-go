# EnvelopeCreateIdentifier9


## Supported Types

### 

```go
envelopeCreateIdentifier9 := operations.CreateEnvelopeCreateIdentifier9Str(string{/* values here */})
```

### 

```go
envelopeCreateIdentifier9 := operations.CreateEnvelopeCreateIdentifier9Number(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch envelopeCreateIdentifier9.Type {
	case operations.EnvelopeCreateIdentifier9TypeStr:
		// envelopeCreateIdentifier9.Str is populated
	case operations.EnvelopeCreateIdentifier9TypeNumber:
		// envelopeCreateIdentifier9.Number is populated
}
```
