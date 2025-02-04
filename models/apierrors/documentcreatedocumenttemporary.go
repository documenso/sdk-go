// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package apierrors

import (
	"encoding/json"
	"github.com/documenso/sdk-go/models/components"
)

type DocumentCreateDocumentTemporaryDocumentsIssues struct {
	Message string `json:"message"`
}

func (o *DocumentCreateDocumentTemporaryDocumentsIssues) GetMessage() string {
	if o == nil {
		return ""
	}
	return o.Message
}

// DocumentCreateDocumentTemporaryDocumentsResponseBody - Internal server error
type DocumentCreateDocumentTemporaryDocumentsResponseBody struct {
	Message  string                                           `json:"message"`
	Code     string                                           `json:"code"`
	Issues   []DocumentCreateDocumentTemporaryDocumentsIssues `json:"issues,omitempty"`
	HTTPMeta components.HTTPMetadata                          `json:"-"`
}

var _ error = &DocumentCreateDocumentTemporaryDocumentsResponseBody{}

func (e *DocumentCreateDocumentTemporaryDocumentsResponseBody) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

type DocumentCreateDocumentTemporaryIssues struct {
	Message string `json:"message"`
}

func (o *DocumentCreateDocumentTemporaryIssues) GetMessage() string {
	if o == nil {
		return ""
	}
	return o.Message
}

// DocumentCreateDocumentTemporaryResponseBody - Invalid input data
type DocumentCreateDocumentTemporaryResponseBody struct {
	Message  string                                  `json:"message"`
	Code     string                                  `json:"code"`
	Issues   []DocumentCreateDocumentTemporaryIssues `json:"issues,omitempty"`
	HTTPMeta components.HTTPMetadata                 `json:"-"`
}

var _ error = &DocumentCreateDocumentTemporaryResponseBody{}

func (e *DocumentCreateDocumentTemporaryResponseBody) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
