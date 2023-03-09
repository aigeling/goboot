package web

import "go.uber.org/zap"

var logger *zap.SugaredLogger = zap.L().Sugar()
