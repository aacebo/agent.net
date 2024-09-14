package models

import (
	"encoding/gob"
)

func Register() {
	gob.Register(Agent{})
	gob.Register(AgentSettings{})
	gob.Register(Map[any]{})
}
