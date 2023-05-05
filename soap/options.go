package soap

import "github.com/hooklift/gowsdl/soap"

type SaveAction string

const (
	SaveActionAddOnly SaveAction = "AddOnly"

	SaveActionDefault SaveAction = "Default"

	SaveActionNothing SaveAction = "Nothing"

	SaveActionUpdateAdd SaveAction = "UpdateAdd"

	SaveActionUpdateOnly SaveAction = "UpdateOnly"

	SaveActionDelete SaveAction = "Delete"
)

type SaveOption struct {
	PropertyName string `xml:"PropertyName,omitempty" json:"PropertyName,omitempty"`

	SaveAction *SaveAction `xml:"SaveAction,omitempty" json:"SaveAction,omitempty"`

	TrackChanges bool `xml:"TrackChanges,omitempty" json:"TrackChanges,omitempty"`
}

type Options struct {
	Client *ClientID `xml:"Client,omitempty" json:"Client,omitempty"`

	SendResponseTo []*AsyncResponse `xml:"SendResponseTo,omitempty" json:"SendResponseTo,omitempty"`

	SaveOptions struct {
		SaveOption []*SaveOption `xml:"SaveOption,omitempty" json:"SaveOption,omitempty"`
	} `xml:"SaveOptions,omitempty" json:"SaveOptions,omitempty"`

	Priority int8 `xml:"Priority,omitempty" json:"Priority,omitempty"`

	ConversationID string `xml:"ConversationID,omitempty" json:"ConversationID,omitempty"`

	SequenceCode int32 `xml:"SequenceCode,omitempty" json:"SequenceCode,omitempty"`

	CallsInConversation int32 `xml:"CallsInConversation,omitempty" json:"CallsInConversation,omitempty"`

	ScheduledTime soap.XSDDateTime `xml:"ScheduledTime,omitempty" json:"ScheduledTime,omitempty"`

	RequestType *RequestType `xml:"RequestType,omitempty" json:"RequestType,omitempty"`

	QueuePriority *Priority `xml:"QueuePriority,omitempty" json:"QueuePriority,omitempty"`
}
