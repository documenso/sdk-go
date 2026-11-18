# FieldCreateDocumentFieldsFieldUnion


## Supported Types

### FieldCreateDocumentFieldsFieldSignature

```go
fieldCreateDocumentFieldsFieldUnion := operations.CreateFieldCreateDocumentFieldsFieldUnionFieldCreateDocumentFieldsFieldSignature(operations.FieldCreateDocumentFieldsFieldSignature{/* values here */})
```

### FieldCreateDocumentFieldsFieldFreeSignature

```go
fieldCreateDocumentFieldsFieldUnion := operations.CreateFieldCreateDocumentFieldsFieldUnionFieldCreateDocumentFieldsFieldFreeSignature(operations.FieldCreateDocumentFieldsFieldFreeSignature{/* values here */})
```

### FieldCreateDocumentFieldsFieldInitials

```go
fieldCreateDocumentFieldsFieldUnion := operations.CreateFieldCreateDocumentFieldsFieldUnionFieldCreateDocumentFieldsFieldInitials(operations.FieldCreateDocumentFieldsFieldInitials{/* values here */})
```

### FieldCreateDocumentFieldsFieldName

```go
fieldCreateDocumentFieldsFieldUnion := operations.CreateFieldCreateDocumentFieldsFieldUnionFieldCreateDocumentFieldsFieldName(operations.FieldCreateDocumentFieldsFieldName{/* values here */})
```

### FieldCreateDocumentFieldsFieldEmail

```go
fieldCreateDocumentFieldsFieldUnion := operations.CreateFieldCreateDocumentFieldsFieldUnionFieldCreateDocumentFieldsFieldEmail(operations.FieldCreateDocumentFieldsFieldEmail{/* values here */})
```

### FieldCreateDocumentFieldsFieldDate

```go
fieldCreateDocumentFieldsFieldUnion := operations.CreateFieldCreateDocumentFieldsFieldUnionFieldCreateDocumentFieldsFieldDate(operations.FieldCreateDocumentFieldsFieldDate{/* values here */})
```

### FieldCreateDocumentFieldsFieldText

```go
fieldCreateDocumentFieldsFieldUnion := operations.CreateFieldCreateDocumentFieldsFieldUnionFieldCreateDocumentFieldsFieldText(operations.FieldCreateDocumentFieldsFieldText{/* values here */})
```

### FieldCreateDocumentFieldsFieldNumber

```go
fieldCreateDocumentFieldsFieldUnion := operations.CreateFieldCreateDocumentFieldsFieldUnionFieldCreateDocumentFieldsFieldNumber(operations.FieldCreateDocumentFieldsFieldNumber{/* values here */})
```

### FieldCreateDocumentFieldsFieldRadio

```go
fieldCreateDocumentFieldsFieldUnion := operations.CreateFieldCreateDocumentFieldsFieldUnionFieldCreateDocumentFieldsFieldRadio(operations.FieldCreateDocumentFieldsFieldRadio{/* values here */})
```

### FieldCreateDocumentFieldsFieldCheckbox

```go
fieldCreateDocumentFieldsFieldUnion := operations.CreateFieldCreateDocumentFieldsFieldUnionFieldCreateDocumentFieldsFieldCheckbox(operations.FieldCreateDocumentFieldsFieldCheckbox{/* values here */})
```

### FieldCreateDocumentFieldsFieldDropdown

```go
fieldCreateDocumentFieldsFieldUnion := operations.CreateFieldCreateDocumentFieldsFieldUnionFieldCreateDocumentFieldsFieldDropdown(operations.FieldCreateDocumentFieldsFieldDropdown{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch fieldCreateDocumentFieldsFieldUnion.Type {
	case operations.FieldCreateDocumentFieldsFieldUnionTypeFieldCreateDocumentFieldsFieldSignature:
		// fieldCreateDocumentFieldsFieldUnion.FieldCreateDocumentFieldsFieldSignature is populated
	case operations.FieldCreateDocumentFieldsFieldUnionTypeFieldCreateDocumentFieldsFieldFreeSignature:
		// fieldCreateDocumentFieldsFieldUnion.FieldCreateDocumentFieldsFieldFreeSignature is populated
	case operations.FieldCreateDocumentFieldsFieldUnionTypeFieldCreateDocumentFieldsFieldInitials:
		// fieldCreateDocumentFieldsFieldUnion.FieldCreateDocumentFieldsFieldInitials is populated
	case operations.FieldCreateDocumentFieldsFieldUnionTypeFieldCreateDocumentFieldsFieldName:
		// fieldCreateDocumentFieldsFieldUnion.FieldCreateDocumentFieldsFieldName is populated
	case operations.FieldCreateDocumentFieldsFieldUnionTypeFieldCreateDocumentFieldsFieldEmail:
		// fieldCreateDocumentFieldsFieldUnion.FieldCreateDocumentFieldsFieldEmail is populated
	case operations.FieldCreateDocumentFieldsFieldUnionTypeFieldCreateDocumentFieldsFieldDate:
		// fieldCreateDocumentFieldsFieldUnion.FieldCreateDocumentFieldsFieldDate is populated
	case operations.FieldCreateDocumentFieldsFieldUnionTypeFieldCreateDocumentFieldsFieldText:
		// fieldCreateDocumentFieldsFieldUnion.FieldCreateDocumentFieldsFieldText is populated
	case operations.FieldCreateDocumentFieldsFieldUnionTypeFieldCreateDocumentFieldsFieldNumber:
		// fieldCreateDocumentFieldsFieldUnion.FieldCreateDocumentFieldsFieldNumber is populated
	case operations.FieldCreateDocumentFieldsFieldUnionTypeFieldCreateDocumentFieldsFieldRadio:
		// fieldCreateDocumentFieldsFieldUnion.FieldCreateDocumentFieldsFieldRadio is populated
	case operations.FieldCreateDocumentFieldsFieldUnionTypeFieldCreateDocumentFieldsFieldCheckbox:
		// fieldCreateDocumentFieldsFieldUnion.FieldCreateDocumentFieldsFieldCheckbox is populated
	case operations.FieldCreateDocumentFieldsFieldUnionTypeFieldCreateDocumentFieldsFieldDropdown:
		// fieldCreateDocumentFieldsFieldUnion.FieldCreateDocumentFieldsFieldDropdown is populated
}
```
