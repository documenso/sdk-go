# FieldUpdateTemplateFieldsFieldUnion


## Supported Types

### FieldUpdateTemplateFieldsFieldSignature

```go
fieldUpdateTemplateFieldsFieldUnion := operations.CreateFieldUpdateTemplateFieldsFieldUnionFieldUpdateTemplateFieldsFieldSignature(operations.FieldUpdateTemplateFieldsFieldSignature{/* values here */})
```

### FieldUpdateTemplateFieldsFieldFreeSignature

```go
fieldUpdateTemplateFieldsFieldUnion := operations.CreateFieldUpdateTemplateFieldsFieldUnionFieldUpdateTemplateFieldsFieldFreeSignature(operations.FieldUpdateTemplateFieldsFieldFreeSignature{/* values here */})
```

### FieldUpdateTemplateFieldsFieldInitials

```go
fieldUpdateTemplateFieldsFieldUnion := operations.CreateFieldUpdateTemplateFieldsFieldUnionFieldUpdateTemplateFieldsFieldInitials(operations.FieldUpdateTemplateFieldsFieldInitials{/* values here */})
```

### FieldUpdateTemplateFieldsFieldName

```go
fieldUpdateTemplateFieldsFieldUnion := operations.CreateFieldUpdateTemplateFieldsFieldUnionFieldUpdateTemplateFieldsFieldName(operations.FieldUpdateTemplateFieldsFieldName{/* values here */})
```

### FieldUpdateTemplateFieldsFieldEmail

```go
fieldUpdateTemplateFieldsFieldUnion := operations.CreateFieldUpdateTemplateFieldsFieldUnionFieldUpdateTemplateFieldsFieldEmail(operations.FieldUpdateTemplateFieldsFieldEmail{/* values here */})
```

### FieldUpdateTemplateFieldsFieldDate

```go
fieldUpdateTemplateFieldsFieldUnion := operations.CreateFieldUpdateTemplateFieldsFieldUnionFieldUpdateTemplateFieldsFieldDate(operations.FieldUpdateTemplateFieldsFieldDate{/* values here */})
```

### FieldUpdateTemplateFieldsFieldText

```go
fieldUpdateTemplateFieldsFieldUnion := operations.CreateFieldUpdateTemplateFieldsFieldUnionFieldUpdateTemplateFieldsFieldText(operations.FieldUpdateTemplateFieldsFieldText{/* values here */})
```

### FieldUpdateTemplateFieldsFieldNumber

```go
fieldUpdateTemplateFieldsFieldUnion := operations.CreateFieldUpdateTemplateFieldsFieldUnionFieldUpdateTemplateFieldsFieldNumber(operations.FieldUpdateTemplateFieldsFieldNumber{/* values here */})
```

### FieldUpdateTemplateFieldsFieldRadio

```go
fieldUpdateTemplateFieldsFieldUnion := operations.CreateFieldUpdateTemplateFieldsFieldUnionFieldUpdateTemplateFieldsFieldRadio(operations.FieldUpdateTemplateFieldsFieldRadio{/* values here */})
```

### FieldUpdateTemplateFieldsFieldCheckbox

```go
fieldUpdateTemplateFieldsFieldUnion := operations.CreateFieldUpdateTemplateFieldsFieldUnionFieldUpdateTemplateFieldsFieldCheckbox(operations.FieldUpdateTemplateFieldsFieldCheckbox{/* values here */})
```

### FieldUpdateTemplateFieldsFieldDropdown

```go
fieldUpdateTemplateFieldsFieldUnion := operations.CreateFieldUpdateTemplateFieldsFieldUnionFieldUpdateTemplateFieldsFieldDropdown(operations.FieldUpdateTemplateFieldsFieldDropdown{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch fieldUpdateTemplateFieldsFieldUnion.Type {
	case operations.FieldUpdateTemplateFieldsFieldUnionTypeFieldUpdateTemplateFieldsFieldSignature:
		// fieldUpdateTemplateFieldsFieldUnion.FieldUpdateTemplateFieldsFieldSignature is populated
	case operations.FieldUpdateTemplateFieldsFieldUnionTypeFieldUpdateTemplateFieldsFieldFreeSignature:
		// fieldUpdateTemplateFieldsFieldUnion.FieldUpdateTemplateFieldsFieldFreeSignature is populated
	case operations.FieldUpdateTemplateFieldsFieldUnionTypeFieldUpdateTemplateFieldsFieldInitials:
		// fieldUpdateTemplateFieldsFieldUnion.FieldUpdateTemplateFieldsFieldInitials is populated
	case operations.FieldUpdateTemplateFieldsFieldUnionTypeFieldUpdateTemplateFieldsFieldName:
		// fieldUpdateTemplateFieldsFieldUnion.FieldUpdateTemplateFieldsFieldName is populated
	case operations.FieldUpdateTemplateFieldsFieldUnionTypeFieldUpdateTemplateFieldsFieldEmail:
		// fieldUpdateTemplateFieldsFieldUnion.FieldUpdateTemplateFieldsFieldEmail is populated
	case operations.FieldUpdateTemplateFieldsFieldUnionTypeFieldUpdateTemplateFieldsFieldDate:
		// fieldUpdateTemplateFieldsFieldUnion.FieldUpdateTemplateFieldsFieldDate is populated
	case operations.FieldUpdateTemplateFieldsFieldUnionTypeFieldUpdateTemplateFieldsFieldText:
		// fieldUpdateTemplateFieldsFieldUnion.FieldUpdateTemplateFieldsFieldText is populated
	case operations.FieldUpdateTemplateFieldsFieldUnionTypeFieldUpdateTemplateFieldsFieldNumber:
		// fieldUpdateTemplateFieldsFieldUnion.FieldUpdateTemplateFieldsFieldNumber is populated
	case operations.FieldUpdateTemplateFieldsFieldUnionTypeFieldUpdateTemplateFieldsFieldRadio:
		// fieldUpdateTemplateFieldsFieldUnion.FieldUpdateTemplateFieldsFieldRadio is populated
	case operations.FieldUpdateTemplateFieldsFieldUnionTypeFieldUpdateTemplateFieldsFieldCheckbox:
		// fieldUpdateTemplateFieldsFieldUnion.FieldUpdateTemplateFieldsFieldCheckbox is populated
	case operations.FieldUpdateTemplateFieldsFieldUnionTypeFieldUpdateTemplateFieldsFieldDropdown:
		// fieldUpdateTemplateFieldsFieldUnion.FieldUpdateTemplateFieldsFieldDropdown is populated
}
```
