package ws

type MessageType string

const (
	MESSAGE_EVENT_TYPE MessageType = "message" // send a message
	COMMAND_EVENT_TYPE MessageType = "command" // execute a command (write) operation
	QUERY_EVENT_TYPE   MessageType = "query"   // execute a query (read) operation
)

func (self MessageType) Valid() bool {
	switch self {
	case MESSAGE_EVENT_TYPE:
		fallthrough
	case COMMAND_EVENT_TYPE:
		fallthrough
	case QUERY_EVENT_TYPE:
		return true
	}

	return false
}
