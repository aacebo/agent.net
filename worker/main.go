package main

import (
	"encoding/gob"
	"os"
	"sync"

	"github.com/aacebo/agent.net/amqp"
	"github.com/aacebo/agent.net/core"
	"github.com/aacebo/agent.net/core/models"
	"github.com/aacebo/agent.net/core/repos"
	"github.com/aacebo/agent.net/postgres"
	"github.com/aacebo/agent.net/worker/routes"
)

func main() {
	os.Setenv("TZ", "") // UTC
	gob.Register(map[string]any{})
	models.Register()
	wg := sync.WaitGroup{}
	wg.Add(1)

	pg := postgres.New()
	defer pg.Close()

	amqp := amqp.New()
	defer amqp.Close()

	ctx := core.Context{
		"amqp":         amqp,
		"pg":           pg,
		"repos.agents": repos.Agents(pg),
	}

	routes.New(ctx)
	wg.Wait()
}
