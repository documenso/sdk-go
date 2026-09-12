# EnvelopeGetManyEnvelopeExpirationPeriodUnion


## Supported Types

### EnvelopeGetManyEnvelopeExpirationPeriod1

```go
envelopeGetManyEnvelopeExpirationPeriodUnion := operations.CreateEnvelopeGetManyEnvelopeExpirationPeriodUnionEnvelopeGetManyEnvelopeExpirationPeriod1(operations.EnvelopeGetManyEnvelopeExpirationPeriod1{/* values here */})
```

### EnvelopeGetManyEnvelopeExpirationPeriod2

```go
envelopeGetManyEnvelopeExpirationPeriodUnion := operations.CreateEnvelopeGetManyEnvelopeExpirationPeriodUnionEnvelopeGetManyEnvelopeExpirationPeriod2(operations.EnvelopeGetManyEnvelopeExpirationPeriod2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch envelopeGetManyEnvelopeExpirationPeriodUnion.Type {
	case operations.EnvelopeGetManyEnvelopeExpirationPeriodUnionTypeEnvelopeGetManyEnvelopeExpirationPeriod1:
		// envelopeGetManyEnvelopeExpirationPeriodUnion.EnvelopeGetManyEnvelopeExpirationPeriod1 is populated
	case operations.EnvelopeGetManyEnvelopeExpirationPeriodUnionTypeEnvelopeGetManyEnvelopeExpirationPeriod2:
		// envelopeGetManyEnvelopeExpirationPeriodUnion.EnvelopeGetManyEnvelopeExpirationPeriod2 is populated
}
```
