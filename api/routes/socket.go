package routes

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/aacebo/agent.net/core/logger"
	"github.com/aacebo/agent.net/core/models"
	"github.com/aacebo/agent.net/core/repos"
	"github.com/aacebo/agent.net/ws"

	"github.com/go-chi/render"
	"github.com/gorilla/websocket"
)

func Socket(ctx context.Context) http.HandlerFunc {
	onStatResponse := onStatResponse(ctx)
	agents := ctx.Value("repos.agents").(repos.IAgentsRepository)
	sockets := ctx.Value("sockets").(*ws.Sockets)
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	return func(w http.ResponseWriter, r *http.Request) {
		clientId := r.Header.Get("client_id")
		clientSecret := r.Header.Get("client_secret")

		agent, exists := agents.GetByCredentials(clientId, clientSecret)

		if !exists {
			render.Status(r, http.StatusUnauthorized)
			render.JSON(w, r, "unauthorized")
		}

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
			agent, exists := agents.GetByCredentials(clientId, clientSecret)

			if !exists {
				return
			}

			agent.Status = models.AGENT_STATUS_DOWN
			agent = agents.Update(agent)
		}()

		go func() {
			for range time.Tick(5 * time.Second) {
				if err := socket.Send(ws.NewQueryMessage("stat")); err != nil {
					return
				}
			}
		}()

		for {
			message, err := socket.Read()

			if err != nil {
				return
			}

			switch message.Type {
			case ws.QUERY_RESPONSE_MESSAGE_TYPE:
				body := message.Body.(ws.QueryResponseMessageBody)

				switch body.Name {
				case "stat":
					onStatResponse(agent, body.Body, socket)
				}
			}
		}
	}
}

func onStatResponse(ctx context.Context) func(models.Agent, any, *ws.Socket) {
	log := logger.New("agent.net/api/sockets")
	agents := ctx.Value("repos.agents").(repos.IAgentsRepository)

	return func(agent models.Agent, data any, socket *ws.Socket) {
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

		ipAddress := socket.IPAddress()
		agent.Status = models.AGENT_STATUS_UP
		agent.Address = &ipAddress
		agent = agents.Update(agent)
	}
}
