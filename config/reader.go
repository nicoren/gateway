package config

import (
	"fmt"
	"io/ioutil"
	"utils"

	"gopkg.in/yaml.v2"
)

type Reader struct {
	Path string
	Ext  string
}

func (r Reader) Read() interface{} {
	files, err := utils.Glob(r.Path, r.Ext)
	if err != nil {
		fmt.Println(err)
	}
	var yamlConfiguration map[string]interface{}

	for _, file := range files {
		bs, err := ioutil.ReadFile(r.Path + "/" + file)
		if err != nil {
			panic(err)
		}
		if err := yaml.Unmarshal(bs, &yamlConfiguration); err != nil {
			panic(err)
		}
	}
	return yamlConfiguration

}
