package environment

import (
	"log"

	"github.com/joho/godotenv"
)

type Environment struct {
	GithubClientID string
	GithubSecretID string

	GiteeClientID string
	GiteeSecretID string

	GitlabClientID string
	GitlabSecretID string
}

var Env Environment

func loadEnv() {
	envMap, err := godotenv.Read(".env")
	if err != nil {
		log.Fatalln("load env file failed")
	}

	Env.GithubClientID = envMap["GITHUB_CLIENT_ID"]
	Env.GithubSecretID = envMap["GITHUB_CLIENT_SECRET"]
	Env.GiteeClientID = envMap["GITEE_CLIENT_ID"]
	Env.GiteeSecretID = envMap["GITEE_CLIENT_SECRET"]
	Env.GitlabClientID = envMap["GITLAB_CLIENT_ID"]
	Env.GitlabSecretID = envMap["GITLAB_CLIENT_SECRET"]
}

func init() {
	loadEnv()
}
