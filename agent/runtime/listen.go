package runtime

import (
	"fmt"
	"net/http"
)

func (self *Agent) Listen() {
	header := http.Header{}
	header.Add("X_AGENT_ID", self.ID)
	header.Add("X_AGENT_NAME", self.Name)
	header.Add("X_AGENT_DESCRIPTION", self.Description)
	header.Add("X_CLIENT_ID", self.ClientID)
	header.Add("X_CLIENT_SECRET", self.ClientSecret)

	if err := self.parent.Connect(fmt.Sprintf("wss://%s/v1/sockets", self.Address), header); err != nil {
		panic(err)
	}

	defer self.parent.Close()
	self.log.Info("connected...", nil)

	for {
		message, err := self.parent.Read()

		if err != nil || !message.Type.Valid() {
			self.log.Warn(err.Error(), nil)
			return
		}
	}
}
