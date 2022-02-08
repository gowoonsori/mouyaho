package env

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_GetRootPath(t *testing.T) {
	d := GetRootPath()
	assert.Equal(t, rootPath, d)
}
