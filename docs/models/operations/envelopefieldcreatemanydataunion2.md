# EnvelopeFieldCreateManyDataUnion2


## Supported Types

### EnvelopeFieldCreateManyDataSignature1

```go
envelopeFieldCreateManyDataUnion2 := operations.CreateEnvelopeFieldCreateManyDataUnion2EnvelopeFieldCreateManyDataSignature1(operations.EnvelopeFieldCreateManyDataSignature1{/* values here */})
```

### EnvelopeFieldCreateManyDataFreeSignature1

```go
envelopeFieldCreateManyDataUnion2 := operations.CreateEnvelopeFieldCreateManyDataUnion2EnvelopeFieldCreateManyDataFreeSignature1(operations.EnvelopeFieldCreateManyDataFreeSignature1{/* values here */})
```

### EnvelopeFieldCreateManyDataInitials1

```go
envelopeFieldCreateManyDataUnion2 := operations.CreateEnvelopeFieldCreateManyDataUnion2EnvelopeFieldCreateManyDataInitials1(operations.EnvelopeFieldCreateManyDataInitials1{/* values here */})
```

### EnvelopeFieldCreateManyDataName1

```go
envelopeFieldCreateManyDataUnion2 := operations.CreateEnvelopeFieldCreateManyDataUnion2EnvelopeFieldCreateManyDataName1(operations.EnvelopeFieldCreateManyDataName1{/* values here */})
```

### EnvelopeFieldCreateManyDataEmail1

```go
envelopeFieldCreateManyDataUnion2 := operations.CreateEnvelopeFieldCreateManyDataUnion2EnvelopeFieldCreateManyDataEmail1(operations.EnvelopeFieldCreateManyDataEmail1{/* values here */})
```

### EnvelopeFieldCreateManyDataDate1

```go
envelopeFieldCreateManyDataUnion2 := operations.CreateEnvelopeFieldCreateManyDataUnion2EnvelopeFieldCreateManyDataDate1(operations.EnvelopeFieldCreateManyDataDate1{/* values here */})
```

### EnvelopeFieldCreateManyDataText1

```go
envelopeFieldCreateManyDataUnion2 := operations.CreateEnvelopeFieldCreateManyDataUnion2EnvelopeFieldCreateManyDataText1(operations.EnvelopeFieldCreateManyDataText1{/* values here */})
```

### EnvelopeFieldCreateManyDataNumber1

```go
envelopeFieldCreateManyDataUnion2 := operations.CreateEnvelopeFieldCreateManyDataUnion2EnvelopeFieldCreateManyDataNumber1(operations.EnvelopeFieldCreateManyDataNumber1{/* values here */})
```

### EnvelopeFieldCreateManyDataRadio1

```go
envelopeFieldCreateManyDataUnion2 := operations.CreateEnvelopeFieldCreateManyDataUnion2EnvelopeFieldCreateManyDataRadio1(operations.EnvelopeFieldCreateManyDataRadio1{/* values here */})
```

### EnvelopeFieldCreateManyDataCheckbox1

```go
envelopeFieldCreateManyDataUnion2 := operations.CreateEnvelopeFieldCreateManyDataUnion2EnvelopeFieldCreateManyDataCheckbox1(operations.EnvelopeFieldCreateManyDataCheckbox1{/* values here */})
```

### EnvelopeFieldCreateManyDataDropdown1

```go
envelopeFieldCreateManyDataUnion2 := operations.CreateEnvelopeFieldCreateManyDataUnion2EnvelopeFieldCreateManyDataDropdown1(operations.EnvelopeFieldCreateManyDataDropdown1{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch envelopeFieldCreateManyDataUnion2.Type {
	case operations.EnvelopeFieldCreateManyDataUnion2TypeEnvelopeFieldCreateManyDataSignature1:
		// envelopeFieldCreateManyDataUnion2.EnvelopeFieldCreateManyDataSignature1 is populated
	case operations.EnvelopeFieldCreateManyDataUnion2TypeEnvelopeFieldCreateManyDataFreeSignature1:
		// envelopeFieldCreateManyDataUnion2.EnvelopeFieldCreateManyDataFreeSignature1 is populated
	case operations.EnvelopeFieldCreateManyDataUnion2TypeEnvelopeFieldCreateManyDataInitials1:
		// envelopeFieldCreateManyDataUnion2.EnvelopeFieldCreateManyDataInitials1 is populated
	case operations.EnvelopeFieldCreateManyDataUnion2TypeEnvelopeFieldCreateManyDataName1:
		// envelopeFieldCreateManyDataUnion2.EnvelopeFieldCreateManyDataName1 is populated
	case operations.EnvelopeFieldCreateManyDataUnion2TypeEnvelopeFieldCreateManyDataEmail1:
		// envelopeFieldCreateManyDataUnion2.EnvelopeFieldCreateManyDataEmail1 is populated
	case operations.EnvelopeFieldCreateManyDataUnion2TypeEnvelopeFieldCreateManyDataDate1:
		// envelopeFieldCreateManyDataUnion2.EnvelopeFieldCreateManyDataDate1 is populated
	case operations.EnvelopeFieldCreateManyDataUnion2TypeEnvelopeFieldCreateManyDataText1:
		// envelopeFieldCreateManyDataUnion2.EnvelopeFieldCreateManyDataText1 is populated
	case operations.EnvelopeFieldCreateManyDataUnion2TypeEnvelopeFieldCreateManyDataNumber1:
		// envelopeFieldCreateManyDataUnion2.EnvelopeFieldCreateManyDataNumber1 is populated
	case operations.EnvelopeFieldCreateManyDataUnion2TypeEnvelopeFieldCreateManyDataRadio1:
		// envelopeFieldCreateManyDataUnion2.EnvelopeFieldCreateManyDataRadio1 is populated
	case operations.EnvelopeFieldCreateManyDataUnion2TypeEnvelopeFieldCreateManyDataCheckbox1:
		// envelopeFieldCreateManyDataUnion2.EnvelopeFieldCreateManyDataCheckbox1 is populated
	case operations.EnvelopeFieldCreateManyDataUnion2TypeEnvelopeFieldCreateManyDataDropdown1:
		// envelopeFieldCreateManyDataUnion2.EnvelopeFieldCreateManyDataDropdown1 is populated
}
```
