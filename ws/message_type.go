package ws

type MessageType string

const (
	TEXT_MESSAGE_TYPE         MessageType = "text"
	CONNECTED_MESSAGE_TYPE    MessageType = "connected"
	DISCONNECTED_MESSAGE_TYPE MessageType = "disconnected"
)

func (self MessageType) Valid() bool {
	switch self {
	case TEXT_MESSAGE_TYPE:
		fallthrough
	case CONNECTED_MESSAGE_TYPE:
		fallthrough
	case DISCONNECTED_MESSAGE_TYPE:
		return true
	}

	return false
}
