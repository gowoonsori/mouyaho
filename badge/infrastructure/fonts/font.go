package fonts

import (
	"io/ioutil"
	"path/filepath"
	"runtime"
)

var (
	arialFont []byte
)

func GetArialFont() ([]byte, error) {
	if arialFont == nil {
		_, b, _, _ := runtime.Caller(0)
		b = filepath.Dir(b) + "/ARIAL.TTF"

		var err error
		arialFont, err = ioutil.ReadFile(b)
		return arialFont, err
	}
	return arialFont, nil
}
