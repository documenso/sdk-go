# Source


## Supported Types

### EnvelopeAuditLogFindSourceDocument

```go
source := operations.CreateSourceEnvelopeAuditLogFindSourceDocument(operations.EnvelopeAuditLogFindSourceDocument{/* values here */})
```

### SourceTemplate

```go
source := operations.CreateSourceSourceTemplate(operations.SourceTemplate{/* values here */})
```

### SourceTemplateDirectLink

```go
source := operations.CreateSourceSourceTemplateDirectLink(operations.SourceTemplateDirectLink{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch source.Type {
	case operations.SourceTypeEnvelopeAuditLogFindSourceDocument:
		// source.EnvelopeAuditLogFindSourceDocument is populated
	case operations.SourceTypeSourceTemplate:
		// source.SourceTemplate is populated
	case operations.SourceTypeSourceTemplateDirectLink:
		// source.SourceTemplateDirectLink is populated
}
```
