package store

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testStorageMap = &storageMap{}

func init() {
	s := InitializeMap()
	testStorageMap = &s
}

func TestCanInitMap(t *testing.T) {
	service := InitializeMap()
	if service.db == nil {
		t.Errorf("Map init failed. db is `nil`")
	}
}

func TestCanSaveMapping(t *testing.T) {
	testStorageMap.SaveMapping("short", "full")
	v := testStorageMap.db["short"]
	assert.Equal(t, "full", v)
}

func TestCanRetrieveURL(t *testing.T) {
	testStorageMap.SaveMapping("short", "full")
	v, _ := testStorageMap.RetrieveURL("short")
	assert.Equal(t, "full", v)
}
