package soap

type Result struct {
	StatusCode string `xml:"StatusCode,omitempty" json:"StatusCode,omitempty"`

	StatusMessage string `xml:"StatusMessage,omitempty" json:"StatusMessage,omitempty"`

	OrdinalID int32 `xml:"OrdinalID,omitempty" json:"OrdinalID,omitempty"`

	ErrorCode int32 `xml:"ErrorCode,omitempty" json:"ErrorCode,omitempty"`

	RequestID string `xml:"RequestID,omitempty" json:"RequestID,omitempty"`

	ConversationID string `xml:"ConversationID,omitempty" json:"ConversationID,omitempty"`

	OverallStatusCode string `xml:"OverallStatusCode,omitempty" json:"OverallStatusCode,omitempty"`

	RequestType *RequestType `xml:"RequestType,omitempty" json:"RequestType,omitempty"`

	ResultType string `xml:"ResultType,omitempty" json:"ResultType,omitempty"`

	ResultDetailXML string `xml:"ResultDetailXML,omitempty" json:"ResultDetailXML,omitempty"`
}

func (r *Result) Success() bool {
	return r.StatusCode == "OK"
}
