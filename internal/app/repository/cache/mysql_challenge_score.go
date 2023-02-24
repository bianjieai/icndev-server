package cache

import (
	"github.com/bianjieai/icndev-server/internal/app/model/entity"
	"github.com/bianjieai/icndev-server/internal/app/repository"
)

type ChallengeScoreCacheRepo struct {
	dbr *repository.ChallengeScoreRepo
}

func NewChallengeScoreCacheRepo(dbr *repository.ChallengeScoreRepo) *ChallengeScoreCacheRepo {
	return &ChallengeScoreCacheRepo{dbr: dbr}
}

func (repo *ChallengeScoreCacheRepo) FindByLimit(offset, limit int, useTotal bool) ([]*entity.ChallengeScore, int64, error) {
	return repo.dbr.FindByLimit(offset, limit, useTotal)
}
