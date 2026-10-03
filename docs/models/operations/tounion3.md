# ToUnion3


## Supported Types

### 

```go
toUnion3 := operations.CreateToUnion3Str(string{/* values here */})
```

### 

```go
toUnion3 := operations.CreateToUnion3ArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch toUnion3.Type {
	case operations.ToUnion3TypeStr:
		// toUnion3.Str is populated
	case operations.ToUnion3TypeArrayOfStr:
		// toUnion3.ArrayOfStr is populated
}
```
