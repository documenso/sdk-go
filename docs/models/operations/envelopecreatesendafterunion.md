# EnvelopeCreateSendAfterUnion


## Supported Types

### EnvelopeCreateSendAfter1

```go
envelopeCreateSendAfterUnion := operations.CreateEnvelopeCreateSendAfterUnionEnvelopeCreateSendAfter1(operations.EnvelopeCreateSendAfter1{/* values here */})
```

### EnvelopeCreateSendAfter2

```go
envelopeCreateSendAfterUnion := operations.CreateEnvelopeCreateSendAfterUnionEnvelopeCreateSendAfter2(operations.EnvelopeCreateSendAfter2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch envelopeCreateSendAfterUnion.Type {
	case operations.EnvelopeCreateSendAfterUnionTypeEnvelopeCreateSendAfter1:
		// envelopeCreateSendAfterUnion.EnvelopeCreateSendAfter1 is populated
	case operations.EnvelopeCreateSendAfterUnionTypeEnvelopeCreateSendAfter2:
		// envelopeCreateSendAfterUnion.EnvelopeCreateSendAfter2 is populated
}
```
