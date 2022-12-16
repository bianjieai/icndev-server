package cache

import (
	"github.com/bianjieai/icndev-server/internal/app/model/entity"
	"github.com/bianjieai/icndev-server/internal/app/repository"
)

type SubscribeEmailCacheRepo struct {
	dbr *repository.SubscribeEmailRepo
}

func NewSubscribeEmailCacheRepo(dbr *repository.SubscribeEmailRepo) *SubscribeEmailCacheRepo {
	return &SubscribeEmailCacheRepo{dbr: dbr}
}

func (repo *SubscribeEmailCacheRepo) Create(se *entity.SubscribeEmail) error {
	return repo.dbr.Create(se)
}
