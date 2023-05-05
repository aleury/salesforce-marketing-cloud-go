package datafolders_test

import (
	"salesforce-marketing-cloud-go/datafolders"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContentTypesAreWellDefined(t *testing.T) {
	type test struct {
		input datafolders.ContentType
		want  string
	}

	tests := []test{
		{datafolders.DataExtension, "dataextension"},
		{datafolders.UserInitiatedSends, "userinitiatedsends"},
	}
	for _, tc := range tests {
		assert.Equal(t, tc.want, tc.input.String())
	}
}
