package config

import (
	"encoding/base64"
	"os"
)

var (
	GithubAPI = "https://api.github.com"
)

type app struct {
	Name       string `yaml:"name"`
	SessionKey string `yaml:"session-key"`
	StateKey   string `yaml:"state-key"`
}

type server struct {
	Port string `yaml:"port"`
}

type github struct {
	AppId        string `yaml:"app-id"`
	ClientId     string `yaml:"client-id"`
	ClientSecret string `yaml:"client-secret"`
	CallbackUrl  string `yaml:"callback-url"`
	PublicKey    string `yaml:"public-key"`
	PrivateKey   string `yaml:"private-key"`
}

var (
	App    = app{}
	Server = server{}
	Github = github{}
)

func init() {
	App.Name = os.Getenv("MH_APP_NAME")
	App.SessionKey = os.Getenv("MH_SESSION_KEY")
	App.StateKey = os.Getenv("MH_STATE_KEY")
	Server.Port = os.Getenv("MH_SERVER_PORT")
	Github.AppId = os.Getenv("MH_GITHUB_APP_ID")
	Github.ClientId = os.Getenv("MH_GITHUB_CLIENT_ID")
	Github.ClientSecret = os.Getenv("MH_GITHUB_CLIENT_SECRET")
	Github.CallbackUrl = os.Getenv("MH_GITHUB_CALLBACK_URL")
	Github.PublicKey = os.Getenv("MH_GITHUB_PUBLIC_KEY")

	//RSA Private key used base64 encode
	pk, _ := base64.URLEncoding.DecodeString(os.Getenv("MH_GITHUB_PRIVATE_KEY"))
	Github.PrivateKey = string(pk)
}
