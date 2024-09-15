package routes

import (
	"context"
	"net/http"

	"github.com/aacebo/agent.net/ws"

	"github.com/go-chi/render"
	"github.com/gorilla/websocket"
)

func Socket(ctx context.Context) http.HandlerFunc {
	sockets := ctx.Value("sockets").(*ws.Sockets)
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	return func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)

		if err != nil {
			render.Status(r, 500)
			render.PlainText(w, r, err.Error())
			return
		}

		socket := sockets.Add(conn)

		defer func() {
			conn.Close()
			sockets.Del(socket.ID)
		}()

		for {
			_, err := socket.Read()

			if err != nil {
				return
			}
		}
	}
}
