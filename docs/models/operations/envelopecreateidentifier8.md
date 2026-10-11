# EnvelopeCreateIdentifier8


## Supported Types

### 

```go
envelopeCreateIdentifier8 := operations.CreateEnvelopeCreateIdentifier8Str(string{/* values here */})
```

### 

```go
envelopeCreateIdentifier8 := operations.CreateEnvelopeCreateIdentifier8Number(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch envelopeCreateIdentifier8.Type {
	case operations.EnvelopeCreateIdentifier8TypeStr:
		// envelopeCreateIdentifier8.Str is populated
	case operations.EnvelopeCreateIdentifier8TypeNumber:
		// envelopeCreateIdentifier8.Number is populated
}
```
