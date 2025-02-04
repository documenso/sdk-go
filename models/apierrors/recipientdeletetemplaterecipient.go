// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package apierrors

import (
	"encoding/json"
	"github.com/documenso/sdk-go/models/components"
)

type RecipientDeleteTemplateRecipientTemplatesRecipientsIssues struct {
	Message string `json:"message"`
}

func (o *RecipientDeleteTemplateRecipientTemplatesRecipientsIssues) GetMessage() string {
	if o == nil {
		return ""
	}
	return o.Message
}

// RecipientDeleteTemplateRecipientTemplatesRecipientsResponseBody - Internal server error
type RecipientDeleteTemplateRecipientTemplatesRecipientsResponseBody struct {
	Message  string                                                      `json:"message"`
	Code     string                                                      `json:"code"`
	Issues   []RecipientDeleteTemplateRecipientTemplatesRecipientsIssues `json:"issues,omitempty"`
	HTTPMeta components.HTTPMetadata                                     `json:"-"`
}

var _ error = &RecipientDeleteTemplateRecipientTemplatesRecipientsResponseBody{}

func (e *RecipientDeleteTemplateRecipientTemplatesRecipientsResponseBody) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

type RecipientDeleteTemplateRecipientIssues struct {
	Message string `json:"message"`
}

func (o *RecipientDeleteTemplateRecipientIssues) GetMessage() string {
	if o == nil {
		return ""
	}
	return o.Message
}

// RecipientDeleteTemplateRecipientResponseBody - Invalid input data
type RecipientDeleteTemplateRecipientResponseBody struct {
	Message  string                                   `json:"message"`
	Code     string                                   `json:"code"`
	Issues   []RecipientDeleteTemplateRecipientIssues `json:"issues,omitempty"`
	HTTPMeta components.HTTPMetadata                  `json:"-"`
}

var _ error = &RecipientDeleteTemplateRecipientResponseBody{}

func (e *RecipientDeleteTemplateRecipientResponseBody) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
