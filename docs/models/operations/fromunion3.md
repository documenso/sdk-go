# FromUnion3


## Supported Types

### 

```go
fromUnion3 := operations.CreateFromUnion3Str(string{/* values here */})
```

### 

```go
fromUnion3 := operations.CreateFromUnion3ArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch fromUnion3.Type {
	case operations.FromUnion3TypeStr:
		// fromUnion3.Str is populated
	case operations.FromUnion3TypeArrayOfStr:
		// fromUnion3.ArrayOfStr is populated
}
```
