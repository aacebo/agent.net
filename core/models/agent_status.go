package models

type AgentStatus string

const (
	AGENT_STATUS_UP   AgentStatus = "up"
	AGENT_STATUS_DOWN AgentStatus = "down"
)

func (self AgentStatus) Valid() bool {
	switch self {
	case AGENT_STATUS_UP:
		fallthrough
	case AGENT_STATUS_DOWN:
		return true
	}

	return false
}
