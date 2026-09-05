# FieldUpdateTemplateFieldFieldUnion


## Supported Types

### FieldUpdateTemplateFieldFieldSignature

```go
fieldUpdateTemplateFieldFieldUnion := operations.CreateFieldUpdateTemplateFieldFieldUnionFieldUpdateTemplateFieldFieldSignature(operations.FieldUpdateTemplateFieldFieldSignature{/* values here */})
```

### FieldUpdateTemplateFieldFieldFreeSignature

```go
fieldUpdateTemplateFieldFieldUnion := operations.CreateFieldUpdateTemplateFieldFieldUnionFieldUpdateTemplateFieldFieldFreeSignature(operations.FieldUpdateTemplateFieldFieldFreeSignature{/* values here */})
```

### FieldUpdateTemplateFieldFieldInitials

```go
fieldUpdateTemplateFieldFieldUnion := operations.CreateFieldUpdateTemplateFieldFieldUnionFieldUpdateTemplateFieldFieldInitials(operations.FieldUpdateTemplateFieldFieldInitials{/* values here */})
```

### FieldUpdateTemplateFieldFieldName

```go
fieldUpdateTemplateFieldFieldUnion := operations.CreateFieldUpdateTemplateFieldFieldUnionFieldUpdateTemplateFieldFieldName(operations.FieldUpdateTemplateFieldFieldName{/* values here */})
```

### FieldUpdateTemplateFieldFieldEmail

```go
fieldUpdateTemplateFieldFieldUnion := operations.CreateFieldUpdateTemplateFieldFieldUnionFieldUpdateTemplateFieldFieldEmail(operations.FieldUpdateTemplateFieldFieldEmail{/* values here */})
```

### FieldUpdateTemplateFieldFieldDate

```go
fieldUpdateTemplateFieldFieldUnion := operations.CreateFieldUpdateTemplateFieldFieldUnionFieldUpdateTemplateFieldFieldDate(operations.FieldUpdateTemplateFieldFieldDate{/* values here */})
```

### FieldUpdateTemplateFieldFieldText

```go
fieldUpdateTemplateFieldFieldUnion := operations.CreateFieldUpdateTemplateFieldFieldUnionFieldUpdateTemplateFieldFieldText(operations.FieldUpdateTemplateFieldFieldText{/* values here */})
```

### FieldUpdateTemplateFieldFieldNumber

```go
fieldUpdateTemplateFieldFieldUnion := operations.CreateFieldUpdateTemplateFieldFieldUnionFieldUpdateTemplateFieldFieldNumber(operations.FieldUpdateTemplateFieldFieldNumber{/* values here */})
```

### FieldUpdateTemplateFieldFieldRadio

```go
fieldUpdateTemplateFieldFieldUnion := operations.CreateFieldUpdateTemplateFieldFieldUnionFieldUpdateTemplateFieldFieldRadio(operations.FieldUpdateTemplateFieldFieldRadio{/* values here */})
```

### FieldUpdateTemplateFieldFieldCheckbox

```go
fieldUpdateTemplateFieldFieldUnion := operations.CreateFieldUpdateTemplateFieldFieldUnionFieldUpdateTemplateFieldFieldCheckbox(operations.FieldUpdateTemplateFieldFieldCheckbox{/* values here */})
```

### FieldUpdateTemplateFieldFieldDropdown

```go
fieldUpdateTemplateFieldFieldUnion := operations.CreateFieldUpdateTemplateFieldFieldUnionFieldUpdateTemplateFieldFieldDropdown(operations.FieldUpdateTemplateFieldFieldDropdown{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch fieldUpdateTemplateFieldFieldUnion.Type {
	case operations.FieldUpdateTemplateFieldFieldUnionTypeFieldUpdateTemplateFieldFieldSignature:
		// fieldUpdateTemplateFieldFieldUnion.FieldUpdateTemplateFieldFieldSignature is populated
	case operations.FieldUpdateTemplateFieldFieldUnionTypeFieldUpdateTemplateFieldFieldFreeSignature:
		// fieldUpdateTemplateFieldFieldUnion.FieldUpdateTemplateFieldFieldFreeSignature is populated
	case operations.FieldUpdateTemplateFieldFieldUnionTypeFieldUpdateTemplateFieldFieldInitials:
		// fieldUpdateTemplateFieldFieldUnion.FieldUpdateTemplateFieldFieldInitials is populated
	case operations.FieldUpdateTemplateFieldFieldUnionTypeFieldUpdateTemplateFieldFieldName:
		// fieldUpdateTemplateFieldFieldUnion.FieldUpdateTemplateFieldFieldName is populated
	case operations.FieldUpdateTemplateFieldFieldUnionTypeFieldUpdateTemplateFieldFieldEmail:
		// fieldUpdateTemplateFieldFieldUnion.FieldUpdateTemplateFieldFieldEmail is populated
	case operations.FieldUpdateTemplateFieldFieldUnionTypeFieldUpdateTemplateFieldFieldDate:
		// fieldUpdateTemplateFieldFieldUnion.FieldUpdateTemplateFieldFieldDate is populated
	case operations.FieldUpdateTemplateFieldFieldUnionTypeFieldUpdateTemplateFieldFieldText:
		// fieldUpdateTemplateFieldFieldUnion.FieldUpdateTemplateFieldFieldText is populated
	case operations.FieldUpdateTemplateFieldFieldUnionTypeFieldUpdateTemplateFieldFieldNumber:
		// fieldUpdateTemplateFieldFieldUnion.FieldUpdateTemplateFieldFieldNumber is populated
	case operations.FieldUpdateTemplateFieldFieldUnionTypeFieldUpdateTemplateFieldFieldRadio:
		// fieldUpdateTemplateFieldFieldUnion.FieldUpdateTemplateFieldFieldRadio is populated
	case operations.FieldUpdateTemplateFieldFieldUnionTypeFieldUpdateTemplateFieldFieldCheckbox:
		// fieldUpdateTemplateFieldFieldUnion.FieldUpdateTemplateFieldFieldCheckbox is populated
	case operations.FieldUpdateTemplateFieldFieldUnionTypeFieldUpdateTemplateFieldFieldDropdown:
		// fieldUpdateTemplateFieldFieldUnion.FieldUpdateTemplateFieldFieldDropdown is populated
}
```
