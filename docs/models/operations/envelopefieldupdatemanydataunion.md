# EnvelopeFieldUpdateManyDataUnion


## Supported Types

### EnvelopeFieldUpdateManyDataSignature

```go
envelopeFieldUpdateManyDataUnion := operations.CreateEnvelopeFieldUpdateManyDataUnionEnvelopeFieldUpdateManyDataSignature(operations.EnvelopeFieldUpdateManyDataSignature{/* values here */})
```

### EnvelopeFieldUpdateManyDataFreeSignature

```go
envelopeFieldUpdateManyDataUnion := operations.CreateEnvelopeFieldUpdateManyDataUnionEnvelopeFieldUpdateManyDataFreeSignature(operations.EnvelopeFieldUpdateManyDataFreeSignature{/* values here */})
```

### EnvelopeFieldUpdateManyDataInitials

```go
envelopeFieldUpdateManyDataUnion := operations.CreateEnvelopeFieldUpdateManyDataUnionEnvelopeFieldUpdateManyDataInitials(operations.EnvelopeFieldUpdateManyDataInitials{/* values here */})
```

### EnvelopeFieldUpdateManyDataName

```go
envelopeFieldUpdateManyDataUnion := operations.CreateEnvelopeFieldUpdateManyDataUnionEnvelopeFieldUpdateManyDataName(operations.EnvelopeFieldUpdateManyDataName{/* values here */})
```

### EnvelopeFieldUpdateManyDataEmail

```go
envelopeFieldUpdateManyDataUnion := operations.CreateEnvelopeFieldUpdateManyDataUnionEnvelopeFieldUpdateManyDataEmail(operations.EnvelopeFieldUpdateManyDataEmail{/* values here */})
```

### EnvelopeFieldUpdateManyDataDate

```go
envelopeFieldUpdateManyDataUnion := operations.CreateEnvelopeFieldUpdateManyDataUnionEnvelopeFieldUpdateManyDataDate(operations.EnvelopeFieldUpdateManyDataDate{/* values here */})
```

### EnvelopeFieldUpdateManyDataText

```go
envelopeFieldUpdateManyDataUnion := operations.CreateEnvelopeFieldUpdateManyDataUnionEnvelopeFieldUpdateManyDataText(operations.EnvelopeFieldUpdateManyDataText{/* values here */})
```

### EnvelopeFieldUpdateManyDataNumber

```go
envelopeFieldUpdateManyDataUnion := operations.CreateEnvelopeFieldUpdateManyDataUnionEnvelopeFieldUpdateManyDataNumber(operations.EnvelopeFieldUpdateManyDataNumber{/* values here */})
```

### EnvelopeFieldUpdateManyDataRadio

```go
envelopeFieldUpdateManyDataUnion := operations.CreateEnvelopeFieldUpdateManyDataUnionEnvelopeFieldUpdateManyDataRadio(operations.EnvelopeFieldUpdateManyDataRadio{/* values here */})
```

### EnvelopeFieldUpdateManyDataCheckbox

```go
envelopeFieldUpdateManyDataUnion := operations.CreateEnvelopeFieldUpdateManyDataUnionEnvelopeFieldUpdateManyDataCheckbox(operations.EnvelopeFieldUpdateManyDataCheckbox{/* values here */})
```

### EnvelopeFieldUpdateManyDataDropdown

```go
envelopeFieldUpdateManyDataUnion := operations.CreateEnvelopeFieldUpdateManyDataUnionEnvelopeFieldUpdateManyDataDropdown(operations.EnvelopeFieldUpdateManyDataDropdown{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch envelopeFieldUpdateManyDataUnion.Type {
	case operations.EnvelopeFieldUpdateManyDataUnionTypeEnvelopeFieldUpdateManyDataSignature:
		// envelopeFieldUpdateManyDataUnion.EnvelopeFieldUpdateManyDataSignature is populated
	case operations.EnvelopeFieldUpdateManyDataUnionTypeEnvelopeFieldUpdateManyDataFreeSignature:
		// envelopeFieldUpdateManyDataUnion.EnvelopeFieldUpdateManyDataFreeSignature is populated
	case operations.EnvelopeFieldUpdateManyDataUnionTypeEnvelopeFieldUpdateManyDataInitials:
		// envelopeFieldUpdateManyDataUnion.EnvelopeFieldUpdateManyDataInitials is populated
	case operations.EnvelopeFieldUpdateManyDataUnionTypeEnvelopeFieldUpdateManyDataName:
		// envelopeFieldUpdateManyDataUnion.EnvelopeFieldUpdateManyDataName is populated
	case operations.EnvelopeFieldUpdateManyDataUnionTypeEnvelopeFieldUpdateManyDataEmail:
		// envelopeFieldUpdateManyDataUnion.EnvelopeFieldUpdateManyDataEmail is populated
	case operations.EnvelopeFieldUpdateManyDataUnionTypeEnvelopeFieldUpdateManyDataDate:
		// envelopeFieldUpdateManyDataUnion.EnvelopeFieldUpdateManyDataDate is populated
	case operations.EnvelopeFieldUpdateManyDataUnionTypeEnvelopeFieldUpdateManyDataText:
		// envelopeFieldUpdateManyDataUnion.EnvelopeFieldUpdateManyDataText is populated
	case operations.EnvelopeFieldUpdateManyDataUnionTypeEnvelopeFieldUpdateManyDataNumber:
		// envelopeFieldUpdateManyDataUnion.EnvelopeFieldUpdateManyDataNumber is populated
	case operations.EnvelopeFieldUpdateManyDataUnionTypeEnvelopeFieldUpdateManyDataRadio:
		// envelopeFieldUpdateManyDataUnion.EnvelopeFieldUpdateManyDataRadio is populated
	case operations.EnvelopeFieldUpdateManyDataUnionTypeEnvelopeFieldUpdateManyDataCheckbox:
		// envelopeFieldUpdateManyDataUnion.EnvelopeFieldUpdateManyDataCheckbox is populated
	case operations.EnvelopeFieldUpdateManyDataUnionTypeEnvelopeFieldUpdateManyDataDropdown:
		// envelopeFieldUpdateManyDataUnion.EnvelopeFieldUpdateManyDataDropdown is populated
}
```
