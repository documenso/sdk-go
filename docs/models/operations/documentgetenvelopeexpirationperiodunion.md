# DocumentGetEnvelopeExpirationPeriodUnion


## Supported Types

### DocumentGetEnvelopeExpirationPeriod1

```go
documentGetEnvelopeExpirationPeriodUnion := operations.CreateDocumentGetEnvelopeExpirationPeriodUnionDocumentGetEnvelopeExpirationPeriod1(operations.DocumentGetEnvelopeExpirationPeriod1{/* values here */})
```

### DocumentGetEnvelopeExpirationPeriod2

```go
documentGetEnvelopeExpirationPeriodUnion := operations.CreateDocumentGetEnvelopeExpirationPeriodUnionDocumentGetEnvelopeExpirationPeriod2(operations.DocumentGetEnvelopeExpirationPeriod2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch documentGetEnvelopeExpirationPeriodUnion.Type {
	case operations.DocumentGetEnvelopeExpirationPeriodUnionTypeDocumentGetEnvelopeExpirationPeriod1:
		// documentGetEnvelopeExpirationPeriodUnion.DocumentGetEnvelopeExpirationPeriod1 is populated
	case operations.DocumentGetEnvelopeExpirationPeriodUnionTypeDocumentGetEnvelopeExpirationPeriod2:
		// documentGetEnvelopeExpirationPeriodUnion.DocumentGetEnvelopeExpirationPeriod2 is populated
}
```
