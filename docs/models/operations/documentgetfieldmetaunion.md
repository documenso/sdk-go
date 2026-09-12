# DocumentGetFieldMetaUnion


## Supported Types

### DocumentGetFieldMetaSignature

```go
documentGetFieldMetaUnion := operations.CreateDocumentGetFieldMetaUnionDocumentGetFieldMetaSignature(operations.DocumentGetFieldMetaSignature{/* values here */})
```

### DocumentGetFieldMetaInitials

```go
documentGetFieldMetaUnion := operations.CreateDocumentGetFieldMetaUnionDocumentGetFieldMetaInitials(operations.DocumentGetFieldMetaInitials{/* values here */})
```

### DocumentGetFieldMetaName

```go
documentGetFieldMetaUnion := operations.CreateDocumentGetFieldMetaUnionDocumentGetFieldMetaName(operations.DocumentGetFieldMetaName{/* values here */})
```

### DocumentGetFieldMetaEmail

```go
documentGetFieldMetaUnion := operations.CreateDocumentGetFieldMetaUnionDocumentGetFieldMetaEmail(operations.DocumentGetFieldMetaEmail{/* values here */})
```

### DocumentGetFieldMetaDate

```go
documentGetFieldMetaUnion := operations.CreateDocumentGetFieldMetaUnionDocumentGetFieldMetaDate(operations.DocumentGetFieldMetaDate{/* values here */})
```

### DocumentGetFieldMetaText

```go
documentGetFieldMetaUnion := operations.CreateDocumentGetFieldMetaUnionDocumentGetFieldMetaText(operations.DocumentGetFieldMetaText{/* values here */})
```

### DocumentGetFieldMetaNumber

```go
documentGetFieldMetaUnion := operations.CreateDocumentGetFieldMetaUnionDocumentGetFieldMetaNumber(operations.DocumentGetFieldMetaNumber{/* values here */})
```

### DocumentGetFieldMetaRadio

```go
documentGetFieldMetaUnion := operations.CreateDocumentGetFieldMetaUnionDocumentGetFieldMetaRadio(operations.DocumentGetFieldMetaRadio{/* values here */})
```

### DocumentGetFieldMetaCheckbox

```go
documentGetFieldMetaUnion := operations.CreateDocumentGetFieldMetaUnionDocumentGetFieldMetaCheckbox(operations.DocumentGetFieldMetaCheckbox{/* values here */})
```

### DocumentGetFieldMetaDropdown

```go
documentGetFieldMetaUnion := operations.CreateDocumentGetFieldMetaUnionDocumentGetFieldMetaDropdown(operations.DocumentGetFieldMetaDropdown{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch documentGetFieldMetaUnion.Type {
	case operations.DocumentGetFieldMetaUnionTypeDocumentGetFieldMetaSignature:
		// documentGetFieldMetaUnion.DocumentGetFieldMetaSignature is populated
	case operations.DocumentGetFieldMetaUnionTypeDocumentGetFieldMetaInitials:
		// documentGetFieldMetaUnion.DocumentGetFieldMetaInitials is populated
	case operations.DocumentGetFieldMetaUnionTypeDocumentGetFieldMetaName:
		// documentGetFieldMetaUnion.DocumentGetFieldMetaName is populated
	case operations.DocumentGetFieldMetaUnionTypeDocumentGetFieldMetaEmail:
		// documentGetFieldMetaUnion.DocumentGetFieldMetaEmail is populated
	case operations.DocumentGetFieldMetaUnionTypeDocumentGetFieldMetaDate:
		// documentGetFieldMetaUnion.DocumentGetFieldMetaDate is populated
	case operations.DocumentGetFieldMetaUnionTypeDocumentGetFieldMetaText:
		// documentGetFieldMetaUnion.DocumentGetFieldMetaText is populated
	case operations.DocumentGetFieldMetaUnionTypeDocumentGetFieldMetaNumber:
		// documentGetFieldMetaUnion.DocumentGetFieldMetaNumber is populated
	case operations.DocumentGetFieldMetaUnionTypeDocumentGetFieldMetaRadio:
		// documentGetFieldMetaUnion.DocumentGetFieldMetaRadio is populated
	case operations.DocumentGetFieldMetaUnionTypeDocumentGetFieldMetaCheckbox:
		// documentGetFieldMetaUnion.DocumentGetFieldMetaCheckbox is populated
	case operations.DocumentGetFieldMetaUnionTypeDocumentGetFieldMetaDropdown:
		// documentGetFieldMetaUnion.DocumentGetFieldMetaDropdown is populated
}
```
