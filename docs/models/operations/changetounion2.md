# ChangeToUnion2


## Supported Types

### 

```go
changeToUnion2 := operations.CreateChangeToUnion2Str(string{/* values here */})
```

### 

```go
changeToUnion2 := operations.CreateChangeToUnion2ArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch changeToUnion2.Type {
	case operations.ChangeToUnion2TypeStr:
		// changeToUnion2.Str is populated
	case operations.ChangeToUnion2TypeArrayOfStr:
		// changeToUnion2.ArrayOfStr is populated
}
```
