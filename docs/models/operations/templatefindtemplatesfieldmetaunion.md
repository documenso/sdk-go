# TemplateFindTemplatesFieldMetaUnion


## Supported Types

### TemplateFindTemplatesFieldMetaSignature

```go
templateFindTemplatesFieldMetaUnion := operations.CreateTemplateFindTemplatesFieldMetaUnionTemplateFindTemplatesFieldMetaSignature(operations.TemplateFindTemplatesFieldMetaSignature{/* values here */})
```

### TemplateFindTemplatesFieldMetaInitials

```go
templateFindTemplatesFieldMetaUnion := operations.CreateTemplateFindTemplatesFieldMetaUnionTemplateFindTemplatesFieldMetaInitials(operations.TemplateFindTemplatesFieldMetaInitials{/* values here */})
```

### TemplateFindTemplatesFieldMetaName

```go
templateFindTemplatesFieldMetaUnion := operations.CreateTemplateFindTemplatesFieldMetaUnionTemplateFindTemplatesFieldMetaName(operations.TemplateFindTemplatesFieldMetaName{/* values here */})
```

### TemplateFindTemplatesFieldMetaEmail

```go
templateFindTemplatesFieldMetaUnion := operations.CreateTemplateFindTemplatesFieldMetaUnionTemplateFindTemplatesFieldMetaEmail(operations.TemplateFindTemplatesFieldMetaEmail{/* values here */})
```

### TemplateFindTemplatesFieldMetaDate

```go
templateFindTemplatesFieldMetaUnion := operations.CreateTemplateFindTemplatesFieldMetaUnionTemplateFindTemplatesFieldMetaDate(operations.TemplateFindTemplatesFieldMetaDate{/* values here */})
```

### TemplateFindTemplatesFieldMetaText

```go
templateFindTemplatesFieldMetaUnion := operations.CreateTemplateFindTemplatesFieldMetaUnionTemplateFindTemplatesFieldMetaText(operations.TemplateFindTemplatesFieldMetaText{/* values here */})
```

### TemplateFindTemplatesFieldMetaNumber

```go
templateFindTemplatesFieldMetaUnion := operations.CreateTemplateFindTemplatesFieldMetaUnionTemplateFindTemplatesFieldMetaNumber(operations.TemplateFindTemplatesFieldMetaNumber{/* values here */})
```

### TemplateFindTemplatesFieldMetaRadio

```go
templateFindTemplatesFieldMetaUnion := operations.CreateTemplateFindTemplatesFieldMetaUnionTemplateFindTemplatesFieldMetaRadio(operations.TemplateFindTemplatesFieldMetaRadio{/* values here */})
```

### TemplateFindTemplatesFieldMetaCheckbox

```go
templateFindTemplatesFieldMetaUnion := operations.CreateTemplateFindTemplatesFieldMetaUnionTemplateFindTemplatesFieldMetaCheckbox(operations.TemplateFindTemplatesFieldMetaCheckbox{/* values here */})
```

### TemplateFindTemplatesFieldMetaDropdown

```go
templateFindTemplatesFieldMetaUnion := operations.CreateTemplateFindTemplatesFieldMetaUnionTemplateFindTemplatesFieldMetaDropdown(operations.TemplateFindTemplatesFieldMetaDropdown{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch templateFindTemplatesFieldMetaUnion.Type {
	case operations.TemplateFindTemplatesFieldMetaUnionTypeTemplateFindTemplatesFieldMetaSignature:
		// templateFindTemplatesFieldMetaUnion.TemplateFindTemplatesFieldMetaSignature is populated
	case operations.TemplateFindTemplatesFieldMetaUnionTypeTemplateFindTemplatesFieldMetaInitials:
		// templateFindTemplatesFieldMetaUnion.TemplateFindTemplatesFieldMetaInitials is populated
	case operations.TemplateFindTemplatesFieldMetaUnionTypeTemplateFindTemplatesFieldMetaName:
		// templateFindTemplatesFieldMetaUnion.TemplateFindTemplatesFieldMetaName is populated
	case operations.TemplateFindTemplatesFieldMetaUnionTypeTemplateFindTemplatesFieldMetaEmail:
		// templateFindTemplatesFieldMetaUnion.TemplateFindTemplatesFieldMetaEmail is populated
	case operations.TemplateFindTemplatesFieldMetaUnionTypeTemplateFindTemplatesFieldMetaDate:
		// templateFindTemplatesFieldMetaUnion.TemplateFindTemplatesFieldMetaDate is populated
	case operations.TemplateFindTemplatesFieldMetaUnionTypeTemplateFindTemplatesFieldMetaText:
		// templateFindTemplatesFieldMetaUnion.TemplateFindTemplatesFieldMetaText is populated
	case operations.TemplateFindTemplatesFieldMetaUnionTypeTemplateFindTemplatesFieldMetaNumber:
		// templateFindTemplatesFieldMetaUnion.TemplateFindTemplatesFieldMetaNumber is populated
	case operations.TemplateFindTemplatesFieldMetaUnionTypeTemplateFindTemplatesFieldMetaRadio:
		// templateFindTemplatesFieldMetaUnion.TemplateFindTemplatesFieldMetaRadio is populated
	case operations.TemplateFindTemplatesFieldMetaUnionTypeTemplateFindTemplatesFieldMetaCheckbox:
		// templateFindTemplatesFieldMetaUnion.TemplateFindTemplatesFieldMetaCheckbox is populated
	case operations.TemplateFindTemplatesFieldMetaUnionTypeTemplateFindTemplatesFieldMetaDropdown:
		// templateFindTemplatesFieldMetaUnion.TemplateFindTemplatesFieldMetaDropdown is populated
}
```
