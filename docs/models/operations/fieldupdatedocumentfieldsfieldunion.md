# FieldUpdateDocumentFieldsFieldUnion


## Supported Types

### FieldUpdateDocumentFieldsFieldSignature

```go
fieldUpdateDocumentFieldsFieldUnion := operations.CreateFieldUpdateDocumentFieldsFieldUnionFieldUpdateDocumentFieldsFieldSignature(operations.FieldUpdateDocumentFieldsFieldSignature{/* values here */})
```

### FieldUpdateDocumentFieldsFieldFreeSignature

```go
fieldUpdateDocumentFieldsFieldUnion := operations.CreateFieldUpdateDocumentFieldsFieldUnionFieldUpdateDocumentFieldsFieldFreeSignature(operations.FieldUpdateDocumentFieldsFieldFreeSignature{/* values here */})
```

### FieldUpdateDocumentFieldsFieldInitials

```go
fieldUpdateDocumentFieldsFieldUnion := operations.CreateFieldUpdateDocumentFieldsFieldUnionFieldUpdateDocumentFieldsFieldInitials(operations.FieldUpdateDocumentFieldsFieldInitials{/* values here */})
```

### FieldUpdateDocumentFieldsFieldName

```go
fieldUpdateDocumentFieldsFieldUnion := operations.CreateFieldUpdateDocumentFieldsFieldUnionFieldUpdateDocumentFieldsFieldName(operations.FieldUpdateDocumentFieldsFieldName{/* values here */})
```

### FieldUpdateDocumentFieldsFieldEmail

```go
fieldUpdateDocumentFieldsFieldUnion := operations.CreateFieldUpdateDocumentFieldsFieldUnionFieldUpdateDocumentFieldsFieldEmail(operations.FieldUpdateDocumentFieldsFieldEmail{/* values here */})
```

### FieldUpdateDocumentFieldsFieldDate

```go
fieldUpdateDocumentFieldsFieldUnion := operations.CreateFieldUpdateDocumentFieldsFieldUnionFieldUpdateDocumentFieldsFieldDate(operations.FieldUpdateDocumentFieldsFieldDate{/* values here */})
```

### FieldUpdateDocumentFieldsFieldText

```go
fieldUpdateDocumentFieldsFieldUnion := operations.CreateFieldUpdateDocumentFieldsFieldUnionFieldUpdateDocumentFieldsFieldText(operations.FieldUpdateDocumentFieldsFieldText{/* values here */})
```

### FieldUpdateDocumentFieldsFieldNumber

```go
fieldUpdateDocumentFieldsFieldUnion := operations.CreateFieldUpdateDocumentFieldsFieldUnionFieldUpdateDocumentFieldsFieldNumber(operations.FieldUpdateDocumentFieldsFieldNumber{/* values here */})
```

### FieldUpdateDocumentFieldsFieldRadio

```go
fieldUpdateDocumentFieldsFieldUnion := operations.CreateFieldUpdateDocumentFieldsFieldUnionFieldUpdateDocumentFieldsFieldRadio(operations.FieldUpdateDocumentFieldsFieldRadio{/* values here */})
```

### FieldUpdateDocumentFieldsFieldCheckbox

```go
fieldUpdateDocumentFieldsFieldUnion := operations.CreateFieldUpdateDocumentFieldsFieldUnionFieldUpdateDocumentFieldsFieldCheckbox(operations.FieldUpdateDocumentFieldsFieldCheckbox{/* values here */})
```

### FieldUpdateDocumentFieldsFieldDropdown

```go
fieldUpdateDocumentFieldsFieldUnion := operations.CreateFieldUpdateDocumentFieldsFieldUnionFieldUpdateDocumentFieldsFieldDropdown(operations.FieldUpdateDocumentFieldsFieldDropdown{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch fieldUpdateDocumentFieldsFieldUnion.Type {
	case operations.FieldUpdateDocumentFieldsFieldUnionTypeFieldUpdateDocumentFieldsFieldSignature:
		// fieldUpdateDocumentFieldsFieldUnion.FieldUpdateDocumentFieldsFieldSignature is populated
	case operations.FieldUpdateDocumentFieldsFieldUnionTypeFieldUpdateDocumentFieldsFieldFreeSignature:
		// fieldUpdateDocumentFieldsFieldUnion.FieldUpdateDocumentFieldsFieldFreeSignature is populated
	case operations.FieldUpdateDocumentFieldsFieldUnionTypeFieldUpdateDocumentFieldsFieldInitials:
		// fieldUpdateDocumentFieldsFieldUnion.FieldUpdateDocumentFieldsFieldInitials is populated
	case operations.FieldUpdateDocumentFieldsFieldUnionTypeFieldUpdateDocumentFieldsFieldName:
		// fieldUpdateDocumentFieldsFieldUnion.FieldUpdateDocumentFieldsFieldName is populated
	case operations.FieldUpdateDocumentFieldsFieldUnionTypeFieldUpdateDocumentFieldsFieldEmail:
		// fieldUpdateDocumentFieldsFieldUnion.FieldUpdateDocumentFieldsFieldEmail is populated
	case operations.FieldUpdateDocumentFieldsFieldUnionTypeFieldUpdateDocumentFieldsFieldDate:
		// fieldUpdateDocumentFieldsFieldUnion.FieldUpdateDocumentFieldsFieldDate is populated
	case operations.FieldUpdateDocumentFieldsFieldUnionTypeFieldUpdateDocumentFieldsFieldText:
		// fieldUpdateDocumentFieldsFieldUnion.FieldUpdateDocumentFieldsFieldText is populated
	case operations.FieldUpdateDocumentFieldsFieldUnionTypeFieldUpdateDocumentFieldsFieldNumber:
		// fieldUpdateDocumentFieldsFieldUnion.FieldUpdateDocumentFieldsFieldNumber is populated
	case operations.FieldUpdateDocumentFieldsFieldUnionTypeFieldUpdateDocumentFieldsFieldRadio:
		// fieldUpdateDocumentFieldsFieldUnion.FieldUpdateDocumentFieldsFieldRadio is populated
	case operations.FieldUpdateDocumentFieldsFieldUnionTypeFieldUpdateDocumentFieldsFieldCheckbox:
		// fieldUpdateDocumentFieldsFieldUnion.FieldUpdateDocumentFieldsFieldCheckbox is populated
	case operations.FieldUpdateDocumentFieldsFieldUnionTypeFieldUpdateDocumentFieldsFieldDropdown:
		// fieldUpdateDocumentFieldsFieldUnion.FieldUpdateDocumentFieldsFieldDropdown is populated
}
```
