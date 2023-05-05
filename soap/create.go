package soap

import (
	"context"
	"encoding/xml"
	"salesforce-marketing-cloud-go/gen"
)

func Create[T any](ctx context.Context, client *Client, obj *T) (int32, error) {
	request := &createRequest[T]{
		Objects: []*T{obj},
	}
	response := new(gen.CreateResponse)
	err := client.CallContext(ctx, "Create", request, response)
	if err != nil {
		return 0, err
	}
	return response.Results[0].NewID, nil
}

type createRequest[T any] struct {
	XMLName xml.Name `xml:"http://exacttarget.com/wsdl/partnerAPI CreateRequest"`

	Options *gen.CreateOptions `xml:"Options,omitempty" json:"Options,omitempty"`

	Objects []*T `xml:"Objects,omitempty" json:"Objects,omitempty"`
}
