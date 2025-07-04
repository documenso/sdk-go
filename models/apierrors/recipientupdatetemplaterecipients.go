// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package apierrors

import (
	"encoding/json"
	"github.com/documenso/sdk-go/models/components"
)

type RecipientUpdateTemplateRecipientsInternalServerErrorIssue struct {
	Message string `json:"message"`
}

func (o *RecipientUpdateTemplateRecipientsInternalServerErrorIssue) GetMessage() string {
	if o == nil {
		return ""
	}
	return o.Message
}

// RecipientUpdateTemplateRecipientsInternalServerError - Internal server error
type RecipientUpdateTemplateRecipientsInternalServerError struct {
	Message  string                                                      `json:"message"`
	Code     string                                                      `json:"code"`
	Issues   []RecipientUpdateTemplateRecipientsInternalServerErrorIssue `json:"issues,omitempty"`
	HTTPMeta components.HTTPMetadata                                     `json:"-"`
}

var _ error = &RecipientUpdateTemplateRecipientsInternalServerError{}

func (e *RecipientUpdateTemplateRecipientsInternalServerError) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

type RecipientUpdateTemplateRecipientsBadRequestIssue struct {
	Message string `json:"message"`
}

func (o *RecipientUpdateTemplateRecipientsBadRequestIssue) GetMessage() string {
	if o == nil {
		return ""
	}
	return o.Message
}

// RecipientUpdateTemplateRecipientsBadRequestError - Invalid input data
type RecipientUpdateTemplateRecipientsBadRequestError struct {
	Message  string                                             `json:"message"`
	Code     string                                             `json:"code"`
	Issues   []RecipientUpdateTemplateRecipientsBadRequestIssue `json:"issues,omitempty"`
	HTTPMeta components.HTTPMetadata                            `json:"-"`
}

var _ error = &RecipientUpdateTemplateRecipientsBadRequestError{}

func (e *RecipientUpdateTemplateRecipientsBadRequestError) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
