package config

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

const configPath = "config_dev.json"

func TestCanGetConfig(t *testing.T) {
	wd, _ := os.Getwd()
	fullPath := filepath.Join(filepath.Dir(wd), configPath)
	data := Get(fullPath)
	assert.Equal(t, "shorturl", data.Db.Name)
	assert.Equal(t, "root", data.Db.User)
}
