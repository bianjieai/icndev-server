package api

import (
	"github.com/bianjieai/icndev-server/internal/app/api/rest"
	"github.com/gin-gonic/gin"
)

func InitSubscribeRouter(r *gin.RouterGroup, ctl *rest.SubscribeController) {
	subscribeRouter := r.Group("/subscribe")

	{
		subscribeRouter.POST("/email", ctl.EmailCreate)
		subscribeRouter.GET("/challenge-score", ctl.ChallengeScore)
	}
}
