// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package apierrors

import (
	"encoding/json"
	"github.com/documenso/sdk-go/models/components"
)

type TemplateMoveTemplateToTeamTemplatesIssues struct {
	Message string `json:"message"`
}

func (o *TemplateMoveTemplateToTeamTemplatesIssues) GetMessage() string {
	if o == nil {
		return ""
	}
	return o.Message
}

// TemplateMoveTemplateToTeamTemplatesResponseBody - Internal server error
type TemplateMoveTemplateToTeamTemplatesResponseBody struct {
	Message  string                                      `json:"message"`
	Code     string                                      `json:"code"`
	Issues   []TemplateMoveTemplateToTeamTemplatesIssues `json:"issues,omitempty"`
	HTTPMeta components.HTTPMetadata                     `json:"-"`
}

var _ error = &TemplateMoveTemplateToTeamTemplatesResponseBody{}

func (e *TemplateMoveTemplateToTeamTemplatesResponseBody) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

type TemplateMoveTemplateToTeamIssues struct {
	Message string `json:"message"`
}

func (o *TemplateMoveTemplateToTeamIssues) GetMessage() string {
	if o == nil {
		return ""
	}
	return o.Message
}

// TemplateMoveTemplateToTeamResponseBody - Invalid input data
type TemplateMoveTemplateToTeamResponseBody struct {
	Message  string                             `json:"message"`
	Code     string                             `json:"code"`
	Issues   []TemplateMoveTemplateToTeamIssues `json:"issues,omitempty"`
	HTTPMeta components.HTTPMetadata            `json:"-"`
}

var _ error = &TemplateMoveTemplateToTeamResponseBody{}

func (e *TemplateMoveTemplateToTeamResponseBody) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
