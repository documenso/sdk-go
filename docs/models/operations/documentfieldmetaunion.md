# DocumentFieldMetaUnion


## Supported Types

### FieldMetaDocumentSignature

```go
documentFieldMetaUnion := operations.CreateDocumentFieldMetaUnionFieldMetaDocumentSignature(operations.FieldMetaDocumentSignature{/* values here */})
```

### FieldMetaDocumentInitials

```go
documentFieldMetaUnion := operations.CreateDocumentFieldMetaUnionFieldMetaDocumentInitials(operations.FieldMetaDocumentInitials{/* values here */})
```

### FieldMetaDocumentName

```go
documentFieldMetaUnion := operations.CreateDocumentFieldMetaUnionFieldMetaDocumentName(operations.FieldMetaDocumentName{/* values here */})
```

### FieldMetaDocumentEmail

```go
documentFieldMetaUnion := operations.CreateDocumentFieldMetaUnionFieldMetaDocumentEmail(operations.FieldMetaDocumentEmail{/* values here */})
```

### FieldMetaDocumentDate

```go
documentFieldMetaUnion := operations.CreateDocumentFieldMetaUnionFieldMetaDocumentDate(operations.FieldMetaDocumentDate{/* values here */})
```

### FieldMetaDocumentText

```go
documentFieldMetaUnion := operations.CreateDocumentFieldMetaUnionFieldMetaDocumentText(operations.FieldMetaDocumentText{/* values here */})
```

### FieldMetaDocumentNumber

```go
documentFieldMetaUnion := operations.CreateDocumentFieldMetaUnionFieldMetaDocumentNumber(operations.FieldMetaDocumentNumber{/* values here */})
```

### FieldMetaDocumentRadio

```go
documentFieldMetaUnion := operations.CreateDocumentFieldMetaUnionFieldMetaDocumentRadio(operations.FieldMetaDocumentRadio{/* values here */})
```

### FieldMetaDocumentCheckbox

```go
documentFieldMetaUnion := operations.CreateDocumentFieldMetaUnionFieldMetaDocumentCheckbox(operations.FieldMetaDocumentCheckbox{/* values here */})
```

### FieldMetaDocumentDropdown

```go
documentFieldMetaUnion := operations.CreateDocumentFieldMetaUnionFieldMetaDocumentDropdown(operations.FieldMetaDocumentDropdown{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch documentFieldMetaUnion.Type {
	case operations.DocumentFieldMetaUnionTypeFieldMetaDocumentSignature:
		// documentFieldMetaUnion.FieldMetaDocumentSignature is populated
	case operations.DocumentFieldMetaUnionTypeFieldMetaDocumentInitials:
		// documentFieldMetaUnion.FieldMetaDocumentInitials is populated
	case operations.DocumentFieldMetaUnionTypeFieldMetaDocumentName:
		// documentFieldMetaUnion.FieldMetaDocumentName is populated
	case operations.DocumentFieldMetaUnionTypeFieldMetaDocumentEmail:
		// documentFieldMetaUnion.FieldMetaDocumentEmail is populated
	case operations.DocumentFieldMetaUnionTypeFieldMetaDocumentDate:
		// documentFieldMetaUnion.FieldMetaDocumentDate is populated
	case operations.DocumentFieldMetaUnionTypeFieldMetaDocumentText:
		// documentFieldMetaUnion.FieldMetaDocumentText is populated
	case operations.DocumentFieldMetaUnionTypeFieldMetaDocumentNumber:
		// documentFieldMetaUnion.FieldMetaDocumentNumber is populated
	case operations.DocumentFieldMetaUnionTypeFieldMetaDocumentRadio:
		// documentFieldMetaUnion.FieldMetaDocumentRadio is populated
	case operations.DocumentFieldMetaUnionTypeFieldMetaDocumentCheckbox:
		// documentFieldMetaUnion.FieldMetaDocumentCheckbox is populated
	case operations.DocumentFieldMetaUnionTypeFieldMetaDocumentDropdown:
		// documentFieldMetaUnion.FieldMetaDocumentDropdown is populated
}
```
