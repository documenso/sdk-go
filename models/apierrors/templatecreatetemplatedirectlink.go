// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package apierrors

import (
	"encoding/json"
	"github.com/documenso/sdk-go/models/components"
)

type TemplateCreateTemplateDirectLinkTemplatesDirectLinkIssues struct {
	Message string `json:"message"`
}

func (o *TemplateCreateTemplateDirectLinkTemplatesDirectLinkIssues) GetMessage() string {
	if o == nil {
		return ""
	}
	return o.Message
}

// TemplateCreateTemplateDirectLinkTemplatesDirectLinkResponseBody - Internal server error
type TemplateCreateTemplateDirectLinkTemplatesDirectLinkResponseBody struct {
	Message  string                                                      `json:"message"`
	Code     string                                                      `json:"code"`
	Issues   []TemplateCreateTemplateDirectLinkTemplatesDirectLinkIssues `json:"issues,omitempty"`
	HTTPMeta components.HTTPMetadata                                     `json:"-"`
}

var _ error = &TemplateCreateTemplateDirectLinkTemplatesDirectLinkResponseBody{}

func (e *TemplateCreateTemplateDirectLinkTemplatesDirectLinkResponseBody) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

type TemplateCreateTemplateDirectLinkIssues struct {
	Message string `json:"message"`
}

func (o *TemplateCreateTemplateDirectLinkIssues) GetMessage() string {
	if o == nil {
		return ""
	}
	return o.Message
}

// TemplateCreateTemplateDirectLinkResponseBody - Invalid input data
type TemplateCreateTemplateDirectLinkResponseBody struct {
	Message  string                                   `json:"message"`
	Code     string                                   `json:"code"`
	Issues   []TemplateCreateTemplateDirectLinkIssues `json:"issues,omitempty"`
	HTTPMeta components.HTTPMetadata                  `json:"-"`
}

var _ error = &TemplateCreateTemplateDirectLinkResponseBody{}

func (e *TemplateCreateTemplateDirectLinkResponseBody) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
