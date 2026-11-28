# EnvelopeAuditLogFindFieldUnion2


## Supported Types

### EnvelopeAuditLogFindFieldInitials2

```go
envelopeAuditLogFindFieldUnion2 := operations.CreateEnvelopeAuditLogFindFieldUnion2EnvelopeAuditLogFindFieldInitials2(operations.EnvelopeAuditLogFindFieldInitials2{/* values here */})
```

### EnvelopeAuditLogFindFieldEmail2

```go
envelopeAuditLogFindFieldUnion2 := operations.CreateEnvelopeAuditLogFindFieldUnion2EnvelopeAuditLogFindFieldEmail2(operations.EnvelopeAuditLogFindFieldEmail2{/* values here */})
```

### EnvelopeAuditLogFindFieldDate2

```go
envelopeAuditLogFindFieldUnion2 := operations.CreateEnvelopeAuditLogFindFieldUnion2EnvelopeAuditLogFindFieldDate2(operations.EnvelopeAuditLogFindFieldDate2{/* values here */})
```

### EnvelopeAuditLogFindFieldName2

```go
envelopeAuditLogFindFieldUnion2 := operations.CreateEnvelopeAuditLogFindFieldUnion2EnvelopeAuditLogFindFieldName2(operations.EnvelopeAuditLogFindFieldName2{/* values here */})
```

### EnvelopeAuditLogFindFieldText2

```go
envelopeAuditLogFindFieldUnion2 := operations.CreateEnvelopeAuditLogFindFieldUnion2EnvelopeAuditLogFindFieldText2(operations.EnvelopeAuditLogFindFieldText2{/* values here */})
```

### EnvelopeAuditLogFindField3

```go
envelopeAuditLogFindFieldUnion2 := operations.CreateEnvelopeAuditLogFindFieldUnion2EnvelopeAuditLogFindField3(operations.EnvelopeAuditLogFindField3{/* values here */})
```

### EnvelopeAuditLogFindFieldRadio2

```go
envelopeAuditLogFindFieldUnion2 := operations.CreateEnvelopeAuditLogFindFieldUnion2EnvelopeAuditLogFindFieldRadio2(operations.EnvelopeAuditLogFindFieldRadio2{/* values here */})
```

### EnvelopeAuditLogFindFieldCheckbox2

```go
envelopeAuditLogFindFieldUnion2 := operations.CreateEnvelopeAuditLogFindFieldUnion2EnvelopeAuditLogFindFieldCheckbox2(operations.EnvelopeAuditLogFindFieldCheckbox2{/* values here */})
```

### EnvelopeAuditLogFindFieldDropdown2

```go
envelopeAuditLogFindFieldUnion2 := operations.CreateEnvelopeAuditLogFindFieldUnion2EnvelopeAuditLogFindFieldDropdown2(operations.EnvelopeAuditLogFindFieldDropdown2{/* values here */})
```

### EnvelopeAuditLogFindFieldNumber2

```go
envelopeAuditLogFindFieldUnion2 := operations.CreateEnvelopeAuditLogFindFieldUnion2EnvelopeAuditLogFindFieldNumber2(operations.EnvelopeAuditLogFindFieldNumber2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch envelopeAuditLogFindFieldUnion2.Type {
	case operations.EnvelopeAuditLogFindFieldUnion2TypeEnvelopeAuditLogFindFieldInitials2:
		// envelopeAuditLogFindFieldUnion2.EnvelopeAuditLogFindFieldInitials2 is populated
	case operations.EnvelopeAuditLogFindFieldUnion2TypeEnvelopeAuditLogFindFieldEmail2:
		// envelopeAuditLogFindFieldUnion2.EnvelopeAuditLogFindFieldEmail2 is populated
	case operations.EnvelopeAuditLogFindFieldUnion2TypeEnvelopeAuditLogFindFieldDate2:
		// envelopeAuditLogFindFieldUnion2.EnvelopeAuditLogFindFieldDate2 is populated
	case operations.EnvelopeAuditLogFindFieldUnion2TypeEnvelopeAuditLogFindFieldName2:
		// envelopeAuditLogFindFieldUnion2.EnvelopeAuditLogFindFieldName2 is populated
	case operations.EnvelopeAuditLogFindFieldUnion2TypeEnvelopeAuditLogFindFieldText2:
		// envelopeAuditLogFindFieldUnion2.EnvelopeAuditLogFindFieldText2 is populated
	case operations.EnvelopeAuditLogFindFieldUnion2TypeEnvelopeAuditLogFindField3:
		// envelopeAuditLogFindFieldUnion2.EnvelopeAuditLogFindField3 is populated
	case operations.EnvelopeAuditLogFindFieldUnion2TypeEnvelopeAuditLogFindFieldRadio2:
		// envelopeAuditLogFindFieldUnion2.EnvelopeAuditLogFindFieldRadio2 is populated
	case operations.EnvelopeAuditLogFindFieldUnion2TypeEnvelopeAuditLogFindFieldCheckbox2:
		// envelopeAuditLogFindFieldUnion2.EnvelopeAuditLogFindFieldCheckbox2 is populated
	case operations.EnvelopeAuditLogFindFieldUnion2TypeEnvelopeAuditLogFindFieldDropdown2:
		// envelopeAuditLogFindFieldUnion2.EnvelopeAuditLogFindFieldDropdown2 is populated
	case operations.EnvelopeAuditLogFindFieldUnion2TypeEnvelopeAuditLogFindFieldNumber2:
		// envelopeAuditLogFindFieldUnion2.EnvelopeAuditLogFindFieldNumber2 is populated
}
```
