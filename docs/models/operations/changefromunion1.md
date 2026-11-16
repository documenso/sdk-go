# ChangeFromUnion1


## Supported Types

### 

```go
changeFromUnion1 := operations.CreateChangeFromUnion1Str(string{/* values here */})
```

### 

```go
changeFromUnion1 := operations.CreateChangeFromUnion1ArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch changeFromUnion1.Type {
	case operations.ChangeFromUnion1TypeStr:
		// changeFromUnion1.Str is populated
	case operations.ChangeFromUnion1TypeArrayOfStr:
		// changeFromUnion1.ArrayOfStr is populated
}
```
