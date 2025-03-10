// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package apierrors

import (
	"encoding/json"
	"github.com/documenso/sdk-go/models/components"
)

type FieldUpdateDocumentFieldsDocumentsFieldsIssues struct {
	Message string `json:"message"`
}

func (o *FieldUpdateDocumentFieldsDocumentsFieldsIssues) GetMessage() string {
	if o == nil {
		return ""
	}
	return o.Message
}

// FieldUpdateDocumentFieldsDocumentsFieldsResponseBody - Internal server error
type FieldUpdateDocumentFieldsDocumentsFieldsResponseBody struct {
	Message  string                                           `json:"message"`
	Code     string                                           `json:"code"`
	Issues   []FieldUpdateDocumentFieldsDocumentsFieldsIssues `json:"issues,omitempty"`
	HTTPMeta components.HTTPMetadata                          `json:"-"`
}

var _ error = &FieldUpdateDocumentFieldsDocumentsFieldsResponseBody{}

func (e *FieldUpdateDocumentFieldsDocumentsFieldsResponseBody) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

type FieldUpdateDocumentFieldsIssues struct {
	Message string `json:"message"`
}

func (o *FieldUpdateDocumentFieldsIssues) GetMessage() string {
	if o == nil {
		return ""
	}
	return o.Message
}

// FieldUpdateDocumentFieldsResponseBody - Invalid input data
type FieldUpdateDocumentFieldsResponseBody struct {
	Message  string                            `json:"message"`
	Code     string                            `json:"code"`
	Issues   []FieldUpdateDocumentFieldsIssues `json:"issues,omitempty"`
	HTTPMeta components.HTTPMetadata           `json:"-"`
}

var _ error = &FieldUpdateDocumentFieldsResponseBody{}

func (e *FieldUpdateDocumentFieldsResponseBody) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
