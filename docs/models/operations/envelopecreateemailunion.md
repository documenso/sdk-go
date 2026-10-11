# EnvelopeCreateEmailUnion


## Supported Types

### EnvelopeCreateEmailEnum

```go
envelopeCreateEmailUnion := operations.CreateEnvelopeCreateEmailUnionEnvelopeCreateEmailEnum(operations.EnvelopeCreateEmailEnum{/* values here */})
```

### 

```go
envelopeCreateEmailUnion := operations.CreateEnvelopeCreateEmailUnionStr(string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch envelopeCreateEmailUnion.Type {
	case operations.EnvelopeCreateEmailUnionTypeEnvelopeCreateEmailEnum:
		// envelopeCreateEmailUnion.EnvelopeCreateEmailEnum is populated
	case operations.EnvelopeCreateEmailUnionTypeStr:
		// envelopeCreateEmailUnion.Str is populated
}
```
