# EnvelopeRecipientGetFieldMetaUnion


## Supported Types

### EnvelopeRecipientGetFieldMetaSignature

```go
envelopeRecipientGetFieldMetaUnion := operations.CreateEnvelopeRecipientGetFieldMetaUnionEnvelopeRecipientGetFieldMetaSignature(operations.EnvelopeRecipientGetFieldMetaSignature{/* values here */})
```

### EnvelopeRecipientGetFieldMetaInitials

```go
envelopeRecipientGetFieldMetaUnion := operations.CreateEnvelopeRecipientGetFieldMetaUnionEnvelopeRecipientGetFieldMetaInitials(operations.EnvelopeRecipientGetFieldMetaInitials{/* values here */})
```

### EnvelopeRecipientGetFieldMetaName

```go
envelopeRecipientGetFieldMetaUnion := operations.CreateEnvelopeRecipientGetFieldMetaUnionEnvelopeRecipientGetFieldMetaName(operations.EnvelopeRecipientGetFieldMetaName{/* values here */})
```

### EnvelopeRecipientGetFieldMetaEmail

```go
envelopeRecipientGetFieldMetaUnion := operations.CreateEnvelopeRecipientGetFieldMetaUnionEnvelopeRecipientGetFieldMetaEmail(operations.EnvelopeRecipientGetFieldMetaEmail{/* values here */})
```

### EnvelopeRecipientGetFieldMetaDate

```go
envelopeRecipientGetFieldMetaUnion := operations.CreateEnvelopeRecipientGetFieldMetaUnionEnvelopeRecipientGetFieldMetaDate(operations.EnvelopeRecipientGetFieldMetaDate{/* values here */})
```

### EnvelopeRecipientGetFieldMetaText

```go
envelopeRecipientGetFieldMetaUnion := operations.CreateEnvelopeRecipientGetFieldMetaUnionEnvelopeRecipientGetFieldMetaText(operations.EnvelopeRecipientGetFieldMetaText{/* values here */})
```

### EnvelopeRecipientGetFieldMetaNumber

```go
envelopeRecipientGetFieldMetaUnion := operations.CreateEnvelopeRecipientGetFieldMetaUnionEnvelopeRecipientGetFieldMetaNumber(operations.EnvelopeRecipientGetFieldMetaNumber{/* values here */})
```

### EnvelopeRecipientGetFieldMetaRadio

```go
envelopeRecipientGetFieldMetaUnion := operations.CreateEnvelopeRecipientGetFieldMetaUnionEnvelopeRecipientGetFieldMetaRadio(operations.EnvelopeRecipientGetFieldMetaRadio{/* values here */})
```

### EnvelopeRecipientGetFieldMetaCheckbox

```go
envelopeRecipientGetFieldMetaUnion := operations.CreateEnvelopeRecipientGetFieldMetaUnionEnvelopeRecipientGetFieldMetaCheckbox(operations.EnvelopeRecipientGetFieldMetaCheckbox{/* values here */})
```

### EnvelopeRecipientGetFieldMetaDropdown

```go
envelopeRecipientGetFieldMetaUnion := operations.CreateEnvelopeRecipientGetFieldMetaUnionEnvelopeRecipientGetFieldMetaDropdown(operations.EnvelopeRecipientGetFieldMetaDropdown{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch envelopeRecipientGetFieldMetaUnion.Type {
	case operations.EnvelopeRecipientGetFieldMetaUnionTypeEnvelopeRecipientGetFieldMetaSignature:
		// envelopeRecipientGetFieldMetaUnion.EnvelopeRecipientGetFieldMetaSignature is populated
	case operations.EnvelopeRecipientGetFieldMetaUnionTypeEnvelopeRecipientGetFieldMetaInitials:
		// envelopeRecipientGetFieldMetaUnion.EnvelopeRecipientGetFieldMetaInitials is populated
	case operations.EnvelopeRecipientGetFieldMetaUnionTypeEnvelopeRecipientGetFieldMetaName:
		// envelopeRecipientGetFieldMetaUnion.EnvelopeRecipientGetFieldMetaName is populated
	case operations.EnvelopeRecipientGetFieldMetaUnionTypeEnvelopeRecipientGetFieldMetaEmail:
		// envelopeRecipientGetFieldMetaUnion.EnvelopeRecipientGetFieldMetaEmail is populated
	case operations.EnvelopeRecipientGetFieldMetaUnionTypeEnvelopeRecipientGetFieldMetaDate:
		// envelopeRecipientGetFieldMetaUnion.EnvelopeRecipientGetFieldMetaDate is populated
	case operations.EnvelopeRecipientGetFieldMetaUnionTypeEnvelopeRecipientGetFieldMetaText:
		// envelopeRecipientGetFieldMetaUnion.EnvelopeRecipientGetFieldMetaText is populated
	case operations.EnvelopeRecipientGetFieldMetaUnionTypeEnvelopeRecipientGetFieldMetaNumber:
		// envelopeRecipientGetFieldMetaUnion.EnvelopeRecipientGetFieldMetaNumber is populated
	case operations.EnvelopeRecipientGetFieldMetaUnionTypeEnvelopeRecipientGetFieldMetaRadio:
		// envelopeRecipientGetFieldMetaUnion.EnvelopeRecipientGetFieldMetaRadio is populated
	case operations.EnvelopeRecipientGetFieldMetaUnionTypeEnvelopeRecipientGetFieldMetaCheckbox:
		// envelopeRecipientGetFieldMetaUnion.EnvelopeRecipientGetFieldMetaCheckbox is populated
	case operations.EnvelopeRecipientGetFieldMetaUnionTypeEnvelopeRecipientGetFieldMetaDropdown:
		// envelopeRecipientGetFieldMetaUnion.EnvelopeRecipientGetFieldMetaDropdown is populated
}
```
