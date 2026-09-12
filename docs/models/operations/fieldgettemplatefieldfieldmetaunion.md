# FieldGetTemplateFieldFieldMetaUnion


## Supported Types

### FieldGetTemplateFieldFieldMetaSignature

```go
fieldGetTemplateFieldFieldMetaUnion := operations.CreateFieldGetTemplateFieldFieldMetaUnionFieldGetTemplateFieldFieldMetaSignature(operations.FieldGetTemplateFieldFieldMetaSignature{/* values here */})
```

### FieldGetTemplateFieldFieldMetaInitials

```go
fieldGetTemplateFieldFieldMetaUnion := operations.CreateFieldGetTemplateFieldFieldMetaUnionFieldGetTemplateFieldFieldMetaInitials(operations.FieldGetTemplateFieldFieldMetaInitials{/* values here */})
```

### FieldGetTemplateFieldFieldMetaName

```go
fieldGetTemplateFieldFieldMetaUnion := operations.CreateFieldGetTemplateFieldFieldMetaUnionFieldGetTemplateFieldFieldMetaName(operations.FieldGetTemplateFieldFieldMetaName{/* values here */})
```

### FieldGetTemplateFieldFieldMetaEmail

```go
fieldGetTemplateFieldFieldMetaUnion := operations.CreateFieldGetTemplateFieldFieldMetaUnionFieldGetTemplateFieldFieldMetaEmail(operations.FieldGetTemplateFieldFieldMetaEmail{/* values here */})
```

### FieldGetTemplateFieldFieldMetaDate

```go
fieldGetTemplateFieldFieldMetaUnion := operations.CreateFieldGetTemplateFieldFieldMetaUnionFieldGetTemplateFieldFieldMetaDate(operations.FieldGetTemplateFieldFieldMetaDate{/* values here */})
```

### FieldGetTemplateFieldFieldMetaText

```go
fieldGetTemplateFieldFieldMetaUnion := operations.CreateFieldGetTemplateFieldFieldMetaUnionFieldGetTemplateFieldFieldMetaText(operations.FieldGetTemplateFieldFieldMetaText{/* values here */})
```

### FieldGetTemplateFieldFieldMetaNumber

```go
fieldGetTemplateFieldFieldMetaUnion := operations.CreateFieldGetTemplateFieldFieldMetaUnionFieldGetTemplateFieldFieldMetaNumber(operations.FieldGetTemplateFieldFieldMetaNumber{/* values here */})
```

### FieldGetTemplateFieldFieldMetaRadio

```go
fieldGetTemplateFieldFieldMetaUnion := operations.CreateFieldGetTemplateFieldFieldMetaUnionFieldGetTemplateFieldFieldMetaRadio(operations.FieldGetTemplateFieldFieldMetaRadio{/* values here */})
```

### FieldGetTemplateFieldFieldMetaCheckbox

```go
fieldGetTemplateFieldFieldMetaUnion := operations.CreateFieldGetTemplateFieldFieldMetaUnionFieldGetTemplateFieldFieldMetaCheckbox(operations.FieldGetTemplateFieldFieldMetaCheckbox{/* values here */})
```

### FieldGetTemplateFieldFieldMetaDropdown

```go
fieldGetTemplateFieldFieldMetaUnion := operations.CreateFieldGetTemplateFieldFieldMetaUnionFieldGetTemplateFieldFieldMetaDropdown(operations.FieldGetTemplateFieldFieldMetaDropdown{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch fieldGetTemplateFieldFieldMetaUnion.Type {
	case operations.FieldGetTemplateFieldFieldMetaUnionTypeFieldGetTemplateFieldFieldMetaSignature:
		// fieldGetTemplateFieldFieldMetaUnion.FieldGetTemplateFieldFieldMetaSignature is populated
	case operations.FieldGetTemplateFieldFieldMetaUnionTypeFieldGetTemplateFieldFieldMetaInitials:
		// fieldGetTemplateFieldFieldMetaUnion.FieldGetTemplateFieldFieldMetaInitials is populated
	case operations.FieldGetTemplateFieldFieldMetaUnionTypeFieldGetTemplateFieldFieldMetaName:
		// fieldGetTemplateFieldFieldMetaUnion.FieldGetTemplateFieldFieldMetaName is populated
	case operations.FieldGetTemplateFieldFieldMetaUnionTypeFieldGetTemplateFieldFieldMetaEmail:
		// fieldGetTemplateFieldFieldMetaUnion.FieldGetTemplateFieldFieldMetaEmail is populated
	case operations.FieldGetTemplateFieldFieldMetaUnionTypeFieldGetTemplateFieldFieldMetaDate:
		// fieldGetTemplateFieldFieldMetaUnion.FieldGetTemplateFieldFieldMetaDate is populated
	case operations.FieldGetTemplateFieldFieldMetaUnionTypeFieldGetTemplateFieldFieldMetaText:
		// fieldGetTemplateFieldFieldMetaUnion.FieldGetTemplateFieldFieldMetaText is populated
	case operations.FieldGetTemplateFieldFieldMetaUnionTypeFieldGetTemplateFieldFieldMetaNumber:
		// fieldGetTemplateFieldFieldMetaUnion.FieldGetTemplateFieldFieldMetaNumber is populated
	case operations.FieldGetTemplateFieldFieldMetaUnionTypeFieldGetTemplateFieldFieldMetaRadio:
		// fieldGetTemplateFieldFieldMetaUnion.FieldGetTemplateFieldFieldMetaRadio is populated
	case operations.FieldGetTemplateFieldFieldMetaUnionTypeFieldGetTemplateFieldFieldMetaCheckbox:
		// fieldGetTemplateFieldFieldMetaUnion.FieldGetTemplateFieldFieldMetaCheckbox is populated
	case operations.FieldGetTemplateFieldFieldMetaUnionTypeFieldGetTemplateFieldFieldMetaDropdown:
		// fieldGetTemplateFieldFieldMetaUnion.FieldGetTemplateFieldFieldMetaDropdown is populated
}
```
