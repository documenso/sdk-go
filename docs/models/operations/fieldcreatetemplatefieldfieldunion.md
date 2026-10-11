# FieldCreateTemplateFieldFieldUnion


## Supported Types

### FieldCreateTemplateFieldFieldSignature

```go
fieldCreateTemplateFieldFieldUnion := operations.CreateFieldCreateTemplateFieldFieldUnionFieldCreateTemplateFieldFieldSignature(operations.FieldCreateTemplateFieldFieldSignature{/* values here */})
```

### FieldCreateTemplateFieldFieldFreeSignature

```go
fieldCreateTemplateFieldFieldUnion := operations.CreateFieldCreateTemplateFieldFieldUnionFieldCreateTemplateFieldFieldFreeSignature(operations.FieldCreateTemplateFieldFieldFreeSignature{/* values here */})
```

### FieldCreateTemplateFieldFieldInitials

```go
fieldCreateTemplateFieldFieldUnion := operations.CreateFieldCreateTemplateFieldFieldUnionFieldCreateTemplateFieldFieldInitials(operations.FieldCreateTemplateFieldFieldInitials{/* values here */})
```

### FieldCreateTemplateFieldFieldName

```go
fieldCreateTemplateFieldFieldUnion := operations.CreateFieldCreateTemplateFieldFieldUnionFieldCreateTemplateFieldFieldName(operations.FieldCreateTemplateFieldFieldName{/* values here */})
```

### FieldCreateTemplateFieldFieldEmail

```go
fieldCreateTemplateFieldFieldUnion := operations.CreateFieldCreateTemplateFieldFieldUnionFieldCreateTemplateFieldFieldEmail(operations.FieldCreateTemplateFieldFieldEmail{/* values here */})
```

### FieldCreateTemplateFieldFieldDate

```go
fieldCreateTemplateFieldFieldUnion := operations.CreateFieldCreateTemplateFieldFieldUnionFieldCreateTemplateFieldFieldDate(operations.FieldCreateTemplateFieldFieldDate{/* values here */})
```

### FieldCreateTemplateFieldFieldText

```go
fieldCreateTemplateFieldFieldUnion := operations.CreateFieldCreateTemplateFieldFieldUnionFieldCreateTemplateFieldFieldText(operations.FieldCreateTemplateFieldFieldText{/* values here */})
```

### FieldCreateTemplateFieldFieldNumber

```go
fieldCreateTemplateFieldFieldUnion := operations.CreateFieldCreateTemplateFieldFieldUnionFieldCreateTemplateFieldFieldNumber(operations.FieldCreateTemplateFieldFieldNumber{/* values here */})
```

### FieldCreateTemplateFieldFieldRadio

```go
fieldCreateTemplateFieldFieldUnion := operations.CreateFieldCreateTemplateFieldFieldUnionFieldCreateTemplateFieldFieldRadio(operations.FieldCreateTemplateFieldFieldRadio{/* values here */})
```

### FieldCreateTemplateFieldFieldCheckbox

```go
fieldCreateTemplateFieldFieldUnion := operations.CreateFieldCreateTemplateFieldFieldUnionFieldCreateTemplateFieldFieldCheckbox(operations.FieldCreateTemplateFieldFieldCheckbox{/* values here */})
```

### FieldCreateTemplateFieldFieldDropdown

```go
fieldCreateTemplateFieldFieldUnion := operations.CreateFieldCreateTemplateFieldFieldUnionFieldCreateTemplateFieldFieldDropdown(operations.FieldCreateTemplateFieldFieldDropdown{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch fieldCreateTemplateFieldFieldUnion.Type {
	case operations.FieldCreateTemplateFieldFieldUnionTypeFieldCreateTemplateFieldFieldSignature:
		// fieldCreateTemplateFieldFieldUnion.FieldCreateTemplateFieldFieldSignature is populated
	case operations.FieldCreateTemplateFieldFieldUnionTypeFieldCreateTemplateFieldFieldFreeSignature:
		// fieldCreateTemplateFieldFieldUnion.FieldCreateTemplateFieldFieldFreeSignature is populated
	case operations.FieldCreateTemplateFieldFieldUnionTypeFieldCreateTemplateFieldFieldInitials:
		// fieldCreateTemplateFieldFieldUnion.FieldCreateTemplateFieldFieldInitials is populated
	case operations.FieldCreateTemplateFieldFieldUnionTypeFieldCreateTemplateFieldFieldName:
		// fieldCreateTemplateFieldFieldUnion.FieldCreateTemplateFieldFieldName is populated
	case operations.FieldCreateTemplateFieldFieldUnionTypeFieldCreateTemplateFieldFieldEmail:
		// fieldCreateTemplateFieldFieldUnion.FieldCreateTemplateFieldFieldEmail is populated
	case operations.FieldCreateTemplateFieldFieldUnionTypeFieldCreateTemplateFieldFieldDate:
		// fieldCreateTemplateFieldFieldUnion.FieldCreateTemplateFieldFieldDate is populated
	case operations.FieldCreateTemplateFieldFieldUnionTypeFieldCreateTemplateFieldFieldText:
		// fieldCreateTemplateFieldFieldUnion.FieldCreateTemplateFieldFieldText is populated
	case operations.FieldCreateTemplateFieldFieldUnionTypeFieldCreateTemplateFieldFieldNumber:
		// fieldCreateTemplateFieldFieldUnion.FieldCreateTemplateFieldFieldNumber is populated
	case operations.FieldCreateTemplateFieldFieldUnionTypeFieldCreateTemplateFieldFieldRadio:
		// fieldCreateTemplateFieldFieldUnion.FieldCreateTemplateFieldFieldRadio is populated
	case operations.FieldCreateTemplateFieldFieldUnionTypeFieldCreateTemplateFieldFieldCheckbox:
		// fieldCreateTemplateFieldFieldUnion.FieldCreateTemplateFieldFieldCheckbox is populated
	case operations.FieldCreateTemplateFieldFieldUnionTypeFieldCreateTemplateFieldFieldDropdown:
		// fieldCreateTemplateFieldFieldUnion.FieldCreateTemplateFieldFieldDropdown is populated
}
```
