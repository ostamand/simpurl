package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Params struct {
	Db ParamsDB `json:"db"`
}

type ParamsDB struct {
	Instance  string `json:"instance"`
	SocketDir string `json:"socketDir"`
	Port      string `json:"port"`
	User      string `json:"user"`
	Pass      string `json:"pass"`
	Name      string `json:"name"`
	Addr      string `json:"addr"`
}

func Get(filePath string) *Params {
	log.Printf("loading config file: %s", filePath)
	jsonFile, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()
	var data *Params
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &data)
	return data
}
