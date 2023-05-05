package datafolders

import (
	"context"
	"fmt"
	"salesforce-marketing-cloud-go/gen"
	"salesforce-marketing-cloud-go/soap"
	"salesforce-marketing-cloud-go/soap/filters"
)

func GetById(ctx context.Context, client *soap.Client, folderId int32) (*gen.DataFolder, error) {
	properties := []string{"ID", "Name", "ContentType"}
	filter := filters.Equals("ID", string(folderId))

	results, err := soap.Retrieve[gen.DataFolder](ctx, client, "DataFolder", properties, filter)
	if err != nil {
		return nil, err
	}

	if len(results) == 0 {
		return nil, fmt.Errorf("data folder with id '%d' not found", folderId)
	}

	if len(results) > 1 {
		return nil, fmt.Errorf("more than one result was returned for id '%d'", folderId)
	}

	return results[0], nil
}
