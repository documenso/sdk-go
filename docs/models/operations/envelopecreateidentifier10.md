# EnvelopeCreateIdentifier10


## Supported Types

### 

```go
envelopeCreateIdentifier10 := operations.CreateEnvelopeCreateIdentifier10Str(string{/* values here */})
```

### 

```go
envelopeCreateIdentifier10 := operations.CreateEnvelopeCreateIdentifier10Number(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch envelopeCreateIdentifier10.Type {
	case operations.EnvelopeCreateIdentifier10TypeStr:
		// envelopeCreateIdentifier10.Str is populated
	case operations.EnvelopeCreateIdentifier10TypeNumber:
		// envelopeCreateIdentifier10.Number is populated
}
```
