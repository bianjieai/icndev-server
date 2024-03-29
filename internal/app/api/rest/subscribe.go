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
	c.JSON(http.StatusOK, response.BuildWithMessage("You are now subscribed.", nil))
}

func (ctl *SubscribeController) ChallengeScore(c *gin.Context) {
	var req vo.ChallengeScoreReq
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusOK, response.BuildErr(errors.Wrap(err)))
		return
	}

	res, total, err := ctl.SubscribeService.ChallengeScore(req)
	if err != nil {
		c.JSON(http.StatusOK, response.BuildErr(errors.Wrap(err)))
		return
	}
	if res == nil {
		c.JSON(http.StatusOK, response.BuildErr(errors.WrapDetail(errors.EC40004, "no data find")))
		return
	}
	c.JSON(http.StatusOK, response.BuildWithPage(res, *req.Page, int(total)))
}

func (ctl *SubscribeController) SpecialAwards(c *gin.Context) {
	res, err := ctl.SubscribeService.SpecialAwards()
	if err != nil {
		c.JSON(http.StatusOK, response.BuildErr(errors.Wrap(err)))
		return
	}

	c.JSON(http.StatusOK, response.Build(res))
}
