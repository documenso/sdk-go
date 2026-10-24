# FieldTypeUnion2


## Supported Types

### EnvelopeAuditLogFindTypeSignature2

```go
fieldTypeUnion2 := operations.CreateFieldTypeUnion2EnvelopeAuditLogFindTypeSignature2(operations.EnvelopeAuditLogFindTypeSignature2{/* values here */})
```

### EnvelopeAuditLogFindTypeFreeSignature2

```go
fieldTypeUnion2 := operations.CreateFieldTypeUnion2EnvelopeAuditLogFindTypeFreeSignature2(operations.EnvelopeAuditLogFindTypeFreeSignature2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch fieldTypeUnion2.Type {
	case operations.FieldTypeUnion2TypeEnvelopeAuditLogFindTypeSignature2:
		// fieldTypeUnion2.EnvelopeAuditLogFindTypeSignature2 is populated
	case operations.FieldTypeUnion2TypeEnvelopeAuditLogFindTypeFreeSignature2:
		// fieldTypeUnion2.EnvelopeAuditLogFindTypeFreeSignature2 is populated
}
```
