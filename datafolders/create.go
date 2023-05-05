package datafolders

import (
	"context"
	"salesforce-marketing-cloud-go/gen"
	"salesforce-marketing-cloud-go/soap"
)

type CreateParams struct {
	Name         string
	Description  string
	ContentType  ContentType
	ParentFolder *gen.DataFolder
}

func Create(ctx context.Context, client *soap.Client, params CreateParams) (int32, error) {
	df := &gen.DataFolder{
		Name:          params.Name,
		Description:   params.Description,
		ContentType:   params.ContentType.String(),
		ParentFolder:  params.ParentFolder,
		AllowChildren: true,
	}
	return soap.Create(ctx, client, df)
}
