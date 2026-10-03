# TemplateGetManyFieldMetaUnion


## Supported Types

### TemplateGetManyFieldMetaSignature

```go
templateGetManyFieldMetaUnion := operations.CreateTemplateGetManyFieldMetaUnionTemplateGetManyFieldMetaSignature(operations.TemplateGetManyFieldMetaSignature{/* values here */})
```

### TemplateGetManyFieldMetaInitials

```go
templateGetManyFieldMetaUnion := operations.CreateTemplateGetManyFieldMetaUnionTemplateGetManyFieldMetaInitials(operations.TemplateGetManyFieldMetaInitials{/* values here */})
```

### TemplateGetManyFieldMetaName

```go
templateGetManyFieldMetaUnion := operations.CreateTemplateGetManyFieldMetaUnionTemplateGetManyFieldMetaName(operations.TemplateGetManyFieldMetaName{/* values here */})
```

### TemplateGetManyFieldMetaEmail

```go
templateGetManyFieldMetaUnion := operations.CreateTemplateGetManyFieldMetaUnionTemplateGetManyFieldMetaEmail(operations.TemplateGetManyFieldMetaEmail{/* values here */})
```

### TemplateGetManyFieldMetaDate

```go
templateGetManyFieldMetaUnion := operations.CreateTemplateGetManyFieldMetaUnionTemplateGetManyFieldMetaDate(operations.TemplateGetManyFieldMetaDate{/* values here */})
```

### TemplateGetManyFieldMetaText

```go
templateGetManyFieldMetaUnion := operations.CreateTemplateGetManyFieldMetaUnionTemplateGetManyFieldMetaText(operations.TemplateGetManyFieldMetaText{/* values here */})
```

### TemplateGetManyFieldMetaNumber

```go
templateGetManyFieldMetaUnion := operations.CreateTemplateGetManyFieldMetaUnionTemplateGetManyFieldMetaNumber(operations.TemplateGetManyFieldMetaNumber{/* values here */})
```

### TemplateGetManyFieldMetaRadio

```go
templateGetManyFieldMetaUnion := operations.CreateTemplateGetManyFieldMetaUnionTemplateGetManyFieldMetaRadio(operations.TemplateGetManyFieldMetaRadio{/* values here */})
```

### TemplateGetManyFieldMetaCheckbox

```go
templateGetManyFieldMetaUnion := operations.CreateTemplateGetManyFieldMetaUnionTemplateGetManyFieldMetaCheckbox(operations.TemplateGetManyFieldMetaCheckbox{/* values here */})
```

### TemplateGetManyFieldMetaDropdown

```go
templateGetManyFieldMetaUnion := operations.CreateTemplateGetManyFieldMetaUnionTemplateGetManyFieldMetaDropdown(operations.TemplateGetManyFieldMetaDropdown{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch templateGetManyFieldMetaUnion.Type {
	case operations.TemplateGetManyFieldMetaUnionTypeTemplateGetManyFieldMetaSignature:
		// templateGetManyFieldMetaUnion.TemplateGetManyFieldMetaSignature is populated
	case operations.TemplateGetManyFieldMetaUnionTypeTemplateGetManyFieldMetaInitials:
		// templateGetManyFieldMetaUnion.TemplateGetManyFieldMetaInitials is populated
	case operations.TemplateGetManyFieldMetaUnionTypeTemplateGetManyFieldMetaName:
		// templateGetManyFieldMetaUnion.TemplateGetManyFieldMetaName is populated
	case operations.TemplateGetManyFieldMetaUnionTypeTemplateGetManyFieldMetaEmail:
		// templateGetManyFieldMetaUnion.TemplateGetManyFieldMetaEmail is populated
	case operations.TemplateGetManyFieldMetaUnionTypeTemplateGetManyFieldMetaDate:
		// templateGetManyFieldMetaUnion.TemplateGetManyFieldMetaDate is populated
	case operations.TemplateGetManyFieldMetaUnionTypeTemplateGetManyFieldMetaText:
		// templateGetManyFieldMetaUnion.TemplateGetManyFieldMetaText is populated
	case operations.TemplateGetManyFieldMetaUnionTypeTemplateGetManyFieldMetaNumber:
		// templateGetManyFieldMetaUnion.TemplateGetManyFieldMetaNumber is populated
	case operations.TemplateGetManyFieldMetaUnionTypeTemplateGetManyFieldMetaRadio:
		// templateGetManyFieldMetaUnion.TemplateGetManyFieldMetaRadio is populated
	case operations.TemplateGetManyFieldMetaUnionTypeTemplateGetManyFieldMetaCheckbox:
		// templateGetManyFieldMetaUnion.TemplateGetManyFieldMetaCheckbox is populated
	case operations.TemplateGetManyFieldMetaUnionTypeTemplateGetManyFieldMetaDropdown:
		// templateGetManyFieldMetaUnion.TemplateGetManyFieldMetaDropdown is populated
}
```
