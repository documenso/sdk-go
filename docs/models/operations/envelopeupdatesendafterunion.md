# EnvelopeUpdateSendAfterUnion


## Supported Types

### EnvelopeUpdateSendAfter1

```go
envelopeUpdateSendAfterUnion := operations.CreateEnvelopeUpdateSendAfterUnionEnvelopeUpdateSendAfter1(operations.EnvelopeUpdateSendAfter1{/* values here */})
```

### EnvelopeUpdateSendAfter2

```go
envelopeUpdateSendAfterUnion := operations.CreateEnvelopeUpdateSendAfterUnionEnvelopeUpdateSendAfter2(operations.EnvelopeUpdateSendAfter2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch envelopeUpdateSendAfterUnion.Type {
	case operations.EnvelopeUpdateSendAfterUnionTypeEnvelopeUpdateSendAfter1:
		// envelopeUpdateSendAfterUnion.EnvelopeUpdateSendAfter1 is populated
	case operations.EnvelopeUpdateSendAfterUnionTypeEnvelopeUpdateSendAfter2:
		// envelopeUpdateSendAfterUnion.EnvelopeUpdateSendAfter2 is populated
}
```
