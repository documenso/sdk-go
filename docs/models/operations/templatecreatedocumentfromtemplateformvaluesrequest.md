# TemplateCreateDocumentFromTemplateFormValuesRequest


## Supported Types

### 

```go
templateCreateDocumentFromTemplateFormValuesRequest := operations.CreateTemplateCreateDocumentFromTemplateFormValuesRequestStr(string{/* values here */})
```

### 

```go
templateCreateDocumentFromTemplateFormValuesRequest := operations.CreateTemplateCreateDocumentFromTemplateFormValuesRequestBoolean(bool{/* values here */})
```

### 

```go
templateCreateDocumentFromTemplateFormValuesRequest := operations.CreateTemplateCreateDocumentFromTemplateFormValuesRequestNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch templateCreateDocumentFromTemplateFormValuesRequest.Type {
	case operations.TemplateCreateDocumentFromTemplateFormValuesRequestTypeStr:
		// templateCreateDocumentFromTemplateFormValuesRequest.Str is populated
	case operations.TemplateCreateDocumentFromTemplateFormValuesRequestTypeBoolean:
		// templateCreateDocumentFromTemplateFormValuesRequest.Boolean is populated
	case operations.TemplateCreateDocumentFromTemplateFormValuesRequestTypeNumber:
		// templateCreateDocumentFromTemplateFormValuesRequest.Number is populated
}
```
