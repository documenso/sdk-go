# EnvelopeUseEmailUnion


## Supported Types

### EnvelopeUseEmailEnum

```go
envelopeUseEmailUnion := operations.CreateEnvelopeUseEmailUnionEnvelopeUseEmailEnum(operations.EnvelopeUseEmailEnum{/* values here */})
```

### 

```go
envelopeUseEmailUnion := operations.CreateEnvelopeUseEmailUnionStr(string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch envelopeUseEmailUnion.Type {
	case operations.EnvelopeUseEmailUnionTypeEnvelopeUseEmailEnum:
		// envelopeUseEmailUnion.EnvelopeUseEmailEnum is populated
	case operations.EnvelopeUseEmailUnionTypeStr:
		// envelopeUseEmailUnion.Str is populated
}
```
