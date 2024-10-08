package main

import (
	"encoding/gob"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/aacebo/agent.net/amqp"
	"github.com/aacebo/agent.net/api/routes"
	"github.com/aacebo/agent.net/api/schemas"
	"github.com/aacebo/agent.net/core"
	"github.com/aacebo/agent.net/core/logger"
	"github.com/aacebo/agent.net/core/models"
	"github.com/aacebo/agent.net/core/repos"
	"github.com/aacebo/agent.net/core/utils"
	"github.com/aacebo/agent.net/postgres"
	"github.com/aacebo/agent.net/ws"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/httprate"
	"github.com/go-chi/render"
)

func main() {
	startedAt := time.Now()
	os.Setenv("TZ", "") // UTC
	gob.Register(map[string]any{})
	gob.Register(time.Time{})
	models.Register()

	pg := postgres.New()
	defer pg.Close()

	amqp := amqp.New()
	defer amqp.Close()

	schemas, err := schemas.Load()

	if err != nil {
		panic(err)
	}

	ctx := core.Context{
		"amqp":             amqp,
		"pg":               pg,
		"schemas":          schemas,
		"sockets":          ws.NewSockets(),
		"repos.agents":     repos.Agents(pg),
		"repos.agent_logs": repos.AgentLogs(pg),
		"repos.chats":      repos.Chats(pg),
		"repos.messages":   repos.Messages(pg),
	}

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(logger.Request(logger.New("http")))
	r.Use(middleware.Recoverer)
	r.Use(middleware.NoCache)
	r.Use(middleware.Timeout(60 * time.Second))
	r.Use(httprate.LimitByRealIP(600, 1*time.Minute))
	r.Use(render.SetContentType(render.ContentTypeJSON))
	r.Use(cors.AllowAll().Handler)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		render.JSON(w, r, map[string]any{
			"started_at": startedAt.UnixNano(),
		})
	})

	r.Mount("/v1", routes.New(ctx))
	err = http.ListenAndServe(fmt.Sprintf(":%s", utils.GetEnv("PORT", "3000")), r)

	if err != nil {
		panic(err)
	}
}
