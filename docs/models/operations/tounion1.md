# ToUnion1


## Supported Types

### 

```go
toUnion1 := operations.CreateToUnion1Str(string{/* values here */})
```

### 

```go
toUnion1 := operations.CreateToUnion1ArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch toUnion1.Type {
	case operations.ToUnion1TypeStr:
		// toUnion1.Str is populated
	case operations.ToUnion1TypeArrayOfStr:
		// toUnion1.ArrayOfStr is populated
}
```
