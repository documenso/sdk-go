# FieldCreateTemplateFieldsFieldUnion


## Supported Types

### FieldCreateTemplateFieldsFieldSignature

```go
fieldCreateTemplateFieldsFieldUnion := operations.CreateFieldCreateTemplateFieldsFieldUnionFieldCreateTemplateFieldsFieldSignature(operations.FieldCreateTemplateFieldsFieldSignature{/* values here */})
```

### FieldCreateTemplateFieldsFieldFreeSignature

```go
fieldCreateTemplateFieldsFieldUnion := operations.CreateFieldCreateTemplateFieldsFieldUnionFieldCreateTemplateFieldsFieldFreeSignature(operations.FieldCreateTemplateFieldsFieldFreeSignature{/* values here */})
```

### FieldCreateTemplateFieldsFieldInitials

```go
fieldCreateTemplateFieldsFieldUnion := operations.CreateFieldCreateTemplateFieldsFieldUnionFieldCreateTemplateFieldsFieldInitials(operations.FieldCreateTemplateFieldsFieldInitials{/* values here */})
```

### FieldCreateTemplateFieldsFieldName

```go
fieldCreateTemplateFieldsFieldUnion := operations.CreateFieldCreateTemplateFieldsFieldUnionFieldCreateTemplateFieldsFieldName(operations.FieldCreateTemplateFieldsFieldName{/* values here */})
```

### FieldCreateTemplateFieldsFieldEmail

```go
fieldCreateTemplateFieldsFieldUnion := operations.CreateFieldCreateTemplateFieldsFieldUnionFieldCreateTemplateFieldsFieldEmail(operations.FieldCreateTemplateFieldsFieldEmail{/* values here */})
```

### FieldCreateTemplateFieldsFieldDate

```go
fieldCreateTemplateFieldsFieldUnion := operations.CreateFieldCreateTemplateFieldsFieldUnionFieldCreateTemplateFieldsFieldDate(operations.FieldCreateTemplateFieldsFieldDate{/* values here */})
```

### FieldCreateTemplateFieldsFieldText

```go
fieldCreateTemplateFieldsFieldUnion := operations.CreateFieldCreateTemplateFieldsFieldUnionFieldCreateTemplateFieldsFieldText(operations.FieldCreateTemplateFieldsFieldText{/* values here */})
```

### FieldCreateTemplateFieldsFieldNumber

```go
fieldCreateTemplateFieldsFieldUnion := operations.CreateFieldCreateTemplateFieldsFieldUnionFieldCreateTemplateFieldsFieldNumber(operations.FieldCreateTemplateFieldsFieldNumber{/* values here */})
```

### FieldCreateTemplateFieldsFieldRadio

```go
fieldCreateTemplateFieldsFieldUnion := operations.CreateFieldCreateTemplateFieldsFieldUnionFieldCreateTemplateFieldsFieldRadio(operations.FieldCreateTemplateFieldsFieldRadio{/* values here */})
```

### FieldCreateTemplateFieldsFieldCheckbox

```go
fieldCreateTemplateFieldsFieldUnion := operations.CreateFieldCreateTemplateFieldsFieldUnionFieldCreateTemplateFieldsFieldCheckbox(operations.FieldCreateTemplateFieldsFieldCheckbox{/* values here */})
```

### FieldCreateTemplateFieldsFieldDropdown

```go
fieldCreateTemplateFieldsFieldUnion := operations.CreateFieldCreateTemplateFieldsFieldUnionFieldCreateTemplateFieldsFieldDropdown(operations.FieldCreateTemplateFieldsFieldDropdown{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch fieldCreateTemplateFieldsFieldUnion.Type {
	case operations.FieldCreateTemplateFieldsFieldUnionTypeFieldCreateTemplateFieldsFieldSignature:
		// fieldCreateTemplateFieldsFieldUnion.FieldCreateTemplateFieldsFieldSignature is populated
	case operations.FieldCreateTemplateFieldsFieldUnionTypeFieldCreateTemplateFieldsFieldFreeSignature:
		// fieldCreateTemplateFieldsFieldUnion.FieldCreateTemplateFieldsFieldFreeSignature is populated
	case operations.FieldCreateTemplateFieldsFieldUnionTypeFieldCreateTemplateFieldsFieldInitials:
		// fieldCreateTemplateFieldsFieldUnion.FieldCreateTemplateFieldsFieldInitials is populated
	case operations.FieldCreateTemplateFieldsFieldUnionTypeFieldCreateTemplateFieldsFieldName:
		// fieldCreateTemplateFieldsFieldUnion.FieldCreateTemplateFieldsFieldName is populated
	case operations.FieldCreateTemplateFieldsFieldUnionTypeFieldCreateTemplateFieldsFieldEmail:
		// fieldCreateTemplateFieldsFieldUnion.FieldCreateTemplateFieldsFieldEmail is populated
	case operations.FieldCreateTemplateFieldsFieldUnionTypeFieldCreateTemplateFieldsFieldDate:
		// fieldCreateTemplateFieldsFieldUnion.FieldCreateTemplateFieldsFieldDate is populated
	case operations.FieldCreateTemplateFieldsFieldUnionTypeFieldCreateTemplateFieldsFieldText:
		// fieldCreateTemplateFieldsFieldUnion.FieldCreateTemplateFieldsFieldText is populated
	case operations.FieldCreateTemplateFieldsFieldUnionTypeFieldCreateTemplateFieldsFieldNumber:
		// fieldCreateTemplateFieldsFieldUnion.FieldCreateTemplateFieldsFieldNumber is populated
	case operations.FieldCreateTemplateFieldsFieldUnionTypeFieldCreateTemplateFieldsFieldRadio:
		// fieldCreateTemplateFieldsFieldUnion.FieldCreateTemplateFieldsFieldRadio is populated
	case operations.FieldCreateTemplateFieldsFieldUnionTypeFieldCreateTemplateFieldsFieldCheckbox:
		// fieldCreateTemplateFieldsFieldUnion.FieldCreateTemplateFieldsFieldCheckbox is populated
	case operations.FieldCreateTemplateFieldsFieldUnionTypeFieldCreateTemplateFieldsFieldDropdown:
		// fieldCreateTemplateFieldsFieldUnion.FieldCreateTemplateFieldsFieldDropdown is populated
}
```
