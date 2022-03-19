package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"path/filepath"
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

type config struct {
	App    app    `yaml:"app"`
	Server server `yaml:"server"`
	Github github `yaml:"github"`
}

var (
	App    = app{}
	Server = server{}
	Github = github{}
)

func init() {
	setLocal()
}

func setLocal() {
	var conf config
	c, _ := filepath.Abs("config/config.yaml")
	ymlFile, err := ioutil.ReadFile(c)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(ymlFile, &conf)
	if err != nil {
		panic(err)
	}

	App = conf.App
	Server = conf.Server
	Github = conf.Github
}

func setProd() {
	App.Name = os.Getenv("MH_APP_NAME")
	App.SessionKey = os.Getenv("MH_SESSION_KEY")
	App.StateKey = os.Getenv("MH_STATE_KEY")
	Server.Port = os.Getenv("MH_SERVER_PORT")
	Github.AppId = os.Getenv("MH_GITHUB_APP_ID")
	Github.ClientId = os.Getenv("MH_GITHUB_CLIENT_ID")
	Github.ClientSecret = os.Getenv("MH_GITHUB_CLIENT_SECRET")
	Github.CallbackUrl = os.Getenv("MH_GITHUB_CALLBACK_URL")
	Github.PublicKey = os.Getenv("MH_GITHUB_PUBLIC_KEY")
	Github.PrivateKey = os.Getenv("MH_GITHUB_PRIVATE_KEY")
}
