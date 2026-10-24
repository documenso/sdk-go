# EnvelopeUseIdentifier


## Supported Types

### 

```go
envelopeUseIdentifier := operations.CreateEnvelopeUseIdentifierStr(string{/* values here */})
```

### 

```go
envelopeUseIdentifier := operations.CreateEnvelopeUseIdentifierNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch envelopeUseIdentifier.Type {
	case operations.EnvelopeUseIdentifierTypeStr:
		// envelopeUseIdentifier.Str is populated
	case operations.EnvelopeUseIdentifierTypeNumber:
		// envelopeUseIdentifier.Number is populated
}
```
