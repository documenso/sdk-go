// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package apierrors

import (
	"encoding/json"
	"github.com/documenso/sdk-go/models/components"
)

type DocumentDeleteDocumentInternalServerErrorIssue struct {
	Message string `json:"message"`
}

func (o *DocumentDeleteDocumentInternalServerErrorIssue) GetMessage() string {
	if o == nil {
		return ""
	}
	return o.Message
}

// DocumentDeleteDocumentInternalServerError - Internal server error
type DocumentDeleteDocumentInternalServerError struct {
	Message  string                                           `json:"message"`
	Code     string                                           `json:"code"`
	Issues   []DocumentDeleteDocumentInternalServerErrorIssue `json:"issues,omitempty"`
	HTTPMeta components.HTTPMetadata                          `json:"-"`
}

var _ error = &DocumentDeleteDocumentInternalServerError{}

func (e *DocumentDeleteDocumentInternalServerError) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

type DocumentDeleteDocumentBadRequestIssue struct {
	Message string `json:"message"`
}

func (o *DocumentDeleteDocumentBadRequestIssue) GetMessage() string {
	if o == nil {
		return ""
	}
	return o.Message
}

// DocumentDeleteDocumentBadRequestError - Invalid input data
type DocumentDeleteDocumentBadRequestError struct {
	Message  string                                  `json:"message"`
	Code     string                                  `json:"code"`
	Issues   []DocumentDeleteDocumentBadRequestIssue `json:"issues,omitempty"`
	HTTPMeta components.HTTPMetadata                 `json:"-"`
}

var _ error = &DocumentDeleteDocumentBadRequestError{}

func (e *DocumentDeleteDocumentBadRequestError) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
