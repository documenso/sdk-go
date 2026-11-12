# EnvelopeGetFieldMetaUnion


## Supported Types

### EnvelopeGetFieldMetaSignature

```go
envelopeGetFieldMetaUnion := operations.CreateEnvelopeGetFieldMetaUnionEnvelopeGetFieldMetaSignature(operations.EnvelopeGetFieldMetaSignature{/* values here */})
```

### EnvelopeGetFieldMetaInitials

```go
envelopeGetFieldMetaUnion := operations.CreateEnvelopeGetFieldMetaUnionEnvelopeGetFieldMetaInitials(operations.EnvelopeGetFieldMetaInitials{/* values here */})
```

### EnvelopeGetFieldMetaName

```go
envelopeGetFieldMetaUnion := operations.CreateEnvelopeGetFieldMetaUnionEnvelopeGetFieldMetaName(operations.EnvelopeGetFieldMetaName{/* values here */})
```

### EnvelopeGetFieldMetaEmail

```go
envelopeGetFieldMetaUnion := operations.CreateEnvelopeGetFieldMetaUnionEnvelopeGetFieldMetaEmail(operations.EnvelopeGetFieldMetaEmail{/* values here */})
```

### EnvelopeGetFieldMetaDate

```go
envelopeGetFieldMetaUnion := operations.CreateEnvelopeGetFieldMetaUnionEnvelopeGetFieldMetaDate(operations.EnvelopeGetFieldMetaDate{/* values here */})
```

### EnvelopeGetFieldMetaText

```go
envelopeGetFieldMetaUnion := operations.CreateEnvelopeGetFieldMetaUnionEnvelopeGetFieldMetaText(operations.EnvelopeGetFieldMetaText{/* values here */})
```

### EnvelopeGetFieldMetaNumber

```go
envelopeGetFieldMetaUnion := operations.CreateEnvelopeGetFieldMetaUnionEnvelopeGetFieldMetaNumber(operations.EnvelopeGetFieldMetaNumber{/* values here */})
```

### EnvelopeGetFieldMetaRadio

```go
envelopeGetFieldMetaUnion := operations.CreateEnvelopeGetFieldMetaUnionEnvelopeGetFieldMetaRadio(operations.EnvelopeGetFieldMetaRadio{/* values here */})
```

### EnvelopeGetFieldMetaCheckbox

```go
envelopeGetFieldMetaUnion := operations.CreateEnvelopeGetFieldMetaUnionEnvelopeGetFieldMetaCheckbox(operations.EnvelopeGetFieldMetaCheckbox{/* values here */})
```

### EnvelopeGetFieldMetaDropdown

```go
envelopeGetFieldMetaUnion := operations.CreateEnvelopeGetFieldMetaUnionEnvelopeGetFieldMetaDropdown(operations.EnvelopeGetFieldMetaDropdown{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch envelopeGetFieldMetaUnion.Type {
	case operations.EnvelopeGetFieldMetaUnionTypeEnvelopeGetFieldMetaSignature:
		// envelopeGetFieldMetaUnion.EnvelopeGetFieldMetaSignature is populated
	case operations.EnvelopeGetFieldMetaUnionTypeEnvelopeGetFieldMetaInitials:
		// envelopeGetFieldMetaUnion.EnvelopeGetFieldMetaInitials is populated
	case operations.EnvelopeGetFieldMetaUnionTypeEnvelopeGetFieldMetaName:
		// envelopeGetFieldMetaUnion.EnvelopeGetFieldMetaName is populated
	case operations.EnvelopeGetFieldMetaUnionTypeEnvelopeGetFieldMetaEmail:
		// envelopeGetFieldMetaUnion.EnvelopeGetFieldMetaEmail is populated
	case operations.EnvelopeGetFieldMetaUnionTypeEnvelopeGetFieldMetaDate:
		// envelopeGetFieldMetaUnion.EnvelopeGetFieldMetaDate is populated
	case operations.EnvelopeGetFieldMetaUnionTypeEnvelopeGetFieldMetaText:
		// envelopeGetFieldMetaUnion.EnvelopeGetFieldMetaText is populated
	case operations.EnvelopeGetFieldMetaUnionTypeEnvelopeGetFieldMetaNumber:
		// envelopeGetFieldMetaUnion.EnvelopeGetFieldMetaNumber is populated
	case operations.EnvelopeGetFieldMetaUnionTypeEnvelopeGetFieldMetaRadio:
		// envelopeGetFieldMetaUnion.EnvelopeGetFieldMetaRadio is populated
	case operations.EnvelopeGetFieldMetaUnionTypeEnvelopeGetFieldMetaCheckbox:
		// envelopeGetFieldMetaUnion.EnvelopeGetFieldMetaCheckbox is populated
	case operations.EnvelopeGetFieldMetaUnionTypeEnvelopeGetFieldMetaDropdown:
		// envelopeGetFieldMetaUnion.EnvelopeGetFieldMetaDropdown is populated
}
```
