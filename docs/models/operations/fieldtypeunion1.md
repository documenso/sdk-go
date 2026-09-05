# FieldTypeUnion1


## Supported Types

### EnvelopeAuditLogFindTypeSignature1

```go
fieldTypeUnion1 := operations.CreateFieldTypeUnion1EnvelopeAuditLogFindTypeSignature1(operations.EnvelopeAuditLogFindTypeSignature1{/* values here */})
```

### EnvelopeAuditLogFindTypeFreeSignature1

```go
fieldTypeUnion1 := operations.CreateFieldTypeUnion1EnvelopeAuditLogFindTypeFreeSignature1(operations.EnvelopeAuditLogFindTypeFreeSignature1{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch fieldTypeUnion1.Type {
	case operations.FieldTypeUnion1TypeEnvelopeAuditLogFindTypeSignature1:
		// fieldTypeUnion1.EnvelopeAuditLogFindTypeSignature1 is populated
	case operations.FieldTypeUnion1TypeEnvelopeAuditLogFindTypeFreeSignature1:
		// fieldTypeUnion1.EnvelopeAuditLogFindTypeFreeSignature1 is populated
}
```
