# EnvelopeCreateRepeatEveryUnion


## Supported Types

### EnvelopeCreateRepeatEvery1

```go
envelopeCreateRepeatEveryUnion := operations.CreateEnvelopeCreateRepeatEveryUnionEnvelopeCreateRepeatEvery1(operations.EnvelopeCreateRepeatEvery1{/* values here */})
```

### EnvelopeCreateRepeatEvery2

```go
envelopeCreateRepeatEveryUnion := operations.CreateEnvelopeCreateRepeatEveryUnionEnvelopeCreateRepeatEvery2(operations.EnvelopeCreateRepeatEvery2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch envelopeCreateRepeatEveryUnion.Type {
	case operations.EnvelopeCreateRepeatEveryUnionTypeEnvelopeCreateRepeatEvery1:
		// envelopeCreateRepeatEveryUnion.EnvelopeCreateRepeatEvery1 is populated
	case operations.EnvelopeCreateRepeatEveryUnionTypeEnvelopeCreateRepeatEvery2:
		// envelopeCreateRepeatEveryUnion.EnvelopeCreateRepeatEvery2 is populated
}
```
