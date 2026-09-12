# ChangeUnion2


## Supported Types

### ChangeDimension

```go
changeUnion2 := operations.CreateChangeUnion2ChangeDimension(operations.ChangeDimension{/* values here */})
```

### ChangePosition

```go
changeUnion2 := operations.CreateChangeUnion2ChangePosition(operations.ChangePosition{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch changeUnion2.Type {
	case operations.ChangeUnion2TypeChangeDimension:
		// changeUnion2.ChangeDimension is populated
	case operations.ChangeUnion2TypeChangePosition:
		// changeUnion2.ChangePosition is populated
}
```
