# TemplateCreateDocumentFromTemplateEmailUnion


## Supported Types

### TemplateCreateDocumentFromTemplateEmailEnum

```go
templateCreateDocumentFromTemplateEmailUnion := operations.CreateTemplateCreateDocumentFromTemplateEmailUnionTemplateCreateDocumentFromTemplateEmailEnum(operations.TemplateCreateDocumentFromTemplateEmailEnum{/* values here */})
```

### 

```go
templateCreateDocumentFromTemplateEmailUnion := operations.CreateTemplateCreateDocumentFromTemplateEmailUnionStr(string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch templateCreateDocumentFromTemplateEmailUnion.Type {
	case operations.TemplateCreateDocumentFromTemplateEmailUnionTypeTemplateCreateDocumentFromTemplateEmailEnum:
		// templateCreateDocumentFromTemplateEmailUnion.TemplateCreateDocumentFromTemplateEmailEnum is populated
	case operations.TemplateCreateDocumentFromTemplateEmailUnionTypeStr:
		// templateCreateDocumentFromTemplateEmailUnion.Str is populated
}
```
