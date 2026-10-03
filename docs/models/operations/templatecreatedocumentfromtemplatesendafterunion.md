# TemplateCreateDocumentFromTemplateSendAfterUnion


## Supported Types

### TemplateCreateDocumentFromTemplateSendAfter1

```go
templateCreateDocumentFromTemplateSendAfterUnion := operations.CreateTemplateCreateDocumentFromTemplateSendAfterUnionTemplateCreateDocumentFromTemplateSendAfter1(operations.TemplateCreateDocumentFromTemplateSendAfter1{/* values here */})
```

### TemplateCreateDocumentFromTemplateSendAfter2

```go
templateCreateDocumentFromTemplateSendAfterUnion := operations.CreateTemplateCreateDocumentFromTemplateSendAfterUnionTemplateCreateDocumentFromTemplateSendAfter2(operations.TemplateCreateDocumentFromTemplateSendAfter2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch templateCreateDocumentFromTemplateSendAfterUnion.Type {
	case operations.TemplateCreateDocumentFromTemplateSendAfterUnionTypeTemplateCreateDocumentFromTemplateSendAfter1:
		// templateCreateDocumentFromTemplateSendAfterUnion.TemplateCreateDocumentFromTemplateSendAfter1 is populated
	case operations.TemplateCreateDocumentFromTemplateSendAfterUnionTypeTemplateCreateDocumentFromTemplateSendAfter2:
		// templateCreateDocumentFromTemplateSendAfterUnion.TemplateCreateDocumentFromTemplateSendAfter2 is populated
}
```
