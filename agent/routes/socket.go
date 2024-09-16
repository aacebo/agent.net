package routes

import (
	"context"
	"net/http"

	"github.com/aacebo/agent.net/agent/client"
	"github.com/aacebo/agent.net/core/models"
	"github.com/aacebo/agent.net/ws"

	"github.com/go-chi/render"
	"github.com/gorilla/websocket"
)

func Socket(ctx context.Context) http.HandlerFunc {
	c := ctx.Value("socket").(*client.Client)
	sockets := ctx.Value("sockets").(*ws.Sockets)
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	return func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)

		if err != nil {
			render.Status(r, 500)
			render.JSON(w, r, err.Error())
			return
		}

		socket := sockets.Add(conn)

		defer func() {
			conn.Close()
			sockets.Del(socket.ID)
		}()

		go func() {
			socket.Send(ws.NewQueryMessage("stat"))
		}()

		for {
			message, err := socket.Read()

			if err != nil || !message.Type.Valid() {
				return
			}

			switch message.Type {
			case ws.QUERY_RESPONSE_MESSAGE_TYPE:
				body := message.Body.(ws.QueryResponseMessageBody)

				switch body.Name {
				case "stat":
					stat := body.Body.(models.Stat)
					c.SetAgent(models.AgentStat{
						ID:          stat.ID,
						SocketID:    socket.ID,
						Description: stat.Description,
						IPAddress:   socket.IPAddress(),
						StartedAt:   stat.StartedAt,
					})

					c.Send(ws.NewQueryResponseMessage("stat", c.Stat()))
				}
			}
		}
	}
}
