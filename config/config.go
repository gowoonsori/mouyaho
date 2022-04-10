package config

import (
	"encoding/base64"
	"flag"
	"os"
)

type app struct {
	Name      string `yaml:"name"`
	CipherKey string `yaml:"ciper-key"`
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
	App.Name = *flag.String("name", "mouyaho", "App Name")
	App.CipherKey = os.Getenv("MH_CIPHER_KEY")
	Server.Port = *flag.String("port", ":8100", "Host port for the server")
	Github.AppId = os.Getenv("MH_GITHUB_APP_ID")
	Github.ClientId = os.Getenv("MH_GITHUB_CLIENT_ID")
	Github.ClientSecret = os.Getenv("MH_GITHUB_CLIENT_SECRET")
	Github.CallbackUrl = os.Getenv("MH_GITHUB_CALLBACK_URL")
	Github.PublicKey = os.Getenv("MH_GITHUB_PUBLIC_KEY")

	//RSA Private key used base64 encode
	pk, _ := base64.URLEncoding.DecodeString(os.Getenv("MH_GITHUB_PRIVATE_KEY"))
	Github.PrivateKey = string(pk)
}
