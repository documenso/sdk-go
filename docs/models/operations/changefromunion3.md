# ChangeFromUnion3


## Supported Types

### 

```go
changeFromUnion3 := operations.CreateChangeFromUnion3Str(string{/* values here */})
```

### 

```go
changeFromUnion3 := operations.CreateChangeFromUnion3ArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch changeFromUnion3.Type {
	case operations.ChangeFromUnion3TypeStr:
		// changeFromUnion3.Str is populated
	case operations.ChangeFromUnion3TypeArrayOfStr:
		// changeFromUnion3.ArrayOfStr is populated
}
```
