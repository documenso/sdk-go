# EnvelopeCreateIdentifier5


## Supported Types

### 

```go
envelopeCreateIdentifier5 := operations.CreateEnvelopeCreateIdentifier5Str(string{/* values here */})
```

### 

```go
envelopeCreateIdentifier5 := operations.CreateEnvelopeCreateIdentifier5Number(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch envelopeCreateIdentifier5.Type {
	case operations.EnvelopeCreateIdentifier5TypeStr:
		// envelopeCreateIdentifier5.Str is populated
	case operations.EnvelopeCreateIdentifier5TypeNumber:
		// envelopeCreateIdentifier5.Number is populated
}
```
