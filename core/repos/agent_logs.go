package repos

import (
	"database/sql"
	"log/slog"
	"time"

	"github.com/aacebo/agent.net/core/logger"
	"github.com/aacebo/agent.net/core/models"
)

type IAgentLogsRepository interface {
	GetByAgentID(agentId string) []models.AgentLog

	Create(value models.AgentLog) models.AgentLog
}

type AgentLogsRepository struct {
	pg  *sql.DB
	log *slog.Logger
}

func AgentLogs(pg *sql.DB) AgentLogsRepository {
	return AgentLogsRepository{
		pg:  pg,
		log: logger.New("repos/agent_logs"),
	}
}

func (self AgentLogsRepository) GetByAgentID(agentId string) []models.AgentLog {
	rows, err := self.pg.Query(
		`
			SELECT
				id,
				agent_id,
				level,
				text,
				data,
				created_at
			FROM agent_logs
			WHERE parent_id IS NULL
		`,
	)

	if err != nil {
		self.log.Error(err.Error())
		return []models.AgentLog{}
	}

	defer rows.Close()
	arr := []models.AgentLog{}

	for rows.Next() {
		v := models.AgentLog{}
		err := rows.Scan(
			&v.ID,
			&v.AgentID,
			&v.Level,
			&v.Text,
			&v.Data,
			&v.CreatedAt,
		)

		if err != nil {
			self.log.Error(err.Error())
			return arr
		}

		arr = append(arr, v)
	}

	return arr
}

func (self AgentLogsRepository) Create(value models.AgentLog) models.AgentLog {
	now := time.Now()
	value.CreatedAt = now
	_, err := self.pg.Exec(
		`
			INSERT INTO agent_logs (
				id,
				agent_id,
				level,
				text,
				data,
				created_at
			) VALUES (
				$1,
				$2,
				$3,
				$4,
				$5,
				$6
			)
		`,
		value.ID,
		value.AgentID,
		value.Level,
		value.Text,
		value.Data,
		value.CreatedAt,
	)

	if err != nil {
		panic(err)
	}

	return value
}
