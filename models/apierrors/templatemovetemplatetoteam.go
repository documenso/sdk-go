// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package apierrors

import (
	"encoding/json"
	"github.com/documenso/sdk-go/models/components"
)

type TemplateMoveTemplateToTeamInternalServerErrorIssue struct {
	Message string `json:"message"`
}

func (o *TemplateMoveTemplateToTeamInternalServerErrorIssue) GetMessage() string {
	if o == nil {
		return ""
	}
	return o.Message
}

// TemplateMoveTemplateToTeamInternalServerError - Internal server error
type TemplateMoveTemplateToTeamInternalServerError struct {
	Message  string                                               `json:"message"`
	Code     string                                               `json:"code"`
	Issues   []TemplateMoveTemplateToTeamInternalServerErrorIssue `json:"issues,omitempty"`
	HTTPMeta components.HTTPMetadata                              `json:"-"`
}

var _ error = &TemplateMoveTemplateToTeamInternalServerError{}

func (e *TemplateMoveTemplateToTeamInternalServerError) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

type TemplateMoveTemplateToTeamBadRequestIssue struct {
	Message string `json:"message"`
}

func (o *TemplateMoveTemplateToTeamBadRequestIssue) GetMessage() string {
	if o == nil {
		return ""
	}
	return o.Message
}

// TemplateMoveTemplateToTeamBadRequestError - Invalid input data
type TemplateMoveTemplateToTeamBadRequestError struct {
	Message  string                                      `json:"message"`
	Code     string                                      `json:"code"`
	Issues   []TemplateMoveTemplateToTeamBadRequestIssue `json:"issues,omitempty"`
	HTTPMeta components.HTTPMetadata                     `json:"-"`
}

var _ error = &TemplateMoveTemplateToTeamBadRequestError{}

func (e *TemplateMoveTemplateToTeamBadRequestError) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
