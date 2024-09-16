package repos

import (
	"database/sql"
	"log/slog"
	"time"

	"github.com/aacebo/agent.net/core/logger"
	"github.com/aacebo/agent.net/core/models"
)

type IAgentsRepository interface {
	GetByID(id string) (models.Agent, bool)

	Create(value models.Agent) models.Agent
	Update(value models.Agent) models.Agent
	Delete(id string)
}

type AgentsRepository struct {
	pg  *sql.DB
	log *slog.Logger
}

func Agents(pg *sql.DB) AgentsRepository {
	return AgentsRepository{
		pg:  pg,
		log: logger.New("agent.net/repos/agents"),
	}
}

func (self AgentsRepository) GetByID(id string) (models.Agent, bool) {
	v := models.Agent{}
	err := self.pg.QueryRow(
		`
			SELECT
				id,
				parent_id,
				container_id,
				status,
				description,
				instructions,
				address,
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
		&v.ContainerID,
		&v.Status,
		&v.Description,
		&v.Instructions,
		&v.Address,
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

func (self AgentsRepository) Create(value models.Agent) models.Agent {
	now := time.Now()
	value.CreatedAt = now
	value.UpdatedAt = now
	_, err := self.pg.Exec(
		`
			INSERT INTO agents (
				id,
				parent_id,
				container_id,
				status,
				description,
				instructions,
				address,
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
				$9,
				$10,
				$11,
				$12
			)
		`,
		value.ID,
		value.ParentID,
		value.ContainerID,
		value.Status,
		value.Description,
		value.Instructions,
		value.Address,
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

func (self AgentsRepository) Update(value models.Agent) models.Agent {
	now := time.Now()
	value.UpdatedAt = now
	_, err := self.pg.Exec(
		`
			UPDATE agents SET
				container_id = $2,
				status = $3,
				description = $4,
				instructions = $5,
				address = $6,
				settings = $7,
				updated_at = $8
			WHERE id = $1
		`,
		value.ID,
		value.ContainerID,
		value.Status,
		value.Description,
		value.Instructions,
		value.Address,
		value.Settings,
		value.UpdatedAt,
	)

	if err != nil {
		panic(err)
	}

	return value
}

func (self AgentsRepository) Delete(id string) {
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
