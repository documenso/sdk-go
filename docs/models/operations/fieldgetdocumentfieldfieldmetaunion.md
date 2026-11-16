# FieldGetDocumentFieldFieldMetaUnion


## Supported Types

### FieldGetDocumentFieldFieldMetaSignature

```go
fieldGetDocumentFieldFieldMetaUnion := operations.CreateFieldGetDocumentFieldFieldMetaUnionFieldGetDocumentFieldFieldMetaSignature(operations.FieldGetDocumentFieldFieldMetaSignature{/* values here */})
```

### FieldGetDocumentFieldFieldMetaInitials

```go
fieldGetDocumentFieldFieldMetaUnion := operations.CreateFieldGetDocumentFieldFieldMetaUnionFieldGetDocumentFieldFieldMetaInitials(operations.FieldGetDocumentFieldFieldMetaInitials{/* values here */})
```

### FieldGetDocumentFieldFieldMetaName

```go
fieldGetDocumentFieldFieldMetaUnion := operations.CreateFieldGetDocumentFieldFieldMetaUnionFieldGetDocumentFieldFieldMetaName(operations.FieldGetDocumentFieldFieldMetaName{/* values here */})
```

### FieldGetDocumentFieldFieldMetaEmail

```go
fieldGetDocumentFieldFieldMetaUnion := operations.CreateFieldGetDocumentFieldFieldMetaUnionFieldGetDocumentFieldFieldMetaEmail(operations.FieldGetDocumentFieldFieldMetaEmail{/* values here */})
```

### FieldGetDocumentFieldFieldMetaDate

```go
fieldGetDocumentFieldFieldMetaUnion := operations.CreateFieldGetDocumentFieldFieldMetaUnionFieldGetDocumentFieldFieldMetaDate(operations.FieldGetDocumentFieldFieldMetaDate{/* values here */})
```

### FieldGetDocumentFieldFieldMetaText

```go
fieldGetDocumentFieldFieldMetaUnion := operations.CreateFieldGetDocumentFieldFieldMetaUnionFieldGetDocumentFieldFieldMetaText(operations.FieldGetDocumentFieldFieldMetaText{/* values here */})
```

### FieldGetDocumentFieldFieldMetaNumber

```go
fieldGetDocumentFieldFieldMetaUnion := operations.CreateFieldGetDocumentFieldFieldMetaUnionFieldGetDocumentFieldFieldMetaNumber(operations.FieldGetDocumentFieldFieldMetaNumber{/* values here */})
```

### FieldGetDocumentFieldFieldMetaRadio

```go
fieldGetDocumentFieldFieldMetaUnion := operations.CreateFieldGetDocumentFieldFieldMetaUnionFieldGetDocumentFieldFieldMetaRadio(operations.FieldGetDocumentFieldFieldMetaRadio{/* values here */})
```

### FieldGetDocumentFieldFieldMetaCheckbox

```go
fieldGetDocumentFieldFieldMetaUnion := operations.CreateFieldGetDocumentFieldFieldMetaUnionFieldGetDocumentFieldFieldMetaCheckbox(operations.FieldGetDocumentFieldFieldMetaCheckbox{/* values here */})
```

### FieldGetDocumentFieldFieldMetaDropdown

```go
fieldGetDocumentFieldFieldMetaUnion := operations.CreateFieldGetDocumentFieldFieldMetaUnionFieldGetDocumentFieldFieldMetaDropdown(operations.FieldGetDocumentFieldFieldMetaDropdown{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch fieldGetDocumentFieldFieldMetaUnion.Type {
	case operations.FieldGetDocumentFieldFieldMetaUnionTypeFieldGetDocumentFieldFieldMetaSignature:
		// fieldGetDocumentFieldFieldMetaUnion.FieldGetDocumentFieldFieldMetaSignature is populated
	case operations.FieldGetDocumentFieldFieldMetaUnionTypeFieldGetDocumentFieldFieldMetaInitials:
		// fieldGetDocumentFieldFieldMetaUnion.FieldGetDocumentFieldFieldMetaInitials is populated
	case operations.FieldGetDocumentFieldFieldMetaUnionTypeFieldGetDocumentFieldFieldMetaName:
		// fieldGetDocumentFieldFieldMetaUnion.FieldGetDocumentFieldFieldMetaName is populated
	case operations.FieldGetDocumentFieldFieldMetaUnionTypeFieldGetDocumentFieldFieldMetaEmail:
		// fieldGetDocumentFieldFieldMetaUnion.FieldGetDocumentFieldFieldMetaEmail is populated
	case operations.FieldGetDocumentFieldFieldMetaUnionTypeFieldGetDocumentFieldFieldMetaDate:
		// fieldGetDocumentFieldFieldMetaUnion.FieldGetDocumentFieldFieldMetaDate is populated
	case operations.FieldGetDocumentFieldFieldMetaUnionTypeFieldGetDocumentFieldFieldMetaText:
		// fieldGetDocumentFieldFieldMetaUnion.FieldGetDocumentFieldFieldMetaText is populated
	case operations.FieldGetDocumentFieldFieldMetaUnionTypeFieldGetDocumentFieldFieldMetaNumber:
		// fieldGetDocumentFieldFieldMetaUnion.FieldGetDocumentFieldFieldMetaNumber is populated
	case operations.FieldGetDocumentFieldFieldMetaUnionTypeFieldGetDocumentFieldFieldMetaRadio:
		// fieldGetDocumentFieldFieldMetaUnion.FieldGetDocumentFieldFieldMetaRadio is populated
	case operations.FieldGetDocumentFieldFieldMetaUnionTypeFieldGetDocumentFieldFieldMetaCheckbox:
		// fieldGetDocumentFieldFieldMetaUnion.FieldGetDocumentFieldFieldMetaCheckbox is populated
	case operations.FieldGetDocumentFieldFieldMetaUnionTypeFieldGetDocumentFieldFieldMetaDropdown:
		// fieldGetDocumentFieldFieldMetaUnion.FieldGetDocumentFieldFieldMetaDropdown is populated
}
```
