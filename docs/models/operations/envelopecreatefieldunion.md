# EnvelopeCreateFieldUnion


## Supported Types

### EnvelopeCreateFieldSignature

```go
envelopeCreateFieldUnion := operations.CreateEnvelopeCreateFieldUnionEnvelopeCreateFieldSignature(operations.EnvelopeCreateFieldSignature{/* values here */})
```

### EnvelopeCreateFieldFreeSignature

```go
envelopeCreateFieldUnion := operations.CreateEnvelopeCreateFieldUnionEnvelopeCreateFieldFreeSignature(operations.EnvelopeCreateFieldFreeSignature{/* values here */})
```

### EnvelopeCreateFieldInitials

```go
envelopeCreateFieldUnion := operations.CreateEnvelopeCreateFieldUnionEnvelopeCreateFieldInitials(operations.EnvelopeCreateFieldInitials{/* values here */})
```

### EnvelopeCreateFieldName

```go
envelopeCreateFieldUnion := operations.CreateEnvelopeCreateFieldUnionEnvelopeCreateFieldName(operations.EnvelopeCreateFieldName{/* values here */})
```

### EnvelopeCreateFieldEmail

```go
envelopeCreateFieldUnion := operations.CreateEnvelopeCreateFieldUnionEnvelopeCreateFieldEmail(operations.EnvelopeCreateFieldEmail{/* values here */})
```

### EnvelopeCreateFieldDate

```go
envelopeCreateFieldUnion := operations.CreateEnvelopeCreateFieldUnionEnvelopeCreateFieldDate(operations.EnvelopeCreateFieldDate{/* values here */})
```

### EnvelopeCreateFieldText

```go
envelopeCreateFieldUnion := operations.CreateEnvelopeCreateFieldUnionEnvelopeCreateFieldText(operations.EnvelopeCreateFieldText{/* values here */})
```

### EnvelopeCreateFieldNumber

```go
envelopeCreateFieldUnion := operations.CreateEnvelopeCreateFieldUnionEnvelopeCreateFieldNumber(operations.EnvelopeCreateFieldNumber{/* values here */})
```

### EnvelopeCreateFieldRadio

```go
envelopeCreateFieldUnion := operations.CreateEnvelopeCreateFieldUnionEnvelopeCreateFieldRadio(operations.EnvelopeCreateFieldRadio{/* values here */})
```

### EnvelopeCreateFieldCheckbox

```go
envelopeCreateFieldUnion := operations.CreateEnvelopeCreateFieldUnionEnvelopeCreateFieldCheckbox(operations.EnvelopeCreateFieldCheckbox{/* values here */})
```

### EnvelopeCreateFieldDropdown

```go
envelopeCreateFieldUnion := operations.CreateEnvelopeCreateFieldUnionEnvelopeCreateFieldDropdown(operations.EnvelopeCreateFieldDropdown{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch envelopeCreateFieldUnion.Type {
	case operations.EnvelopeCreateFieldUnionTypeEnvelopeCreateFieldSignature:
		// envelopeCreateFieldUnion.EnvelopeCreateFieldSignature is populated
	case operations.EnvelopeCreateFieldUnionTypeEnvelopeCreateFieldFreeSignature:
		// envelopeCreateFieldUnion.EnvelopeCreateFieldFreeSignature is populated
	case operations.EnvelopeCreateFieldUnionTypeEnvelopeCreateFieldInitials:
		// envelopeCreateFieldUnion.EnvelopeCreateFieldInitials is populated
	case operations.EnvelopeCreateFieldUnionTypeEnvelopeCreateFieldName:
		// envelopeCreateFieldUnion.EnvelopeCreateFieldName is populated
	case operations.EnvelopeCreateFieldUnionTypeEnvelopeCreateFieldEmail:
		// envelopeCreateFieldUnion.EnvelopeCreateFieldEmail is populated
	case operations.EnvelopeCreateFieldUnionTypeEnvelopeCreateFieldDate:
		// envelopeCreateFieldUnion.EnvelopeCreateFieldDate is populated
	case operations.EnvelopeCreateFieldUnionTypeEnvelopeCreateFieldText:
		// envelopeCreateFieldUnion.EnvelopeCreateFieldText is populated
	case operations.EnvelopeCreateFieldUnionTypeEnvelopeCreateFieldNumber:
		// envelopeCreateFieldUnion.EnvelopeCreateFieldNumber is populated
	case operations.EnvelopeCreateFieldUnionTypeEnvelopeCreateFieldRadio:
		// envelopeCreateFieldUnion.EnvelopeCreateFieldRadio is populated
	case operations.EnvelopeCreateFieldUnionTypeEnvelopeCreateFieldCheckbox:
		// envelopeCreateFieldUnion.EnvelopeCreateFieldCheckbox is populated
	case operations.EnvelopeCreateFieldUnionTypeEnvelopeCreateFieldDropdown:
		// envelopeCreateFieldUnion.EnvelopeCreateFieldDropdown is populated
}
```
