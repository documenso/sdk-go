// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/documenso/sdk-go/models/components"
)

type DocumentDeleteDocumentRequestBody struct {
	DocumentID float64 `json:"documentId"`
}

func (o *DocumentDeleteDocumentRequestBody) GetDocumentID() float64 {
	if o == nil {
		return 0.0
	}
	return o.DocumentID
}

// DocumentDeleteDocumentResponseBody - Successful response
type DocumentDeleteDocumentResponseBody struct {
	Success bool `json:"success"`
}

func (o *DocumentDeleteDocumentResponseBody) GetSuccess() bool {
	if o == nil {
		return false
	}
	return o.Success
}

type DocumentDeleteDocumentResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Successful response
	Object *DocumentDeleteDocumentResponseBody
}

func (o *DocumentDeleteDocumentResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *DocumentDeleteDocumentResponse) GetObject() *DocumentDeleteDocumentResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
