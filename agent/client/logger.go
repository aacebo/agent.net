package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/aacebo/agent.net/core/logger"
	"github.com/aacebo/agent.net/core/models"
)

type LoggerClient struct {
	id      string
	baseUrl string
	client  *http.Client
	log     *slog.Logger
}

func NewLogger(id string, name string, baseUrl string) *LoggerClient {
	client := http.Client{}

	return &LoggerClient{
		id:      id,
		baseUrl: baseUrl,
		client:  &client,
		log:     logger.New(name),
	}
}

func (self LoggerClient) Info(text string, data map[string]any) error {
	return self.Log(models.LOG_LEVEL_INFO, text, data)
}

func (self LoggerClient) Warn(text string, data map[string]any) error {
	return self.Log(models.LOG_LEVEL_WARN, text, data)
}

func (self LoggerClient) Error(text string, data map[string]any) error {
	return self.Log(models.LOG_LEVEL_ERROR, text, data)
}

func (self LoggerClient) Debug(text string, data map[string]any) error {
	return self.Log(models.LOG_LEVEL_DEBUG, text, data)
}

func (self LoggerClient) Log(level models.LogLevel, text string, data map[string]any) error {
	log := self.log.Info

	switch level {
	case models.LOG_LEVEL_WARN:
		log = self.log.Warn
	case models.LOG_LEVEL_ERROR:
		log = self.log.Error
	case models.LOG_LEVEL_DEBUG:
		log = self.log.Debug
	}

	if data == nil {
		data = map[string]any{}
	}

	log(text, slog.Any("data", data))
	b, err := json.Marshal(map[string]any{
		"level": level,
		"text":  text,
		"data":  data,
	})

	if err != nil {
		return err
	}

	buf := bytes.NewBuffer(b)
	res, err := self.client.Post(
		fmt.Sprintf(
			"%s/v1/agents/%s/logs",
			self.baseUrl,
			self.id,
		),
		"application/json",
		buf,
	)

	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusCreated {
		body, _ := io.ReadAll(res.Body)
		self.log.Error(string(body))
		return errors.New(string(body))
	}

	return nil
}
