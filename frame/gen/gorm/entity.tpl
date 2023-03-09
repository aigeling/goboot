package entity

import (
	"reflect"
	"strings"
    "{{GenImportTime .Fields}}"
	
    "{{.PackageBasePath}}/model"

	"github.com/aigeling/goboot/frame/web"
)

{{$objectName := .ObjectName}}

//{{.Remark}}  
type {{.ObjectName}} struct {

	web.PageInfo

    {{- range $k,$v := .Fields}} 
    //{{$v.Remark}}
    {{$v.FieldName}}  *{{$v.FieldType}}  `json:"{{$v.FieldName}},omitempty" {{if eq $v.Nullable false }}validate:"required"{{end}}`
    {{- end}}
}

// 转换为map格式，以tag中的json name为key，若json name为空则为字段名
func (o *{{.ObjectName}}) ToRespMap() map[string]any {

	t := reflect.TypeOf(*o)
	v := reflect.ValueOf(*o)
	m := make(map[string]any, t.NumField())
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		name := strings.Split(f.Tag.Get("json"), ",")[0]
		if name == "" {
			name = f.Name
		}

		if !v.Field(i).IsNil() {
			m[name] = v.Field(i).Elem().Interface()
		}
	}
	return m

}

// 转为map格式,以字段名为key
func (o *{{.ObjectName}}) ToMap() map[string]any{
	m := make(map[string]any)

    {{- range $k,$v := .Fields}} 
	if o.{{$v.FieldName}} != nil {
		m["{{$v.FieldName}}"] = *o.{{$v.FieldName}}
	}
    {{- end}}

	return m
}

//转为model格式
func (o *{{.ObjectName}}) ToModel() *model.{{.ObjectName}} {
	m := &model.{{.ObjectName}}{}

	 {{- range $k,$v := .Fields}} 
	if o.{{$v.FieldName}} != nil {
		m.{{$v.FieldName}} = *o.{{$v.FieldName}}
	}
    {{- end}}

	return m

}


{{range $k,$v := .Fields}} 
func (o *{{$objectName}}) Get{{$v.FieldName}}() {{$v.FieldType}}{
	return *o.{{$v.FieldName}}
}

{{end}}

{{range $k,$v := .Fields}} 
func (o *{{$objectName}}) Set{{$v.FieldName}}({{LowerFirstChar $v.FieldName}} {{$v.FieldType}}){
	o.{{$v.FieldName}} = &{{LowerFirstChar $v.FieldName}}
}
{{end}}