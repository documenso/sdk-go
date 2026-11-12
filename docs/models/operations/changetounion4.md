# ChangeToUnion4


## Supported Types

### 

```go
changeToUnion4 := operations.CreateChangeToUnion4Str(string{/* values here */})
```

### 

```go
changeToUnion4 := operations.CreateChangeToUnion4ArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch changeToUnion4.Type {
	case operations.ChangeToUnion4TypeStr:
		// changeToUnion4.Str is populated
	case operations.ChangeToUnion4TypeArrayOfStr:
		// changeToUnion4.ArrayOfStr is populated
}
```
