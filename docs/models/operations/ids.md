# Ids


## Supported Types

### IdsEnvelopeID

```go
ids := operations.CreateIdsIdsEnvelopeID(operations.IdsEnvelopeID{/* values here */})
```

### IdsDocumentID

```go
ids := operations.CreateIdsIdsDocumentID(operations.IdsDocumentID{/* values here */})
```

### IdsTemplateID

```go
ids := operations.CreateIdsIdsTemplateID(operations.IdsTemplateID{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch ids.Type {
	case operations.IdsTypeIdsEnvelopeID:
		// ids.IdsEnvelopeID is populated
	case operations.IdsTypeIdsDocumentID:
		// ids.IdsDocumentID is populated
	case operations.IdsTypeIdsTemplateID:
		// ids.IdsTemplateID is populated
}
```
