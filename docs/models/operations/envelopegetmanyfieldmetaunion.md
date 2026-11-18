# EnvelopeGetManyFieldMetaUnion


## Supported Types

### EnvelopeGetManyFieldMetaSignature

```go
envelopeGetManyFieldMetaUnion := operations.CreateEnvelopeGetManyFieldMetaUnionEnvelopeGetManyFieldMetaSignature(operations.EnvelopeGetManyFieldMetaSignature{/* values here */})
```

### EnvelopeGetManyFieldMetaInitials

```go
envelopeGetManyFieldMetaUnion := operations.CreateEnvelopeGetManyFieldMetaUnionEnvelopeGetManyFieldMetaInitials(operations.EnvelopeGetManyFieldMetaInitials{/* values here */})
```

### EnvelopeGetManyFieldMetaName

```go
envelopeGetManyFieldMetaUnion := operations.CreateEnvelopeGetManyFieldMetaUnionEnvelopeGetManyFieldMetaName(operations.EnvelopeGetManyFieldMetaName{/* values here */})
```

### EnvelopeGetManyFieldMetaEmail

```go
envelopeGetManyFieldMetaUnion := operations.CreateEnvelopeGetManyFieldMetaUnionEnvelopeGetManyFieldMetaEmail(operations.EnvelopeGetManyFieldMetaEmail{/* values here */})
```

### EnvelopeGetManyFieldMetaDate

```go
envelopeGetManyFieldMetaUnion := operations.CreateEnvelopeGetManyFieldMetaUnionEnvelopeGetManyFieldMetaDate(operations.EnvelopeGetManyFieldMetaDate{/* values here */})
```

### EnvelopeGetManyFieldMetaText

```go
envelopeGetManyFieldMetaUnion := operations.CreateEnvelopeGetManyFieldMetaUnionEnvelopeGetManyFieldMetaText(operations.EnvelopeGetManyFieldMetaText{/* values here */})
```

### EnvelopeGetManyFieldMetaNumber

```go
envelopeGetManyFieldMetaUnion := operations.CreateEnvelopeGetManyFieldMetaUnionEnvelopeGetManyFieldMetaNumber(operations.EnvelopeGetManyFieldMetaNumber{/* values here */})
```

### EnvelopeGetManyFieldMetaRadio

```go
envelopeGetManyFieldMetaUnion := operations.CreateEnvelopeGetManyFieldMetaUnionEnvelopeGetManyFieldMetaRadio(operations.EnvelopeGetManyFieldMetaRadio{/* values here */})
```

### EnvelopeGetManyFieldMetaCheckbox

```go
envelopeGetManyFieldMetaUnion := operations.CreateEnvelopeGetManyFieldMetaUnionEnvelopeGetManyFieldMetaCheckbox(operations.EnvelopeGetManyFieldMetaCheckbox{/* values here */})
```

### EnvelopeGetManyFieldMetaDropdown

```go
envelopeGetManyFieldMetaUnion := operations.CreateEnvelopeGetManyFieldMetaUnionEnvelopeGetManyFieldMetaDropdown(operations.EnvelopeGetManyFieldMetaDropdown{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch envelopeGetManyFieldMetaUnion.Type {
	case operations.EnvelopeGetManyFieldMetaUnionTypeEnvelopeGetManyFieldMetaSignature:
		// envelopeGetManyFieldMetaUnion.EnvelopeGetManyFieldMetaSignature is populated
	case operations.EnvelopeGetManyFieldMetaUnionTypeEnvelopeGetManyFieldMetaInitials:
		// envelopeGetManyFieldMetaUnion.EnvelopeGetManyFieldMetaInitials is populated
	case operations.EnvelopeGetManyFieldMetaUnionTypeEnvelopeGetManyFieldMetaName:
		// envelopeGetManyFieldMetaUnion.EnvelopeGetManyFieldMetaName is populated
	case operations.EnvelopeGetManyFieldMetaUnionTypeEnvelopeGetManyFieldMetaEmail:
		// envelopeGetManyFieldMetaUnion.EnvelopeGetManyFieldMetaEmail is populated
	case operations.EnvelopeGetManyFieldMetaUnionTypeEnvelopeGetManyFieldMetaDate:
		// envelopeGetManyFieldMetaUnion.EnvelopeGetManyFieldMetaDate is populated
	case operations.EnvelopeGetManyFieldMetaUnionTypeEnvelopeGetManyFieldMetaText:
		// envelopeGetManyFieldMetaUnion.EnvelopeGetManyFieldMetaText is populated
	case operations.EnvelopeGetManyFieldMetaUnionTypeEnvelopeGetManyFieldMetaNumber:
		// envelopeGetManyFieldMetaUnion.EnvelopeGetManyFieldMetaNumber is populated
	case operations.EnvelopeGetManyFieldMetaUnionTypeEnvelopeGetManyFieldMetaRadio:
		// envelopeGetManyFieldMetaUnion.EnvelopeGetManyFieldMetaRadio is populated
	case operations.EnvelopeGetManyFieldMetaUnionTypeEnvelopeGetManyFieldMetaCheckbox:
		// envelopeGetManyFieldMetaUnion.EnvelopeGetManyFieldMetaCheckbox is populated
	case operations.EnvelopeGetManyFieldMetaUnionTypeEnvelopeGetManyFieldMetaDropdown:
		// envelopeGetManyFieldMetaUnion.EnvelopeGetManyFieldMetaDropdown is populated
}
```
