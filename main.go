package main

import (
	"fmt"

	toml "github.com/pelletier/go-toml"
)

type configuration struct {
	url        string               `toml:"superset_url"`
	path       string               `toml:"superset_dashboards_path"`
	separator  string               `toml:"param_values_separator"`
	dashboards map[string]dashboard `toml:"dashboards"`
}

type dashboard struct {
	format string            `toml:"type"`
	fields map[string]fields `toml:"fields"`
}

type fields struct {
	id         string `toml:"id"`
	columnName string `toml:"actual_column_name"`
	def        string `toml:"default"`
}

func main() {
	fmt.Println("starting")

	field_one := fields{
		id:         "f1id",
		columnName: "col1 name",
		def:        "default",
	}
	field_two := fields{
		id:         "f2id",
		columnName: "col2name",
		def:        "default",
	}
	fields_map := make(map[string]fields)
	fields_map["one"] = field_one
	fields_map["two"] = field_two

	dash_one := dashboard{
		format: "one",
		fields: fields_map,
	}
	dash_two := dashboard{
		format: "two",
		fields: fields_map,
	}
	dash_map := make(map[string]dashboard)
	dash_map["unos"] = dash_one
	dash_map["dos"] = dash_two
	config := configuration{
		url:        "uurrll",
		path:       "ppaatthh",
		separator:  ",",
		dashboards: dash_map,
	}

	b, err := toml.Marshal(config)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))

	fmt.Println("finished")

}
