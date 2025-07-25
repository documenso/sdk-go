// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/documenso/sdk-go/internal/utils"
	"github.com/documenso/sdk-go/models/components"
)

// DocumentSendDocumentDateFormat - The date format to use for date fields and signing the document.
type DocumentSendDocumentDateFormat string

const (
	DocumentSendDocumentDateFormatYyyyMmDdHhMmA         DocumentSendDocumentDateFormat = "yyyy-MM-dd hh:mm a"
	DocumentSendDocumentDateFormatYyyyMmDd              DocumentSendDocumentDateFormat = "yyyy-MM-dd"
	DocumentSendDocumentDateFormatDdMmYyyyHhMmA         DocumentSendDocumentDateFormat = "dd/MM/yyyy hh:mm a"
	DocumentSendDocumentDateFormatMmDdYyyyHhMmA         DocumentSendDocumentDateFormat = "MM/dd/yyyy hh:mm a"
	DocumentSendDocumentDateFormatYyyyMmDdHhMm          DocumentSendDocumentDateFormat = "yyyy-MM-dd HH:mm"
	DocumentSendDocumentDateFormatYyMmDdHhMmA           DocumentSendDocumentDateFormat = "yy-MM-dd hh:mm a"
	DocumentSendDocumentDateFormatYyyyMmDdHhMmSs        DocumentSendDocumentDateFormat = "yyyy-MM-dd HH:mm:ss"
	DocumentSendDocumentDateFormatMmmmDdYyyyHhMmA       DocumentSendDocumentDateFormat = "MMMM dd, yyyy hh:mm a"
	DocumentSendDocumentDateFormatEeeeMmmmDdYyyyHhMmA   DocumentSendDocumentDateFormat = "EEEE, MMMM dd, yyyy hh:mm a"
	DocumentSendDocumentDateFormatYyyyMmDdTHhMmSsSssxxx DocumentSendDocumentDateFormat = "yyyy-MM-dd'T'HH:mm:ss.SSSXXX"
)

func (e DocumentSendDocumentDateFormat) ToPointer() *DocumentSendDocumentDateFormat {
	return &e
}
func (e *DocumentSendDocumentDateFormat) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "yyyy-MM-dd hh:mm a":
		fallthrough
	case "yyyy-MM-dd":
		fallthrough
	case "dd/MM/yyyy hh:mm a":
		fallthrough
	case "MM/dd/yyyy hh:mm a":
		fallthrough
	case "yyyy-MM-dd HH:mm":
		fallthrough
	case "yy-MM-dd hh:mm a":
		fallthrough
	case "yyyy-MM-dd HH:mm:ss":
		fallthrough
	case "MMMM dd, yyyy hh:mm a":
		fallthrough
	case "EEEE, MMMM dd, yyyy hh:mm a":
		fallthrough
	case "yyyy-MM-dd'T'HH:mm:ss.SSSXXX":
		*e = DocumentSendDocumentDateFormat(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DocumentSendDocumentDateFormat: %v", v)
	}
}

// DocumentSendDocumentDistributionMethod - The distribution method to use when sending the document to the recipients.
type DocumentSendDocumentDistributionMethod string

const (
	DocumentSendDocumentDistributionMethodEmail DocumentSendDocumentDistributionMethod = "EMAIL"
	DocumentSendDocumentDistributionMethodNone  DocumentSendDocumentDistributionMethod = "NONE"
)

func (e DocumentSendDocumentDistributionMethod) ToPointer() *DocumentSendDocumentDistributionMethod {
	return &e
}
func (e *DocumentSendDocumentDistributionMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "EMAIL":
		fallthrough
	case "NONE":
		*e = DocumentSendDocumentDistributionMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DocumentSendDocumentDistributionMethod: %v", v)
	}
}

// DocumentSendDocumentLanguage - The language to use for email communications with recipients.
type DocumentSendDocumentLanguage string

const (
	DocumentSendDocumentLanguageDe DocumentSendDocumentLanguage = "de"
	DocumentSendDocumentLanguageEn DocumentSendDocumentLanguage = "en"
	DocumentSendDocumentLanguageFr DocumentSendDocumentLanguage = "fr"
	DocumentSendDocumentLanguageEs DocumentSendDocumentLanguage = "es"
	DocumentSendDocumentLanguageIt DocumentSendDocumentLanguage = "it"
	DocumentSendDocumentLanguagePl DocumentSendDocumentLanguage = "pl"
)

func (e DocumentSendDocumentLanguage) ToPointer() *DocumentSendDocumentLanguage {
	return &e
}
func (e *DocumentSendDocumentLanguage) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "de":
		fallthrough
	case "en":
		fallthrough
	case "fr":
		fallthrough
	case "es":
		fallthrough
	case "it":
		fallthrough
	case "pl":
		*e = DocumentSendDocumentLanguage(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DocumentSendDocumentLanguage: %v", v)
	}
}

type DocumentSendDocumentEmailSettings struct {
	// Whether to send an email to all recipients that the document is ready for them to sign.
	RecipientSigningRequest *bool `default:"true" json:"recipientSigningRequest"`
	// Whether to send an email to the recipient who was removed from a pending document.
	RecipientRemoved *bool `default:"true" json:"recipientRemoved"`
	// Whether to send an email to the document owner when a recipient has signed the document.
	RecipientSigned *bool `default:"true" json:"recipientSigned"`
	// Whether to send an email to the recipient who has just signed the document indicating that there are still other recipients who need to sign the document. This will only be sent if the document is still pending after the recipient has signed.
	DocumentPending *bool `default:"true" json:"documentPending"`
	// Whether to send an email to all recipients when the document is complete.
	DocumentCompleted *bool `default:"true" json:"documentCompleted"`
	// Whether to send an email to all recipients if a pending document has been deleted.
	DocumentDeleted *bool `default:"true" json:"documentDeleted"`
	// Whether to send an email to the document owner when the document is complete.
	OwnerDocumentCompleted *bool `default:"true" json:"ownerDocumentCompleted"`
}

func (d DocumentSendDocumentEmailSettings) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DocumentSendDocumentEmailSettings) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *DocumentSendDocumentEmailSettings) GetRecipientSigningRequest() *bool {
	if o == nil {
		return nil
	}
	return o.RecipientSigningRequest
}

func (o *DocumentSendDocumentEmailSettings) GetRecipientRemoved() *bool {
	if o == nil {
		return nil
	}
	return o.RecipientRemoved
}

func (o *DocumentSendDocumentEmailSettings) GetRecipientSigned() *bool {
	if o == nil {
		return nil
	}
	return o.RecipientSigned
}

func (o *DocumentSendDocumentEmailSettings) GetDocumentPending() *bool {
	if o == nil {
		return nil
	}
	return o.DocumentPending
}

func (o *DocumentSendDocumentEmailSettings) GetDocumentCompleted() *bool {
	if o == nil {
		return nil
	}
	return o.DocumentCompleted
}

func (o *DocumentSendDocumentEmailSettings) GetDocumentDeleted() *bool {
	if o == nil {
		return nil
	}
	return o.DocumentDeleted
}

func (o *DocumentSendDocumentEmailSettings) GetOwnerDocumentCompleted() *bool {
	if o == nil {
		return nil
	}
	return o.OwnerDocumentCompleted
}

type DocumentSendDocumentMeta struct {
	// The subject of the email that will be sent to the recipients.
	Subject *string `json:"subject,omitempty"`
	// The message of the email that will be sent to the recipients.
	Message *string `json:"message,omitempty"`
	// The timezone to use for date fields and signing the document. Example Etc/UTC, Australia/Melbourne
	Timezone *string `json:"timezone,omitempty"`
	// The date format to use for date fields and signing the document.
	DateFormat *DocumentSendDocumentDateFormat `json:"dateFormat,omitempty"`
	// The distribution method to use when sending the document to the recipients.
	DistributionMethod *DocumentSendDocumentDistributionMethod `json:"distributionMethod,omitempty"`
	// The URL to which the recipient should be redirected after signing the document.
	RedirectURL *string `json:"redirectUrl,omitempty"`
	// The language to use for email communications with recipients.
	Language      *DocumentSendDocumentLanguage      `json:"language,omitempty"`
	EmailSettings *DocumentSendDocumentEmailSettings `json:"emailSettings,omitempty"`
}

func (o *DocumentSendDocumentMeta) GetSubject() *string {
	if o == nil {
		return nil
	}
	return o.Subject
}

func (o *DocumentSendDocumentMeta) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *DocumentSendDocumentMeta) GetTimezone() *string {
	if o == nil {
		return nil
	}
	return o.Timezone
}

func (o *DocumentSendDocumentMeta) GetDateFormat() *DocumentSendDocumentDateFormat {
	if o == nil {
		return nil
	}
	return o.DateFormat
}

func (o *DocumentSendDocumentMeta) GetDistributionMethod() *DocumentSendDocumentDistributionMethod {
	if o == nil {
		return nil
	}
	return o.DistributionMethod
}

func (o *DocumentSendDocumentMeta) GetRedirectURL() *string {
	if o == nil {
		return nil
	}
	return o.RedirectURL
}

func (o *DocumentSendDocumentMeta) GetLanguage() *DocumentSendDocumentLanguage {
	if o == nil {
		return nil
	}
	return o.Language
}

func (o *DocumentSendDocumentMeta) GetEmailSettings() *DocumentSendDocumentEmailSettings {
	if o == nil {
		return nil
	}
	return o.EmailSettings
}

type DocumentSendDocumentRequest struct {
	// The ID of the document to send.
	DocumentID float64                   `json:"documentId"`
	Meta       *DocumentSendDocumentMeta `json:"meta,omitempty"`
}

func (o *DocumentSendDocumentRequest) GetDocumentID() float64 {
	if o == nil {
		return 0.0
	}
	return o.DocumentID
}

func (o *DocumentSendDocumentRequest) GetMeta() *DocumentSendDocumentMeta {
	if o == nil {
		return nil
	}
	return o.Meta
}

type DocumentSendDocumentVisibility string

const (
	DocumentSendDocumentVisibilityEveryone        DocumentSendDocumentVisibility = "EVERYONE"
	DocumentSendDocumentVisibilityManagerAndAbove DocumentSendDocumentVisibility = "MANAGER_AND_ABOVE"
	DocumentSendDocumentVisibilityAdmin           DocumentSendDocumentVisibility = "ADMIN"
)

func (e DocumentSendDocumentVisibility) ToPointer() *DocumentSendDocumentVisibility {
	return &e
}
func (e *DocumentSendDocumentVisibility) UnmarshalJSON(data []byte) error {
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
		*e = DocumentSendDocumentVisibility(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DocumentSendDocumentVisibility: %v", v)
	}
}

type DocumentSendDocumentStatus string

const (
	DocumentSendDocumentStatusDraft     DocumentSendDocumentStatus = "DRAFT"
	DocumentSendDocumentStatusPending   DocumentSendDocumentStatus = "PENDING"
	DocumentSendDocumentStatusCompleted DocumentSendDocumentStatus = "COMPLETED"
	DocumentSendDocumentStatusRejected  DocumentSendDocumentStatus = "REJECTED"
)

func (e DocumentSendDocumentStatus) ToPointer() *DocumentSendDocumentStatus {
	return &e
}
func (e *DocumentSendDocumentStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "DRAFT":
		fallthrough
	case "PENDING":
		fallthrough
	case "COMPLETED":
		fallthrough
	case "REJECTED":
		*e = DocumentSendDocumentStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DocumentSendDocumentStatus: %v", v)
	}
}

type DocumentSendDocumentSource string

const (
	DocumentSendDocumentSourceDocument           DocumentSendDocumentSource = "DOCUMENT"
	DocumentSendDocumentSourceTemplate           DocumentSendDocumentSource = "TEMPLATE"
	DocumentSendDocumentSourceTemplateDirectLink DocumentSendDocumentSource = "TEMPLATE_DIRECT_LINK"
)

func (e DocumentSendDocumentSource) ToPointer() *DocumentSendDocumentSource {
	return &e
}
func (e *DocumentSendDocumentSource) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "DOCUMENT":
		fallthrough
	case "TEMPLATE":
		fallthrough
	case "TEMPLATE_DIRECT_LINK":
		*e = DocumentSendDocumentSource(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DocumentSendDocumentSource: %v", v)
	}
}

// DocumentSendDocumentGlobalAccessAuth - The type of authentication required for the recipient to access the document.
type DocumentSendDocumentGlobalAccessAuth string

const (
	DocumentSendDocumentGlobalAccessAuthAccount DocumentSendDocumentGlobalAccessAuth = "ACCOUNT"
)

func (e DocumentSendDocumentGlobalAccessAuth) ToPointer() *DocumentSendDocumentGlobalAccessAuth {
	return &e
}
func (e *DocumentSendDocumentGlobalAccessAuth) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ACCOUNT":
		*e = DocumentSendDocumentGlobalAccessAuth(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DocumentSendDocumentGlobalAccessAuth: %v", v)
	}
}

// DocumentSendDocumentGlobalActionAuth - The type of authentication required for the recipient to sign the document. This field is restricted to Enterprise plan users only.
type DocumentSendDocumentGlobalActionAuth string

const (
	DocumentSendDocumentGlobalActionAuthAccount       DocumentSendDocumentGlobalActionAuth = "ACCOUNT"
	DocumentSendDocumentGlobalActionAuthPasskey       DocumentSendDocumentGlobalActionAuth = "PASSKEY"
	DocumentSendDocumentGlobalActionAuthTwoFactorAuth DocumentSendDocumentGlobalActionAuth = "TWO_FACTOR_AUTH"
)

func (e DocumentSendDocumentGlobalActionAuth) ToPointer() *DocumentSendDocumentGlobalActionAuth {
	return &e
}
func (e *DocumentSendDocumentGlobalActionAuth) UnmarshalJSON(data []byte) error {
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
		*e = DocumentSendDocumentGlobalActionAuth(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DocumentSendDocumentGlobalActionAuth: %v", v)
	}
}

type DocumentSendDocumentAuthOptions struct {
	// The type of authentication required for the recipient to access the document.
	GlobalAccessAuth *DocumentSendDocumentGlobalAccessAuth `json:"globalAccessAuth"`
	// The type of authentication required for the recipient to sign the document. This field is restricted to Enterprise plan users only.
	GlobalActionAuth *DocumentSendDocumentGlobalActionAuth `json:"globalActionAuth"`
}

func (o *DocumentSendDocumentAuthOptions) GetGlobalAccessAuth() *DocumentSendDocumentGlobalAccessAuth {
	if o == nil {
		return nil
	}
	return o.GlobalAccessAuth
}

func (o *DocumentSendDocumentAuthOptions) GetGlobalActionAuth() *DocumentSendDocumentGlobalActionAuth {
	if o == nil {
		return nil
	}
	return o.GlobalActionAuth
}

type DocumentSendDocumentFormValuesType string

const (
	DocumentSendDocumentFormValuesTypeStr     DocumentSendDocumentFormValuesType = "str"
	DocumentSendDocumentFormValuesTypeBoolean DocumentSendDocumentFormValuesType = "boolean"
	DocumentSendDocumentFormValuesTypeNumber  DocumentSendDocumentFormValuesType = "number"
)

type DocumentSendDocumentFormValues struct {
	Str     *string  `queryParam:"inline"`
	Boolean *bool    `queryParam:"inline"`
	Number  *float64 `queryParam:"inline"`

	Type DocumentSendDocumentFormValuesType
}

func CreateDocumentSendDocumentFormValuesStr(str string) DocumentSendDocumentFormValues {
	typ := DocumentSendDocumentFormValuesTypeStr

	return DocumentSendDocumentFormValues{
		Str:  &str,
		Type: typ,
	}
}

func CreateDocumentSendDocumentFormValuesBoolean(boolean bool) DocumentSendDocumentFormValues {
	typ := DocumentSendDocumentFormValuesTypeBoolean

	return DocumentSendDocumentFormValues{
		Boolean: &boolean,
		Type:    typ,
	}
}

func CreateDocumentSendDocumentFormValuesNumber(number float64) DocumentSendDocumentFormValues {
	typ := DocumentSendDocumentFormValuesTypeNumber

	return DocumentSendDocumentFormValues{
		Number: &number,
		Type:   typ,
	}
}

func (u *DocumentSendDocumentFormValues) UnmarshalJSON(data []byte) error {

	var str string = ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = DocumentSendDocumentFormValuesTypeStr
		return nil
	}

	var boolean bool = false
	if err := utils.UnmarshalJSON(data, &boolean, "", true, true); err == nil {
		u.Boolean = &boolean
		u.Type = DocumentSendDocumentFormValuesTypeBoolean
		return nil
	}

	var number float64 = float64(0)
	if err := utils.UnmarshalJSON(data, &number, "", true, true); err == nil {
		u.Number = &number
		u.Type = DocumentSendDocumentFormValuesTypeNumber
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for DocumentSendDocumentFormValues", string(data))
}

func (u DocumentSendDocumentFormValues) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.Boolean != nil {
		return utils.MarshalJSON(u.Boolean, "", true)
	}

	if u.Number != nil {
		return utils.MarshalJSON(u.Number, "", true)
	}

	return nil, errors.New("could not marshal union type DocumentSendDocumentFormValues: all fields are null")
}

// DocumentSendDocumentResponseBody - Successful response
type DocumentSendDocumentResponseBody struct {
	Visibility DocumentSendDocumentVisibility `json:"visibility"`
	Status     DocumentSendDocumentStatus     `json:"status"`
	Source     DocumentSendDocumentSource     `json:"source"`
	ID         float64                        `json:"id"`
	// A custom external ID you can use to identify the document.
	ExternalID *string `json:"externalId"`
	// The ID of the user that created this document.
	UserID         float64                                   `json:"userId"`
	AuthOptions    *DocumentSendDocumentAuthOptions          `json:"authOptions"`
	FormValues     map[string]DocumentSendDocumentFormValues `json:"formValues"`
	Title          string                                    `json:"title"`
	DocumentDataID string                                    `json:"documentDataId"`
	CreatedAt      string                                    `json:"createdAt"`
	UpdatedAt      string                                    `json:"updatedAt"`
	CompletedAt    *string                                   `json:"completedAt"`
	DeletedAt      *string                                   `json:"deletedAt"`
	TeamID         *float64                                  `json:"teamId"`
	TemplateID     *float64                                  `json:"templateId"`
}

func (o *DocumentSendDocumentResponseBody) GetVisibility() DocumentSendDocumentVisibility {
	if o == nil {
		return DocumentSendDocumentVisibility("")
	}
	return o.Visibility
}

func (o *DocumentSendDocumentResponseBody) GetStatus() DocumentSendDocumentStatus {
	if o == nil {
		return DocumentSendDocumentStatus("")
	}
	return o.Status
}

func (o *DocumentSendDocumentResponseBody) GetSource() DocumentSendDocumentSource {
	if o == nil {
		return DocumentSendDocumentSource("")
	}
	return o.Source
}

func (o *DocumentSendDocumentResponseBody) GetID() float64 {
	if o == nil {
		return 0.0
	}
	return o.ID
}

func (o *DocumentSendDocumentResponseBody) GetExternalID() *string {
	if o == nil {
		return nil
	}
	return o.ExternalID
}

func (o *DocumentSendDocumentResponseBody) GetUserID() float64 {
	if o == nil {
		return 0.0
	}
	return o.UserID
}

func (o *DocumentSendDocumentResponseBody) GetAuthOptions() *DocumentSendDocumentAuthOptions {
	if o == nil {
		return nil
	}
	return o.AuthOptions
}

func (o *DocumentSendDocumentResponseBody) GetFormValues() map[string]DocumentSendDocumentFormValues {
	if o == nil {
		return nil
	}
	return o.FormValues
}

func (o *DocumentSendDocumentResponseBody) GetTitle() string {
	if o == nil {
		return ""
	}
	return o.Title
}

func (o *DocumentSendDocumentResponseBody) GetDocumentDataID() string {
	if o == nil {
		return ""
	}
	return o.DocumentDataID
}

func (o *DocumentSendDocumentResponseBody) GetCreatedAt() string {
	if o == nil {
		return ""
	}
	return o.CreatedAt
}

func (o *DocumentSendDocumentResponseBody) GetUpdatedAt() string {
	if o == nil {
		return ""
	}
	return o.UpdatedAt
}

func (o *DocumentSendDocumentResponseBody) GetCompletedAt() *string {
	if o == nil {
		return nil
	}
	return o.CompletedAt
}

func (o *DocumentSendDocumentResponseBody) GetDeletedAt() *string {
	if o == nil {
		return nil
	}
	return o.DeletedAt
}

func (o *DocumentSendDocumentResponseBody) GetTeamID() *float64 {
	if o == nil {
		return nil
	}
	return o.TeamID
}

func (o *DocumentSendDocumentResponseBody) GetTemplateID() *float64 {
	if o == nil {
		return nil
	}
	return o.TemplateID
}

type DocumentSendDocumentResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Successful response
	Object *DocumentSendDocumentResponseBody
}

func (o *DocumentSendDocumentResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *DocumentSendDocumentResponse) GetObject() *DocumentSendDocumentResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
