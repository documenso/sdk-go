# DocumentUpdateEnvelopeExpirationPeriodUnion


## Supported Types

### DocumentUpdateEnvelopeExpirationPeriod1

```go
documentUpdateEnvelopeExpirationPeriodUnion := operations.CreateDocumentUpdateEnvelopeExpirationPeriodUnionDocumentUpdateEnvelopeExpirationPeriod1(operations.DocumentUpdateEnvelopeExpirationPeriod1{/* values here */})
```

### DocumentUpdateEnvelopeExpirationPeriod2

```go
documentUpdateEnvelopeExpirationPeriodUnion := operations.CreateDocumentUpdateEnvelopeExpirationPeriodUnionDocumentUpdateEnvelopeExpirationPeriod2(operations.DocumentUpdateEnvelopeExpirationPeriod2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch documentUpdateEnvelopeExpirationPeriodUnion.Type {
	case operations.DocumentUpdateEnvelopeExpirationPeriodUnionTypeDocumentUpdateEnvelopeExpirationPeriod1:
		// documentUpdateEnvelopeExpirationPeriodUnion.DocumentUpdateEnvelopeExpirationPeriod1 is populated
	case operations.DocumentUpdateEnvelopeExpirationPeriodUnionTypeDocumentUpdateEnvelopeExpirationPeriod2:
		// documentUpdateEnvelopeExpirationPeriodUnion.DocumentUpdateEnvelopeExpirationPeriod2 is populated
}
```
