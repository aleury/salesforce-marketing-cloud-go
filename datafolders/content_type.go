package datafolders

type ContentType string

const (
	DataExtension      ContentType = "dataextension"
	UserInitiatedSends ContentType = "userinitiatedsends"
)

func (c ContentType) String() string {
	return string(c)
}
