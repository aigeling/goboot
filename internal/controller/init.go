package controller

import (
	"github.com/aigeling/goboot/frame/utils"
	"go.uber.org/zap"
)

var logger *zap.SugaredLogger = zap.L().Sugar()

var binder *utils.GinBinder = utils.NewGinBinder()
