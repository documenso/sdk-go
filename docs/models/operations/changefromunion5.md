# ChangeFromUnion5


## Supported Types

### 

```go
changeFromUnion5 := operations.CreateChangeFromUnion5Str(string{/* values here */})
```

### 

```go
changeFromUnion5 := operations.CreateChangeFromUnion5ArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch changeFromUnion5.Type {
	case operations.ChangeFromUnion5TypeStr:
		// changeFromUnion5.Str is populated
	case operations.ChangeFromUnion5TypeArrayOfStr:
		// changeFromUnion5.ArrayOfStr is populated
}
```
