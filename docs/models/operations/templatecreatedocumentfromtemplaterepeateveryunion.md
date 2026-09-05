# TemplateCreateDocumentFromTemplateRepeatEveryUnion


## Supported Types

### TemplateCreateDocumentFromTemplateRepeatEvery1

```go
templateCreateDocumentFromTemplateRepeatEveryUnion := operations.CreateTemplateCreateDocumentFromTemplateRepeatEveryUnionTemplateCreateDocumentFromTemplateRepeatEvery1(operations.TemplateCreateDocumentFromTemplateRepeatEvery1{/* values here */})
```

### TemplateCreateDocumentFromTemplateRepeatEvery2

```go
templateCreateDocumentFromTemplateRepeatEveryUnion := operations.CreateTemplateCreateDocumentFromTemplateRepeatEveryUnionTemplateCreateDocumentFromTemplateRepeatEvery2(operations.TemplateCreateDocumentFromTemplateRepeatEvery2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch templateCreateDocumentFromTemplateRepeatEveryUnion.Type {
	case operations.TemplateCreateDocumentFromTemplateRepeatEveryUnionTypeTemplateCreateDocumentFromTemplateRepeatEvery1:
		// templateCreateDocumentFromTemplateRepeatEveryUnion.TemplateCreateDocumentFromTemplateRepeatEvery1 is populated
	case operations.TemplateCreateDocumentFromTemplateRepeatEveryUnionTypeTemplateCreateDocumentFromTemplateRepeatEvery2:
		// templateCreateDocumentFromTemplateRepeatEveryUnion.TemplateCreateDocumentFromTemplateRepeatEvery2 is populated
}
```
