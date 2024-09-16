package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aacebo/agent.net/core/logger"
	"github.com/aacebo/agent.net/core/models"
	"github.com/aacebo/agent.net/core/repos"
	"github.com/aacebo/agent.net/ws"

	"github.com/go-chi/render"
	"github.com/gorilla/websocket"
)

func Socket(ctx context.Context) http.HandlerFunc {
	onStatResponse := onStatResponse(ctx)
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

		socket.Send(ws.NewQueryMessage("stat"))

		for {
			message, err := socket.Read()

			if err != nil {
				return
			}

			fmt.Println(message)

			switch message.Type {
			case ws.QUERY_RESPONSE_MESSAGE_TYPE:
				body := message.Body.(ws.QueryResponseMessageBody)

				switch body.Name {
				case "stat":
					onStatResponse(body.Body, socket)
				}
			}
		}
	}
}

func onStatResponse(ctx context.Context) func(any, *ws.Socket) {
	log := logger.New("agent.net/api/sockets")
	agents := ctx.Value("repos.agents").(repos.IAgentsRepository)

	return func(data any, socket *ws.Socket) {
		b, err := json.Marshal(data)

		if err != nil {
			log.Error(err.Error())
			return
		}

		stat := models.Stat{}

		if err = json.Unmarshal(b, &stat); err != nil {
			log.Error(err.Error())
			return
		}

		log.Debug(stat.String())

		ipAddress := socket.IPAddress()
		agent, exists := agents.GetByID(stat.ID)

		if !exists {
			return
		}

		agent.Status = models.AGENT_STATUS_UP
		agent.URL = &ipAddress
		agent = agents.Update(agent)
	}
}
