package aconf
// 自动生成模板{{.StructName}}

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
    "math/rand"
	"sync"
)

 {{- range .Structs}}
type {{.StructName}} struct {
    {{- range .Fields}}
    {{- if eq .FieldName "Id" }}
    {{- else }}
    {{.FieldName}}  {{.FieldType}} `json:"{{.FieldJson}}"` // {{.FieldDesc}}
    {{- end }} {{- end }}
}
{{- end }}

type {{.StructName}} struct {
    {{- range .Fields}}
    {{- if eq .FieldType "bool" }}
    {{.FieldName}}  *{{.FieldType}} `json:"{{.FieldJson}}"` // {{.FieldDesc}}
    {{- else }}
    {{.FieldName}}  {{.FieldType}} `json:"{{.FieldJson}}"` // {{.FieldDesc}}
    {{- end }} {{- end }}
}

var {{.StructName}}Map map[{{.IDType}}]{{.StructName}}
var {{.StructName}}Ary []{{.StructName}}

func init{{.StructName}}() {
	v := viper.New()
	v.SetConfigFile("./conf/game/{{.StructName}}.json")
	v.SetConfigType("json")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		if err := v.Unmarshal(&{{.StructName}}Map); err != nil {
			fmt.Println(err)
		}
	})
	if err := v.Unmarshal(&{{.StructName}}Map); err != nil {
		panic(err)
	}
	{{.StructName}}Ary = make([]{{.StructName}}, 0, len({{.StructName}}Map))
	for _, item := range {{.StructName}}Map {
        {{.StructName}}Ary = append({{.StructName}}Ary, item)
    }
}

var load{{.StructName}} sync.Once

func Get{{.StructName}}Map() map[{{.IDType}}]{{.StructName}} {
	load{{.StructName}}.Do(init{{.StructName}})
	return {{.StructName}}Map
}

func Get{{.StructName}}Ary() []{{.StructName}} {
	load{{.StructName}}.Do(init{{.StructName}})
	return {{.StructName}}Ary
}

func Get{{.StructName}}ByID(id {{.IDType}}) ({{.StructName}}, bool) {
	load{{.StructName}}.Do(init{{.StructName}})
	item,ok := {{.StructName}}Map[id]
	return item,ok
}

func Get{{.StructName}}ByIndex(idx int) (item {{.StructName}},ok bool) {
	load{{.StructName}}.Do(init{{.StructName}})
	lens := len({{.StructName}}Ary)
	if lens <=0 || idx >= lens {
	    return
	}
	return {{.StructName}}Ary[idx],true
}

func GetRand{{.StructName}}() (item {{.StructName}},ok bool) {
	load{{.StructName}}.Do(init{{.StructName}})
	lens := len({{.StructName}}Ary)
	if lens <=0 {
	    return
	}
	return {{.StructName}}Ary[rand.Intn(lens)],true
}
