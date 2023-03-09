package service

import (
	"errors"

	"github.com/aigeling/goboot/internal/entity"
	"github.com/aigeling/goboot/internal/model"
	"gorm.io/gorm"
)

type userService struct {
	db *gorm.DB
}

// 获得Service实例
func NewUserService(db *gorm.DB) *userService {
	return &userService{
		db: db,
	}
}

// 数据插入
func (s *userService) Insert(en *entity.User) error {
	return s.db.Create(en.ToModel()).Error
}

// 按主键更新
func (s *userService) UpdateByPk(en *entity.User) error {
	mo := en.ToMap()
	return s.db.Model(&model.User{}).Where("id=?", mo["Id"]).Updates(mo).Error
}

// 按主键删除
func (s *userService) DeleteByPk(en *entity.User) error {
	mo := en.ToModel()
	return s.db.Where("id=?", mo.Id).Delete(&model.User{}).Error
}

// 按主键查询一条记录
func (s *userService) GetByPk(en *entity.User) (*entity.User, error) {
	mo := en.ToModel()
	rst := &entity.User{}
	err := s.db.Model(&model.User{}).Where("id=?", mo.Id).First(rst).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rst, nil
}

// 按条件查询数据列表
func (s *userService) QueryList(en *entity.User) ([]entity.User, error) {

	rst := []entity.User{}
	stmt := s.db.Model(&model.User{})
	if en.Id != nil {
		stmt.Where("id=?", en.Id)
	}
	if en.Code != nil {
		stmt.Where("code=?", en.Code)
	}
	if en.LoginName != nil {
		stmt.Where("login_name=?", en.LoginName)
	}
	if en.Password != nil {
		stmt.Where("password=?", en.Password)
	}
	if en.Name != nil {
		stmt.Where("name=?", en.Name)
	}
	if en.GroupId != nil {
		stmt.Where("group_id=?", en.GroupId)
	}
	if en.CreatedAt != nil {
		stmt.Where("created_at=?", en.CreatedAt)
	}
	if en.UpdatedAt != nil {
		stmt.Where("updated_at=?", en.UpdatedAt)
	}
	err := stmt.Scan(&rst).Error
	if err != nil {
		return nil, err
	}
	return rst, nil
}

// 按条件查询数据列表
func (s *userService) FindPage(en *entity.User) (rst []map[string]any, count int64, err error) {

	stmt := s.db.Model(&model.User{})
	if en.Id != nil {
		stmt.Where("id=?", en.Id)
	}
	if en.Code != nil {
		stmt.Where("code=?", en.Code)
	}
	if en.LoginName != nil {
		stmt.Where("login_name=?", en.LoginName)
	}
	if en.Password != nil {
		stmt.Where("password=?", en.Password)
	}
	if en.Name != nil {
		stmt.Where("name=?", en.Name)
	}
	if en.GroupId != nil {
		stmt.Where("group_id=?", en.GroupId)
	}
	if en.CreatedAt != nil {
		stmt.Where("created_at=?", en.CreatedAt)
	}
	if en.UpdatedAt != nil {
		stmt.Where("updated_at=?", en.UpdatedAt)
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
func (s *userService) ExecuteQuery(sql string, params ...any) (rst []map[string]any, err error) {
	rst = make([]map[string]any, 0, 10)
	err = s.db.Raw(sql, params...).Scan(&rst).Error
	return
}

// 执行sql更新
func (s *userService) ExecuteUpdate(sql string, params ...any) error {
	return s.db.Raw(sql, params...).Error
}

// // 按条件查询数据列表
// func (s *userService) FindPage(en *web.PageInfo[entity.User]) (rst []map[string]any, count int64, err error) {

// 	stmt := s.db.Model(&model.User{})
// 	if en.QueryParam.Id != nil {
// 		stmt.Where("id=?", en.QueryParam.Id)
// 	}
// 	if en.QueryParam.Code != nil {
// 		stmt.Where("code=?", en.QueryParam.Code)
// 	}
// 	if en.QueryParam.LoginName != nil {
// 		stmt.Where("login_name=?", en.QueryParam.LoginName)
// 	}
// 	if en.QueryParam.Password != nil {
// 		stmt.Where("password=?", en.QueryParam.Password)
// 	}
// 	if en.QueryParam.Name != nil {
// 		stmt.Where("name=?", en.QueryParam.Name)
// 	}
// 	if en.QueryParam.GroupId != nil {
// 		stmt.Where("group_id=?", en.QueryParam.GroupId)
// 	}
// 	if en.QueryParam.CreatedAt != nil {
// 		stmt.Where("created_at=?", en.QueryParam.CreatedAt)
// 	}
// 	if en.QueryParam.UpdatedAt != nil {
// 		stmt.Where("updated_at=?", en.QueryParam.UpdatedAt)
// 	}
// 	err = stmt.Count(&count).Error
// 	if err != nil {

// 		return
// 	}

// 	offset, limit := en.GetOffSetAndLimit()

// 	err = stmt.Offset(offset).Limit(limit).Scan(&rst).Error

// 	return
// }
