package service

import (
	"github.com/aigeling/goboot/internal/model"
	"gorm.io/gorm"
)

type roleService struct {
	db *gorm.DB
}

// 获得Service实例
func NewRoleService(db *gorm.DB) *roleService {
	return &roleService{
		db: db,
	}
}

// 数据插入
func (s *roleService) Insert(mo *model.Role) error {
	return s.db.Create(mo).Error
}

// 按主键更新
func (s *roleService) UpdateByPk(mo map[string]any) error {
	return s.db.Model(&model.Role{}).Where("id=?", mo["Id"]).Updates(mo).Error
}

// 按主键删除
func (s *roleService) DeleteByPk(mo *model.Role) error {
	return s.db.Where("id=?", mo.Id).Delete(&model.Role{}).Error
}

// 按主键查询一条记录
func (s *roleService) GetByPk(mo *model.Role) (*model.Role, error) {
	rst := &model.Role{}
	err := s.db.Model(&model.Role{}).Where("id=?", mo.Id).Scan(rst).Error
	if err != nil {
		return nil, err
	}
	return rst, nil
}

// 按条件查询数据列表
func (s *roleService) QueryList(mo map[string]any) ([]model.Role, error) {

	rst := []model.Role{}
	stmt := s.db.Model(&model.Role{})
	if mo["Id"] != nil {
		stmt.Where("id=?", mo["Id"])
	}
	if mo["Code"] != nil {
		stmt.Where("code=?", mo["Code"])
	}
	if mo["Name"] != nil {
		stmt.Where("name=?", mo["Name"])
	}
	if mo["GroupId"] != nil {
		stmt.Where("group_id=?", mo["GroupId"])
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
func (s *roleService) ExecuteQuery(sql string, params ...any) (rst map[string]any, err error) {
	rst = make(map[string]any)
	err = s.db.Raw(sql, params...).Scan(&rst).Error
	return
}

// 执行sql更新
func (s *roleService) ExecuteUpdate(sql string, params ...any) error {
	return s.db.Raw(sql, params...).Error
}
