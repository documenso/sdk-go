# ChangeUnion1


## Supported Types

### Change2

```go
changeUnion1 := operations.CreateChangeUnion1Change2(operations.Change2{/* values here */})
```

### ChangePassword

```go
changeUnion1 := operations.CreateChangeUnion1ChangePassword(operations.ChangePassword{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch changeUnion1.Type {
	case operations.ChangeUnion1TypeChange2:
		// changeUnion1.Change2 is populated
	case operations.ChangeUnion1TypeChangePassword:
		// changeUnion1.ChangePassword is populated
}
```
