# FromUnion2


## Supported Types

### 

```go
fromUnion2 := operations.CreateFromUnion2Str(string{/* values here */})
```

### 

```go
fromUnion2 := operations.CreateFromUnion2ArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch fromUnion2.Type {
	case operations.FromUnion2TypeStr:
		// fromUnion2.Str is populated
	case operations.FromUnion2TypeArrayOfStr:
		// fromUnion2.ArrayOfStr is populated
}
```
