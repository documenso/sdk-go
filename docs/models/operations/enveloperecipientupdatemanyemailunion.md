# EnvelopeRecipientUpdateManyEmailUnion


## Supported Types

### EnvelopeRecipientUpdateManyEmailEnum

```go
envelopeRecipientUpdateManyEmailUnion := operations.CreateEnvelopeRecipientUpdateManyEmailUnionEnvelopeRecipientUpdateManyEmailEnum(operations.EnvelopeRecipientUpdateManyEmailEnum{/* values here */})
```

### 

```go
envelopeRecipientUpdateManyEmailUnion := operations.CreateEnvelopeRecipientUpdateManyEmailUnionStr(string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch envelopeRecipientUpdateManyEmailUnion.Type {
	case operations.EnvelopeRecipientUpdateManyEmailUnionTypeEnvelopeRecipientUpdateManyEmailEnum:
		// envelopeRecipientUpdateManyEmailUnion.EnvelopeRecipientUpdateManyEmailEnum is populated
	case operations.EnvelopeRecipientUpdateManyEmailUnionTypeStr:
		// envelopeRecipientUpdateManyEmailUnion.Str is populated
}
```
