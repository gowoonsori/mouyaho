package main

import (
	"fmt"
	"likeIt/env"
)

func main() {
	fmt.Printf(env.Config.App.Name)
}
