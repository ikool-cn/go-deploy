package config

import (
	"path/filepath"
	"io/ioutil"
	"encoding/json"
	"net"
)

type Config struct {
	ListenHttp string `json:"listen_http"`
	ListenTcp  string `json:"listen_tcp"`
	Apps       []Apps `json:"apps"`
}

type Apps struct {
	GroupId string `json:"groupid,omitempty"`
	Name    string `json:"name"`
	Node    []Node `json:"node"`
}

type Node struct {
	Ip          string   `json:"ip"`
	Type        string   `json:"type"`
	Path        string   `json:"path"`
	BeforDeploy string   `json:"befor_deploy"`
	AfterDeploy string   `json:"after_deploy"`
	Online      bool     `json:"online,omitempty"`
	Conn        net.Conn `json:"-"`
}

func New(file string) *Config {
	if !filepath.IsAbs(file) {
		absFile, err := filepath.Abs(file)
		if err != nil {
			panic(err)
		}
		file = absFile
	}

	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}

	c := &Config{}
	return ParseJson(bytes, c)
}

func ParseJson(bytes []byte, v *Config) *Config {
	err := json.Unmarshal(bytes, &v)
	if err != nil {
		panic(err)
	}
	return v
}
