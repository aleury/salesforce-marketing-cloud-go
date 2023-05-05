package soap

type RespondWhen string

const (
	RespondWhenNever RespondWhen = "Never"

	RespondWhenOnError RespondWhen = "OnError"

	RespondWhenAlways RespondWhen = "Always"

	RespondWhenOnConversationError RespondWhen = "OnConversationError"

	RespondWhenOnConversationComplete RespondWhen = "OnConversationComplete"

	RespondWhenOnCallComplete RespondWhen = "OnCallComplete"
)

type AsyncResponseType string

const (
	AsyncResponseTypeNone AsyncResponseType = "None"

	AsyncResponseTypeEmail AsyncResponseType = "email"

	AsyncResponseTypeFTP AsyncResponseType = "FTP"

	AsyncResponseTypeHTTPPost AsyncResponseType = "HTTPPost"
)

type AsyncResponse struct {
	ResponseType *AsyncResponseType `xml:"ResponseType,omitempty" json:"ResponseType,omitempty"`

	ResponseAddress string `xml:"ResponseAddress,omitempty" json:"ResponseAddress,omitempty"`

	RespondWhen *RespondWhen `xml:"RespondWhen,omitempty" json:"RespondWhen,omitempty"`

	IncludeResults bool `xml:"IncludeResults,omitempty" json:"IncludeResults,omitempty"`

	IncludeObjects bool `xml:"IncludeObjects,omitempty" json:"IncludeObjects,omitempty"`

	OnlyIncludeBase bool `xml:"OnlyIncludeBase,omitempty" json:"OnlyIncludeBase,omitempty"`
}
