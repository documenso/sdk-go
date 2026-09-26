# EnvelopeUpdateRepeatEveryUnion


## Supported Types

### EnvelopeUpdateRepeatEvery1

```go
envelopeUpdateRepeatEveryUnion := operations.CreateEnvelopeUpdateRepeatEveryUnionEnvelopeUpdateRepeatEvery1(operations.EnvelopeUpdateRepeatEvery1{/* values here */})
```

### EnvelopeUpdateRepeatEvery2

```go
envelopeUpdateRepeatEveryUnion := operations.CreateEnvelopeUpdateRepeatEveryUnionEnvelopeUpdateRepeatEvery2(operations.EnvelopeUpdateRepeatEvery2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch envelopeUpdateRepeatEveryUnion.Type {
	case operations.EnvelopeUpdateRepeatEveryUnionTypeEnvelopeUpdateRepeatEvery1:
		// envelopeUpdateRepeatEveryUnion.EnvelopeUpdateRepeatEvery1 is populated
	case operations.EnvelopeUpdateRepeatEveryUnionTypeEnvelopeUpdateRepeatEvery2:
		// envelopeUpdateRepeatEveryUnion.EnvelopeUpdateRepeatEvery2 is populated
}
```
