# FromUnion1


## Supported Types

### 

```go
fromUnion1 := operations.CreateFromUnion1Str(string{/* values here */})
```

### 

```go
fromUnion1 := operations.CreateFromUnion1ArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch fromUnion1.Type {
	case operations.FromUnion1TypeStr:
		// fromUnion1.Str is populated
	case operations.FromUnion1TypeArrayOfStr:
		// fromUnion1.ArrayOfStr is populated
}
```
