# ChangeFromUnion2


## Supported Types

### 

```go
changeFromUnion2 := operations.CreateChangeFromUnion2Str(string{/* values here */})
```

### 

```go
changeFromUnion2 := operations.CreateChangeFromUnion2ArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch changeFromUnion2.Type {
	case operations.ChangeFromUnion2TypeStr:
		// changeFromUnion2.Str is populated
	case operations.ChangeFromUnion2TypeArrayOfStr:
		// changeFromUnion2.ArrayOfStr is populated
}
```
