package env

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"path/filepath"
)

type app struct {
	Name string `yaml:"name"`
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
	Config.App.Name = os.Getenv("WC_APP_NAME")
	Config.Server.Port = os.Getenv("WC_SERVER_PORT")
	Config.Github.AppId = os.Getenv("WC_GITHUB_APP_ID")
	Config.Github.ClientId = os.Getenv("WC_GITHUB_CLIENT_ID")
	Config.Github.ClientSecret = os.Getenv("WC_GITHUB_CLIENT_SECRET")
	Config.Github.CallbackUrl = os.Getenv("WC_GITHUB_CALLBACK_URL")
	Config.Github.PublicKey = os.Getenv("WC_GITHUB_PUBLIC_KEY")
	Config.Github.PrivateKey = os.Getenv("WC_GITHUB_PRIVATE_KEY")
}
