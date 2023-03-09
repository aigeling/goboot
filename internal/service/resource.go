package service

import (
	"github.com/aigeling/goboot/internal/model"
	"gorm.io/gorm"
)

type resourceService struct {
	db *gorm.DB
}

// 获得Service实例
func NewResourceService(db *gorm.DB) *resourceService {
	return &resourceService{
		db: db,
	}
}

// 数据插入
func (s *resourceService) Insert(mo *model.Resource) error {
	return s.db.Create(mo).Error
}

// 按主键更新
func (s *resourceService) UpdateByPk(mo map[string]any) error {
	return s.db.Model(&model.Resource{}).Where("id=?", mo["Id"]).Where("obj=?", mo["Obj"]).Where("act=?", mo["Act"]).Updates(mo).Error
}

// 按主键删除
func (s *resourceService) DeleteByPk(mo *model.Resource) error {
	return s.db.Where("id=?", mo.Id).Where("obj=?", mo.Obj).Where("act=?", mo.Act).Delete(&model.Resource{}).Error
}

// 按主键查询一条记录
func (s *resourceService) GetByPk(mo *model.Resource) (*model.Resource, error) {
	rst := &model.Resource{}
	err := s.db.Model(&model.Resource{}).Where("id=?", mo.Id).Where("obj=?", mo.Obj).Where("act=?", mo.Act).Scan(rst).Error
	if err != nil {
		return nil, err
	}
	return rst, nil
}

// 按条件查询数据列表
func (s *resourceService) QueryList(mo map[string]any) ([]model.Resource, error) {

	rst := []model.Resource{}
	stmt := s.db.Model(&model.Resource{})
	if mo["Id"] != nil {
		stmt.Where("id=?", mo["Id"])
	}
	if mo["Code"] != nil {
		stmt.Where("code=?", mo["Code"])
	}
	if mo["Name"] != nil {
		stmt.Where("name=?", mo["Name"])
	}
	if mo["Obj"] != nil {
		stmt.Where("obj=?", mo["Obj"])
	}
	if mo["Act"] != nil {
		stmt.Where("act=?", mo["Act"])
	}
	if mo["CreatedAt"] != nil {
		stmt.Where("created_at=?", mo["CreatedAt"])
	}
	if mo["UpdatedAt"] != nil {
		stmt.Where("updated_at=?", mo["UpdatedAt"])
	}
	err := stmt.Scan(&rst).Error
	if err != nil {
		return nil, err
	}
	return rst, nil
}

// 执行sql查询，返回map切片数据格式
func (s *resourceService) ExecuteQuery(sql string, params ...any) (rst map[string]any, err error) {
	rst = make(map[string]any)
	err = s.db.Raw(sql, params...).Scan(&rst).Error
	return
}

// 执行sql更新
func (s *resourceService) ExecuteUpdate(sql string, params ...any) error {
	return s.db.Raw(sql, params...).Error
}
