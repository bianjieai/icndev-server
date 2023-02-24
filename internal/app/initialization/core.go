package initialization

import (
	"github.com/bianjieai/icndev-server/internal/app/api/rest"
	"github.com/bianjieai/icndev-server/internal/app/model"
	"github.com/bianjieai/icndev-server/internal/app/repository"
	"github.com/bianjieai/icndev-server/internal/app/repository/cache"
	"github.com/bianjieai/icndev-server/internal/app/service"
	"gorm.io/gorm"
)

func NewRepositories(mysqlIcndev *gorm.DB) *model.Repositories {
	return &model.Repositories{
		SubscribeRepo: repository.NewSubscribeEmailRepo(mysqlIcndev),
		ChallengeRepo: repository.NewChallengeScoreRepo(mysqlIcndev),
	}
}

func NewCacheRepositories(r *model.Repositories) *model.CacheRepositories {
	return &model.CacheRepositories{
		SubscribeCacheRepo: cache.NewSubscribeEmailCacheRepo(r.SubscribeRepo),
		ChallengeCacheRepo: cache.NewChallengeScoreCacheRepo(r.ChallengeRepo),
	}
}

func NewServices(r *model.CacheRepositories) *model.Services {
	return &model.Services{
		SubscribeService: service.NewSubscribeService(r.SubscribeCacheRepo, r.ChallengeCacheRepo),
	}
}

func NewControllers(s *model.Services) *model.Controllers {
	return &model.Controllers{
		SubscribeController: rest.NewSubscribeController(s.SubscribeService),
	}
}
