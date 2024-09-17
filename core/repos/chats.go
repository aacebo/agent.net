package repos

import (
	"database/sql"
	"log/slog"
	"time"

	"github.com/aacebo/agent.net/core/logger"
	"github.com/aacebo/agent.net/core/models"
)

type IChatsRepository interface {
	GetByAgentID(agentId string) []models.Chat
	GetByID(id string) (models.Chat, bool)

	Create(value models.Chat) models.Chat
	Delete(id string)
}

type ChatsRepository struct {
	pg  *sql.DB
	log *slog.Logger
}

func Chats(pg *sql.DB) ChatsRepository {
	return ChatsRepository{
		pg:  pg,
		log: logger.New("agent.net/repos/chats"),
	}
}

func (self ChatsRepository) GetByAgentID(agentId string) []models.Chat {
	rows, err := self.pg.Query(
		`
			SELECT
				id,
				agent_id,
				created_at,
				updated_at
			FROM chats
			WHERE agent_id = $1
			ORDER BY created_at DESC
		`,
		agentId,
	)

	if err != nil {
		self.log.Error(err.Error())
		return []models.Chat{}
	}

	defer rows.Close()
	arr := []models.Chat{}

	for rows.Next() {
		v := models.Chat{}
		err := rows.Scan(
			&v.ID,
			&v.AgentID,
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

func (self ChatsRepository) GetByID(id string) (models.Chat, bool) {
	v := models.Chat{}
	err := self.pg.QueryRow(
		`
			SELECT
				id,
				agent_id,
				created_at,
				updated_at
			FROM chats
			WHERE id = $1
		`,
		id,
	).Scan(
		&v.ID,
		&v.AgentID,
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

func (self ChatsRepository) Create(value models.Chat) models.Chat {
	now := time.Now()
	value.CreatedAt = now
	value.UpdatedAt = now
	_, err := self.pg.Exec(
		`
			INSERT INTO chats (
				id,
				agent_id,
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
		value.AgentID,
		value.CreatedAt,
		value.UpdatedAt,
	)

	if err != nil {
		panic(err)
	}

	return value
}

func (self ChatsRepository) Delete(id string) {
	_, err := self.pg.Exec(
		`
			DELETE FROM chats
			WHERE id = $1
		`,
		id,
	)

	if err != nil {
		panic(err)
	}
}
