# EnvelopeUpdateEnvelopeExpirationPeriodUnion


## Supported Types

### EnvelopeUpdateEnvelopeExpirationPeriod1

```go
envelopeUpdateEnvelopeExpirationPeriodUnion := operations.CreateEnvelopeUpdateEnvelopeExpirationPeriodUnionEnvelopeUpdateEnvelopeExpirationPeriod1(operations.EnvelopeUpdateEnvelopeExpirationPeriod1{/* values here */})
```

### EnvelopeUpdateEnvelopeExpirationPeriod2

```go
envelopeUpdateEnvelopeExpirationPeriodUnion := operations.CreateEnvelopeUpdateEnvelopeExpirationPeriodUnionEnvelopeUpdateEnvelopeExpirationPeriod2(operations.EnvelopeUpdateEnvelopeExpirationPeriod2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch envelopeUpdateEnvelopeExpirationPeriodUnion.Type {
	case operations.EnvelopeUpdateEnvelopeExpirationPeriodUnionTypeEnvelopeUpdateEnvelopeExpirationPeriod1:
		// envelopeUpdateEnvelopeExpirationPeriodUnion.EnvelopeUpdateEnvelopeExpirationPeriod1 is populated
	case operations.EnvelopeUpdateEnvelopeExpirationPeriodUnionTypeEnvelopeUpdateEnvelopeExpirationPeriod2:
		// envelopeUpdateEnvelopeExpirationPeriodUnion.EnvelopeUpdateEnvelopeExpirationPeriod2 is populated
}
```
