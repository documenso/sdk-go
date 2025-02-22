// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package apierrors

import (
	"encoding/json"
	"github.com/documenso/sdk-go/models/components"
)

type DocumentDeleteDocumentDocumentsIssues struct {
	Message string `json:"message"`
}

func (o *DocumentDeleteDocumentDocumentsIssues) GetMessage() string {
	if o == nil {
		return ""
	}
	return o.Message
}

// DocumentDeleteDocumentDocumentsResponseBody - Internal server error
type DocumentDeleteDocumentDocumentsResponseBody struct {
	Message  string                                  `json:"message"`
	Code     string                                  `json:"code"`
	Issues   []DocumentDeleteDocumentDocumentsIssues `json:"issues,omitempty"`
	HTTPMeta components.HTTPMetadata                 `json:"-"`
}

var _ error = &DocumentDeleteDocumentDocumentsResponseBody{}

func (e *DocumentDeleteDocumentDocumentsResponseBody) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

type DocumentDeleteDocumentIssues struct {
	Message string `json:"message"`
}

func (o *DocumentDeleteDocumentIssues) GetMessage() string {
	if o == nil {
		return ""
	}
	return o.Message
}

// DocumentDeleteDocumentResponseBody - Invalid input data
type DocumentDeleteDocumentResponseBody struct {
	Message  string                         `json:"message"`
	Code     string                         `json:"code"`
	Issues   []DocumentDeleteDocumentIssues `json:"issues,omitempty"`
	HTTPMeta components.HTTPMetadata        `json:"-"`
}

var _ error = &DocumentDeleteDocumentResponseBody{}

func (e *DocumentDeleteDocumentResponseBody) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
