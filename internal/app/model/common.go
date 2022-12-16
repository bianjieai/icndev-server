package model

import (
	"github.com/bianjieai/icndev-server/internal/app/api/rest"
	"github.com/bianjieai/icndev-server/internal/app/repository"
	"github.com/bianjieai/icndev-server/internal/app/repository/cache"
	"github.com/bianjieai/icndev-server/internal/app/service"
)

type Repositories struct {
	SubscribeRepo *repository.SubscribeEmailRepo
}

type CacheRepositories struct {
	SubscribeCacheRepo *cache.SubscribeEmailCacheRepo
}

type Services struct {
	SubscribeService *service.SubscribeService
}

type Controllers struct {
	SubscribeController *rest.SubscribeController
}
