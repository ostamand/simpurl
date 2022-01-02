package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
)

type Params struct {
	Db      ParamsDB  `json:"db"`
	General GeneralDB `json:"general"`
}

type GeneralDB struct {
	AdminOnly bool `json:"adminOnly"`
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

func FindInParent(dir string, fileName string) (p string, ok bool) {
	ok = true
	p = path.Join(dir, fileName)
	if _, err := os.Stat(p); err != nil {
		if strings.Compare(dir, "/") != 0 {
			p, ok = FindInParent(path.Dir(dir), fileName)
		} else {
			return dir, false
		}
	}
	return p, ok
}
