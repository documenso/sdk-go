# FieldUpdateDocumentFieldFieldUnion


## Supported Types

### FieldUpdateDocumentFieldFieldSignature

```go
fieldUpdateDocumentFieldFieldUnion := operations.CreateFieldUpdateDocumentFieldFieldUnionFieldUpdateDocumentFieldFieldSignature(operations.FieldUpdateDocumentFieldFieldSignature{/* values here */})
```

### FieldUpdateDocumentFieldFieldFreeSignature

```go
fieldUpdateDocumentFieldFieldUnion := operations.CreateFieldUpdateDocumentFieldFieldUnionFieldUpdateDocumentFieldFieldFreeSignature(operations.FieldUpdateDocumentFieldFieldFreeSignature{/* values here */})
```

### FieldUpdateDocumentFieldFieldInitials

```go
fieldUpdateDocumentFieldFieldUnion := operations.CreateFieldUpdateDocumentFieldFieldUnionFieldUpdateDocumentFieldFieldInitials(operations.FieldUpdateDocumentFieldFieldInitials{/* values here */})
```

### FieldUpdateDocumentFieldFieldName

```go
fieldUpdateDocumentFieldFieldUnion := operations.CreateFieldUpdateDocumentFieldFieldUnionFieldUpdateDocumentFieldFieldName(operations.FieldUpdateDocumentFieldFieldName{/* values here */})
```

### FieldUpdateDocumentFieldFieldEmail

```go
fieldUpdateDocumentFieldFieldUnion := operations.CreateFieldUpdateDocumentFieldFieldUnionFieldUpdateDocumentFieldFieldEmail(operations.FieldUpdateDocumentFieldFieldEmail{/* values here */})
```

### FieldUpdateDocumentFieldFieldDate

```go
fieldUpdateDocumentFieldFieldUnion := operations.CreateFieldUpdateDocumentFieldFieldUnionFieldUpdateDocumentFieldFieldDate(operations.FieldUpdateDocumentFieldFieldDate{/* values here */})
```

### FieldUpdateDocumentFieldFieldText

```go
fieldUpdateDocumentFieldFieldUnion := operations.CreateFieldUpdateDocumentFieldFieldUnionFieldUpdateDocumentFieldFieldText(operations.FieldUpdateDocumentFieldFieldText{/* values here */})
```

### FieldUpdateDocumentFieldFieldNumber

```go
fieldUpdateDocumentFieldFieldUnion := operations.CreateFieldUpdateDocumentFieldFieldUnionFieldUpdateDocumentFieldFieldNumber(operations.FieldUpdateDocumentFieldFieldNumber{/* values here */})
```

### FieldUpdateDocumentFieldFieldRadio

```go
fieldUpdateDocumentFieldFieldUnion := operations.CreateFieldUpdateDocumentFieldFieldUnionFieldUpdateDocumentFieldFieldRadio(operations.FieldUpdateDocumentFieldFieldRadio{/* values here */})
```

### FieldUpdateDocumentFieldFieldCheckbox

```go
fieldUpdateDocumentFieldFieldUnion := operations.CreateFieldUpdateDocumentFieldFieldUnionFieldUpdateDocumentFieldFieldCheckbox(operations.FieldUpdateDocumentFieldFieldCheckbox{/* values here */})
```

### FieldUpdateDocumentFieldFieldDropdown

```go
fieldUpdateDocumentFieldFieldUnion := operations.CreateFieldUpdateDocumentFieldFieldUnionFieldUpdateDocumentFieldFieldDropdown(operations.FieldUpdateDocumentFieldFieldDropdown{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch fieldUpdateDocumentFieldFieldUnion.Type {
	case operations.FieldUpdateDocumentFieldFieldUnionTypeFieldUpdateDocumentFieldFieldSignature:
		// fieldUpdateDocumentFieldFieldUnion.FieldUpdateDocumentFieldFieldSignature is populated
	case operations.FieldUpdateDocumentFieldFieldUnionTypeFieldUpdateDocumentFieldFieldFreeSignature:
		// fieldUpdateDocumentFieldFieldUnion.FieldUpdateDocumentFieldFieldFreeSignature is populated
	case operations.FieldUpdateDocumentFieldFieldUnionTypeFieldUpdateDocumentFieldFieldInitials:
		// fieldUpdateDocumentFieldFieldUnion.FieldUpdateDocumentFieldFieldInitials is populated
	case operations.FieldUpdateDocumentFieldFieldUnionTypeFieldUpdateDocumentFieldFieldName:
		// fieldUpdateDocumentFieldFieldUnion.FieldUpdateDocumentFieldFieldName is populated
	case operations.FieldUpdateDocumentFieldFieldUnionTypeFieldUpdateDocumentFieldFieldEmail:
		// fieldUpdateDocumentFieldFieldUnion.FieldUpdateDocumentFieldFieldEmail is populated
	case operations.FieldUpdateDocumentFieldFieldUnionTypeFieldUpdateDocumentFieldFieldDate:
		// fieldUpdateDocumentFieldFieldUnion.FieldUpdateDocumentFieldFieldDate is populated
	case operations.FieldUpdateDocumentFieldFieldUnionTypeFieldUpdateDocumentFieldFieldText:
		// fieldUpdateDocumentFieldFieldUnion.FieldUpdateDocumentFieldFieldText is populated
	case operations.FieldUpdateDocumentFieldFieldUnionTypeFieldUpdateDocumentFieldFieldNumber:
		// fieldUpdateDocumentFieldFieldUnion.FieldUpdateDocumentFieldFieldNumber is populated
	case operations.FieldUpdateDocumentFieldFieldUnionTypeFieldUpdateDocumentFieldFieldRadio:
		// fieldUpdateDocumentFieldFieldUnion.FieldUpdateDocumentFieldFieldRadio is populated
	case operations.FieldUpdateDocumentFieldFieldUnionTypeFieldUpdateDocumentFieldFieldCheckbox:
		// fieldUpdateDocumentFieldFieldUnion.FieldUpdateDocumentFieldFieldCheckbox is populated
	case operations.FieldUpdateDocumentFieldFieldUnionTypeFieldUpdateDocumentFieldFieldDropdown:
		// fieldUpdateDocumentFieldFieldUnion.FieldUpdateDocumentFieldFieldDropdown is populated
}
```
