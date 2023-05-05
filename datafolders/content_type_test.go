package datafolders_test

import (
	"salesforce-marketing-cloud-go/datafolders"
	"testing"
)

func TestDataExtensionContentType(t *testing.T) {
	expected := "dataextension"
	actual := datafolders.DataExtension.String()
	if expected != actual {
		t.Fatalf("expected: %s, got: %s", expected, actual)
	}
}

func TestUserInitiatedSendContentType(t *testing.T) {
	expected := "userinitiatedsends"
	actual := datafolders.UserInitiatedSends.String()
	if expected != actual {
		t.Fatalf("expected: %s, got: %s", expected, actual)
	}
}
