package config

import (
	"path/filepath"
	"io/ioutil"
	"encoding/json"
	"os"
	"github.com/bwmarrin/snowflake"
	"flag"
	"fmt"
)

type Config struct {
	ListenHttp string          `json:"listen_http"`
	Apps       []Apps          `json:"apps"`
	UniqAddr   map[string]bool `json:"-"`
}

type Apps struct {
	GroupId      string `json:"groupid,omitempty"`
	Name         string `json:"name"`
	Type         string `json:"type"`
	Url          string `json:"url"`
	Node         []Node `json:"node"`
	Fetchlogpath string `json:"fetchlogpath"`
}

type Node struct {
	Alias       string `json:"alias"`
	Addr        string `json:"addr"`
	Type        string `json:"type,omitempty"`
	Path        string `json:"path"`
	BeforDeploy string `json:"befor_deploy"`
	AfterDeploy string `json:"after_deploy"`
	Online      bool   `json:"online,omitempty"`
}

var C *Config
var configUsage = `Usage: /pathto/server -c /pathto/config.json`

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, configUsage+"\n")
	}
	file := flag.String("c", "", configUsage)
	flag.Parse()
	if *file == "" {
		flag.Usage()
		os.Exit(1)
	}

	C = New(*file)
	node, err := snowflake.NewNode(1)
	if err != nil {
		panic(err)
	}
	for key, val := range C.Apps {
		C.Apps[key].GroupId = node.Generate().String()
		for k, v := range val.Node {
			C.UniqAddr[v.Addr] = true
			C.Apps[key].Node[k].Type = val.Type
		}
	}
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

	c := &Config{UniqAddr: make(map[string]bool)}
	return ParseJson(bytes, c)
}

func ParseJson(bytes []byte, v *Config) *Config {
	err := json.Unmarshal(bytes, &v)
	if err != nil {
		panic(err)
	}
	return v
}
