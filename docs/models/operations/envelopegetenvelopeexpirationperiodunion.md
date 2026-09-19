# EnvelopeGetEnvelopeExpirationPeriodUnion


## Supported Types

### EnvelopeGetEnvelopeExpirationPeriod1

```go
envelopeGetEnvelopeExpirationPeriodUnion := operations.CreateEnvelopeGetEnvelopeExpirationPeriodUnionEnvelopeGetEnvelopeExpirationPeriod1(operations.EnvelopeGetEnvelopeExpirationPeriod1{/* values here */})
```

### EnvelopeGetEnvelopeExpirationPeriod2

```go
envelopeGetEnvelopeExpirationPeriodUnion := operations.CreateEnvelopeGetEnvelopeExpirationPeriodUnionEnvelopeGetEnvelopeExpirationPeriod2(operations.EnvelopeGetEnvelopeExpirationPeriod2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch envelopeGetEnvelopeExpirationPeriodUnion.Type {
	case operations.EnvelopeGetEnvelopeExpirationPeriodUnionTypeEnvelopeGetEnvelopeExpirationPeriod1:
		// envelopeGetEnvelopeExpirationPeriodUnion.EnvelopeGetEnvelopeExpirationPeriod1 is populated
	case operations.EnvelopeGetEnvelopeExpirationPeriodUnionTypeEnvelopeGetEnvelopeExpirationPeriod2:
		// envelopeGetEnvelopeExpirationPeriodUnion.EnvelopeGetEnvelopeExpirationPeriod2 is populated
}
```
