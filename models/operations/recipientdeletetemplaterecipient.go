// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/documenso/sdk-go/models/components"
)

type RecipientDeleteTemplateRecipientRequest struct {
	RecipientID float64 `json:"recipientId"`
}

func (o *RecipientDeleteTemplateRecipientRequest) GetRecipientID() float64 {
	if o == nil {
		return 0.0
	}
	return o.RecipientID
}

// RecipientDeleteTemplateRecipientResponseBody - Successful response
type RecipientDeleteTemplateRecipientResponseBody struct {
	Success bool `json:"success"`
}

func (o *RecipientDeleteTemplateRecipientResponseBody) GetSuccess() bool {
	if o == nil {
		return false
	}
	return o.Success
}

type RecipientDeleteTemplateRecipientResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Successful response
	Object *RecipientDeleteTemplateRecipientResponseBody
}

func (o *RecipientDeleteTemplateRecipientResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *RecipientDeleteTemplateRecipientResponse) GetObject() *RecipientDeleteTemplateRecipientResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
