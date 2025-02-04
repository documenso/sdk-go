// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package apierrors

import (
	"encoding/json"
	"github.com/documenso/sdk-go/models/components"
)

type TemplateDuplicateTemplateTemplatesIssues struct {
	Message string `json:"message"`
}

func (o *TemplateDuplicateTemplateTemplatesIssues) GetMessage() string {
	if o == nil {
		return ""
	}
	return o.Message
}

// TemplateDuplicateTemplateTemplatesResponseBody - Internal server error
type TemplateDuplicateTemplateTemplatesResponseBody struct {
	Message  string                                     `json:"message"`
	Code     string                                     `json:"code"`
	Issues   []TemplateDuplicateTemplateTemplatesIssues `json:"issues,omitempty"`
	HTTPMeta components.HTTPMetadata                    `json:"-"`
}

var _ error = &TemplateDuplicateTemplateTemplatesResponseBody{}

func (e *TemplateDuplicateTemplateTemplatesResponseBody) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

type TemplateDuplicateTemplateIssues struct {
	Message string `json:"message"`
}

func (o *TemplateDuplicateTemplateIssues) GetMessage() string {
	if o == nil {
		return ""
	}
	return o.Message
}

// TemplateDuplicateTemplateResponseBody - Invalid input data
type TemplateDuplicateTemplateResponseBody struct {
	Message  string                            `json:"message"`
	Code     string                            `json:"code"`
	Issues   []TemplateDuplicateTemplateIssues `json:"issues,omitempty"`
	HTTPMeta components.HTTPMetadata           `json:"-"`
}

var _ error = &TemplateDuplicateTemplateResponseBody{}

func (e *TemplateDuplicateTemplateResponseBody) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
