# EnvelopeAuditLogFindFieldUnion1


## Supported Types

### EnvelopeAuditLogFindFieldInitials1

```go
envelopeAuditLogFindFieldUnion1 := operations.CreateEnvelopeAuditLogFindFieldUnion1EnvelopeAuditLogFindFieldInitials1(operations.EnvelopeAuditLogFindFieldInitials1{/* values here */})
```

### EnvelopeAuditLogFindFieldEmail1

```go
envelopeAuditLogFindFieldUnion1 := operations.CreateEnvelopeAuditLogFindFieldUnion1EnvelopeAuditLogFindFieldEmail1(operations.EnvelopeAuditLogFindFieldEmail1{/* values here */})
```

### EnvelopeAuditLogFindFieldDate1

```go
envelopeAuditLogFindFieldUnion1 := operations.CreateEnvelopeAuditLogFindFieldUnion1EnvelopeAuditLogFindFieldDate1(operations.EnvelopeAuditLogFindFieldDate1{/* values here */})
```

### EnvelopeAuditLogFindFieldName1

```go
envelopeAuditLogFindFieldUnion1 := operations.CreateEnvelopeAuditLogFindFieldUnion1EnvelopeAuditLogFindFieldName1(operations.EnvelopeAuditLogFindFieldName1{/* values here */})
```

### EnvelopeAuditLogFindFieldText1

```go
envelopeAuditLogFindFieldUnion1 := operations.CreateEnvelopeAuditLogFindFieldUnion1EnvelopeAuditLogFindFieldText1(operations.EnvelopeAuditLogFindFieldText1{/* values here */})
```

### EnvelopeAuditLogFindField2

```go
envelopeAuditLogFindFieldUnion1 := operations.CreateEnvelopeAuditLogFindFieldUnion1EnvelopeAuditLogFindField2(operations.EnvelopeAuditLogFindField2{/* values here */})
```

### EnvelopeAuditLogFindFieldRadio1

```go
envelopeAuditLogFindFieldUnion1 := operations.CreateEnvelopeAuditLogFindFieldUnion1EnvelopeAuditLogFindFieldRadio1(operations.EnvelopeAuditLogFindFieldRadio1{/* values here */})
```

### EnvelopeAuditLogFindFieldCheckbox1

```go
envelopeAuditLogFindFieldUnion1 := operations.CreateEnvelopeAuditLogFindFieldUnion1EnvelopeAuditLogFindFieldCheckbox1(operations.EnvelopeAuditLogFindFieldCheckbox1{/* values here */})
```

### EnvelopeAuditLogFindFieldDropdown1

```go
envelopeAuditLogFindFieldUnion1 := operations.CreateEnvelopeAuditLogFindFieldUnion1EnvelopeAuditLogFindFieldDropdown1(operations.EnvelopeAuditLogFindFieldDropdown1{/* values here */})
```

### EnvelopeAuditLogFindFieldNumber1

```go
envelopeAuditLogFindFieldUnion1 := operations.CreateEnvelopeAuditLogFindFieldUnion1EnvelopeAuditLogFindFieldNumber1(operations.EnvelopeAuditLogFindFieldNumber1{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch envelopeAuditLogFindFieldUnion1.Type {
	case operations.EnvelopeAuditLogFindFieldUnion1TypeEnvelopeAuditLogFindFieldInitials1:
		// envelopeAuditLogFindFieldUnion1.EnvelopeAuditLogFindFieldInitials1 is populated
	case operations.EnvelopeAuditLogFindFieldUnion1TypeEnvelopeAuditLogFindFieldEmail1:
		// envelopeAuditLogFindFieldUnion1.EnvelopeAuditLogFindFieldEmail1 is populated
	case operations.EnvelopeAuditLogFindFieldUnion1TypeEnvelopeAuditLogFindFieldDate1:
		// envelopeAuditLogFindFieldUnion1.EnvelopeAuditLogFindFieldDate1 is populated
	case operations.EnvelopeAuditLogFindFieldUnion1TypeEnvelopeAuditLogFindFieldName1:
		// envelopeAuditLogFindFieldUnion1.EnvelopeAuditLogFindFieldName1 is populated
	case operations.EnvelopeAuditLogFindFieldUnion1TypeEnvelopeAuditLogFindFieldText1:
		// envelopeAuditLogFindFieldUnion1.EnvelopeAuditLogFindFieldText1 is populated
	case operations.EnvelopeAuditLogFindFieldUnion1TypeEnvelopeAuditLogFindField2:
		// envelopeAuditLogFindFieldUnion1.EnvelopeAuditLogFindField2 is populated
	case operations.EnvelopeAuditLogFindFieldUnion1TypeEnvelopeAuditLogFindFieldRadio1:
		// envelopeAuditLogFindFieldUnion1.EnvelopeAuditLogFindFieldRadio1 is populated
	case operations.EnvelopeAuditLogFindFieldUnion1TypeEnvelopeAuditLogFindFieldCheckbox1:
		// envelopeAuditLogFindFieldUnion1.EnvelopeAuditLogFindFieldCheckbox1 is populated
	case operations.EnvelopeAuditLogFindFieldUnion1TypeEnvelopeAuditLogFindFieldDropdown1:
		// envelopeAuditLogFindFieldUnion1.EnvelopeAuditLogFindFieldDropdown1 is populated
	case operations.EnvelopeAuditLogFindFieldUnion1TypeEnvelopeAuditLogFindFieldNumber1:
		// envelopeAuditLogFindFieldUnion1.EnvelopeAuditLogFindFieldNumber1 is populated
}
```
