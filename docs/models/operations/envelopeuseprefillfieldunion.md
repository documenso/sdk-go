# EnvelopeUsePrefillFieldUnion


## Supported Types

### EnvelopeUsePrefillFieldText

```go
envelopeUsePrefillFieldUnion := operations.CreateEnvelopeUsePrefillFieldUnionEnvelopeUsePrefillFieldText(operations.EnvelopeUsePrefillFieldText{/* values here */})
```

### EnvelopeUsePrefillFieldNumber

```go
envelopeUsePrefillFieldUnion := operations.CreateEnvelopeUsePrefillFieldUnionEnvelopeUsePrefillFieldNumber(operations.EnvelopeUsePrefillFieldNumber{/* values here */})
```

### EnvelopeUsePrefillFieldRadio

```go
envelopeUsePrefillFieldUnion := operations.CreateEnvelopeUsePrefillFieldUnionEnvelopeUsePrefillFieldRadio(operations.EnvelopeUsePrefillFieldRadio{/* values here */})
```

### EnvelopeUsePrefillFieldCheckbox

```go
envelopeUsePrefillFieldUnion := operations.CreateEnvelopeUsePrefillFieldUnionEnvelopeUsePrefillFieldCheckbox(operations.EnvelopeUsePrefillFieldCheckbox{/* values here */})
```

### EnvelopeUsePrefillFieldDropdown

```go
envelopeUsePrefillFieldUnion := operations.CreateEnvelopeUsePrefillFieldUnionEnvelopeUsePrefillFieldDropdown(operations.EnvelopeUsePrefillFieldDropdown{/* values here */})
```

### EnvelopeUsePrefillFieldDate

```go
envelopeUsePrefillFieldUnion := operations.CreateEnvelopeUsePrefillFieldUnionEnvelopeUsePrefillFieldDate(operations.EnvelopeUsePrefillFieldDate{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch envelopeUsePrefillFieldUnion.Type {
	case operations.EnvelopeUsePrefillFieldUnionTypeEnvelopeUsePrefillFieldText:
		// envelopeUsePrefillFieldUnion.EnvelopeUsePrefillFieldText is populated
	case operations.EnvelopeUsePrefillFieldUnionTypeEnvelopeUsePrefillFieldNumber:
		// envelopeUsePrefillFieldUnion.EnvelopeUsePrefillFieldNumber is populated
	case operations.EnvelopeUsePrefillFieldUnionTypeEnvelopeUsePrefillFieldRadio:
		// envelopeUsePrefillFieldUnion.EnvelopeUsePrefillFieldRadio is populated
	case operations.EnvelopeUsePrefillFieldUnionTypeEnvelopeUsePrefillFieldCheckbox:
		// envelopeUsePrefillFieldUnion.EnvelopeUsePrefillFieldCheckbox is populated
	case operations.EnvelopeUsePrefillFieldUnionTypeEnvelopeUsePrefillFieldDropdown:
		// envelopeUsePrefillFieldUnion.EnvelopeUsePrefillFieldDropdown is populated
	case operations.EnvelopeUsePrefillFieldUnionTypeEnvelopeUsePrefillFieldDate:
		// envelopeUsePrefillFieldUnion.EnvelopeUsePrefillFieldDate is populated
}
```
