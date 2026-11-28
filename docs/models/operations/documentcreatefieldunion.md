# DocumentCreateFieldUnion


## Supported Types

### DocumentCreateFieldSignature

```go
documentCreateFieldUnion := operations.CreateDocumentCreateFieldUnionDocumentCreateFieldSignature(operations.DocumentCreateFieldSignature{/* values here */})
```

### DocumentCreateFieldFreeSignature

```go
documentCreateFieldUnion := operations.CreateDocumentCreateFieldUnionDocumentCreateFieldFreeSignature(operations.DocumentCreateFieldFreeSignature{/* values here */})
```

### DocumentCreateFieldInitials

```go
documentCreateFieldUnion := operations.CreateDocumentCreateFieldUnionDocumentCreateFieldInitials(operations.DocumentCreateFieldInitials{/* values here */})
```

### DocumentCreateFieldName

```go
documentCreateFieldUnion := operations.CreateDocumentCreateFieldUnionDocumentCreateFieldName(operations.DocumentCreateFieldName{/* values here */})
```

### DocumentCreateFieldEmail

```go
documentCreateFieldUnion := operations.CreateDocumentCreateFieldUnionDocumentCreateFieldEmail(operations.DocumentCreateFieldEmail{/* values here */})
```

### DocumentCreateFieldDate

```go
documentCreateFieldUnion := operations.CreateDocumentCreateFieldUnionDocumentCreateFieldDate(operations.DocumentCreateFieldDate{/* values here */})
```

### DocumentCreateFieldText

```go
documentCreateFieldUnion := operations.CreateDocumentCreateFieldUnionDocumentCreateFieldText(operations.DocumentCreateFieldText{/* values here */})
```

### DocumentCreateFieldNumber

```go
documentCreateFieldUnion := operations.CreateDocumentCreateFieldUnionDocumentCreateFieldNumber(operations.DocumentCreateFieldNumber{/* values here */})
```

### DocumentCreateFieldRadio

```go
documentCreateFieldUnion := operations.CreateDocumentCreateFieldUnionDocumentCreateFieldRadio(operations.DocumentCreateFieldRadio{/* values here */})
```

### DocumentCreateFieldCheckbox

```go
documentCreateFieldUnion := operations.CreateDocumentCreateFieldUnionDocumentCreateFieldCheckbox(operations.DocumentCreateFieldCheckbox{/* values here */})
```

### DocumentCreateFieldDropdown

```go
documentCreateFieldUnion := operations.CreateDocumentCreateFieldUnionDocumentCreateFieldDropdown(operations.DocumentCreateFieldDropdown{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch documentCreateFieldUnion.Type {
	case operations.DocumentCreateFieldUnionTypeDocumentCreateFieldSignature:
		// documentCreateFieldUnion.DocumentCreateFieldSignature is populated
	case operations.DocumentCreateFieldUnionTypeDocumentCreateFieldFreeSignature:
		// documentCreateFieldUnion.DocumentCreateFieldFreeSignature is populated
	case operations.DocumentCreateFieldUnionTypeDocumentCreateFieldInitials:
		// documentCreateFieldUnion.DocumentCreateFieldInitials is populated
	case operations.DocumentCreateFieldUnionTypeDocumentCreateFieldName:
		// documentCreateFieldUnion.DocumentCreateFieldName is populated
	case operations.DocumentCreateFieldUnionTypeDocumentCreateFieldEmail:
		// documentCreateFieldUnion.DocumentCreateFieldEmail is populated
	case operations.DocumentCreateFieldUnionTypeDocumentCreateFieldDate:
		// documentCreateFieldUnion.DocumentCreateFieldDate is populated
	case operations.DocumentCreateFieldUnionTypeDocumentCreateFieldText:
		// documentCreateFieldUnion.DocumentCreateFieldText is populated
	case operations.DocumentCreateFieldUnionTypeDocumentCreateFieldNumber:
		// documentCreateFieldUnion.DocumentCreateFieldNumber is populated
	case operations.DocumentCreateFieldUnionTypeDocumentCreateFieldRadio:
		// documentCreateFieldUnion.DocumentCreateFieldRadio is populated
	case operations.DocumentCreateFieldUnionTypeDocumentCreateFieldCheckbox:
		// documentCreateFieldUnion.DocumentCreateFieldCheckbox is populated
	case operations.DocumentCreateFieldUnionTypeDocumentCreateFieldDropdown:
		// documentCreateFieldUnion.DocumentCreateFieldDropdown is populated
}
```
