# ToUnion2


## Supported Types

### 

```go
toUnion2 := operations.CreateToUnion2Str(string{/* values here */})
```

### 

```go
toUnion2 := operations.CreateToUnion2ArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch toUnion2.Type {
	case operations.ToUnion2TypeStr:
		// toUnion2.Str is populated
	case operations.ToUnion2TypeArrayOfStr:
		// toUnion2.ArrayOfStr is populated
}
```
