// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package apierrors

import (
	"encoding/json"
	"github.com/documenso/sdk-go/models/components"
)

type TemplateGetTemplateByIDInternalServerErrorIssue struct {
	Message string `json:"message"`
}

func (o *TemplateGetTemplateByIDInternalServerErrorIssue) GetMessage() string {
	if o == nil {
		return ""
	}
	return o.Message
}

// TemplateGetTemplateByIDInternalServerError - Internal server error
type TemplateGetTemplateByIDInternalServerError struct {
	Message  string                                            `json:"message"`
	Code     string                                            `json:"code"`
	Issues   []TemplateGetTemplateByIDInternalServerErrorIssue `json:"issues,omitempty"`
	HTTPMeta components.HTTPMetadata                           `json:"-"`
}

var _ error = &TemplateGetTemplateByIDInternalServerError{}

func (e *TemplateGetTemplateByIDInternalServerError) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

type TemplateGetTemplateByIDNotFoundIssue struct {
	Message string `json:"message"`
}

func (o *TemplateGetTemplateByIDNotFoundIssue) GetMessage() string {
	if o == nil {
		return ""
	}
	return o.Message
}

// TemplateGetTemplateByIDNotFoundError - Not found
type TemplateGetTemplateByIDNotFoundError struct {
	Message  string                                 `json:"message"`
	Code     string                                 `json:"code"`
	Issues   []TemplateGetTemplateByIDNotFoundIssue `json:"issues,omitempty"`
	HTTPMeta components.HTTPMetadata                `json:"-"`
}

var _ error = &TemplateGetTemplateByIDNotFoundError{}

func (e *TemplateGetTemplateByIDNotFoundError) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

type TemplateGetTemplateByIDBadRequestIssue struct {
	Message string `json:"message"`
}

func (o *TemplateGetTemplateByIDBadRequestIssue) GetMessage() string {
	if o == nil {
		return ""
	}
	return o.Message
}

// TemplateGetTemplateByIDBadRequestError - Invalid input data
type TemplateGetTemplateByIDBadRequestError struct {
	Message  string                                   `json:"message"`
	Code     string                                   `json:"code"`
	Issues   []TemplateGetTemplateByIDBadRequestIssue `json:"issues,omitempty"`
	HTTPMeta components.HTTPMetadata                  `json:"-"`
}

var _ error = &TemplateGetTemplateByIDBadRequestError{}

func (e *TemplateGetTemplateByIDBadRequestError) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
