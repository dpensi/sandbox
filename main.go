package main

import (
	"encoding/json"
	"fmt"

	toml "github.com/pelletier/go-toml"
	yaml "gopkg.in/yaml.v3"
)

type configuration struct {
	Url        string               `json:"superset_url"`
	Path       string               `json:"superset_dashboards_path"`
	Separator  string               `json:"param_values_separator"`
	Dashboards map[string]dashboard `json:"dashboards"`
}

type dashboard struct {
	Format string            `json:"type"`
	Fields map[string]fields `json:"fields"`
}

type fields struct {
	Id         string `json:"id"`
	ColumnName string `json:"actual_column_name"`
	Def        string `json:"default"`
}

func main() {
	fmt.Println("starting")

	field_one := fields{
		Id:         "f1id",
		ColumnName: "col1 name",
		Def:        "default",
	}
	field_two := fields{
		Id:         "f2id",
		ColumnName: "col2name",
		Def:        "default",
	}
	fields_map := make(map[string]fields)
	fields_map["one"] = field_one
	fields_map["two"] = field_two

	dash_one := dashboard{
		Format: "one",
		Fields: fields_map,
	}
	dash_two := dashboard{
		Format: "two",
		Fields: fields_map,
	}
	dash_map := make(map[string]dashboard)
	dash_map["unos"] = dash_one
	dash_map["dos"] = dash_two
	config := configuration{
		Url:        "uurrll",
		Path:       "ppaatthh",
		Separator:  ",",
		Dashboards: dash_map,
	}

	b, err := toml.Marshal(config)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))

	fmt.Println("finished toml")

	d, err := yaml.Marshal(config)
	if err != nil {
		panic("oh no")
	}
	fmt.Printf("yaml:\n%s\n\n", string(d))

	j, err := json.Marshal(config)
	if err != nil {
		panic("oh no")
	}
	fmt.Printf("json:\n%s\n\n", string(j))

}
