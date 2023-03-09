package service

import (
	"errors"

	"github.com/aigeling/goboot/internal/entity"
	"github.com/aigeling/goboot/internal/model"
	"gorm.io/gorm"
)

type officeService struct {
	db *gorm.DB
}

// 获得Service实例
func NewOfficeService(db *gorm.DB) *officeService {
	return &officeService{
		db: db,
	}
}

// 数据插入
func (s *officeService) Insert(en *entity.Office) error {
	return s.db.Create(en.ToModel()).Error
}

// 按主键更新
func (s *officeService) UpdateByPk(en *entity.Office) error {
	mo := en.ToMap()
	return s.db.Model(&model.Office{}).Where("id=?", mo["Id"]).Updates(mo).Error
}

// 按主键删除
func (s *officeService) DeleteByPk(en *entity.Office) error {
	mo := en.ToModel()
	return s.db.Where("id=?", mo.Id).Delete(&model.Office{}).Error
}

// 按主键查询一条记录
func (s *officeService) GetByPk(en *entity.Office) (*entity.Office, error) {
	mo := en.ToModel()
	rst := &entity.Office{}
	err := s.db.Model(&model.Office{}).Where("id=?", mo.Id).First(rst).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rst, nil
}

// 按条件查询数据列表
func (s *officeService) QueryList(en *entity.Office) ([]entity.Office, error) {

	rst := []entity.Office{}
	stmt := s.db.Model(&model.Office{})
	if en.Id != nil {
		stmt.Where("id=?", en.Id)
	}
	if en.Code != nil {
		stmt.Where("code=?", en.Code)
	}
	if en.ParentId != nil {
		stmt.Where("parent_id=?", en.ParentId)
	}
	if en.Sname != nil {
		stmt.Where("sname=?", en.Sname)
	}
	if en.Lname != nil {
		stmt.Where("lname=?", en.Lname)
	}
	if en.CreateAt != nil {
		stmt.Where("create_at=?", en.CreateAt)
	}
	if en.UpdateAt != nil {
		stmt.Where("update_at=?", en.UpdateAt)
	}
	err := stmt.Scan(&rst).Error
	if err != nil {
		return nil, err
	}
	return rst, nil
}

// 按条件查询数据列表
func (s *officeService) FindPage(en *entity.Office) (rst []map[string]any, count int64, err error) {

	stmt := s.db.Model(&model.Office{})
	if en.Id != nil {
		stmt.Where("id=?", en.Id)
	}
	if en.Code != nil {
		stmt.Where("code=?", en.Code)
	}
	if en.ParentId != nil {
		stmt.Where("parent_id=?", en.ParentId)
	}
	if en.Sname != nil {
		stmt.Where("sname=?", en.Sname)
	}
	if en.Lname != nil {
		stmt.Where("lname=?", en.Lname)
	}
	if en.CreateAt != nil {
		stmt.Where("create_at=?", en.CreateAt)
	}
	if en.UpdateAt != nil {
		stmt.Where("update_at=?", en.UpdateAt)
	}
	err = stmt.Count(&count).Error
	if err != nil {
		return
	}

	offset, limit := en.GetOffSetAndLimit()

	err = stmt.Offset(offset).Limit(limit).Scan(&rst).Error

	return
}

// 执行sql查询，返回map切片数据格式
func (s *officeService) ExecuteQuery(sql string, params ...any) (rst []map[string]any, err error) {
	rst = make([]map[string]any, 0, 10)
	err = s.db.Raw(sql, params...).Scan(&rst).Error
	return
}

// 执行sql更新
func (s *officeService) ExecuteUpdate(sql string, params ...any) error {
	return s.db.Raw(sql, params...).Error
}
