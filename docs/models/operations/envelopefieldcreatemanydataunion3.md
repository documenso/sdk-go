# EnvelopeFieldCreateManyDataUnion3


## Supported Types

### EnvelopeFieldCreateManyDataSignature2

```go
envelopeFieldCreateManyDataUnion3 := operations.CreateEnvelopeFieldCreateManyDataUnion3EnvelopeFieldCreateManyDataSignature2(operations.EnvelopeFieldCreateManyDataSignature2{/* values here */})
```

### EnvelopeFieldCreateManyDataFreeSignature2

```go
envelopeFieldCreateManyDataUnion3 := operations.CreateEnvelopeFieldCreateManyDataUnion3EnvelopeFieldCreateManyDataFreeSignature2(operations.EnvelopeFieldCreateManyDataFreeSignature2{/* values here */})
```

### EnvelopeFieldCreateManyDataInitials2

```go
envelopeFieldCreateManyDataUnion3 := operations.CreateEnvelopeFieldCreateManyDataUnion3EnvelopeFieldCreateManyDataInitials2(operations.EnvelopeFieldCreateManyDataInitials2{/* values here */})
```

### EnvelopeFieldCreateManyDataName2

```go
envelopeFieldCreateManyDataUnion3 := operations.CreateEnvelopeFieldCreateManyDataUnion3EnvelopeFieldCreateManyDataName2(operations.EnvelopeFieldCreateManyDataName2{/* values here */})
```

### EnvelopeFieldCreateManyDataEmail2

```go
envelopeFieldCreateManyDataUnion3 := operations.CreateEnvelopeFieldCreateManyDataUnion3EnvelopeFieldCreateManyDataEmail2(operations.EnvelopeFieldCreateManyDataEmail2{/* values here */})
```

### EnvelopeFieldCreateManyDataDate2

```go
envelopeFieldCreateManyDataUnion3 := operations.CreateEnvelopeFieldCreateManyDataUnion3EnvelopeFieldCreateManyDataDate2(operations.EnvelopeFieldCreateManyDataDate2{/* values here */})
```

### EnvelopeFieldCreateManyDataText2

```go
envelopeFieldCreateManyDataUnion3 := operations.CreateEnvelopeFieldCreateManyDataUnion3EnvelopeFieldCreateManyDataText2(operations.EnvelopeFieldCreateManyDataText2{/* values here */})
```

### EnvelopeFieldCreateManyDataNumber2

```go
envelopeFieldCreateManyDataUnion3 := operations.CreateEnvelopeFieldCreateManyDataUnion3EnvelopeFieldCreateManyDataNumber2(operations.EnvelopeFieldCreateManyDataNumber2{/* values here */})
```

### EnvelopeFieldCreateManyDataRadio2

```go
envelopeFieldCreateManyDataUnion3 := operations.CreateEnvelopeFieldCreateManyDataUnion3EnvelopeFieldCreateManyDataRadio2(operations.EnvelopeFieldCreateManyDataRadio2{/* values here */})
```

### EnvelopeFieldCreateManyDataCheckbox2

```go
envelopeFieldCreateManyDataUnion3 := operations.CreateEnvelopeFieldCreateManyDataUnion3EnvelopeFieldCreateManyDataCheckbox2(operations.EnvelopeFieldCreateManyDataCheckbox2{/* values here */})
```

### EnvelopeFieldCreateManyDataDropdown2

```go
envelopeFieldCreateManyDataUnion3 := operations.CreateEnvelopeFieldCreateManyDataUnion3EnvelopeFieldCreateManyDataDropdown2(operations.EnvelopeFieldCreateManyDataDropdown2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch envelopeFieldCreateManyDataUnion3.Type {
	case operations.EnvelopeFieldCreateManyDataUnion3TypeEnvelopeFieldCreateManyDataSignature2:
		// envelopeFieldCreateManyDataUnion3.EnvelopeFieldCreateManyDataSignature2 is populated
	case operations.EnvelopeFieldCreateManyDataUnion3TypeEnvelopeFieldCreateManyDataFreeSignature2:
		// envelopeFieldCreateManyDataUnion3.EnvelopeFieldCreateManyDataFreeSignature2 is populated
	case operations.EnvelopeFieldCreateManyDataUnion3TypeEnvelopeFieldCreateManyDataInitials2:
		// envelopeFieldCreateManyDataUnion3.EnvelopeFieldCreateManyDataInitials2 is populated
	case operations.EnvelopeFieldCreateManyDataUnion3TypeEnvelopeFieldCreateManyDataName2:
		// envelopeFieldCreateManyDataUnion3.EnvelopeFieldCreateManyDataName2 is populated
	case operations.EnvelopeFieldCreateManyDataUnion3TypeEnvelopeFieldCreateManyDataEmail2:
		// envelopeFieldCreateManyDataUnion3.EnvelopeFieldCreateManyDataEmail2 is populated
	case operations.EnvelopeFieldCreateManyDataUnion3TypeEnvelopeFieldCreateManyDataDate2:
		// envelopeFieldCreateManyDataUnion3.EnvelopeFieldCreateManyDataDate2 is populated
	case operations.EnvelopeFieldCreateManyDataUnion3TypeEnvelopeFieldCreateManyDataText2:
		// envelopeFieldCreateManyDataUnion3.EnvelopeFieldCreateManyDataText2 is populated
	case operations.EnvelopeFieldCreateManyDataUnion3TypeEnvelopeFieldCreateManyDataNumber2:
		// envelopeFieldCreateManyDataUnion3.EnvelopeFieldCreateManyDataNumber2 is populated
	case operations.EnvelopeFieldCreateManyDataUnion3TypeEnvelopeFieldCreateManyDataRadio2:
		// envelopeFieldCreateManyDataUnion3.EnvelopeFieldCreateManyDataRadio2 is populated
	case operations.EnvelopeFieldCreateManyDataUnion3TypeEnvelopeFieldCreateManyDataCheckbox2:
		// envelopeFieldCreateManyDataUnion3.EnvelopeFieldCreateManyDataCheckbox2 is populated
	case operations.EnvelopeFieldCreateManyDataUnion3TypeEnvelopeFieldCreateManyDataDropdown2:
		// envelopeFieldCreateManyDataUnion3.EnvelopeFieldCreateManyDataDropdown2 is populated
}
```
