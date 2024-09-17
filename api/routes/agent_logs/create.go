package agentLogs

import (
	"context"
	"net/http"

	"github.com/aacebo/agent.net/core/models"
	"github.com/aacebo/agent.net/core/repos"
	"github.com/go-chi/render"
)

type CreateBody struct {
	Level models.LogLevel `json:"level"`
	Text  string          `json:"text"`
	Data  map[string]any  `json:"data"`
}

func Create(ctx context.Context) http.HandlerFunc {
	agentLogs := ctx.Value("repos.agent_logs").(repos.IAgentLogsRepository)

	return func(w http.ResponseWriter, r *http.Request) {
		agent := r.Context().Value("agent").(models.Agent)
		body := r.Context().Value("body").(CreateBody)

		log := models.NewAgentLog()
		log.AgentID = agent.ID
		log.Level = body.Level
		log.Text = body.Text
		log.Data = body.Data
		log = agentLogs.Create(log)

		render.Status(r, http.StatusCreated)
		render.JSON(w, r, log)
	}
}
