# EnvelopeCreateEnvelopeExpirationPeriodUnion


## Supported Types

### EnvelopeCreateEnvelopeExpirationPeriod1

```go
envelopeCreateEnvelopeExpirationPeriodUnion := operations.CreateEnvelopeCreateEnvelopeExpirationPeriodUnionEnvelopeCreateEnvelopeExpirationPeriod1(operations.EnvelopeCreateEnvelopeExpirationPeriod1{/* values here */})
```

### EnvelopeCreateEnvelopeExpirationPeriod2

```go
envelopeCreateEnvelopeExpirationPeriodUnion := operations.CreateEnvelopeCreateEnvelopeExpirationPeriodUnionEnvelopeCreateEnvelopeExpirationPeriod2(operations.EnvelopeCreateEnvelopeExpirationPeriod2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch envelopeCreateEnvelopeExpirationPeriodUnion.Type {
	case operations.EnvelopeCreateEnvelopeExpirationPeriodUnionTypeEnvelopeCreateEnvelopeExpirationPeriod1:
		// envelopeCreateEnvelopeExpirationPeriodUnion.EnvelopeCreateEnvelopeExpirationPeriod1 is populated
	case operations.EnvelopeCreateEnvelopeExpirationPeriodUnionTypeEnvelopeCreateEnvelopeExpirationPeriod2:
		// envelopeCreateEnvelopeExpirationPeriodUnion.EnvelopeCreateEnvelopeExpirationPeriod2 is populated
}
```
