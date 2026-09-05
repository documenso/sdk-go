# EnvelopeCreateIdentifier11


## Supported Types

### 

```go
envelopeCreateIdentifier11 := operations.CreateEnvelopeCreateIdentifier11Str(string{/* values here */})
```

### 

```go
envelopeCreateIdentifier11 := operations.CreateEnvelopeCreateIdentifier11Number(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch envelopeCreateIdentifier11.Type {
	case operations.EnvelopeCreateIdentifier11TypeStr:
		// envelopeCreateIdentifier11.Str is populated
	case operations.EnvelopeCreateIdentifier11TypeNumber:
		// envelopeCreateIdentifier11.Number is populated
}
```
