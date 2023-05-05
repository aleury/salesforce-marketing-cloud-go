package soap

import (
	"context"
	"encoding/xml"
)

func Retrieve[T, F any](ctx context.Context, client *Client, objectType string, properties []string, filter *F) ([]*T, error) {
	request := &retrieveRequestMsg[F]{
		RetrieveRequest: &retrieveRequest[F]{
			ObjectType: objectType,
			Properties: properties,
			Filter:     filter,
		},
	}
	response := new(retrieveResponseMsg[T])
	err := client.CallContext(ctx, "Retrieve", request, response)
	if err != nil {
		return nil, err
	}
	return response.Results, nil
}

type retrieveRequestMsg[F any] struct {
	XMLName xml.Name `xml:"http://exacttarget.com/wsdl/partnerAPI RetrieveRequestMsg"`

	RetrieveRequest *retrieveRequest[F] `xml:"RetrieveRequest,omitempty" json:"RetrieveRequest,omitempty"`
}

type retrieveResponseMsg[T any] struct {
	XMLName xml.Name `xml:"http://exacttarget.com/wsdl/partnerAPI RetrieveResponseMsg"`

	OverallStatus string `xml:"OverallStatus,omitempty" json:"OverallStatus,omitempty"`

	RequestID string `xml:"RequestID,omitempty" json:"RequestID,omitempty"`

	Results []*T `xml:"Results,omitempty" json:"Results,omitempty"`
}

type retrieveRequest[F any] struct {
	ClientIDs []*ClientID `xml:"ClientIDs,omitempty" json:"ClientIDs,omitempty"`

	ObjectType string `xml:"ObjectType,omitempty" json:"ObjectType,omitempty"`

	Properties []string `xml:"Properties,omitempty" json:"Properties,omitempty"`

	Filter *F `xml:"Filter,omitempty" json:"Filter,omitempty"`

	RespondTo []*AsyncResponse `xml:"RespondTo,omitempty" json:"RespondTo,omitempty"`

	PartnerProperties []*APIProperty `xml:"PartnerProperties,omitempty" json:"PartnerProperties,omitempty"`

	ContinueRequest string `xml:"ContinueRequest,omitempty" json:"ContinueRequest,omitempty"`

	QueryAllAccounts bool `xml:"QueryAllAccounts,omitempty" json:"QueryAllAccounts,omitempty"`

	RetrieveAllSinceLastBatch bool `xml:"RetrieveAllSinceLastBatch,omitempty" json:"RetrieveAllSinceLastBatch,omitempty"`

	RepeatLastResult bool `xml:"RepeatLastResult,omitempty" json:"RepeatLastResult,omitempty"`

	Retrieves struct {
		Request []*Request `xml:"Request,omitempty" json:"Request,omitempty"`
	} `xml:"Retrieves,omitempty" json:"Retrieves,omitempty"`

	Options *retrieveOptions `xml:"Options,omitempty" json:"Options,omitempty"`
}

type retrieveOptions struct {
	*Options

	BatchSize int32 `xml:"BatchSize,omitempty" json:"BatchSize,omitempty"`

	IncludeObjects bool `xml:"IncludeObjects,omitempty" json:"IncludeObjects,omitempty"`

	OnlyIncludeBase bool `xml:"OnlyIncludeBase,omitempty" json:"OnlyIncludeBase,omitempty"`
}
