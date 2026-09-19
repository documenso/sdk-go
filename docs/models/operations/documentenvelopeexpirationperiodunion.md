# DocumentEnvelopeExpirationPeriodUnion


## Supported Types

### EnvelopeExpirationPeriodDocument1

```go
documentEnvelopeExpirationPeriodUnion := operations.CreateDocumentEnvelopeExpirationPeriodUnionEnvelopeExpirationPeriodDocument1(operations.EnvelopeExpirationPeriodDocument1{/* values here */})
```

### EnvelopeExpirationPeriodDocument2

```go
documentEnvelopeExpirationPeriodUnion := operations.CreateDocumentEnvelopeExpirationPeriodUnionEnvelopeExpirationPeriodDocument2(operations.EnvelopeExpirationPeriodDocument2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch documentEnvelopeExpirationPeriodUnion.Type {
	case operations.DocumentEnvelopeExpirationPeriodUnionTypeEnvelopeExpirationPeriodDocument1:
		// documentEnvelopeExpirationPeriodUnion.EnvelopeExpirationPeriodDocument1 is populated
	case operations.DocumentEnvelopeExpirationPeriodUnionTypeEnvelopeExpirationPeriodDocument2:
		// documentEnvelopeExpirationPeriodUnion.EnvelopeExpirationPeriodDocument2 is populated
}
```
