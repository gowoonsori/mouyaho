package config

import (
	"fmt"
	"testing"
)

func Test_Env(t *testing.T) {
	fmt.Println(Config.App.Name)
	fmt.Println(Config.Server.Port)
	fmt.Println(Config.Github.AppId)
	fmt.Println(Config.Github.ClientId)
	fmt.Println(Config.Github.ClientSecret)
	fmt.Println(Config.Github.CallbackUrl)
	fmt.Println(Config.Github.PrivateKey)
}
