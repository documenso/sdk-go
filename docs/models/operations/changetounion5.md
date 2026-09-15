# ChangeToUnion5


## Supported Types

### 

```go
changeToUnion5 := operations.CreateChangeToUnion5Str(string{/* values here */})
```

### 

```go
changeToUnion5 := operations.CreateChangeToUnion5ArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch changeToUnion5.Type {
	case operations.ChangeToUnion5TypeStr:
		// changeToUnion5.Str is populated
	case operations.ChangeToUnion5TypeArrayOfStr:
		// changeToUnion5.ArrayOfStr is populated
}
```
