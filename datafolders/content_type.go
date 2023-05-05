package datafolders

type ContentType string

const (
	DataExtension     ContentType = "dataextension"
	UserInitiatedSend ContentType = "userinitiatedsends"
)

func (c ContentType) String() string {
	return string(c)
}
