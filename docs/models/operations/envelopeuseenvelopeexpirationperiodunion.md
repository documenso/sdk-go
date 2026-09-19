# EnvelopeUseEnvelopeExpirationPeriodUnion


## Supported Types

### EnvelopeUseEnvelopeExpirationPeriod1

```go
envelopeUseEnvelopeExpirationPeriodUnion := operations.CreateEnvelopeUseEnvelopeExpirationPeriodUnionEnvelopeUseEnvelopeExpirationPeriod1(operations.EnvelopeUseEnvelopeExpirationPeriod1{/* values here */})
```

### EnvelopeUseEnvelopeExpirationPeriod2

```go
envelopeUseEnvelopeExpirationPeriodUnion := operations.CreateEnvelopeUseEnvelopeExpirationPeriodUnionEnvelopeUseEnvelopeExpirationPeriod2(operations.EnvelopeUseEnvelopeExpirationPeriod2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch envelopeUseEnvelopeExpirationPeriodUnion.Type {
	case operations.EnvelopeUseEnvelopeExpirationPeriodUnionTypeEnvelopeUseEnvelopeExpirationPeriod1:
		// envelopeUseEnvelopeExpirationPeriodUnion.EnvelopeUseEnvelopeExpirationPeriod1 is populated
	case operations.EnvelopeUseEnvelopeExpirationPeriodUnionTypeEnvelopeUseEnvelopeExpirationPeriod2:
		// envelopeUseEnvelopeExpirationPeriodUnion.EnvelopeUseEnvelopeExpirationPeriod2 is populated
}
```
