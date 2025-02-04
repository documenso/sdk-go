// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package apierrors

import (
	"encoding/json"
	"github.com/documenso/sdk-go/models/components"
)

type DocumentGetDocumentWithDetailsByIDDocumentsResponseIssues struct {
	Message string `json:"message"`
}

func (o *DocumentGetDocumentWithDetailsByIDDocumentsResponseIssues) GetMessage() string {
	if o == nil {
		return ""
	}
	return o.Message
}

// DocumentGetDocumentWithDetailsByIDDocumentsResponseResponseBody - Internal server error
type DocumentGetDocumentWithDetailsByIDDocumentsResponseResponseBody struct {
	Message  string                                                      `json:"message"`
	Code     string                                                      `json:"code"`
	Issues   []DocumentGetDocumentWithDetailsByIDDocumentsResponseIssues `json:"issues,omitempty"`
	HTTPMeta components.HTTPMetadata                                     `json:"-"`
}

var _ error = &DocumentGetDocumentWithDetailsByIDDocumentsResponseResponseBody{}

func (e *DocumentGetDocumentWithDetailsByIDDocumentsResponseResponseBody) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

type DocumentGetDocumentWithDetailsByIDDocumentsIssues struct {
	Message string `json:"message"`
}

func (o *DocumentGetDocumentWithDetailsByIDDocumentsIssues) GetMessage() string {
	if o == nil {
		return ""
	}
	return o.Message
}

// DocumentGetDocumentWithDetailsByIDDocumentsResponseBody - Not found
type DocumentGetDocumentWithDetailsByIDDocumentsResponseBody struct {
	Message  string                                              `json:"message"`
	Code     string                                              `json:"code"`
	Issues   []DocumentGetDocumentWithDetailsByIDDocumentsIssues `json:"issues,omitempty"`
	HTTPMeta components.HTTPMetadata                             `json:"-"`
}

var _ error = &DocumentGetDocumentWithDetailsByIDDocumentsResponseBody{}

func (e *DocumentGetDocumentWithDetailsByIDDocumentsResponseBody) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

type DocumentGetDocumentWithDetailsByIDIssues struct {
	Message string `json:"message"`
}

func (o *DocumentGetDocumentWithDetailsByIDIssues) GetMessage() string {
	if o == nil {
		return ""
	}
	return o.Message
}

// DocumentGetDocumentWithDetailsByIDResponseBody - Invalid input data
type DocumentGetDocumentWithDetailsByIDResponseBody struct {
	Message  string                                     `json:"message"`
	Code     string                                     `json:"code"`
	Issues   []DocumentGetDocumentWithDetailsByIDIssues `json:"issues,omitempty"`
	HTTPMeta components.HTTPMetadata                    `json:"-"`
}

var _ error = &DocumentGetDocumentWithDetailsByIDResponseBody{}

func (e *DocumentGetDocumentWithDetailsByIDResponseBody) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
