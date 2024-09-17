package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"

	"github.com/aacebo/agent.net/core/logger"
	"github.com/aacebo/agent.net/core/models"
)

type LoggerClient struct {
	clientId     string
	clientSecret string
	baseUrl      string
	client       *http.Client
	log          *slog.Logger
}

func NewLogger(name string, baseUrl string) *LoggerClient {
	client := http.Client{}

	return &LoggerClient{
		clientId:     os.Getenv("AGENT_CLIENT_ID"),
		clientSecret: os.Getenv("AGENT_CLIENT_SECRET"),
		baseUrl:      baseUrl,
		client:       &client,
		log:          logger.New(name),
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
	req, err := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("%s/v1/agents/logs", self.baseUrl),
		buf,
	)

	if err != nil {
		return err
	}

	req.Header.Add("X_CLIENT_ID", self.clientId)
	req.Header.Add("X_CLIENT_SECRET", self.clientSecret)
	res, err := self.client.Do(req)

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
