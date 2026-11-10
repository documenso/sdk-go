# EnvelopeFieldGetFieldMetaUnion


## Supported Types

### EnvelopeFieldGetFieldMetaSignature

```go
envelopeFieldGetFieldMetaUnion := operations.CreateEnvelopeFieldGetFieldMetaUnionEnvelopeFieldGetFieldMetaSignature(operations.EnvelopeFieldGetFieldMetaSignature{/* values here */})
```

### EnvelopeFieldGetFieldMetaInitials

```go
envelopeFieldGetFieldMetaUnion := operations.CreateEnvelopeFieldGetFieldMetaUnionEnvelopeFieldGetFieldMetaInitials(operations.EnvelopeFieldGetFieldMetaInitials{/* values here */})
```

### EnvelopeFieldGetFieldMetaName

```go
envelopeFieldGetFieldMetaUnion := operations.CreateEnvelopeFieldGetFieldMetaUnionEnvelopeFieldGetFieldMetaName(operations.EnvelopeFieldGetFieldMetaName{/* values here */})
```

### EnvelopeFieldGetFieldMetaEmail

```go
envelopeFieldGetFieldMetaUnion := operations.CreateEnvelopeFieldGetFieldMetaUnionEnvelopeFieldGetFieldMetaEmail(operations.EnvelopeFieldGetFieldMetaEmail{/* values here */})
```

### EnvelopeFieldGetFieldMetaDate

```go
envelopeFieldGetFieldMetaUnion := operations.CreateEnvelopeFieldGetFieldMetaUnionEnvelopeFieldGetFieldMetaDate(operations.EnvelopeFieldGetFieldMetaDate{/* values here */})
```

### EnvelopeFieldGetFieldMetaText

```go
envelopeFieldGetFieldMetaUnion := operations.CreateEnvelopeFieldGetFieldMetaUnionEnvelopeFieldGetFieldMetaText(operations.EnvelopeFieldGetFieldMetaText{/* values here */})
```

### EnvelopeFieldGetFieldMetaNumber

```go
envelopeFieldGetFieldMetaUnion := operations.CreateEnvelopeFieldGetFieldMetaUnionEnvelopeFieldGetFieldMetaNumber(operations.EnvelopeFieldGetFieldMetaNumber{/* values here */})
```

### EnvelopeFieldGetFieldMetaRadio

```go
envelopeFieldGetFieldMetaUnion := operations.CreateEnvelopeFieldGetFieldMetaUnionEnvelopeFieldGetFieldMetaRadio(operations.EnvelopeFieldGetFieldMetaRadio{/* values here */})
```

### EnvelopeFieldGetFieldMetaCheckbox

```go
envelopeFieldGetFieldMetaUnion := operations.CreateEnvelopeFieldGetFieldMetaUnionEnvelopeFieldGetFieldMetaCheckbox(operations.EnvelopeFieldGetFieldMetaCheckbox{/* values here */})
```

### EnvelopeFieldGetFieldMetaDropdown

```go
envelopeFieldGetFieldMetaUnion := operations.CreateEnvelopeFieldGetFieldMetaUnionEnvelopeFieldGetFieldMetaDropdown(operations.EnvelopeFieldGetFieldMetaDropdown{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch envelopeFieldGetFieldMetaUnion.Type {
	case operations.EnvelopeFieldGetFieldMetaUnionTypeEnvelopeFieldGetFieldMetaSignature:
		// envelopeFieldGetFieldMetaUnion.EnvelopeFieldGetFieldMetaSignature is populated
	case operations.EnvelopeFieldGetFieldMetaUnionTypeEnvelopeFieldGetFieldMetaInitials:
		// envelopeFieldGetFieldMetaUnion.EnvelopeFieldGetFieldMetaInitials is populated
	case operations.EnvelopeFieldGetFieldMetaUnionTypeEnvelopeFieldGetFieldMetaName:
		// envelopeFieldGetFieldMetaUnion.EnvelopeFieldGetFieldMetaName is populated
	case operations.EnvelopeFieldGetFieldMetaUnionTypeEnvelopeFieldGetFieldMetaEmail:
		// envelopeFieldGetFieldMetaUnion.EnvelopeFieldGetFieldMetaEmail is populated
	case operations.EnvelopeFieldGetFieldMetaUnionTypeEnvelopeFieldGetFieldMetaDate:
		// envelopeFieldGetFieldMetaUnion.EnvelopeFieldGetFieldMetaDate is populated
	case operations.EnvelopeFieldGetFieldMetaUnionTypeEnvelopeFieldGetFieldMetaText:
		// envelopeFieldGetFieldMetaUnion.EnvelopeFieldGetFieldMetaText is populated
	case operations.EnvelopeFieldGetFieldMetaUnionTypeEnvelopeFieldGetFieldMetaNumber:
		// envelopeFieldGetFieldMetaUnion.EnvelopeFieldGetFieldMetaNumber is populated
	case operations.EnvelopeFieldGetFieldMetaUnionTypeEnvelopeFieldGetFieldMetaRadio:
		// envelopeFieldGetFieldMetaUnion.EnvelopeFieldGetFieldMetaRadio is populated
	case operations.EnvelopeFieldGetFieldMetaUnionTypeEnvelopeFieldGetFieldMetaCheckbox:
		// envelopeFieldGetFieldMetaUnion.EnvelopeFieldGetFieldMetaCheckbox is populated
	case operations.EnvelopeFieldGetFieldMetaUnionTypeEnvelopeFieldGetFieldMetaDropdown:
		// envelopeFieldGetFieldMetaUnion.EnvelopeFieldGetFieldMetaDropdown is populated
}
```
