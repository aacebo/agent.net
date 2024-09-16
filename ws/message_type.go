package ws

type MessageType string

const (
	TEXT_MESSAGE_TYPE           MessageType = "text"           // send a message
	QUERY_MESSAGE_TYPE          MessageType = "query"          // execute a query (read) operation
	QUERY_RESPONSE_MESSAGE_TYPE MessageType = "query-response" // a query response operation
)

func (self MessageType) Valid() bool {
	switch self {
	case TEXT_MESSAGE_TYPE:
		fallthrough
	case QUERY_MESSAGE_TYPE:
		fallthrough
	case QUERY_RESPONSE_MESSAGE_TYPE:
		return true
	}

	return false
}
