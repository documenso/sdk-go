# ChangeFromUnion4


## Supported Types

### 

```go
changeFromUnion4 := operations.CreateChangeFromUnion4Str(string{/* values here */})
```

### 

```go
changeFromUnion4 := operations.CreateChangeFromUnion4ArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch changeFromUnion4.Type {
	case operations.ChangeFromUnion4TypeStr:
		// changeFromUnion4.Str is populated
	case operations.ChangeFromUnion4TypeArrayOfStr:
		// changeFromUnion4.ArrayOfStr is populated
}
```
