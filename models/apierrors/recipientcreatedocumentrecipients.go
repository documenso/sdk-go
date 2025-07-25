// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package apierrors

import (
	"encoding/json"
	"github.com/documenso/sdk-go/models/components"
)

type RecipientCreateDocumentRecipientsInternalServerErrorIssue struct {
	Message string `json:"message"`
}

func (o *RecipientCreateDocumentRecipientsInternalServerErrorIssue) GetMessage() string {
	if o == nil {
		return ""
	}
	return o.Message
}

// RecipientCreateDocumentRecipientsInternalServerError - Internal server error
type RecipientCreateDocumentRecipientsInternalServerError struct {
	Message  string                                                      `json:"message"`
	Code     string                                                      `json:"code"`
	Issues   []RecipientCreateDocumentRecipientsInternalServerErrorIssue `json:"issues,omitempty"`
	HTTPMeta components.HTTPMetadata                                     `json:"-"`
}

var _ error = &RecipientCreateDocumentRecipientsInternalServerError{}

func (e *RecipientCreateDocumentRecipientsInternalServerError) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

type RecipientCreateDocumentRecipientsBadRequestIssue struct {
	Message string `json:"message"`
}

func (o *RecipientCreateDocumentRecipientsBadRequestIssue) GetMessage() string {
	if o == nil {
		return ""
	}
	return o.Message
}

// RecipientCreateDocumentRecipientsBadRequestError - Invalid input data
type RecipientCreateDocumentRecipientsBadRequestError struct {
	Message  string                                             `json:"message"`
	Code     string                                             `json:"code"`
	Issues   []RecipientCreateDocumentRecipientsBadRequestIssue `json:"issues,omitempty"`
	HTTPMeta components.HTTPMetadata                            `json:"-"`
}

var _ error = &RecipientCreateDocumentRecipientsBadRequestError{}

func (e *RecipientCreateDocumentRecipientsBadRequestError) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
