package system

import "os"

func isEnvExist(key string) bool {
	if _, ok := os.LookupEnv(key); ok {
		return true
	}
	return false
}

func OpenAIOrganization() string {
	return os.Getenv("OPENAI_ORGANIZATION")
}

func OpenAIApiKey() string {
	return os.Getenv("OPENAI_APIKEY")
}

func CheckOpenAIConfiguration() bool {
	return isEnvExist("OPENAI_ORGANIZATION") && isEnvExist("OPENAI_APIKEY")
}
