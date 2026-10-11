# DocumentCreateEnvelopeExpirationPeriodUnion


## Supported Types

### DocumentCreateEnvelopeExpirationPeriod1

```go
documentCreateEnvelopeExpirationPeriodUnion := operations.CreateDocumentCreateEnvelopeExpirationPeriodUnionDocumentCreateEnvelopeExpirationPeriod1(operations.DocumentCreateEnvelopeExpirationPeriod1{/* values here */})
```

### DocumentCreateEnvelopeExpirationPeriod2

```go
documentCreateEnvelopeExpirationPeriodUnion := operations.CreateDocumentCreateEnvelopeExpirationPeriodUnionDocumentCreateEnvelopeExpirationPeriod2(operations.DocumentCreateEnvelopeExpirationPeriod2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch documentCreateEnvelopeExpirationPeriodUnion.Type {
	case operations.DocumentCreateEnvelopeExpirationPeriodUnionTypeDocumentCreateEnvelopeExpirationPeriod1:
		// documentCreateEnvelopeExpirationPeriodUnion.DocumentCreateEnvelopeExpirationPeriod1 is populated
	case operations.DocumentCreateEnvelopeExpirationPeriodUnionTypeDocumentCreateEnvelopeExpirationPeriod2:
		// documentCreateEnvelopeExpirationPeriodUnion.DocumentCreateEnvelopeExpirationPeriod2 is populated
}
```
