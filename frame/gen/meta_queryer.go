package gen

import (
	"errors"
	"fmt"

	"github.com/aigeling/goboot/frame/db"
	"github.com/aigeling/goboot/global"
)

// 初始化数据表的元数据，不同数据库需要不同的实现
type MetaQueryer interface {
	//查询表中的字段信息
	QueryMetaData(schema, tablename string) ([]Field, error)
}

var metaQueryers map[string]MetaQueryer = make(map[string]MetaQueryer)

func RegisteMetaQuery(dbCode string, metaQuery MetaQueryer) {
	metaQueryers[dbCode] = metaQuery
}

func InitMetaQueryers() {

	metaQueryers = make(map[string]MetaQueryer)
	var mq MetaQueryer
	for k, v := range global.Conf.DBConfigs {
		logger.Info("init gen db :%+v", v)
		if v.DBType == "postgresql" && v.Enabled {
			mq = NewInitMetaDataPostgres(db.GetDB(k))
		} else if v.DBType == "mysql" && v.Enabled {
			mq = NewInitMetaDataMysql(db.GetDB(k))
		}
		if mq != nil {
			metaQueryers[k] = mq
		}
	}
	fmt.Printf("mettaquerys : %v\n", metaQueryers)
}

// 根据指定的数据库代码返回数据库元数据生成器
func GetMetaQueryer(dbCode string) (m MetaQueryer, err error) {

	v, ok := metaQueryers[dbCode]

	if !ok {
		return nil, errors.New("No metaqueryer found:" + dbCode)
	}
	return v, nil
}

func mapTableFieldType(t *TableTemplate) error {

	dbCfg, ok := global.Conf.DBConfigs[t.DBCode]

	if !ok {
		return fmt.Errorf("db config not found : %s", t.DBCode)
	}
	fieldMap, ok := Conf.FieldTypeMap[dbCfg.DBType]

	if !ok {
		return fmt.Errorf("fieldType map not found : %s", dbCfg.DBType)
	}

	for i, f := range t.Fields {
		f.FieldName = toCamelName(f.ColumnName)
		f.FieldType, ok = fieldMap[f.DataType]
		if !ok {
			f.FieldType = "string"
		}
		t.Fields[i] = f
	}

	return nil
}
