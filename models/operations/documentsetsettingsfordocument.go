// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/documenso/sdk-go/internal/utils"
	"github.com/documenso/sdk-go/models/components"
)

// DocumentSetSettingsForDocumentVisibility - The visibility of the document.
type DocumentSetSettingsForDocumentVisibility string

const (
	DocumentSetSettingsForDocumentVisibilityEveryone        DocumentSetSettingsForDocumentVisibility = "EVERYONE"
	DocumentSetSettingsForDocumentVisibilityManagerAndAbove DocumentSetSettingsForDocumentVisibility = "MANAGER_AND_ABOVE"
	DocumentSetSettingsForDocumentVisibilityAdmin           DocumentSetSettingsForDocumentVisibility = "ADMIN"
)

func (e DocumentSetSettingsForDocumentVisibility) ToPointer() *DocumentSetSettingsForDocumentVisibility {
	return &e
}
func (e *DocumentSetSettingsForDocumentVisibility) UnmarshalJSON(data []byte) error {
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
		*e = DocumentSetSettingsForDocumentVisibility(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DocumentSetSettingsForDocumentVisibility: %v", v)
	}
}

// DocumentSetSettingsForDocumentGlobalAccessAuth - The type of authentication required for the recipient to access the document.
type DocumentSetSettingsForDocumentGlobalAccessAuth string

const (
	DocumentSetSettingsForDocumentGlobalAccessAuthAccount DocumentSetSettingsForDocumentGlobalAccessAuth = "ACCOUNT"
)

func (e DocumentSetSettingsForDocumentGlobalAccessAuth) ToPointer() *DocumentSetSettingsForDocumentGlobalAccessAuth {
	return &e
}
func (e *DocumentSetSettingsForDocumentGlobalAccessAuth) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ACCOUNT":
		*e = DocumentSetSettingsForDocumentGlobalAccessAuth(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DocumentSetSettingsForDocumentGlobalAccessAuth: %v", v)
	}
}

// DocumentSetSettingsForDocumentGlobalActionAuth - The type of authentication required for the recipient to sign the document. This field is restricted to Enterprise plan users only.
type DocumentSetSettingsForDocumentGlobalActionAuth string

const (
	DocumentSetSettingsForDocumentGlobalActionAuthAccount       DocumentSetSettingsForDocumentGlobalActionAuth = "ACCOUNT"
	DocumentSetSettingsForDocumentGlobalActionAuthPasskey       DocumentSetSettingsForDocumentGlobalActionAuth = "PASSKEY"
	DocumentSetSettingsForDocumentGlobalActionAuthTwoFactorAuth DocumentSetSettingsForDocumentGlobalActionAuth = "TWO_FACTOR_AUTH"
)

func (e DocumentSetSettingsForDocumentGlobalActionAuth) ToPointer() *DocumentSetSettingsForDocumentGlobalActionAuth {
	return &e
}
func (e *DocumentSetSettingsForDocumentGlobalActionAuth) UnmarshalJSON(data []byte) error {
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
		*e = DocumentSetSettingsForDocumentGlobalActionAuth(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DocumentSetSettingsForDocumentGlobalActionAuth: %v", v)
	}
}

type Data struct {
	// The title of the document.
	Title *string `json:"title,omitempty"`
	// The external ID of the document.
	ExternalID *string `json:"externalId,omitempty"`
	// The visibility of the document.
	Visibility *DocumentSetSettingsForDocumentVisibility `json:"visibility,omitempty"`
	// The type of authentication required for the recipient to access the document.
	GlobalAccessAuth *DocumentSetSettingsForDocumentGlobalAccessAuth `json:"globalAccessAuth,omitempty"`
	// The type of authentication required for the recipient to sign the document. This field is restricted to Enterprise plan users only.
	GlobalActionAuth *DocumentSetSettingsForDocumentGlobalActionAuth `json:"globalActionAuth,omitempty"`
}

func (o *Data) GetTitle() *string {
	if o == nil {
		return nil
	}
	return o.Title
}

func (o *Data) GetExternalID() *string {
	if o == nil {
		return nil
	}
	return o.ExternalID
}

func (o *Data) GetVisibility() *DocumentSetSettingsForDocumentVisibility {
	if o == nil {
		return nil
	}
	return o.Visibility
}

func (o *Data) GetGlobalAccessAuth() *DocumentSetSettingsForDocumentGlobalAccessAuth {
	if o == nil {
		return nil
	}
	return o.GlobalAccessAuth
}

func (o *Data) GetGlobalActionAuth() *DocumentSetSettingsForDocumentGlobalActionAuth {
	if o == nil {
		return nil
	}
	return o.GlobalActionAuth
}

// DocumentSetSettingsForDocumentDateFormat - The date format to use for date fields and signing the document.
type DocumentSetSettingsForDocumentDateFormat string

const (
	DocumentSetSettingsForDocumentDateFormatYyyyMmDdHhMmA         DocumentSetSettingsForDocumentDateFormat = "yyyy-MM-dd hh:mm a"
	DocumentSetSettingsForDocumentDateFormatYyyyMmDd              DocumentSetSettingsForDocumentDateFormat = "yyyy-MM-dd"
	DocumentSetSettingsForDocumentDateFormatDdMmYyyyHhMmA         DocumentSetSettingsForDocumentDateFormat = "dd/MM/yyyy hh:mm a"
	DocumentSetSettingsForDocumentDateFormatMmDdYyyyHhMmA         DocumentSetSettingsForDocumentDateFormat = "MM/dd/yyyy hh:mm a"
	DocumentSetSettingsForDocumentDateFormatYyyyMmDdHhMm          DocumentSetSettingsForDocumentDateFormat = "yyyy-MM-dd HH:mm"
	DocumentSetSettingsForDocumentDateFormatYyMmDdHhMmA           DocumentSetSettingsForDocumentDateFormat = "yy-MM-dd hh:mm a"
	DocumentSetSettingsForDocumentDateFormatYyyyMmDdHhMmSs        DocumentSetSettingsForDocumentDateFormat = "yyyy-MM-dd HH:mm:ss"
	DocumentSetSettingsForDocumentDateFormatMmmmDdYyyyHhMmA       DocumentSetSettingsForDocumentDateFormat = "MMMM dd, yyyy hh:mm a"
	DocumentSetSettingsForDocumentDateFormatEeeeMmmmDdYyyyHhMmA   DocumentSetSettingsForDocumentDateFormat = "EEEE, MMMM dd, yyyy hh:mm a"
	DocumentSetSettingsForDocumentDateFormatYyyyMmDdTHhMmSsSssxxx DocumentSetSettingsForDocumentDateFormat = "yyyy-MM-dd'T'HH:mm:ss.SSSXXX"
)

func (e DocumentSetSettingsForDocumentDateFormat) ToPointer() *DocumentSetSettingsForDocumentDateFormat {
	return &e
}
func (e *DocumentSetSettingsForDocumentDateFormat) UnmarshalJSON(data []byte) error {
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
		*e = DocumentSetSettingsForDocumentDateFormat(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DocumentSetSettingsForDocumentDateFormat: %v", v)
	}
}

// DocumentSetSettingsForDocumentDistributionMethod - The distribution method to use when sending the document to the recipients.
type DocumentSetSettingsForDocumentDistributionMethod string

const (
	DocumentSetSettingsForDocumentDistributionMethodEmail DocumentSetSettingsForDocumentDistributionMethod = "EMAIL"
	DocumentSetSettingsForDocumentDistributionMethodNone  DocumentSetSettingsForDocumentDistributionMethod = "NONE"
)

func (e DocumentSetSettingsForDocumentDistributionMethod) ToPointer() *DocumentSetSettingsForDocumentDistributionMethod {
	return &e
}
func (e *DocumentSetSettingsForDocumentDistributionMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "EMAIL":
		fallthrough
	case "NONE":
		*e = DocumentSetSettingsForDocumentDistributionMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DocumentSetSettingsForDocumentDistributionMethod: %v", v)
	}
}

type DocumentSetSettingsForDocumentSigningOrder string

const (
	DocumentSetSettingsForDocumentSigningOrderParallel   DocumentSetSettingsForDocumentSigningOrder = "PARALLEL"
	DocumentSetSettingsForDocumentSigningOrderSequential DocumentSetSettingsForDocumentSigningOrder = "SEQUENTIAL"
)

func (e DocumentSetSettingsForDocumentSigningOrder) ToPointer() *DocumentSetSettingsForDocumentSigningOrder {
	return &e
}
func (e *DocumentSetSettingsForDocumentSigningOrder) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "PARALLEL":
		fallthrough
	case "SEQUENTIAL":
		*e = DocumentSetSettingsForDocumentSigningOrder(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DocumentSetSettingsForDocumentSigningOrder: %v", v)
	}
}

// DocumentSetSettingsForDocumentLanguage - The language to use for email communications with recipients.
type DocumentSetSettingsForDocumentLanguage string

const (
	DocumentSetSettingsForDocumentLanguageDe DocumentSetSettingsForDocumentLanguage = "de"
	DocumentSetSettingsForDocumentLanguageEn DocumentSetSettingsForDocumentLanguage = "en"
	DocumentSetSettingsForDocumentLanguageFr DocumentSetSettingsForDocumentLanguage = "fr"
	DocumentSetSettingsForDocumentLanguageEs DocumentSetSettingsForDocumentLanguage = "es"
)

func (e DocumentSetSettingsForDocumentLanguage) ToPointer() *DocumentSetSettingsForDocumentLanguage {
	return &e
}
func (e *DocumentSetSettingsForDocumentLanguage) UnmarshalJSON(data []byte) error {
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
		*e = DocumentSetSettingsForDocumentLanguage(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DocumentSetSettingsForDocumentLanguage: %v", v)
	}
}

type DocumentSetSettingsForDocumentEmailSettings struct {
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

func (d DocumentSetSettingsForDocumentEmailSettings) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DocumentSetSettingsForDocumentEmailSettings) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *DocumentSetSettingsForDocumentEmailSettings) GetRecipientSigningRequest() *bool {
	if o == nil {
		return nil
	}
	return o.RecipientSigningRequest
}

func (o *DocumentSetSettingsForDocumentEmailSettings) GetRecipientRemoved() *bool {
	if o == nil {
		return nil
	}
	return o.RecipientRemoved
}

func (o *DocumentSetSettingsForDocumentEmailSettings) GetRecipientSigned() *bool {
	if o == nil {
		return nil
	}
	return o.RecipientSigned
}

func (o *DocumentSetSettingsForDocumentEmailSettings) GetDocumentPending() *bool {
	if o == nil {
		return nil
	}
	return o.DocumentPending
}

func (o *DocumentSetSettingsForDocumentEmailSettings) GetDocumentCompleted() *bool {
	if o == nil {
		return nil
	}
	return o.DocumentCompleted
}

func (o *DocumentSetSettingsForDocumentEmailSettings) GetDocumentDeleted() *bool {
	if o == nil {
		return nil
	}
	return o.DocumentDeleted
}

func (o *DocumentSetSettingsForDocumentEmailSettings) GetOwnerDocumentCompleted() *bool {
	if o == nil {
		return nil
	}
	return o.OwnerDocumentCompleted
}

type DocumentSetSettingsForDocumentMeta struct {
	// The subject of the email that will be sent to the recipients.
	Subject *string `json:"subject,omitempty"`
	// The message of the email that will be sent to the recipients.
	Message *string `json:"message,omitempty"`
	// The timezone to use for date fields and signing the document. Example Etc/UTC, Australia/Melbourne
	Timezone *string `json:"timezone,omitempty"`
	// The date format to use for date fields and signing the document.
	DateFormat *DocumentSetSettingsForDocumentDateFormat `json:"dateFormat,omitempty"`
	// The distribution method to use when sending the document to the recipients.
	DistributionMethod *DocumentSetSettingsForDocumentDistributionMethod `json:"distributionMethod,omitempty"`
	SigningOrder       *DocumentSetSettingsForDocumentSigningOrder       `json:"signingOrder,omitempty"`
	// The URL to which the recipient should be redirected after signing the document.
	RedirectURL *string `json:"redirectUrl,omitempty"`
	// The language to use for email communications with recipients.
	Language *DocumentSetSettingsForDocumentLanguage `json:"language,omitempty"`
	// Whether to allow recipients to sign using a typed signature.
	TypedSignatureEnabled *bool                                        `json:"typedSignatureEnabled,omitempty"`
	EmailSettings         *DocumentSetSettingsForDocumentEmailSettings `json:"emailSettings,omitempty"`
}

func (o *DocumentSetSettingsForDocumentMeta) GetSubject() *string {
	if o == nil {
		return nil
	}
	return o.Subject
}

func (o *DocumentSetSettingsForDocumentMeta) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *DocumentSetSettingsForDocumentMeta) GetTimezone() *string {
	if o == nil {
		return nil
	}
	return o.Timezone
}

func (o *DocumentSetSettingsForDocumentMeta) GetDateFormat() *DocumentSetSettingsForDocumentDateFormat {
	if o == nil {
		return nil
	}
	return o.DateFormat
}

func (o *DocumentSetSettingsForDocumentMeta) GetDistributionMethod() *DocumentSetSettingsForDocumentDistributionMethod {
	if o == nil {
		return nil
	}
	return o.DistributionMethod
}

func (o *DocumentSetSettingsForDocumentMeta) GetSigningOrder() *DocumentSetSettingsForDocumentSigningOrder {
	if o == nil {
		return nil
	}
	return o.SigningOrder
}

func (o *DocumentSetSettingsForDocumentMeta) GetRedirectURL() *string {
	if o == nil {
		return nil
	}
	return o.RedirectURL
}

func (o *DocumentSetSettingsForDocumentMeta) GetLanguage() *DocumentSetSettingsForDocumentLanguage {
	if o == nil {
		return nil
	}
	return o.Language
}

func (o *DocumentSetSettingsForDocumentMeta) GetTypedSignatureEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.TypedSignatureEnabled
}

func (o *DocumentSetSettingsForDocumentMeta) GetEmailSettings() *DocumentSetSettingsForDocumentEmailSettings {
	if o == nil {
		return nil
	}
	return o.EmailSettings
}

type DocumentSetSettingsForDocumentRequestBody struct {
	DocumentID float64                             `json:"documentId"`
	Data       *Data                               `json:"data,omitempty"`
	Meta       *DocumentSetSettingsForDocumentMeta `json:"meta,omitempty"`
}

func (o *DocumentSetSettingsForDocumentRequestBody) GetDocumentID() float64 {
	if o == nil {
		return 0.0
	}
	return o.DocumentID
}

func (o *DocumentSetSettingsForDocumentRequestBody) GetData() *Data {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *DocumentSetSettingsForDocumentRequestBody) GetMeta() *DocumentSetSettingsForDocumentMeta {
	if o == nil {
		return nil
	}
	return o.Meta
}

type DocumentSetSettingsForDocumentDocumentsVisibility string

const (
	DocumentSetSettingsForDocumentDocumentsVisibilityEveryone        DocumentSetSettingsForDocumentDocumentsVisibility = "EVERYONE"
	DocumentSetSettingsForDocumentDocumentsVisibilityManagerAndAbove DocumentSetSettingsForDocumentDocumentsVisibility = "MANAGER_AND_ABOVE"
	DocumentSetSettingsForDocumentDocumentsVisibilityAdmin           DocumentSetSettingsForDocumentDocumentsVisibility = "ADMIN"
)

func (e DocumentSetSettingsForDocumentDocumentsVisibility) ToPointer() *DocumentSetSettingsForDocumentDocumentsVisibility {
	return &e
}
func (e *DocumentSetSettingsForDocumentDocumentsVisibility) UnmarshalJSON(data []byte) error {
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
		*e = DocumentSetSettingsForDocumentDocumentsVisibility(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DocumentSetSettingsForDocumentDocumentsVisibility: %v", v)
	}
}

type DocumentSetSettingsForDocumentStatus string

const (
	DocumentSetSettingsForDocumentStatusDraft     DocumentSetSettingsForDocumentStatus = "DRAFT"
	DocumentSetSettingsForDocumentStatusPending   DocumentSetSettingsForDocumentStatus = "PENDING"
	DocumentSetSettingsForDocumentStatusCompleted DocumentSetSettingsForDocumentStatus = "COMPLETED"
)

func (e DocumentSetSettingsForDocumentStatus) ToPointer() *DocumentSetSettingsForDocumentStatus {
	return &e
}
func (e *DocumentSetSettingsForDocumentStatus) UnmarshalJSON(data []byte) error {
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
		*e = DocumentSetSettingsForDocumentStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DocumentSetSettingsForDocumentStatus: %v", v)
	}
}

type DocumentSetSettingsForDocumentSource string

const (
	DocumentSetSettingsForDocumentSourceDocument           DocumentSetSettingsForDocumentSource = "DOCUMENT"
	DocumentSetSettingsForDocumentSourceTemplate           DocumentSetSettingsForDocumentSource = "TEMPLATE"
	DocumentSetSettingsForDocumentSourceTemplateDirectLink DocumentSetSettingsForDocumentSource = "TEMPLATE_DIRECT_LINK"
)

func (e DocumentSetSettingsForDocumentSource) ToPointer() *DocumentSetSettingsForDocumentSource {
	return &e
}
func (e *DocumentSetSettingsForDocumentSource) UnmarshalJSON(data []byte) error {
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
		*e = DocumentSetSettingsForDocumentSource(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DocumentSetSettingsForDocumentSource: %v", v)
	}
}

// DocumentSetSettingsForDocumentDocumentsGlobalAccessAuth - The type of authentication required for the recipient to access the document.
type DocumentSetSettingsForDocumentDocumentsGlobalAccessAuth string

const (
	DocumentSetSettingsForDocumentDocumentsGlobalAccessAuthAccount DocumentSetSettingsForDocumentDocumentsGlobalAccessAuth = "ACCOUNT"
)

func (e DocumentSetSettingsForDocumentDocumentsGlobalAccessAuth) ToPointer() *DocumentSetSettingsForDocumentDocumentsGlobalAccessAuth {
	return &e
}
func (e *DocumentSetSettingsForDocumentDocumentsGlobalAccessAuth) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ACCOUNT":
		*e = DocumentSetSettingsForDocumentDocumentsGlobalAccessAuth(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DocumentSetSettingsForDocumentDocumentsGlobalAccessAuth: %v", v)
	}
}

// DocumentSetSettingsForDocumentDocumentsGlobalActionAuth - The type of authentication required for the recipient to sign the document. This field is restricted to Enterprise plan users only.
type DocumentSetSettingsForDocumentDocumentsGlobalActionAuth string

const (
	DocumentSetSettingsForDocumentDocumentsGlobalActionAuthAccount       DocumentSetSettingsForDocumentDocumentsGlobalActionAuth = "ACCOUNT"
	DocumentSetSettingsForDocumentDocumentsGlobalActionAuthPasskey       DocumentSetSettingsForDocumentDocumentsGlobalActionAuth = "PASSKEY"
	DocumentSetSettingsForDocumentDocumentsGlobalActionAuthTwoFactorAuth DocumentSetSettingsForDocumentDocumentsGlobalActionAuth = "TWO_FACTOR_AUTH"
)

func (e DocumentSetSettingsForDocumentDocumentsGlobalActionAuth) ToPointer() *DocumentSetSettingsForDocumentDocumentsGlobalActionAuth {
	return &e
}
func (e *DocumentSetSettingsForDocumentDocumentsGlobalActionAuth) UnmarshalJSON(data []byte) error {
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
		*e = DocumentSetSettingsForDocumentDocumentsGlobalActionAuth(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DocumentSetSettingsForDocumentDocumentsGlobalActionAuth: %v", v)
	}
}

type DocumentSetSettingsForDocumentAuthOptions struct {
	// The type of authentication required for the recipient to access the document.
	GlobalAccessAuth *DocumentSetSettingsForDocumentDocumentsGlobalAccessAuth `json:"globalAccessAuth"`
	// The type of authentication required for the recipient to sign the document. This field is restricted to Enterprise plan users only.
	GlobalActionAuth *DocumentSetSettingsForDocumentDocumentsGlobalActionAuth `json:"globalActionAuth"`
}

func (o *DocumentSetSettingsForDocumentAuthOptions) GetGlobalAccessAuth() *DocumentSetSettingsForDocumentDocumentsGlobalAccessAuth {
	if o == nil {
		return nil
	}
	return o.GlobalAccessAuth
}

func (o *DocumentSetSettingsForDocumentAuthOptions) GetGlobalActionAuth() *DocumentSetSettingsForDocumentDocumentsGlobalActionAuth {
	if o == nil {
		return nil
	}
	return o.GlobalActionAuth
}

type DocumentSetSettingsForDocumentFormValuesType string

const (
	DocumentSetSettingsForDocumentFormValuesTypeStr     DocumentSetSettingsForDocumentFormValuesType = "str"
	DocumentSetSettingsForDocumentFormValuesTypeBoolean DocumentSetSettingsForDocumentFormValuesType = "boolean"
	DocumentSetSettingsForDocumentFormValuesTypeNumber  DocumentSetSettingsForDocumentFormValuesType = "number"
)

type DocumentSetSettingsForDocumentFormValues struct {
	Str     *string  `queryParam:"inline"`
	Boolean *bool    `queryParam:"inline"`
	Number  *float64 `queryParam:"inline"`

	Type DocumentSetSettingsForDocumentFormValuesType
}

func CreateDocumentSetSettingsForDocumentFormValuesStr(str string) DocumentSetSettingsForDocumentFormValues {
	typ := DocumentSetSettingsForDocumentFormValuesTypeStr

	return DocumentSetSettingsForDocumentFormValues{
		Str:  &str,
		Type: typ,
	}
}

func CreateDocumentSetSettingsForDocumentFormValuesBoolean(boolean bool) DocumentSetSettingsForDocumentFormValues {
	typ := DocumentSetSettingsForDocumentFormValuesTypeBoolean

	return DocumentSetSettingsForDocumentFormValues{
		Boolean: &boolean,
		Type:    typ,
	}
}

func CreateDocumentSetSettingsForDocumentFormValuesNumber(number float64) DocumentSetSettingsForDocumentFormValues {
	typ := DocumentSetSettingsForDocumentFormValuesTypeNumber

	return DocumentSetSettingsForDocumentFormValues{
		Number: &number,
		Type:   typ,
	}
}

func (u *DocumentSetSettingsForDocumentFormValues) UnmarshalJSON(data []byte) error {

	var str string = ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = DocumentSetSettingsForDocumentFormValuesTypeStr
		return nil
	}

	var boolean bool = false
	if err := utils.UnmarshalJSON(data, &boolean, "", true, true); err == nil {
		u.Boolean = &boolean
		u.Type = DocumentSetSettingsForDocumentFormValuesTypeBoolean
		return nil
	}

	var number float64 = float64(0)
	if err := utils.UnmarshalJSON(data, &number, "", true, true); err == nil {
		u.Number = &number
		u.Type = DocumentSetSettingsForDocumentFormValuesTypeNumber
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for DocumentSetSettingsForDocumentFormValues", string(data))
}

func (u DocumentSetSettingsForDocumentFormValues) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.Boolean != nil {
		return utils.MarshalJSON(u.Boolean, "", true)
	}

	if u.Number != nil {
		return utils.MarshalJSON(u.Number, "", true)
	}

	return nil, errors.New("could not marshal union type DocumentSetSettingsForDocumentFormValues: all fields are null")
}

// DocumentSetSettingsForDocumentResponseBody - Successful response
type DocumentSetSettingsForDocumentResponseBody struct {
	Visibility DocumentSetSettingsForDocumentDocumentsVisibility `json:"visibility"`
	Status     DocumentSetSettingsForDocumentStatus              `json:"status"`
	Source     DocumentSetSettingsForDocumentSource              `json:"source"`
	ID         int64                                             `json:"id"`
	// A custom external ID you can use to identify the document.
	ExternalID *string `json:"externalId"`
	// The ID of the user that created this document.
	UserID         float64                                             `json:"userId"`
	AuthOptions    *DocumentSetSettingsForDocumentAuthOptions          `json:"authOptions"`
	FormValues     map[string]DocumentSetSettingsForDocumentFormValues `json:"formValues"`
	Title          string                                              `json:"title"`
	DocumentDataID string                                              `json:"documentDataId"`
	CreatedAt      string                                              `json:"createdAt"`
	UpdatedAt      string                                              `json:"updatedAt"`
	CompletedAt    *string                                             `json:"completedAt"`
	DeletedAt      *string                                             `json:"deletedAt"`
	TeamID         *int64                                              `json:"teamId"`
	TemplateID     *int64                                              `json:"templateId"`
}

func (o *DocumentSetSettingsForDocumentResponseBody) GetVisibility() DocumentSetSettingsForDocumentDocumentsVisibility {
	if o == nil {
		return DocumentSetSettingsForDocumentDocumentsVisibility("")
	}
	return o.Visibility
}

func (o *DocumentSetSettingsForDocumentResponseBody) GetStatus() DocumentSetSettingsForDocumentStatus {
	if o == nil {
		return DocumentSetSettingsForDocumentStatus("")
	}
	return o.Status
}

func (o *DocumentSetSettingsForDocumentResponseBody) GetSource() DocumentSetSettingsForDocumentSource {
	if o == nil {
		return DocumentSetSettingsForDocumentSource("")
	}
	return o.Source
}

func (o *DocumentSetSettingsForDocumentResponseBody) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

func (o *DocumentSetSettingsForDocumentResponseBody) GetExternalID() *string {
	if o == nil {
		return nil
	}
	return o.ExternalID
}

func (o *DocumentSetSettingsForDocumentResponseBody) GetUserID() float64 {
	if o == nil {
		return 0.0
	}
	return o.UserID
}

func (o *DocumentSetSettingsForDocumentResponseBody) GetAuthOptions() *DocumentSetSettingsForDocumentAuthOptions {
	if o == nil {
		return nil
	}
	return o.AuthOptions
}

func (o *DocumentSetSettingsForDocumentResponseBody) GetFormValues() map[string]DocumentSetSettingsForDocumentFormValues {
	if o == nil {
		return nil
	}
	return o.FormValues
}

func (o *DocumentSetSettingsForDocumentResponseBody) GetTitle() string {
	if o == nil {
		return ""
	}
	return o.Title
}

func (o *DocumentSetSettingsForDocumentResponseBody) GetDocumentDataID() string {
	if o == nil {
		return ""
	}
	return o.DocumentDataID
}

func (o *DocumentSetSettingsForDocumentResponseBody) GetCreatedAt() string {
	if o == nil {
		return ""
	}
	return o.CreatedAt
}

func (o *DocumentSetSettingsForDocumentResponseBody) GetUpdatedAt() string {
	if o == nil {
		return ""
	}
	return o.UpdatedAt
}

func (o *DocumentSetSettingsForDocumentResponseBody) GetCompletedAt() *string {
	if o == nil {
		return nil
	}
	return o.CompletedAt
}

func (o *DocumentSetSettingsForDocumentResponseBody) GetDeletedAt() *string {
	if o == nil {
		return nil
	}
	return o.DeletedAt
}

func (o *DocumentSetSettingsForDocumentResponseBody) GetTeamID() *int64 {
	if o == nil {
		return nil
	}
	return o.TeamID
}

func (o *DocumentSetSettingsForDocumentResponseBody) GetTemplateID() *int64 {
	if o == nil {
		return nil
	}
	return o.TemplateID
}

type DocumentSetSettingsForDocumentResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Successful response
	Object *DocumentSetSettingsForDocumentResponseBody
}

func (o *DocumentSetSettingsForDocumentResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *DocumentSetSettingsForDocumentResponse) GetObject() *DocumentSetSettingsForDocumentResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
