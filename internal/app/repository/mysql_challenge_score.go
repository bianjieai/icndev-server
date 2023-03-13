package repository

import (
	"github.com/bianjieai/icndev-server/internal/app/model/entity"
	"gorm.io/gorm"
)

type IChallengeScoreRepo interface {
	FindAll() ([]*entity.ChallengeScore, int64, error)
}

type ChallengeScoreRepo struct {
	db *gorm.DB
}

func NewChallengeScoreRepo(db *gorm.DB) *ChallengeScoreRepo {
	return &ChallengeScoreRepo{db: db}
}

func (repo *ChallengeScoreRepo) FindAll() ([]*entity.ChallengeScore, int64, error) {
	var res []*entity.ChallengeScore
	var total int64
	var err error
	tx := repo.db.Table(entity.TableNameChallengeScore)
	err = tx.Count(&total).Find(&res).Error
	return res, total, err
}
