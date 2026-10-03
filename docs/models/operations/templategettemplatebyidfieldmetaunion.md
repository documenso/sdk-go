# TemplateGetTemplateByIDFieldMetaUnion


## Supported Types

### TemplateGetTemplateByIDFieldMetaSignature

```go
templateGetTemplateByIDFieldMetaUnion := operations.CreateTemplateGetTemplateByIDFieldMetaUnionTemplateGetTemplateByIDFieldMetaSignature(operations.TemplateGetTemplateByIDFieldMetaSignature{/* values here */})
```

### TemplateGetTemplateByIDFieldMetaInitials

```go
templateGetTemplateByIDFieldMetaUnion := operations.CreateTemplateGetTemplateByIDFieldMetaUnionTemplateGetTemplateByIDFieldMetaInitials(operations.TemplateGetTemplateByIDFieldMetaInitials{/* values here */})
```

### TemplateGetTemplateByIDFieldMetaName

```go
templateGetTemplateByIDFieldMetaUnion := operations.CreateTemplateGetTemplateByIDFieldMetaUnionTemplateGetTemplateByIDFieldMetaName(operations.TemplateGetTemplateByIDFieldMetaName{/* values here */})
```

### TemplateGetTemplateByIDFieldMetaEmail

```go
templateGetTemplateByIDFieldMetaUnion := operations.CreateTemplateGetTemplateByIDFieldMetaUnionTemplateGetTemplateByIDFieldMetaEmail(operations.TemplateGetTemplateByIDFieldMetaEmail{/* values here */})
```

### TemplateGetTemplateByIDFieldMetaDate

```go
templateGetTemplateByIDFieldMetaUnion := operations.CreateTemplateGetTemplateByIDFieldMetaUnionTemplateGetTemplateByIDFieldMetaDate(operations.TemplateGetTemplateByIDFieldMetaDate{/* values here */})
```

### TemplateGetTemplateByIDFieldMetaText

```go
templateGetTemplateByIDFieldMetaUnion := operations.CreateTemplateGetTemplateByIDFieldMetaUnionTemplateGetTemplateByIDFieldMetaText(operations.TemplateGetTemplateByIDFieldMetaText{/* values here */})
```

### TemplateGetTemplateByIDFieldMetaNumber

```go
templateGetTemplateByIDFieldMetaUnion := operations.CreateTemplateGetTemplateByIDFieldMetaUnionTemplateGetTemplateByIDFieldMetaNumber(operations.TemplateGetTemplateByIDFieldMetaNumber{/* values here */})
```

### TemplateGetTemplateByIDFieldMetaRadio

```go
templateGetTemplateByIDFieldMetaUnion := operations.CreateTemplateGetTemplateByIDFieldMetaUnionTemplateGetTemplateByIDFieldMetaRadio(operations.TemplateGetTemplateByIDFieldMetaRadio{/* values here */})
```

### TemplateGetTemplateByIDFieldMetaCheckbox

```go
templateGetTemplateByIDFieldMetaUnion := operations.CreateTemplateGetTemplateByIDFieldMetaUnionTemplateGetTemplateByIDFieldMetaCheckbox(operations.TemplateGetTemplateByIDFieldMetaCheckbox{/* values here */})
```

### TemplateGetTemplateByIDFieldMetaDropdown

```go
templateGetTemplateByIDFieldMetaUnion := operations.CreateTemplateGetTemplateByIDFieldMetaUnionTemplateGetTemplateByIDFieldMetaDropdown(operations.TemplateGetTemplateByIDFieldMetaDropdown{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch templateGetTemplateByIDFieldMetaUnion.Type {
	case operations.TemplateGetTemplateByIDFieldMetaUnionTypeTemplateGetTemplateByIDFieldMetaSignature:
		// templateGetTemplateByIDFieldMetaUnion.TemplateGetTemplateByIDFieldMetaSignature is populated
	case operations.TemplateGetTemplateByIDFieldMetaUnionTypeTemplateGetTemplateByIDFieldMetaInitials:
		// templateGetTemplateByIDFieldMetaUnion.TemplateGetTemplateByIDFieldMetaInitials is populated
	case operations.TemplateGetTemplateByIDFieldMetaUnionTypeTemplateGetTemplateByIDFieldMetaName:
		// templateGetTemplateByIDFieldMetaUnion.TemplateGetTemplateByIDFieldMetaName is populated
	case operations.TemplateGetTemplateByIDFieldMetaUnionTypeTemplateGetTemplateByIDFieldMetaEmail:
		// templateGetTemplateByIDFieldMetaUnion.TemplateGetTemplateByIDFieldMetaEmail is populated
	case operations.TemplateGetTemplateByIDFieldMetaUnionTypeTemplateGetTemplateByIDFieldMetaDate:
		// templateGetTemplateByIDFieldMetaUnion.TemplateGetTemplateByIDFieldMetaDate is populated
	case operations.TemplateGetTemplateByIDFieldMetaUnionTypeTemplateGetTemplateByIDFieldMetaText:
		// templateGetTemplateByIDFieldMetaUnion.TemplateGetTemplateByIDFieldMetaText is populated
	case operations.TemplateGetTemplateByIDFieldMetaUnionTypeTemplateGetTemplateByIDFieldMetaNumber:
		// templateGetTemplateByIDFieldMetaUnion.TemplateGetTemplateByIDFieldMetaNumber is populated
	case operations.TemplateGetTemplateByIDFieldMetaUnionTypeTemplateGetTemplateByIDFieldMetaRadio:
		// templateGetTemplateByIDFieldMetaUnion.TemplateGetTemplateByIDFieldMetaRadio is populated
	case operations.TemplateGetTemplateByIDFieldMetaUnionTypeTemplateGetTemplateByIDFieldMetaCheckbox:
		// templateGetTemplateByIDFieldMetaUnion.TemplateGetTemplateByIDFieldMetaCheckbox is populated
	case operations.TemplateGetTemplateByIDFieldMetaUnionTypeTemplateGetTemplateByIDFieldMetaDropdown:
		// templateGetTemplateByIDFieldMetaUnion.TemplateGetTemplateByIDFieldMetaDropdown is populated
}
```
