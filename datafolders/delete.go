package datafolders

import (
	"context"
	"salesforce-marketing-cloud-go/gen"
	"salesforce-marketing-cloud-go/soap"
)

func Delete(ctx context.Context, client *soap.Client, folderId int32) (bool, error) {
	df := new(gen.DataFolder)
	df.ID = folderId

	ok, err := soap.Delete(ctx, client, df)
	if err != nil {
		return false, err
	}

	return ok, nil
}
