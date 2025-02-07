// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package apierrors

import (
	"encoding/json"
	"github.com/documenso/sdk-go/models/components"
)

type FieldUpdateTemplateFieldTemplatesFieldsIssues struct {
	Message string `json:"message"`
}

func (o *FieldUpdateTemplateFieldTemplatesFieldsIssues) GetMessage() string {
	if o == nil {
		return ""
	}
	return o.Message
}

// FieldUpdateTemplateFieldTemplatesFieldsResponseBody - Internal server error
type FieldUpdateTemplateFieldTemplatesFieldsResponseBody struct {
	Message  string                                          `json:"message"`
	Code     string                                          `json:"code"`
	Issues   []FieldUpdateTemplateFieldTemplatesFieldsIssues `json:"issues,omitempty"`
	HTTPMeta components.HTTPMetadata                         `json:"-"`
}

var _ error = &FieldUpdateTemplateFieldTemplatesFieldsResponseBody{}

func (e *FieldUpdateTemplateFieldTemplatesFieldsResponseBody) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

type FieldUpdateTemplateFieldIssues struct {
	Message string `json:"message"`
}

func (o *FieldUpdateTemplateFieldIssues) GetMessage() string {
	if o == nil {
		return ""
	}
	return o.Message
}

// FieldUpdateTemplateFieldResponseBody - Invalid input data
type FieldUpdateTemplateFieldResponseBody struct {
	Message  string                           `json:"message"`
	Code     string                           `json:"code"`
	Issues   []FieldUpdateTemplateFieldIssues `json:"issues,omitempty"`
	HTTPMeta components.HTTPMetadata          `json:"-"`
}

var _ error = &FieldUpdateTemplateFieldResponseBody{}

func (e *FieldUpdateTemplateFieldResponseBody) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
