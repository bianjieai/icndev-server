package initialization

import (
	"github.com/bianjieai/icndev-server/internal/app/api"
	"github.com/bianjieai/icndev-server/internal/app/api/middleware"
	"github.com/bianjieai/icndev-server/internal/app/api/rest"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Routers(Router *gin.Engine, ctl *rest.SubscribeController) {
	Router.Use(middleware.Cors())
	commonGroup := Router.Group("")
	api.InitSubscribeRouter(commonGroup, ctl)
	logrus.Debug("router register success")
}
