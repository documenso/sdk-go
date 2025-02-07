// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package apierrors

import (
	"encoding/json"
	"github.com/documenso/sdk-go/models/components"
)

type FieldDeleteTemplateFieldTemplatesFieldsIssues struct {
	Message string `json:"message"`
}

func (o *FieldDeleteTemplateFieldTemplatesFieldsIssues) GetMessage() string {
	if o == nil {
		return ""
	}
	return o.Message
}

// FieldDeleteTemplateFieldTemplatesFieldsResponseBody - Internal server error
type FieldDeleteTemplateFieldTemplatesFieldsResponseBody struct {
	Message  string                                          `json:"message"`
	Code     string                                          `json:"code"`
	Issues   []FieldDeleteTemplateFieldTemplatesFieldsIssues `json:"issues,omitempty"`
	HTTPMeta components.HTTPMetadata                         `json:"-"`
}

var _ error = &FieldDeleteTemplateFieldTemplatesFieldsResponseBody{}

func (e *FieldDeleteTemplateFieldTemplatesFieldsResponseBody) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

type FieldDeleteTemplateFieldIssues struct {
	Message string `json:"message"`
}

func (o *FieldDeleteTemplateFieldIssues) GetMessage() string {
	if o == nil {
		return ""
	}
	return o.Message
}

// FieldDeleteTemplateFieldResponseBody - Invalid input data
type FieldDeleteTemplateFieldResponseBody struct {
	Message  string                           `json:"message"`
	Code     string                           `json:"code"`
	Issues   []FieldDeleteTemplateFieldIssues `json:"issues,omitempty"`
	HTTPMeta components.HTTPMetadata          `json:"-"`
}

var _ error = &FieldDeleteTemplateFieldResponseBody{}

func (e *FieldDeleteTemplateFieldResponseBody) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
