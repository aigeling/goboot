package rbac

import (
	"testing"

	"github.com/aigeling/goboot/frame/config"
	"github.com/aigeling/goboot/frame/db"
	"github.com/aigeling/goboot/global"
)

func TestCreateModel(t *testing.T) {

	config.InitConfig(global.Conf, "app")
	dbcfg := global.Conf.DBConfigs[global.Conf.DefaultDB]
	db, err := db.CreateDB(&dbcfg, true)
	if err != nil {
		t.Fatal(err)
	}

	CreateModel(db)
}
