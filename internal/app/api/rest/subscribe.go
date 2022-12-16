package rest

import (
	"github.com/bianjieai/icndev-server/internal/app/api/response"
	"github.com/bianjieai/icndev-server/internal/app/errors"
	"github.com/bianjieai/icndev-server/internal/app/model/vo"
	"github.com/bianjieai/icndev-server/internal/app/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type SubscribeController struct {
	SubscribeService *service.SubscribeService
}

func NewSubscribeController(subscribeService *service.SubscribeService) *SubscribeController {
	return &SubscribeController{SubscribeService: subscribeService}
}

// EmailCreate
// @Summary create email
// @Tags EmailCreate
// @Accept  json
// @Produce  json
// @Param	createSubscribeEmail	body	vo.CreateSubscribeEmailReq	true	"新增订阅邮箱"
// @Success 200 {object} []string "success"
// @Router /subscribe/email [post]
func (ctl *SubscribeController) EmailCreate(c *gin.Context) {
	var req vo.CreateSubscribeEmailReq
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusOK, response.BuildErr(errors.WrapBadRequest(err)))
		return
	}
	if err := ctl.SubscribeService.SubscribeEmail(&req); err != nil {
		if strings.Contains(err.Error(), "Duplicate") {
			c.JSON(http.StatusOK, response.BuildSubscribed(errors.Wrap(err)))
		} else {
			c.JSON(http.StatusOK, response.BuildErr(errors.Wrap(err)))
		}
		return
	}
	c.JSON(http.StatusOK, response.BuildWithMessage("Subscription Success", nil))
}
