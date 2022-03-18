package env

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"path/filepath"
)

type app struct {
	Name       string `yaml:"name"`
	SessionKey string `yaml:"session-key"`
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
	Config config
)

func init() {
	setLocal()
}

func setLocal() {
	conf, _ := filepath.Abs("config.yaml")
	ymlFile, err := ioutil.ReadFile(conf)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(ymlFile, &Config)
	if err != nil {
		panic(err)
	}
}

func setProd() {
	Config.App.Name = os.Getenv("MH_APP_NAME")
	Config.App.SessionKey = os.Getenv("MH_SESSION_KEY")
	Config.Server.Port = os.Getenv("MH_SERVER_PORT")
	Config.Github.AppId = os.Getenv("MH_GITHUB_APP_ID")
	Config.Github.ClientId = os.Getenv("MH_GITHUB_CLIENT_ID")
	Config.Github.ClientSecret = os.Getenv("MH_GITHUB_CLIENT_SECRET")
	Config.Github.CallbackUrl = os.Getenv("MH_GITHUB_CALLBACK_URL")
	Config.Github.PublicKey = os.Getenv("MH_GITHUB_PUBLIC_KEY")
	Config.Github.PrivateKey = os.Getenv("MH_GITHUB_PRIVATE_KEY")
}
