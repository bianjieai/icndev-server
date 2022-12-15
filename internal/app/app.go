package app

import (
	"github.com/bianjieai/icndev-server/configs"
	"github.com/bianjieai/icndev-server/internal/app/initialization"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Serve(cfg *configs.Config) {
	mysqlIcndev := initialization.NewMysqlDB(cfg.MysqlIcndev)
	repositories := initialization.NewRepositories(mysqlIcndev)

	cacheRepositories := initialization.NewCacheRepositories(repositories)
	services := initialization.NewServices(cacheRepositories)
	controllers := initialization.NewControllers(services)
	r := gin.Default()
	initialization.Routers(r, controllers.SubscribeController)
	logrus.Fatal(r.Run(cfg.App.Addr))
}
