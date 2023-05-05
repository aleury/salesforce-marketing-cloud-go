package soap

import (
	"github.com/hooklift/gowsdl/soap"
)

type Client = soap.Client

func NewClient(url string, opt ...soap.Option) *Client {
	return soap.NewClient(url, opt...)
}
