package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/aacebo/agent.net/agent/routes"
	"github.com/aacebo/agent.net/agent/sockets"
	"github.com/aacebo/agent.net/amqp"
	"github.com/aacebo/agent.net/core"
	"github.com/aacebo/agent.net/core/logger"
	"github.com/aacebo/agent.net/core/utils"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/httprate"
	"github.com/go-chi/render"
)

func main() {
	startedAt := time.Now()
	log := logger.New("agents.net/agent")
	id := os.Getenv("AGENT_ID")
	parentId := os.Getenv("AGENT_PARENT_ID")
	parentUrl := os.Getenv("AGENT_PARENT_URL")
	clientId := os.Getenv("AGENT_CLIENT_ID")
	clientSecret := os.Getenv("AGENT_CLIENT_SECRET")
	description := os.Getenv("AGENT_DESCRIPTION")
	instructions := os.Getenv("AGENT_INSTRUCTIONS")

	log.Info(id)
	log.Info(parentId)
	log.Info(parentUrl)
	log.Info(clientId)
	log.Info(clientSecret)
	log.Info(description)
	log.Info(instructions)

	amqp := amqp.New()
	defer amqp.Close()

	ctx := core.Context{
		"id":            id,
		"parent_id":     parentId,
		"parent_url":    parentUrl,
		"client_id":     clientId,
		"client_secret": clientSecret,
		"description":   description,
		"instructions":  instructions,
		"amqp":          amqp,
		"sockets":       sockets.New(),
	}

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(logger.Request(logger.New("agent.net/http")))
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
	err := http.ListenAndServe(fmt.Sprintf(":%s", utils.GetEnv("PORT", "8080")), r)

	if err != nil {
		panic(err)
	}
}
