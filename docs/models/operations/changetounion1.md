# ChangeToUnion1


## Supported Types

### 

```go
changeToUnion1 := operations.CreateChangeToUnion1Str(string{/* values here */})
```

### 

```go
changeToUnion1 := operations.CreateChangeToUnion1ArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch changeToUnion1.Type {
	case operations.ChangeToUnion1TypeStr:
		// changeToUnion1.Str is populated
	case operations.ChangeToUnion1TypeArrayOfStr:
		// changeToUnion1.ArrayOfStr is populated
}
```
