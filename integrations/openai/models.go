package openai

import (
	"encoding/json"
	"fmt"
	"openai-cli/system"
)

type ModelResponse struct {
	Object string  `json:"object"`
	Data   []Model `json:"data"`
}

type Model struct {
	Id      string `json:"id"`
	Object  string `json:"object"`
	Created int64  `json:"created"`
	OwnedBy string `json:"owned_by"`
}

func ListModels() []Model {
	listModelsUrlGet := "https://api.openai.com/v1/models"
	headers := map[string]string{
		"Accept":        "application/json",
		"Authorization": fmt.Sprintf("Bearer %s", system.OpenAIApiKey()),
	}

	response := system.Get(listModelsUrlGet, headers)
	ModelResponseJson := new(ModelResponse)
	err := json.Unmarshal([]byte(response), &ModelResponseJson)

	if err != nil {
		println("failed to parse response from ListModels()")
		return make([]Model, 0)
	}

	return ModelResponseJson.Data
}
