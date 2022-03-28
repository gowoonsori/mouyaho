package config

import (
	"fmt"
	"os"
	"testing"
)

func Test_Env(t *testing.T) {
	fmt.Println(os.Getenv("MH_APP_NAME"))
	fmt.Println(os.Getenv("MH_SESSION_KEY"))
	fmt.Println(os.Getenv("MH_STATE_KEY"))
	fmt.Println(os.Getenv("MH_SERVER_PORT"))
	fmt.Println(os.Getenv("MH_GITHUB_APP_ID"))
	fmt.Println(os.Getenv("MH_GITHUB_CLIENT_ID"))
	fmt.Println(os.Getenv("MH_GITHUB_CLIENT_SECRET"))
	fmt.Println(os.Getenv("MH_GITHUB_CALLBACK_URL"))
	fmt.Println(os.Getenv("MH_GITHUB_PUBLIC_KEY"))
	fmt.Println(os.Getenv("MH_GITHUB_PRIVATE_KEY"))
	fmt.Println(Github.PrivateKey)
}
