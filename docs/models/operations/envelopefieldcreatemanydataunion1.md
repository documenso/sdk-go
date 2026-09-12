# EnvelopeFieldCreateManyDataUnion1


## Supported Types

### EnvelopeFieldCreateManyDataUnion2

```go
envelopeFieldCreateManyDataUnion1 := operations.CreateEnvelopeFieldCreateManyDataUnion1EnvelopeFieldCreateManyDataUnion2(operations.EnvelopeFieldCreateManyDataUnion2{/* values here */})
```

### EnvelopeFieldCreateManyDataUnion3

```go
envelopeFieldCreateManyDataUnion1 := operations.CreateEnvelopeFieldCreateManyDataUnion1EnvelopeFieldCreateManyDataUnion3(operations.EnvelopeFieldCreateManyDataUnion3{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch envelopeFieldCreateManyDataUnion1.Type {
	case operations.EnvelopeFieldCreateManyDataUnion1TypeEnvelopeFieldCreateManyDataUnion2:
		// envelopeFieldCreateManyDataUnion1.EnvelopeFieldCreateManyDataUnion2 is populated
	case operations.EnvelopeFieldCreateManyDataUnion1TypeEnvelopeFieldCreateManyDataUnion3:
		// envelopeFieldCreateManyDataUnion1.EnvelopeFieldCreateManyDataUnion3 is populated
}
```
