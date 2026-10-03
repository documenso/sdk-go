# FieldCreateDocumentFieldFieldUnion


## Supported Types

### FieldCreateDocumentFieldFieldSignature

```go
fieldCreateDocumentFieldFieldUnion := operations.CreateFieldCreateDocumentFieldFieldUnionFieldCreateDocumentFieldFieldSignature(operations.FieldCreateDocumentFieldFieldSignature{/* values here */})
```

### FieldCreateDocumentFieldFieldFreeSignature

```go
fieldCreateDocumentFieldFieldUnion := operations.CreateFieldCreateDocumentFieldFieldUnionFieldCreateDocumentFieldFieldFreeSignature(operations.FieldCreateDocumentFieldFieldFreeSignature{/* values here */})
```

### FieldCreateDocumentFieldFieldInitials

```go
fieldCreateDocumentFieldFieldUnion := operations.CreateFieldCreateDocumentFieldFieldUnionFieldCreateDocumentFieldFieldInitials(operations.FieldCreateDocumentFieldFieldInitials{/* values here */})
```

### FieldCreateDocumentFieldFieldName

```go
fieldCreateDocumentFieldFieldUnion := operations.CreateFieldCreateDocumentFieldFieldUnionFieldCreateDocumentFieldFieldName(operations.FieldCreateDocumentFieldFieldName{/* values here */})
```

### FieldCreateDocumentFieldFieldEmail

```go
fieldCreateDocumentFieldFieldUnion := operations.CreateFieldCreateDocumentFieldFieldUnionFieldCreateDocumentFieldFieldEmail(operations.FieldCreateDocumentFieldFieldEmail{/* values here */})
```

### FieldCreateDocumentFieldFieldDate

```go
fieldCreateDocumentFieldFieldUnion := operations.CreateFieldCreateDocumentFieldFieldUnionFieldCreateDocumentFieldFieldDate(operations.FieldCreateDocumentFieldFieldDate{/* values here */})
```

### FieldCreateDocumentFieldFieldText

```go
fieldCreateDocumentFieldFieldUnion := operations.CreateFieldCreateDocumentFieldFieldUnionFieldCreateDocumentFieldFieldText(operations.FieldCreateDocumentFieldFieldText{/* values here */})
```

### FieldCreateDocumentFieldFieldNumber

```go
fieldCreateDocumentFieldFieldUnion := operations.CreateFieldCreateDocumentFieldFieldUnionFieldCreateDocumentFieldFieldNumber(operations.FieldCreateDocumentFieldFieldNumber{/* values here */})
```

### FieldCreateDocumentFieldFieldRadio

```go
fieldCreateDocumentFieldFieldUnion := operations.CreateFieldCreateDocumentFieldFieldUnionFieldCreateDocumentFieldFieldRadio(operations.FieldCreateDocumentFieldFieldRadio{/* values here */})
```

### FieldCreateDocumentFieldFieldCheckbox

```go
fieldCreateDocumentFieldFieldUnion := operations.CreateFieldCreateDocumentFieldFieldUnionFieldCreateDocumentFieldFieldCheckbox(operations.FieldCreateDocumentFieldFieldCheckbox{/* values here */})
```

### FieldCreateDocumentFieldFieldDropdown

```go
fieldCreateDocumentFieldFieldUnion := operations.CreateFieldCreateDocumentFieldFieldUnionFieldCreateDocumentFieldFieldDropdown(operations.FieldCreateDocumentFieldFieldDropdown{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch fieldCreateDocumentFieldFieldUnion.Type {
	case operations.FieldCreateDocumentFieldFieldUnionTypeFieldCreateDocumentFieldFieldSignature:
		// fieldCreateDocumentFieldFieldUnion.FieldCreateDocumentFieldFieldSignature is populated
	case operations.FieldCreateDocumentFieldFieldUnionTypeFieldCreateDocumentFieldFieldFreeSignature:
		// fieldCreateDocumentFieldFieldUnion.FieldCreateDocumentFieldFieldFreeSignature is populated
	case operations.FieldCreateDocumentFieldFieldUnionTypeFieldCreateDocumentFieldFieldInitials:
		// fieldCreateDocumentFieldFieldUnion.FieldCreateDocumentFieldFieldInitials is populated
	case operations.FieldCreateDocumentFieldFieldUnionTypeFieldCreateDocumentFieldFieldName:
		// fieldCreateDocumentFieldFieldUnion.FieldCreateDocumentFieldFieldName is populated
	case operations.FieldCreateDocumentFieldFieldUnionTypeFieldCreateDocumentFieldFieldEmail:
		// fieldCreateDocumentFieldFieldUnion.FieldCreateDocumentFieldFieldEmail is populated
	case operations.FieldCreateDocumentFieldFieldUnionTypeFieldCreateDocumentFieldFieldDate:
		// fieldCreateDocumentFieldFieldUnion.FieldCreateDocumentFieldFieldDate is populated
	case operations.FieldCreateDocumentFieldFieldUnionTypeFieldCreateDocumentFieldFieldText:
		// fieldCreateDocumentFieldFieldUnion.FieldCreateDocumentFieldFieldText is populated
	case operations.FieldCreateDocumentFieldFieldUnionTypeFieldCreateDocumentFieldFieldNumber:
		// fieldCreateDocumentFieldFieldUnion.FieldCreateDocumentFieldFieldNumber is populated
	case operations.FieldCreateDocumentFieldFieldUnionTypeFieldCreateDocumentFieldFieldRadio:
		// fieldCreateDocumentFieldFieldUnion.FieldCreateDocumentFieldFieldRadio is populated
	case operations.FieldCreateDocumentFieldFieldUnionTypeFieldCreateDocumentFieldFieldCheckbox:
		// fieldCreateDocumentFieldFieldUnion.FieldCreateDocumentFieldFieldCheckbox is populated
	case operations.FieldCreateDocumentFieldFieldUnionTypeFieldCreateDocumentFieldFieldDropdown:
		// fieldCreateDocumentFieldFieldUnion.FieldCreateDocumentFieldFieldDropdown is populated
}
```
