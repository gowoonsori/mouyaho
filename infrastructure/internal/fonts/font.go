package fonts

import (
	"github.com/joho/godotenv"
	"io/ioutil"
	"log"
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
		env := os.Getenv("APP_ENV")
		p := os.Getenv("FONT_PATH_"+env) + "ARIAL.TTF"

		log.Println("font path = " + p)
		var err error
		arialFont, err = ioutil.ReadFile(p)
		return arialFont, err
	}
	return arialFont, nil
}
