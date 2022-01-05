package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const configFile = "config_dev.json"

func TestCanGetConfig(t *testing.T) {
	wd, _ := os.Getwd()
	fullPath, ok := FindIn(wd, configFile)
	assert.True(t, ok)
	data := Get(fullPath)
	assert.Equal(t, "simpurl", data.Db.Name)
	assert.Equal(t, "root", data.Db.User)
}

func TestConfigDoesNotExists(t *testing.T) {
	wd, _ := os.Getwd()
	p, ok := FindIn(wd, "rnd")
	assert.Equal(t, "/", p)
	assert.False(t, ok)
}
