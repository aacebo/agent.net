package repos

import (
	"database/sql"
	"fmt"
	"log/slog"
	"time"

	"github.com/aacebo/agent.net/core/logger"
	"github.com/aacebo/agent.net/core/models"
)

type IAgentsRepository interface {
	Get() []models.Agent
	GetEdges(parentId string) []models.Agent
	GetByID(id string) (models.Agent, bool)
	GetByName(name string) (models.Agent, bool)
	GetByCredentials(clientId string, clientSecret string) (models.Agent, bool)

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

func (self AgentsRepository) Get() []models.Agent {
	rows, err := self.pg.Query(
		`
			SELECT
				id,
				parent_id,
				container_id,
				status,
				name,
				description,
				instructions,
				address,
				client_id,
				client_secret,
				settings,
				created_at,
				updated_at
			FROM agents
			ORDER BY updated_at DESC
		`,
	)

	if err != nil {
		self.log.Error(err.Error())
		return []models.Agent{}
	}

	defer rows.Close()
	arr := []models.Agent{}

	for rows.Next() {
		v := models.Agent{}
		err := rows.Scan(
			&v.ID,
			&v.ParentID,
			&v.ContainerID,
			&v.Status,
			&v.Name,
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
			self.log.Error(err.Error())
			return arr
		}

		arr = append(arr, v)
	}

	return arr
}

func (self AgentsRepository) GetEdges(parentId string) []models.Agent {
	rows, err := self.pg.Query(
		`
			SELECT
				id,
				parent_id,
				container_id,
				status,
				name,
				description,
				instructions,
				address,
				client_id,
				client_secret,
				settings,
				created_at,
				updated_at
			FROM agents
			WHERE parent_id = $1
			ORDER BY updated_at DESC
		`,
		parentId,
	)

	if err != nil {
		self.log.Error(err.Error())
		return []models.Agent{}
	}

	defer rows.Close()
	arr := []models.Agent{}

	for rows.Next() {
		v := models.Agent{}
		err := rows.Scan(
			&v.ID,
			&v.ParentID,
			&v.ContainerID,
			&v.Status,
			&v.Name,
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
			self.log.Error(err.Error())
			return arr
		}

		arr = append(arr, v)
	}

	return arr
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
				name,
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
		&v.Name,
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

func (self AgentsRepository) GetByName(name string) (models.Agent, bool) {
	v := models.Agent{}
	err := self.pg.QueryRow(
		`
			SELECT
				id,
				parent_id,
				container_id,
				status,
				name,
				description,
				instructions,
				address,
				client_id,
				client_secret,
				settings,
				created_at,
				updated_at
			FROM agents
			WHERE name = $1
		`,
		name,
	).Scan(
		&v.ID,
		&v.ParentID,
		&v.ContainerID,
		&v.Status,
		&v.Name,
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

func (self AgentsRepository) GetByCredentials(clientId string, clientSecret string) (models.Agent, bool) {
	v := models.Agent{}
	err := self.pg.QueryRow(
		`
			SELECT
				id,
				parent_id,
				container_id,
				status,
				name,
				description,
				instructions,
				address,
				client_id,
				client_secret,
				settings,
				created_at,
				updated_at
			FROM agents
			WHERE client_id = $1
		`,
		clientId,
	).Scan(
		&v.ID,
		&v.ParentID,
		&v.ContainerID,
		&v.Status,
		&v.Name,
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

	if v.ClientSecret != models.Secret(clientSecret) {
		return v, false
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
				name,
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
				$12,
				$13
			)
		`,
		value.ID,
		value.ParentID,
		value.ContainerID,
		value.Status,
		value.Name,
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

	if err = self.partition(value); err != nil {
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

func (self AgentsRepository) partition(agent models.Agent) error {
	_, err := self.pg.Exec(
		fmt.Sprintf(
			`
				CREATE TABLE IF NOT EXISTS agent_logs_%s
				PARTITION OF agent_logs FOR VALUES IN ('%s')
				PARTITION BY RANGE (created_at);

				SELECT partman.create_parent(
					p_parent_table => 'public.agent_logs_%s',
					p_control => 'created_at',
					p_type => 'range',
					p_interval => '1 month',
					p_premake => 4
				);
			`,
			agent.Name,
			agent.ID,
			agent.Name,
		),
	)

	return err
}
