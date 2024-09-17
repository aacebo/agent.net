package routes

import (
	"context"
	"net/http"

	"github.com/aacebo/agent.net/agent/runtime"
	"github.com/aacebo/agent.net/ws"

	"github.com/go-chi/render"
	"github.com/gorilla/websocket"
)

func Socket(ctx context.Context) http.HandlerFunc {
	runtime := ctx.Value("runtime").(*runtime.Agent)
	log := runtime.Logger()
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	return func(w http.ResponseWriter, r *http.Request) {
		agentId := r.Header.Get("X_AGENT_ID")
		agentAddress := r.Header.Get("X_AGENT_ADDRESS")
		conn, err := upgrader.Upgrade(w, r, nil)

		if err != nil {
			render.Status(r, 500)
			render.JSON(w, r, err.Error())
			log.Error(err.Error(), nil)
			return
		}

		socket := runtime.Add(r, conn)

		defer func() {
			runtime.Remove(socket.ID)
			runtime.SendToParent(ws.NewDisconnectedMessage(agentId))
		}()

		runtime.SendToParent(ws.NewConnectedMessage(ws.ConnectedMessageBody{
			"id":      agentId,
			"address": agentAddress,
		}))

		for {
			message, err := socket.Read()

			if err != nil || !message.Type.Valid() {
				return
			}

			switch message.Type {
			case ws.CONNECTED_MESSAGE_TYPE:
				runtime.SendToParent(message)
			case ws.DISCONNECTED_MESSAGE_TYPE:
				runtime.SendToParent(message)
			}
		}
	}
}
