# ChangeUnion3


## Supported Types

### ChangeActionAuth

```go
changeUnion3 := operations.CreateChangeUnion3ChangeActionAuth(operations.ChangeActionAuth{/* values here */})
```

### ChangeAccessAuth

```go
changeUnion3 := operations.CreateChangeUnion3ChangeAccessAuth(operations.ChangeAccessAuth{/* values here */})
```

### ChangeName

```go
changeUnion3 := operations.CreateChangeUnion3ChangeName(operations.ChangeName{/* values here */})
```

### ChangeRole

```go
changeUnion3 := operations.CreateChangeUnion3ChangeRole(operations.ChangeRole{/* values here */})
```

### ChangeEmail

```go
changeUnion3 := operations.CreateChangeUnion3ChangeEmail(operations.ChangeEmail{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch changeUnion3.Type {
	case operations.ChangeUnion3TypeChangeActionAuth:
		// changeUnion3.ChangeActionAuth is populated
	case operations.ChangeUnion3TypeChangeAccessAuth:
		// changeUnion3.ChangeAccessAuth is populated
	case operations.ChangeUnion3TypeChangeName:
		// changeUnion3.ChangeName is populated
	case operations.ChangeUnion3TypeChangeRole:
		// changeUnion3.ChangeRole is populated
	case operations.ChangeUnion3TypeChangeEmail:
		// changeUnion3.ChangeEmail is populated
}
```
