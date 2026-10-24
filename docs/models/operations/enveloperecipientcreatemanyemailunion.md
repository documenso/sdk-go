# EnvelopeRecipientCreateManyEmailUnion


## Supported Types

### EnvelopeRecipientCreateManyEmailEnum

```go
envelopeRecipientCreateManyEmailUnion := operations.CreateEnvelopeRecipientCreateManyEmailUnionEnvelopeRecipientCreateManyEmailEnum(operations.EnvelopeRecipientCreateManyEmailEnum{/* values here */})
```

### 

```go
envelopeRecipientCreateManyEmailUnion := operations.CreateEnvelopeRecipientCreateManyEmailUnionStr(string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch envelopeRecipientCreateManyEmailUnion.Type {
	case operations.EnvelopeRecipientCreateManyEmailUnionTypeEnvelopeRecipientCreateManyEmailEnum:
		// envelopeRecipientCreateManyEmailUnion.EnvelopeRecipientCreateManyEmailEnum is populated
	case operations.EnvelopeRecipientCreateManyEmailUnionTypeStr:
		// envelopeRecipientCreateManyEmailUnion.Str is populated
}
```
