package soap

import (
	"context"
	"encoding/xml"
)

func Delete[T any](ctx context.Context, client *Client, obj *T) (bool, error) {
	request := &deleteRequest[T]{
		Objects: []*T{obj},
	}
	response := new(deleteResponse[T])
	err := client.CallContext(ctx, "Delete", request, response)
	if err != nil {
		return false, err
	}

	// TODO(adam): determine how to better bubble up unsuccessful results
	for _, result := range response.Results {
		if !result.Success() {
			return false, nil
		}
	}

	return true, nil
}

type deleteRequest[T any] struct {
	XMLName xml.Name `xml:"http://exacttarget.com/wsdl/partnerAPI DeleteRequest"`

	Options *deleteOptions `xml:"Options,omitempty" json:"Options,omitempty"`

	Objects []*T `xml:"Objects,omitempty" json:"Objects,omitempty"`
}

type deleteOptions struct {
	*Options
}

type deleteResponse[T any] struct {
	XMLName xml.Name `xml:"http://exacttarget.com/wsdl/partnerAPI DeleteResponse"`

	Results []*deleteResult[T] `xml:"Results,omitempty" json:"Results,omitempty"`

	RequestID string `xml:"RequestID,omitempty" json:"RequestID,omitempty"`

	OverallStatus string `xml:"OverallStatus,omitempty" json:"OverallStatus,omitempty"`
}

type deleteResult[T any] struct {
	*Result

	Object *T `xml:"Object,omitempty" json:"Object,omitempty"`
}
