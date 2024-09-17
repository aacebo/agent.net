package routes

import (
	"context"
	"errors"
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
	log := logger.New("agent.net/api/sockets")
	onConnected := onConnected(ctx)
	onDisconnected := onDisconnected(ctx)
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
		agent := r.Context().Value("agent").(models.Agent)
		agentAddress := r.Header.Get("X_AGENT_ADDRESS")
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
			agent, exists := agents.GetByID(agent.ID)

			if !exists {
				return
			}

			agent.Status = models.AGENT_STATUS_DOWN
			agent.Address = nil
			agent = agents.Update(agent)
		}()

		agent.Status = models.AGENT_STATUS_UP
		agent.Address = &agentAddress
		agent = agents.Update(agent)

		for {
			message, err := socket.Read()

			if err != nil {
				return
			}

			switch message.Type {
			case ws.CONNECTED_MESSAGE_TYPE:
				err = onConnected(socket, message)
			case ws.DISCONNECTED_MESSAGE_TYPE:
				err = onDisconnected(socket, message)
			}

			if err != nil {
				log.Warn(err.Error())
				return
			}
		}
	}
}

func onConnected(ctx context.Context) func(*ws.Socket, ws.Message) error {
	log := logger.New("agent.net/api/sockets")
	agents := ctx.Value("repos.agents").(repos.IAgentsRepository)

	return func(socket *ws.Socket, msg ws.Message) error {
		body := ws.ConnectedMessageBody(msg.Body.(map[string]any))
		address := body.Address()
		agent, exists := agents.GetByID(body.ID())

		if !exists {
			return errors.New("agent not found")
		}

		agent.Status = models.AGENT_STATUS_UP
		agent.Address = &address
		agent = agents.Update(agent)

		log.Debug(fmt.Sprintf("agent '%s' connected", agent.Name))
		return nil
	}
}

func onDisconnected(ctx context.Context) func(*ws.Socket, ws.Message) error {
	log := logger.New("agent.net/api/sockets")
	agents := ctx.Value("repos.agents").(repos.IAgentsRepository)

	return func(socket *ws.Socket, msg ws.Message) error {
		id := msg.Body.(string)
		agent, exists := agents.GetByID(id)

		if !exists {
			return errors.New("agent not found")
		}

		agent.Status = models.AGENT_STATUS_DOWN
		agent = agents.Update(agent)

		log.Debug(fmt.Sprintf("agent '%s' disconnected", agent.Name))
		return nil
	}
}
