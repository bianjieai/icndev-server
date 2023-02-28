package repository

import (
	"github.com/bianjieai/icndev-server/internal/app/model/entity"
	"gorm.io/gorm"
)

type IChallengeScoreRepo interface {
	FindByLimit(offset, limit int) ([]*entity.ChallengeScore, int64, error)
}

type ChallengeScoreRepo struct {
	db *gorm.DB
}

func NewChallengeScoreRepo(db *gorm.DB) *ChallengeScoreRepo {
	return &ChallengeScoreRepo{db: db}
}

func (repo *ChallengeScoreRepo) FindByLimit(offset, limit int) ([]*entity.ChallengeScore, int64, error) {
	var res []*entity.ChallengeScore
	var total int64
	var err error
	tx := repo.db.Table(entity.TableNameChallengeScore)
	err = tx.Count(&total).Order("`rank`, team_name DESC").Offset(offset).Limit(limit).Find(&res).Error
	return res, total, err
}
