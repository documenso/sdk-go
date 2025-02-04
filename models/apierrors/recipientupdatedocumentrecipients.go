// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package apierrors

import (
	"encoding/json"
	"github.com/documenso/sdk-go/models/components"
)

type RecipientUpdateDocumentRecipientsDocumentsRecipientsIssues struct {
	Message string `json:"message"`
}

func (o *RecipientUpdateDocumentRecipientsDocumentsRecipientsIssues) GetMessage() string {
	if o == nil {
		return ""
	}
	return o.Message
}

// RecipientUpdateDocumentRecipientsDocumentsRecipientsResponseBody - Internal server error
type RecipientUpdateDocumentRecipientsDocumentsRecipientsResponseBody struct {
	Message  string                                                       `json:"message"`
	Code     string                                                       `json:"code"`
	Issues   []RecipientUpdateDocumentRecipientsDocumentsRecipientsIssues `json:"issues,omitempty"`
	HTTPMeta components.HTTPMetadata                                      `json:"-"`
}

var _ error = &RecipientUpdateDocumentRecipientsDocumentsRecipientsResponseBody{}

func (e *RecipientUpdateDocumentRecipientsDocumentsRecipientsResponseBody) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

type RecipientUpdateDocumentRecipientsIssues struct {
	Message string `json:"message"`
}

func (o *RecipientUpdateDocumentRecipientsIssues) GetMessage() string {
	if o == nil {
		return ""
	}
	return o.Message
}

// RecipientUpdateDocumentRecipientsResponseBody - Invalid input data
type RecipientUpdateDocumentRecipientsResponseBody struct {
	Message  string                                    `json:"message"`
	Code     string                                    `json:"code"`
	Issues   []RecipientUpdateDocumentRecipientsIssues `json:"issues,omitempty"`
	HTTPMeta components.HTTPMetadata                   `json:"-"`
}

var _ error = &RecipientUpdateDocumentRecipientsResponseBody{}

func (e *RecipientUpdateDocumentRecipientsResponseBody) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
