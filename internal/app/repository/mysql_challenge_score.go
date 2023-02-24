package repository

import (
	"github.com/bianjieai/icndev-server/internal/app/model/entity"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type IChallengeScoreRepo interface {
	FindByLimit(offset, limit int, useTotal bool) ([]*entity.ChallengeScore, int64, error)
}

type ChallengeScoreRepo struct {
	db *gorm.DB
}

func NewChallengeScoreRepo(db *gorm.DB) *ChallengeScoreRepo {
	return &ChallengeScoreRepo{db: db}
}

func (repo *ChallengeScoreRepo) FindByLimit(offset, limit int, useTotal bool) ([]*entity.ChallengeScore, int64, error) {
	var res []*entity.ChallengeScore
	var total int64
	var err error
	tx := repo.db.Table(entity.TableNameChallengeScore)
	if useTotal {
		err = tx.Session(&gorm.Session{}).Count(&total).Error
		if err != nil {
			logrus.Errorf("count challenge score failed, err:%s", err.Error())
			return nil, 0, err
		}
		return nil, total, nil
	}
	err = tx.Debug().Order("`rank`").Offset(offset).Limit(limit).Find(&res).Error
	return res, total, err
}
