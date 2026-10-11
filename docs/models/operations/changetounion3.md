# ChangeToUnion3


## Supported Types

### 

```go
changeToUnion3 := operations.CreateChangeToUnion3Str(string{/* values here */})
```

### 

```go
changeToUnion3 := operations.CreateChangeToUnion3ArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch changeToUnion3.Type {
	case operations.ChangeToUnion3TypeStr:
		// changeToUnion3.Str is populated
	case operations.ChangeToUnion3TypeArrayOfStr:
		// changeToUnion3.ArrayOfStr is populated
}
```
