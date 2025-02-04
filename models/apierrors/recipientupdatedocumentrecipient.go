// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package apierrors

import (
	"encoding/json"
	"github.com/documenso/sdk-go/models/components"
)

type RecipientUpdateDocumentRecipientDocumentsRecipientsIssues struct {
	Message string `json:"message"`
}

func (o *RecipientUpdateDocumentRecipientDocumentsRecipientsIssues) GetMessage() string {
	if o == nil {
		return ""
	}
	return o.Message
}

// RecipientUpdateDocumentRecipientDocumentsRecipientsResponseBody - Internal server error
type RecipientUpdateDocumentRecipientDocumentsRecipientsResponseBody struct {
	Message  string                                                      `json:"message"`
	Code     string                                                      `json:"code"`
	Issues   []RecipientUpdateDocumentRecipientDocumentsRecipientsIssues `json:"issues,omitempty"`
	HTTPMeta components.HTTPMetadata                                     `json:"-"`
}

var _ error = &RecipientUpdateDocumentRecipientDocumentsRecipientsResponseBody{}

func (e *RecipientUpdateDocumentRecipientDocumentsRecipientsResponseBody) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

type RecipientUpdateDocumentRecipientIssues struct {
	Message string `json:"message"`
}

func (o *RecipientUpdateDocumentRecipientIssues) GetMessage() string {
	if o == nil {
		return ""
	}
	return o.Message
}

// RecipientUpdateDocumentRecipientResponseBody - Invalid input data
type RecipientUpdateDocumentRecipientResponseBody struct {
	Message  string                                   `json:"message"`
	Code     string                                   `json:"code"`
	Issues   []RecipientUpdateDocumentRecipientIssues `json:"issues,omitempty"`
	HTTPMeta components.HTTPMetadata                  `json:"-"`
}

var _ error = &RecipientUpdateDocumentRecipientResponseBody{}

func (e *RecipientUpdateDocumentRecipientResponseBody) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
