# TemplateCreateDocumentFromTemplateFormValuesResponse


## Supported Types

### 

```go
templateCreateDocumentFromTemplateFormValuesResponse := operations.CreateTemplateCreateDocumentFromTemplateFormValuesResponseStr(string{/* values here */})
```

### 

```go
templateCreateDocumentFromTemplateFormValuesResponse := operations.CreateTemplateCreateDocumentFromTemplateFormValuesResponseBoolean(bool{/* values here */})
```

### 

```go
templateCreateDocumentFromTemplateFormValuesResponse := operations.CreateTemplateCreateDocumentFromTemplateFormValuesResponseNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch templateCreateDocumentFromTemplateFormValuesResponse.Type {
	case operations.TemplateCreateDocumentFromTemplateFormValuesResponseTypeStr:
		// templateCreateDocumentFromTemplateFormValuesResponse.Str is populated
	case operations.TemplateCreateDocumentFromTemplateFormValuesResponseTypeBoolean:
		// templateCreateDocumentFromTemplateFormValuesResponse.Boolean is populated
	case operations.TemplateCreateDocumentFromTemplateFormValuesResponseTypeNumber:
		// templateCreateDocumentFromTemplateFormValuesResponse.Number is populated
}
```
