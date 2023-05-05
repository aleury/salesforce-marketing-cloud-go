package soap

type RequestType string

const (
	RequestTypeSynchronous RequestType = "Synchronous"

	RequestTypeAsynchronous RequestType = "Asynchronous"
)

type Request struct{}
