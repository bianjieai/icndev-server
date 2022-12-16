package service

import (
	"github.com/bianjieai/icndev-server/internal/app/model/entity"
	"github.com/bianjieai/icndev-server/internal/app/model/vo"
	"github.com/bianjieai/icndev-server/internal/app/repository/cache"
	"time"
)

type SubscribeService struct {
	subscribeRepo *cache.SubscribeEmailCacheRepo
}

func NewSubscribeService(subscribeRepo *cache.SubscribeEmailCacheRepo) *SubscribeService {
	return &SubscribeService{subscribeRepo: subscribeRepo}
}

func (svc *SubscribeService) SubscribeEmail(req *vo.CreateSubscribeEmailReq) error {
	var se entity.SubscribeEmail
	se.Email = req.Email
	se.CreateAt = time.Now().Unix()
	err := svc.subscribeRepo.Create(&se)
	return err
}
