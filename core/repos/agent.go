package repos

import (
	"database/sql"
	"log/slog"
	"time"

	"github.com/aacebo/agent.net/core/logger"
	"github.com/aacebo/agent.net/core/models"
)

type IAgentRepository interface {
	GetByID(id string) (models.Agent, bool)

	Create(value models.Agent) models.Agent
	Update(value models.Agent) models.Agent
	Delete(id string)
}

type AgentRepository struct {
	pg  *sql.DB
	log *slog.Logger
}

func NewAgent(pg *sql.DB) AgentRepository {
	return AgentRepository{
		pg:  pg,
		log: logger.New("agent.net/repos/agents"),
	}
}

func (self AgentRepository) GetByID(id string) (models.Agent, bool) {
	v := models.Agent{}
	err := self.pg.QueryRow(
		`
			SELECT
				id,
				parent_id,
				description,
				instructions,
				client_id,
				client_secret,
				settings,
				created_at,
				updated_at
			FROM agents
			WHERE id = $1
		`,
		id,
	).Scan(
		&v.ID,
		&v.ParentID,
		&v.Description,
		&v.Instructions,
		&v.ClientID,
		&v.ClientSecret,
		&v.Settings,
		&v.CreatedAt,
		&v.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return v, false
		}

		panic(err)
	}

	return v, true
}

func (self AgentRepository) Create(value models.Agent) models.Agent {
	now := time.Now()
	value.CreatedAt = now
	value.UpdatedAt = now
	_, err := self.pg.Exec(
		`
			INSERT INTO agents (
				id,
				parent_id,
				description,
				instructions,
				client_id,
				client_secret,
				settings,
				created_at,
				updated_at
			) VALUES (
				$1,
				$2,
				$3,
				$4,
				$5,
				$6,
				$7,
				$8,
				$9
			)
		`,
		value.ID,
		value.ParentID,
		value.Description,
		value.Instructions,
		value.ClientID,
		value.ClientSecret,
		value.Settings,
		value.CreatedAt,
		value.UpdatedAt,
	)

	if err != nil {
		panic(err)
	}

	return value
}

func (self AgentRepository) Update(value models.Agent) models.Agent {
	now := time.Now()
	value.UpdatedAt = now
	_, err := self.pg.Exec(
		`
			UPDATE agents SET
				description = $2,
				instructions = $3,
				settings = $4,
				updated_at = $5
			WHERE id = $1
		`,
		value.ID,
		value.Description,
		value.Instructions,
		value.Settings,
		value.UpdatedAt,
	)

	if err != nil {
		panic(err)
	}

	return value
}

func (self AgentRepository) Delete(id string) {
	_, err := self.pg.Exec(
		`
			DELETE FROM agents
			WHERE id = $1
		`,
		id,
	)

	if err != nil {
		panic(err)
	}
}
