// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/documenso/sdk-go/models/components"
)

type RecipientDeleteDocumentRecipientRequestBody struct {
	RecipientID float64 `json:"recipientId"`
}

func (o *RecipientDeleteDocumentRecipientRequestBody) GetRecipientID() float64 {
	if o == nil {
		return 0.0
	}
	return o.RecipientID
}

// RecipientDeleteDocumentRecipientResponseBody - Successful response
type RecipientDeleteDocumentRecipientResponseBody struct {
	Success bool `json:"success"`
}

func (o *RecipientDeleteDocumentRecipientResponseBody) GetSuccess() bool {
	if o == nil {
		return false
	}
	return o.Success
}

type RecipientDeleteDocumentRecipientResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Successful response
	Object *RecipientDeleteDocumentRecipientResponseBody
}

func (o *RecipientDeleteDocumentRecipientResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *RecipientDeleteDocumentRecipientResponse) GetObject() *RecipientDeleteDocumentRecipientResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
