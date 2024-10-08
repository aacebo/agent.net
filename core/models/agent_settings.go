package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type AgentSettings struct {
	ApiKey           string   `json:"api_key"`
	Model            string   `json:"model"`                       // https://platform.openai.com/docs/api-reference/chat/create#chat-create-model
	FrequencyPenalty *float32 `json:"frequency_penalty,omitempty"` // https://platform.openai.com/docs/api-reference/chat/create#chat-create-frequency_penalty
	LogitBias        Map[any] `json:"logit_bias,omitempty"`        // https://platform.openai.com/docs/api-reference/chat/create#chat-create-logit_bias
	LogProbs         *bool    `json:"logprobs,omitempty"`          // https://platform.openai.com/docs/api-reference/chat/create#chat-create-logprobs
}

func (self AgentSettings) String() string {
	b, _ := json.Marshal(self)
	return string(b)
}

func (self AgentSettings) Value() (driver.Value, error) {
	value, err := json.Marshal(self)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("[models:agent_settings]: %s", err.Error()))
	}

	return value, nil
}

func (self *AgentSettings) Scan(value any) error {
	err := json.Unmarshal(value.([]byte), self)

	if err != nil {
		return errors.New(fmt.Sprintf("[models:agent_settings]: %s", err.Error()))
	}

	return nil
}
