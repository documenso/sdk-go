# EnvelopeFieldCreateManyFieldMetaUnion


## Supported Types

### EnvelopeFieldCreateManyFieldMetaSignatureResponse

```go
envelopeFieldCreateManyFieldMetaUnion := operations.CreateEnvelopeFieldCreateManyFieldMetaUnionEnvelopeFieldCreateManyFieldMetaSignatureResponse(operations.EnvelopeFieldCreateManyFieldMetaSignatureResponse{/* values here */})
```

### EnvelopeFieldCreateManyFieldMetaInitialsResponse

```go
envelopeFieldCreateManyFieldMetaUnion := operations.CreateEnvelopeFieldCreateManyFieldMetaUnionEnvelopeFieldCreateManyFieldMetaInitialsResponse(operations.EnvelopeFieldCreateManyFieldMetaInitialsResponse{/* values here */})
```

### EnvelopeFieldCreateManyFieldMetaNameResponse

```go
envelopeFieldCreateManyFieldMetaUnion := operations.CreateEnvelopeFieldCreateManyFieldMetaUnionEnvelopeFieldCreateManyFieldMetaNameResponse(operations.EnvelopeFieldCreateManyFieldMetaNameResponse{/* values here */})
```

### EnvelopeFieldCreateManyFieldMetaEmailResponse

```go
envelopeFieldCreateManyFieldMetaUnion := operations.CreateEnvelopeFieldCreateManyFieldMetaUnionEnvelopeFieldCreateManyFieldMetaEmailResponse(operations.EnvelopeFieldCreateManyFieldMetaEmailResponse{/* values here */})
```

### EnvelopeFieldCreateManyFieldMetaDateResponse

```go
envelopeFieldCreateManyFieldMetaUnion := operations.CreateEnvelopeFieldCreateManyFieldMetaUnionEnvelopeFieldCreateManyFieldMetaDateResponse(operations.EnvelopeFieldCreateManyFieldMetaDateResponse{/* values here */})
```

### EnvelopeFieldCreateManyFieldMetaTextResponse

```go
envelopeFieldCreateManyFieldMetaUnion := operations.CreateEnvelopeFieldCreateManyFieldMetaUnionEnvelopeFieldCreateManyFieldMetaTextResponse(operations.EnvelopeFieldCreateManyFieldMetaTextResponse{/* values here */})
```

### EnvelopeFieldCreateManyFieldMetaNumberResponse

```go
envelopeFieldCreateManyFieldMetaUnion := operations.CreateEnvelopeFieldCreateManyFieldMetaUnionEnvelopeFieldCreateManyFieldMetaNumberResponse(operations.EnvelopeFieldCreateManyFieldMetaNumberResponse{/* values here */})
```

### EnvelopeFieldCreateManyFieldMetaRadioResponse

```go
envelopeFieldCreateManyFieldMetaUnion := operations.CreateEnvelopeFieldCreateManyFieldMetaUnionEnvelopeFieldCreateManyFieldMetaRadioResponse(operations.EnvelopeFieldCreateManyFieldMetaRadioResponse{/* values here */})
```

### EnvelopeFieldCreateManyFieldMetaCheckboxResponse

```go
envelopeFieldCreateManyFieldMetaUnion := operations.CreateEnvelopeFieldCreateManyFieldMetaUnionEnvelopeFieldCreateManyFieldMetaCheckboxResponse(operations.EnvelopeFieldCreateManyFieldMetaCheckboxResponse{/* values here */})
```

### EnvelopeFieldCreateManyFieldMetaDropdownResponse

```go
envelopeFieldCreateManyFieldMetaUnion := operations.CreateEnvelopeFieldCreateManyFieldMetaUnionEnvelopeFieldCreateManyFieldMetaDropdownResponse(operations.EnvelopeFieldCreateManyFieldMetaDropdownResponse{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch envelopeFieldCreateManyFieldMetaUnion.Type {
	case operations.EnvelopeFieldCreateManyFieldMetaUnionTypeEnvelopeFieldCreateManyFieldMetaSignatureResponse:
		// envelopeFieldCreateManyFieldMetaUnion.EnvelopeFieldCreateManyFieldMetaSignatureResponse is populated
	case operations.EnvelopeFieldCreateManyFieldMetaUnionTypeEnvelopeFieldCreateManyFieldMetaInitialsResponse:
		// envelopeFieldCreateManyFieldMetaUnion.EnvelopeFieldCreateManyFieldMetaInitialsResponse is populated
	case operations.EnvelopeFieldCreateManyFieldMetaUnionTypeEnvelopeFieldCreateManyFieldMetaNameResponse:
		// envelopeFieldCreateManyFieldMetaUnion.EnvelopeFieldCreateManyFieldMetaNameResponse is populated
	case operations.EnvelopeFieldCreateManyFieldMetaUnionTypeEnvelopeFieldCreateManyFieldMetaEmailResponse:
		// envelopeFieldCreateManyFieldMetaUnion.EnvelopeFieldCreateManyFieldMetaEmailResponse is populated
	case operations.EnvelopeFieldCreateManyFieldMetaUnionTypeEnvelopeFieldCreateManyFieldMetaDateResponse:
		// envelopeFieldCreateManyFieldMetaUnion.EnvelopeFieldCreateManyFieldMetaDateResponse is populated
	case operations.EnvelopeFieldCreateManyFieldMetaUnionTypeEnvelopeFieldCreateManyFieldMetaTextResponse:
		// envelopeFieldCreateManyFieldMetaUnion.EnvelopeFieldCreateManyFieldMetaTextResponse is populated
	case operations.EnvelopeFieldCreateManyFieldMetaUnionTypeEnvelopeFieldCreateManyFieldMetaNumberResponse:
		// envelopeFieldCreateManyFieldMetaUnion.EnvelopeFieldCreateManyFieldMetaNumberResponse is populated
	case operations.EnvelopeFieldCreateManyFieldMetaUnionTypeEnvelopeFieldCreateManyFieldMetaRadioResponse:
		// envelopeFieldCreateManyFieldMetaUnion.EnvelopeFieldCreateManyFieldMetaRadioResponse is populated
	case operations.EnvelopeFieldCreateManyFieldMetaUnionTypeEnvelopeFieldCreateManyFieldMetaCheckboxResponse:
		// envelopeFieldCreateManyFieldMetaUnion.EnvelopeFieldCreateManyFieldMetaCheckboxResponse is populated
	case operations.EnvelopeFieldCreateManyFieldMetaUnionTypeEnvelopeFieldCreateManyFieldMetaDropdownResponse:
		// envelopeFieldCreateManyFieldMetaUnion.EnvelopeFieldCreateManyFieldMetaDropdownResponse is populated
}
```
