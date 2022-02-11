package env

import (
	"os"
)

var (
	rootPath = os.Getenv("ROOT_PATH")
)

func GetRootPath() string {
	return rootPath
}
