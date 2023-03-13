package cache

import (
	"github.com/bianjieai/icndev-server/internal/app/model/entity"
	"github.com/bianjieai/icndev-server/internal/app/repository"
)

type SpecialAwardsCacheRepo struct {
	dbr *repository.SpecialAwardsRepo
}

func NewSpecialAwardsCacheRepo(dbr *repository.SpecialAwardsRepo) *SpecialAwardsCacheRepo {
	return &SpecialAwardsCacheRepo{dbr: dbr}
}

func (repo *SpecialAwardsCacheRepo) FindAll() ([]*entity.SpecialAwards, int64, error) {
	return repo.dbr.FindAll()
}
