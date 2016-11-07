package connectors

type Connector interface {
	send(request RequestMessage)
	receive() ResponseMessage
}

type RequestMessage string

type ResponseMessage string
