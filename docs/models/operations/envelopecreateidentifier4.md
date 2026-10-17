# EnvelopeCreateIdentifier4


## Supported Types

### 

```go
envelopeCreateIdentifier4 := operations.CreateEnvelopeCreateIdentifier4Str(string{/* values here */})
```

### 

```go
envelopeCreateIdentifier4 := operations.CreateEnvelopeCreateIdentifier4Number(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch envelopeCreateIdentifier4.Type {
	case operations.EnvelopeCreateIdentifier4TypeStr:
		// envelopeCreateIdentifier4.Str is populated
	case operations.EnvelopeCreateIdentifier4TypeNumber:
		// envelopeCreateIdentifier4.Number is populated
}
```
