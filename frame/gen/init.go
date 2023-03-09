package gen

import (
	"github.com/aigeling/goboot/frame/config"
	"go.uber.org/zap"
)

var logger *zap.SugaredLogger = zap.L().Sugar()

type GenConf struct {
	FieldTypeMap map[string]map[string]string `json:"fieldTypeMap"`
}

var Conf *GenConf = new(GenConf)

func InitGen() {
	config.InitConfig(Conf, "gen")
	InitMetaQueryers()
}
