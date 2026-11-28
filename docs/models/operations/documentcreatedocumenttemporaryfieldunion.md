# DocumentCreateDocumentTemporaryFieldUnion


## Supported Types

### DocumentCreateDocumentTemporaryFieldSignature

```go
documentCreateDocumentTemporaryFieldUnion := operations.CreateDocumentCreateDocumentTemporaryFieldUnionDocumentCreateDocumentTemporaryFieldSignature(operations.DocumentCreateDocumentTemporaryFieldSignature{/* values here */})
```

### DocumentCreateDocumentTemporaryFieldFreeSignature

```go
documentCreateDocumentTemporaryFieldUnion := operations.CreateDocumentCreateDocumentTemporaryFieldUnionDocumentCreateDocumentTemporaryFieldFreeSignature(operations.DocumentCreateDocumentTemporaryFieldFreeSignature{/* values here */})
```

### DocumentCreateDocumentTemporaryFieldInitials

```go
documentCreateDocumentTemporaryFieldUnion := operations.CreateDocumentCreateDocumentTemporaryFieldUnionDocumentCreateDocumentTemporaryFieldInitials(operations.DocumentCreateDocumentTemporaryFieldInitials{/* values here */})
```

### DocumentCreateDocumentTemporaryFieldName

```go
documentCreateDocumentTemporaryFieldUnion := operations.CreateDocumentCreateDocumentTemporaryFieldUnionDocumentCreateDocumentTemporaryFieldName(operations.DocumentCreateDocumentTemporaryFieldName{/* values here */})
```

### DocumentCreateDocumentTemporaryFieldEmail

```go
documentCreateDocumentTemporaryFieldUnion := operations.CreateDocumentCreateDocumentTemporaryFieldUnionDocumentCreateDocumentTemporaryFieldEmail(operations.DocumentCreateDocumentTemporaryFieldEmail{/* values here */})
```

### DocumentCreateDocumentTemporaryFieldDate

```go
documentCreateDocumentTemporaryFieldUnion := operations.CreateDocumentCreateDocumentTemporaryFieldUnionDocumentCreateDocumentTemporaryFieldDate(operations.DocumentCreateDocumentTemporaryFieldDate{/* values here */})
```

### DocumentCreateDocumentTemporaryFieldText

```go
documentCreateDocumentTemporaryFieldUnion := operations.CreateDocumentCreateDocumentTemporaryFieldUnionDocumentCreateDocumentTemporaryFieldText(operations.DocumentCreateDocumentTemporaryFieldText{/* values here */})
```

### DocumentCreateDocumentTemporaryFieldNumber

```go
documentCreateDocumentTemporaryFieldUnion := operations.CreateDocumentCreateDocumentTemporaryFieldUnionDocumentCreateDocumentTemporaryFieldNumber(operations.DocumentCreateDocumentTemporaryFieldNumber{/* values here */})
```

### DocumentCreateDocumentTemporaryFieldRadio

```go
documentCreateDocumentTemporaryFieldUnion := operations.CreateDocumentCreateDocumentTemporaryFieldUnionDocumentCreateDocumentTemporaryFieldRadio(operations.DocumentCreateDocumentTemporaryFieldRadio{/* values here */})
```

### DocumentCreateDocumentTemporaryFieldCheckbox

```go
documentCreateDocumentTemporaryFieldUnion := operations.CreateDocumentCreateDocumentTemporaryFieldUnionDocumentCreateDocumentTemporaryFieldCheckbox(operations.DocumentCreateDocumentTemporaryFieldCheckbox{/* values here */})
```

### DocumentCreateDocumentTemporaryFieldDropdown

```go
documentCreateDocumentTemporaryFieldUnion := operations.CreateDocumentCreateDocumentTemporaryFieldUnionDocumentCreateDocumentTemporaryFieldDropdown(operations.DocumentCreateDocumentTemporaryFieldDropdown{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch documentCreateDocumentTemporaryFieldUnion.Type {
	case operations.DocumentCreateDocumentTemporaryFieldUnionTypeDocumentCreateDocumentTemporaryFieldSignature:
		// documentCreateDocumentTemporaryFieldUnion.DocumentCreateDocumentTemporaryFieldSignature is populated
	case operations.DocumentCreateDocumentTemporaryFieldUnionTypeDocumentCreateDocumentTemporaryFieldFreeSignature:
		// documentCreateDocumentTemporaryFieldUnion.DocumentCreateDocumentTemporaryFieldFreeSignature is populated
	case operations.DocumentCreateDocumentTemporaryFieldUnionTypeDocumentCreateDocumentTemporaryFieldInitials:
		// documentCreateDocumentTemporaryFieldUnion.DocumentCreateDocumentTemporaryFieldInitials is populated
	case operations.DocumentCreateDocumentTemporaryFieldUnionTypeDocumentCreateDocumentTemporaryFieldName:
		// documentCreateDocumentTemporaryFieldUnion.DocumentCreateDocumentTemporaryFieldName is populated
	case operations.DocumentCreateDocumentTemporaryFieldUnionTypeDocumentCreateDocumentTemporaryFieldEmail:
		// documentCreateDocumentTemporaryFieldUnion.DocumentCreateDocumentTemporaryFieldEmail is populated
	case operations.DocumentCreateDocumentTemporaryFieldUnionTypeDocumentCreateDocumentTemporaryFieldDate:
		// documentCreateDocumentTemporaryFieldUnion.DocumentCreateDocumentTemporaryFieldDate is populated
	case operations.DocumentCreateDocumentTemporaryFieldUnionTypeDocumentCreateDocumentTemporaryFieldText:
		// documentCreateDocumentTemporaryFieldUnion.DocumentCreateDocumentTemporaryFieldText is populated
	case operations.DocumentCreateDocumentTemporaryFieldUnionTypeDocumentCreateDocumentTemporaryFieldNumber:
		// documentCreateDocumentTemporaryFieldUnion.DocumentCreateDocumentTemporaryFieldNumber is populated
	case operations.DocumentCreateDocumentTemporaryFieldUnionTypeDocumentCreateDocumentTemporaryFieldRadio:
		// documentCreateDocumentTemporaryFieldUnion.DocumentCreateDocumentTemporaryFieldRadio is populated
	case operations.DocumentCreateDocumentTemporaryFieldUnionTypeDocumentCreateDocumentTemporaryFieldCheckbox:
		// documentCreateDocumentTemporaryFieldUnion.DocumentCreateDocumentTemporaryFieldCheckbox is populated
	case operations.DocumentCreateDocumentTemporaryFieldUnionTypeDocumentCreateDocumentTemporaryFieldDropdown:
		// documentCreateDocumentTemporaryFieldUnion.DocumentCreateDocumentTemporaryFieldDropdown is populated
}
```
