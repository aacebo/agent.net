package main

import (
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/aacebo/agent.net/agent/routes"
	"github.com/aacebo/agent.net/core"
	"github.com/aacebo/agent.net/core/logger"
	"github.com/aacebo/agent.net/core/utils"
	"github.com/aacebo/agent.net/ws"
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
	url := os.Getenv("AGENT_URL")
	clientId := os.Getenv("AGENT_CLIENT_ID")
	clientSecret := os.Getenv("AGENT_CLIENT_SECRET")
	description := os.Getenv("AGENT_DESCRIPTION")
	instructions := os.Getenv("AGENT_INSTRUCTIONS")

	client := ws.NewClient()
	clientHeaders := http.Header{
		"client_id":     {clientId},
		"client_secret": {clientSecret},
	}

	ctx := core.Context{
		"id":            id,
		"url":           url,
		"client_id":     clientId,
		"client_secret": clientSecret,
		"description":   description,
		"instructions":  instructions,
		"socket":        client,
		"sockets":       ws.NewSockets(),
	}

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()

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
	}()

	go func() {
		defer wg.Done()

		if err := client.Connect(url, clientHeaders); err != nil {
			panic(err)
		}

		defer client.Close()
		log.Info("connected...")

		for {
			message, err := client.Read()

			if err != nil {
				return
			}

			log.Debug(message.String())
		}
	}()

	wg.Wait()
}
