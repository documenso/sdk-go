// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package apierrors

import (
	"encoding/json"
	"github.com/documenso/sdk-go/models/components"
)

type FieldDeleteTemplateFieldInternalServerErrorIssue struct {
	Message string `json:"message"`
}

func (o *FieldDeleteTemplateFieldInternalServerErrorIssue) GetMessage() string {
	if o == nil {
		return ""
	}
	return o.Message
}

// FieldDeleteTemplateFieldInternalServerError - Internal server error
type FieldDeleteTemplateFieldInternalServerError struct {
	Message  string                                             `json:"message"`
	Code     string                                             `json:"code"`
	Issues   []FieldDeleteTemplateFieldInternalServerErrorIssue `json:"issues,omitempty"`
	HTTPMeta components.HTTPMetadata                            `json:"-"`
}

var _ error = &FieldDeleteTemplateFieldInternalServerError{}

func (e *FieldDeleteTemplateFieldInternalServerError) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

type FieldDeleteTemplateFieldBadRequestIssue struct {
	Message string `json:"message"`
}

func (o *FieldDeleteTemplateFieldBadRequestIssue) GetMessage() string {
	if o == nil {
		return ""
	}
	return o.Message
}

// FieldDeleteTemplateFieldBadRequestError - Invalid input data
type FieldDeleteTemplateFieldBadRequestError struct {
	Message  string                                    `json:"message"`
	Code     string                                    `json:"code"`
	Issues   []FieldDeleteTemplateFieldBadRequestIssue `json:"issues,omitempty"`
	HTTPMeta components.HTTPMetadata                   `json:"-"`
}

var _ error = &FieldDeleteTemplateFieldBadRequestError{}

func (e *FieldDeleteTemplateFieldBadRequestError) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
