# EnvelopeFieldUpdateManyFieldMetaUnion


## Supported Types

### EnvelopeFieldUpdateManyFieldMetaSignatureResponse

```go
envelopeFieldUpdateManyFieldMetaUnion := operations.CreateEnvelopeFieldUpdateManyFieldMetaUnionEnvelopeFieldUpdateManyFieldMetaSignatureResponse(operations.EnvelopeFieldUpdateManyFieldMetaSignatureResponse{/* values here */})
```

### EnvelopeFieldUpdateManyFieldMetaInitialsResponse

```go
envelopeFieldUpdateManyFieldMetaUnion := operations.CreateEnvelopeFieldUpdateManyFieldMetaUnionEnvelopeFieldUpdateManyFieldMetaInitialsResponse(operations.EnvelopeFieldUpdateManyFieldMetaInitialsResponse{/* values here */})
```

### EnvelopeFieldUpdateManyFieldMetaNameResponse

```go
envelopeFieldUpdateManyFieldMetaUnion := operations.CreateEnvelopeFieldUpdateManyFieldMetaUnionEnvelopeFieldUpdateManyFieldMetaNameResponse(operations.EnvelopeFieldUpdateManyFieldMetaNameResponse{/* values here */})
```

### EnvelopeFieldUpdateManyFieldMetaEmailResponse

```go
envelopeFieldUpdateManyFieldMetaUnion := operations.CreateEnvelopeFieldUpdateManyFieldMetaUnionEnvelopeFieldUpdateManyFieldMetaEmailResponse(operations.EnvelopeFieldUpdateManyFieldMetaEmailResponse{/* values here */})
```

### EnvelopeFieldUpdateManyFieldMetaDateResponse

```go
envelopeFieldUpdateManyFieldMetaUnion := operations.CreateEnvelopeFieldUpdateManyFieldMetaUnionEnvelopeFieldUpdateManyFieldMetaDateResponse(operations.EnvelopeFieldUpdateManyFieldMetaDateResponse{/* values here */})
```

### EnvelopeFieldUpdateManyFieldMetaTextResponse

```go
envelopeFieldUpdateManyFieldMetaUnion := operations.CreateEnvelopeFieldUpdateManyFieldMetaUnionEnvelopeFieldUpdateManyFieldMetaTextResponse(operations.EnvelopeFieldUpdateManyFieldMetaTextResponse{/* values here */})
```

### EnvelopeFieldUpdateManyFieldMetaNumberResponse

```go
envelopeFieldUpdateManyFieldMetaUnion := operations.CreateEnvelopeFieldUpdateManyFieldMetaUnionEnvelopeFieldUpdateManyFieldMetaNumberResponse(operations.EnvelopeFieldUpdateManyFieldMetaNumberResponse{/* values here */})
```

### EnvelopeFieldUpdateManyFieldMetaRadioResponse

```go
envelopeFieldUpdateManyFieldMetaUnion := operations.CreateEnvelopeFieldUpdateManyFieldMetaUnionEnvelopeFieldUpdateManyFieldMetaRadioResponse(operations.EnvelopeFieldUpdateManyFieldMetaRadioResponse{/* values here */})
```

### EnvelopeFieldUpdateManyFieldMetaCheckboxResponse

```go
envelopeFieldUpdateManyFieldMetaUnion := operations.CreateEnvelopeFieldUpdateManyFieldMetaUnionEnvelopeFieldUpdateManyFieldMetaCheckboxResponse(operations.EnvelopeFieldUpdateManyFieldMetaCheckboxResponse{/* values here */})
```

### EnvelopeFieldUpdateManyFieldMetaDropdownResponse

```go
envelopeFieldUpdateManyFieldMetaUnion := operations.CreateEnvelopeFieldUpdateManyFieldMetaUnionEnvelopeFieldUpdateManyFieldMetaDropdownResponse(operations.EnvelopeFieldUpdateManyFieldMetaDropdownResponse{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch envelopeFieldUpdateManyFieldMetaUnion.Type {
	case operations.EnvelopeFieldUpdateManyFieldMetaUnionTypeEnvelopeFieldUpdateManyFieldMetaSignatureResponse:
		// envelopeFieldUpdateManyFieldMetaUnion.EnvelopeFieldUpdateManyFieldMetaSignatureResponse is populated
	case operations.EnvelopeFieldUpdateManyFieldMetaUnionTypeEnvelopeFieldUpdateManyFieldMetaInitialsResponse:
		// envelopeFieldUpdateManyFieldMetaUnion.EnvelopeFieldUpdateManyFieldMetaInitialsResponse is populated
	case operations.EnvelopeFieldUpdateManyFieldMetaUnionTypeEnvelopeFieldUpdateManyFieldMetaNameResponse:
		// envelopeFieldUpdateManyFieldMetaUnion.EnvelopeFieldUpdateManyFieldMetaNameResponse is populated
	case operations.EnvelopeFieldUpdateManyFieldMetaUnionTypeEnvelopeFieldUpdateManyFieldMetaEmailResponse:
		// envelopeFieldUpdateManyFieldMetaUnion.EnvelopeFieldUpdateManyFieldMetaEmailResponse is populated
	case operations.EnvelopeFieldUpdateManyFieldMetaUnionTypeEnvelopeFieldUpdateManyFieldMetaDateResponse:
		// envelopeFieldUpdateManyFieldMetaUnion.EnvelopeFieldUpdateManyFieldMetaDateResponse is populated
	case operations.EnvelopeFieldUpdateManyFieldMetaUnionTypeEnvelopeFieldUpdateManyFieldMetaTextResponse:
		// envelopeFieldUpdateManyFieldMetaUnion.EnvelopeFieldUpdateManyFieldMetaTextResponse is populated
	case operations.EnvelopeFieldUpdateManyFieldMetaUnionTypeEnvelopeFieldUpdateManyFieldMetaNumberResponse:
		// envelopeFieldUpdateManyFieldMetaUnion.EnvelopeFieldUpdateManyFieldMetaNumberResponse is populated
	case operations.EnvelopeFieldUpdateManyFieldMetaUnionTypeEnvelopeFieldUpdateManyFieldMetaRadioResponse:
		// envelopeFieldUpdateManyFieldMetaUnion.EnvelopeFieldUpdateManyFieldMetaRadioResponse is populated
	case operations.EnvelopeFieldUpdateManyFieldMetaUnionTypeEnvelopeFieldUpdateManyFieldMetaCheckboxResponse:
		// envelopeFieldUpdateManyFieldMetaUnion.EnvelopeFieldUpdateManyFieldMetaCheckboxResponse is populated
	case operations.EnvelopeFieldUpdateManyFieldMetaUnionTypeEnvelopeFieldUpdateManyFieldMetaDropdownResponse:
		// envelopeFieldUpdateManyFieldMetaUnion.EnvelopeFieldUpdateManyFieldMetaDropdownResponse is populated
}
```
