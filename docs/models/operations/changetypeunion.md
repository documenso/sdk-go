# ChangeTypeUnion


## Supported Types

### TypeDateFormat

```go
changeTypeUnion := operations.CreateChangeTypeUnionTypeDateFormat(operations.TypeDateFormat{/* values here */})
```

### TypeMessage

```go
changeTypeUnion := operations.CreateChangeTypeUnionTypeMessage(operations.TypeMessage{/* values here */})
```

### TypeRedirectURL

```go
changeTypeUnion := operations.CreateChangeTypeUnionTypeRedirectURL(operations.TypeRedirectURL{/* values here */})
```

### TypeSubject

```go
changeTypeUnion := operations.CreateChangeTypeUnionTypeSubject(operations.TypeSubject{/* values here */})
```

### TypeTimezone

```go
changeTypeUnion := operations.CreateChangeTypeUnionTypeTimezone(operations.TypeTimezone{/* values here */})
```

### TypeEmailID

```go
changeTypeUnion := operations.CreateChangeTypeUnionTypeEmailID(operations.TypeEmailID{/* values here */})
```

### TypeEmailReplyTo

```go
changeTypeUnion := operations.CreateChangeTypeUnionTypeEmailReplyTo(operations.TypeEmailReplyTo{/* values here */})
```

### TypeEmailSettings

```go
changeTypeUnion := operations.CreateChangeTypeUnionTypeEmailSettings(operations.TypeEmailSettings{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch changeTypeUnion.Type {
	case operations.ChangeTypeUnionTypeTypeDateFormat:
		// changeTypeUnion.TypeDateFormat is populated
	case operations.ChangeTypeUnionTypeTypeMessage:
		// changeTypeUnion.TypeMessage is populated
	case operations.ChangeTypeUnionTypeTypeRedirectURL:
		// changeTypeUnion.TypeRedirectURL is populated
	case operations.ChangeTypeUnionTypeTypeSubject:
		// changeTypeUnion.TypeSubject is populated
	case operations.ChangeTypeUnionTypeTypeTimezone:
		// changeTypeUnion.TypeTimezone is populated
	case operations.ChangeTypeUnionTypeTypeEmailID:
		// changeTypeUnion.TypeEmailID is populated
	case operations.ChangeTypeUnionTypeTypeEmailReplyTo:
		// changeTypeUnion.TypeEmailReplyTo is populated
	case operations.ChangeTypeUnionTypeTypeEmailSettings:
		// changeTypeUnion.TypeEmailSettings is populated
}
```
