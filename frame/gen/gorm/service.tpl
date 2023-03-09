package service

import (
	"errors"
	
	"{{.PackageBasePath}}/model"
	"{{.PackageBasePath}}/entity"
	"gorm.io/gorm"
)
{{$serviceName := LowerFirstChar .ObjectName}}

type {{$serviceName}}Service struct {
	db *gorm.DB
}

// 获得Service实例
func New{{.ObjectName}}Service(db *gorm.DB) *{{$serviceName}}Service {
	return &{{$serviceName}}Service{
		db: db,
	}
}

// 数据插入
func (s *{{$serviceName}}Service) Insert(en *entity.{{.ObjectName}}) error {
	return s.db.Create(en.ToModel()).Error
}

// 按主键更新
func (s *{{$serviceName}}Service) UpdateByPk(en *entity.{{.ObjectName}}) error {
	mo := en.ToMap()
	return s.db.Model(&model.{{.ObjectName}}{}){{- range $k,$v := .PrimaryKey}}.Where("{{$v.ColumnName}}=?",mo["{{$v.FieldName}}"]){{end}}.Updates(mo).Error
}

// 按主键删除
func (s *{{$serviceName}}Service) DeleteByPk(en *entity.{{.ObjectName}}) error {
	mo := en.ToModel()
	return s.db{{- range $k,$v := .PrimaryKey}}.Where("{{$v.ColumnName}}=?",mo.{{$v.FieldName}}){{end}}.Delete(&model.{{.ObjectName}}{}).Error
}

// 按主键查询一条记录
func (s *{{$serviceName}}Service) GetByPk(en *entity.{{.ObjectName}}) (*entity.{{.ObjectName}},error) {
	mo := en.ToModel()
	rst := &entity.{{.ObjectName}}{}
	err := s.db.Model(&model.{{.ObjectName}}{}){{- range $k,$v := .PrimaryKey}}.Where("{{$v.ColumnName}}=?",mo.{{$v.FieldName}}){{end}}.First(rst).Error
	if err != nil{
		if errors.Is(err,gorm.ErrRecordNotFound){
			return nil,nil
		}
		return nil,err
	}
	return rst,nil
}

// 按条件查询数据列表
func (s *{{$serviceName}}Service) QueryList(en *entity.{{.ObjectName}}) ([]entity.{{.ObjectName}}, error) {

	rst := []entity.{{.ObjectName}}{}
	stmt := s.db.Model(&model.{{.ObjectName}}{})
	{{- range $k,$v := .Fields}}
	if en.{{$v.FieldName}} != nil {
		stmt.Where("{{$v.ColumnName}}=?",en.{{$v.FieldName}})
	}
	{{- end}}
	err := stmt.Scan(&rst).Error
	if err != nil {
		return nil, err
	}
	return rst, nil
}

// 按条件查询数据列表
func (s *{{$serviceName}}Service) FindPage(en *entity.{{.ObjectName}}) (rst []map[string]any, count int64, err error) {

	
	stmt := s.db.Model(&model.{{.ObjectName}}{})
	{{- range $k,$v := .Fields}}
	if en.{{$v.FieldName}} != nil {
		stmt.Where("{{$v.ColumnName}}=?",en.{{$v.FieldName}})
	}
	{{- end}}
	err = stmt.Count(&count).Error
	if err != nil {
		return
	}

	offset, limit := en.GetOffSetAndLimit()

	err = stmt.Offset(offset).Limit(limit).Scan(&rst).Error

	return
}

// 执行sql查询，返回map切片数据格式
func (s *{{$serviceName}}Service) ExecuteQuery(sql string, params ...any) (rst []map[string]any, err error) {
	rst = make([]map[string]any,0,10)
	err = s.db.Raw(sql, params...).Scan(&rst).Error
	return
}

// 执行sql更新
func (s *{{$serviceName}}Service) ExecuteUpdate(sql string, params ...any) error {
	return s.db.Raw(sql, params...).Error
}
