package fonts

import (
	"github.com/joho/godotenv"
	"io/ioutil"
	"os"
)

var (
	arialFont []byte
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic("Error Loading .env file")
	}
}

func GetArialFont() ([]byte, error) {
	if arialFont == nil {
		p := os.Getenv("FONT_PATH")
		if p == "" {
			env := os.Getenv("APP_ENV")
			p = os.Getenv("FONT_PATH_" + env)
		}
		p += "ARIAL.TTF"

		var err error
		arialFont, err = ioutil.ReadFile(p)
		return arialFont, err
	}
	return arialFont, nil
}
