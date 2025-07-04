// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/documenso/sdk-go/models/components"
)

type TemplateMoveTemplateToTeamRequest struct {
	// The ID of the template to move to.
	TemplateID float64 `json:"templateId"`
	// The ID of the team to move the template to.
	TeamID float64 `json:"teamId"`
}

func (o *TemplateMoveTemplateToTeamRequest) GetTemplateID() float64 {
	if o == nil {
		return 0.0
	}
	return o.TemplateID
}

func (o *TemplateMoveTemplateToTeamRequest) GetTeamID() float64 {
	if o == nil {
		return 0.0
	}
	return o.TeamID
}

type TemplateMoveTemplateToTeamType string

const (
	TemplateMoveTemplateToTeamTypePublic  TemplateMoveTemplateToTeamType = "PUBLIC"
	TemplateMoveTemplateToTeamTypePrivate TemplateMoveTemplateToTeamType = "PRIVATE"
)

func (e TemplateMoveTemplateToTeamType) ToPointer() *TemplateMoveTemplateToTeamType {
	return &e
}
func (e *TemplateMoveTemplateToTeamType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "PUBLIC":
		fallthrough
	case "PRIVATE":
		*e = TemplateMoveTemplateToTeamType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for TemplateMoveTemplateToTeamType: %v", v)
	}
}

type TemplateMoveTemplateToTeamVisibility string

const (
	TemplateMoveTemplateToTeamVisibilityEveryone        TemplateMoveTemplateToTeamVisibility = "EVERYONE"
	TemplateMoveTemplateToTeamVisibilityManagerAndAbove TemplateMoveTemplateToTeamVisibility = "MANAGER_AND_ABOVE"
	TemplateMoveTemplateToTeamVisibilityAdmin           TemplateMoveTemplateToTeamVisibility = "ADMIN"
)

func (e TemplateMoveTemplateToTeamVisibility) ToPointer() *TemplateMoveTemplateToTeamVisibility {
	return &e
}
func (e *TemplateMoveTemplateToTeamVisibility) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "EVERYONE":
		fallthrough
	case "MANAGER_AND_ABOVE":
		fallthrough
	case "ADMIN":
		*e = TemplateMoveTemplateToTeamVisibility(v)
		return nil
	default:
		return fmt.Errorf("invalid value for TemplateMoveTemplateToTeamVisibility: %v", v)
	}
}

// TemplateMoveTemplateToTeamGlobalAccessAuth - The type of authentication required for the recipient to access the document.
type TemplateMoveTemplateToTeamGlobalAccessAuth string

const (
	TemplateMoveTemplateToTeamGlobalAccessAuthAccount TemplateMoveTemplateToTeamGlobalAccessAuth = "ACCOUNT"
)

func (e TemplateMoveTemplateToTeamGlobalAccessAuth) ToPointer() *TemplateMoveTemplateToTeamGlobalAccessAuth {
	return &e
}
func (e *TemplateMoveTemplateToTeamGlobalAccessAuth) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ACCOUNT":
		*e = TemplateMoveTemplateToTeamGlobalAccessAuth(v)
		return nil
	default:
		return fmt.Errorf("invalid value for TemplateMoveTemplateToTeamGlobalAccessAuth: %v", v)
	}
}

// TemplateMoveTemplateToTeamGlobalActionAuth - The type of authentication required for the recipient to sign the document. This field is restricted to Enterprise plan users only.
type TemplateMoveTemplateToTeamGlobalActionAuth string

const (
	TemplateMoveTemplateToTeamGlobalActionAuthAccount       TemplateMoveTemplateToTeamGlobalActionAuth = "ACCOUNT"
	TemplateMoveTemplateToTeamGlobalActionAuthPasskey       TemplateMoveTemplateToTeamGlobalActionAuth = "PASSKEY"
	TemplateMoveTemplateToTeamGlobalActionAuthTwoFactorAuth TemplateMoveTemplateToTeamGlobalActionAuth = "TWO_FACTOR_AUTH"
)

func (e TemplateMoveTemplateToTeamGlobalActionAuth) ToPointer() *TemplateMoveTemplateToTeamGlobalActionAuth {
	return &e
}
func (e *TemplateMoveTemplateToTeamGlobalActionAuth) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ACCOUNT":
		fallthrough
	case "PASSKEY":
		fallthrough
	case "TWO_FACTOR_AUTH":
		*e = TemplateMoveTemplateToTeamGlobalActionAuth(v)
		return nil
	default:
		return fmt.Errorf("invalid value for TemplateMoveTemplateToTeamGlobalActionAuth: %v", v)
	}
}

type TemplateMoveTemplateToTeamAuthOptions struct {
	// The type of authentication required for the recipient to access the document.
	GlobalAccessAuth *TemplateMoveTemplateToTeamGlobalAccessAuth `json:"globalAccessAuth"`
	// The type of authentication required for the recipient to sign the document. This field is restricted to Enterprise plan users only.
	GlobalActionAuth *TemplateMoveTemplateToTeamGlobalActionAuth `json:"globalActionAuth"`
}

func (o *TemplateMoveTemplateToTeamAuthOptions) GetGlobalAccessAuth() *TemplateMoveTemplateToTeamGlobalAccessAuth {
	if o == nil {
		return nil
	}
	return o.GlobalAccessAuth
}

func (o *TemplateMoveTemplateToTeamAuthOptions) GetGlobalActionAuth() *TemplateMoveTemplateToTeamGlobalActionAuth {
	if o == nil {
		return nil
	}
	return o.GlobalActionAuth
}

// TemplateMoveTemplateToTeamResponseBody - Successful response
type TemplateMoveTemplateToTeamResponseBody struct {
	Type                   TemplateMoveTemplateToTeamType         `json:"type"`
	Visibility             TemplateMoveTemplateToTeamVisibility   `json:"visibility"`
	ID                     float64                                `json:"id"`
	ExternalID             *string                                `json:"externalId"`
	Title                  string                                 `json:"title"`
	UserID                 float64                                `json:"userId"`
	TeamID                 *float64                               `json:"teamId"`
	AuthOptions            *TemplateMoveTemplateToTeamAuthOptions `json:"authOptions"`
	TemplateDocumentDataID string                                 `json:"templateDocumentDataId"`
	CreatedAt              string                                 `json:"createdAt"`
	UpdatedAt              string                                 `json:"updatedAt"`
	PublicTitle            string                                 `json:"publicTitle"`
	PublicDescription      string                                 `json:"publicDescription"`
}

func (o *TemplateMoveTemplateToTeamResponseBody) GetType() TemplateMoveTemplateToTeamType {
	if o == nil {
		return TemplateMoveTemplateToTeamType("")
	}
	return o.Type
}

func (o *TemplateMoveTemplateToTeamResponseBody) GetVisibility() TemplateMoveTemplateToTeamVisibility {
	if o == nil {
		return TemplateMoveTemplateToTeamVisibility("")
	}
	return o.Visibility
}

func (o *TemplateMoveTemplateToTeamResponseBody) GetID() float64 {
	if o == nil {
		return 0.0
	}
	return o.ID
}

func (o *TemplateMoveTemplateToTeamResponseBody) GetExternalID() *string {
	if o == nil {
		return nil
	}
	return o.ExternalID
}

func (o *TemplateMoveTemplateToTeamResponseBody) GetTitle() string {
	if o == nil {
		return ""
	}
	return o.Title
}

func (o *TemplateMoveTemplateToTeamResponseBody) GetUserID() float64 {
	if o == nil {
		return 0.0
	}
	return o.UserID
}

func (o *TemplateMoveTemplateToTeamResponseBody) GetTeamID() *float64 {
	if o == nil {
		return nil
	}
	return o.TeamID
}

func (o *TemplateMoveTemplateToTeamResponseBody) GetAuthOptions() *TemplateMoveTemplateToTeamAuthOptions {
	if o == nil {
		return nil
	}
	return o.AuthOptions
}

func (o *TemplateMoveTemplateToTeamResponseBody) GetTemplateDocumentDataID() string {
	if o == nil {
		return ""
	}
	return o.TemplateDocumentDataID
}

func (o *TemplateMoveTemplateToTeamResponseBody) GetCreatedAt() string {
	if o == nil {
		return ""
	}
	return o.CreatedAt
}

func (o *TemplateMoveTemplateToTeamResponseBody) GetUpdatedAt() string {
	if o == nil {
		return ""
	}
	return o.UpdatedAt
}

func (o *TemplateMoveTemplateToTeamResponseBody) GetPublicTitle() string {
	if o == nil {
		return ""
	}
	return o.PublicTitle
}

func (o *TemplateMoveTemplateToTeamResponseBody) GetPublicDescription() string {
	if o == nil {
		return ""
	}
	return o.PublicDescription
}

type TemplateMoveTemplateToTeamResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Successful response
	Object *TemplateMoveTemplateToTeamResponseBody
}

func (o *TemplateMoveTemplateToTeamResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *TemplateMoveTemplateToTeamResponse) GetObject() *TemplateMoveTemplateToTeamResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
