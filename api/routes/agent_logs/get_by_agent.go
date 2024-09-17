package agentLogs

import (
	"context"
	"net/http"

	"github.com/aacebo/agent.net/core/models"
	"github.com/aacebo/agent.net/core/repos"
	"github.com/go-chi/render"
)

func GetByAgent(ctx context.Context) http.HandlerFunc {
	agentLogs := ctx.Value("repos.agent_logs").(repos.IAgentLogsRepository)

	return func(w http.ResponseWriter, r *http.Request) {
		agent := r.Context().Value("agent").(models.Agent)
		logs := agentLogs.GetByAgentID(agent.ID)

		render.JSON(w, r, logs)
	}
}
