package repos

import (
	"database/sql"
	"log/slog"
	"time"

	"github.com/aacebo/agent.net/core/logger"
	"github.com/aacebo/agent.net/core/models"
)

type IMessagesRepository interface {
	GetByChatID(chatId string) []models.Message
	GetByID(id string) (models.Message, bool)

	Create(value models.Message) models.Message
	Delete(id string)
}

type MessagesRepository struct {
	pg  *sql.DB
	log *slog.Logger
}

func Messages(pg *sql.DB) MessagesRepository {
	return MessagesRepository{
		pg:  pg,
		log: logger.New("agent.net/repos/messages"),
	}
}

func (self MessagesRepository) GetByChatID(chatId string) []models.Message {
	rows, err := self.pg.Query(
		`
			SELECT
				id,
				chat_id,
				from_id,
				text,
				created_at,
				updated_at
			FROM messages
			WHERE chat_id = $1
			ORDER BY created_at DESC
		`,
		chatId,
	)

	if err != nil {
		self.log.Error(err.Error())
		return []models.Message{}
	}

	defer rows.Close()
	arr := []models.Message{}

	for rows.Next() {
		v := models.Message{}
		err := rows.Scan(
			&v.ID,
			&v.ChatID,
			&v.FromID,
			&v.Text,
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

func (self MessagesRepository) GetByID(id string) (models.Message, bool) {
	v := models.Message{}
	err := self.pg.QueryRow(
		`
			SELECT
				id,
				chat_id,
				from_id,
				text,
				created_at,
				updated_at
			FROM messages
			WHERE id = $1
		`,
		id,
	).Scan(
		&v.ID,
		&v.ChatID,
		&v.FromID,
		&v.Text,
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

func (self MessagesRepository) Create(value models.Message) models.Message {
	now := time.Now()
	value.CreatedAt = now
	value.UpdatedAt = now
	_, err := self.pg.Exec(
		`
			INSERT INTO messages (
				id,
				chat_id,
				from_id,
				text,
				created_at,
				updated_at
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
		value.ChatID,
		value.FromID,
		value.Text,
		value.CreatedAt,
		value.UpdatedAt,
	)

	if err != nil {
		panic(err)
	}

	return value
}

func (self MessagesRepository) Delete(id string) {
	_, err := self.pg.Exec(
		`
			DELETE FROM messages
			WHERE id = $1
		`,
		id,
	)

	if err != nil {
		panic(err)
	}
}
