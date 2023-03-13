package repository

import (
	"github.com/bianjieai/icndev-server/internal/app/model/entity"
	"gorm.io/gorm"
)

type ISpecialAwardsRepo interface {
	FindAll() ([]*entity.SpecialAwards, int64, error)
}

type SpecialAwardsRepo struct {
	db *gorm.DB
}

func NewSpecialAwardsRepo(db *gorm.DB) *SpecialAwardsRepo {
	return &SpecialAwardsRepo{db: db}
}

func (repo *SpecialAwardsRepo) FindAll() ([]*entity.SpecialAwards, int64, error) {
	var res []*entity.SpecialAwards
	var total int64
	var err error
	tx := repo.db.Table(entity.TableNameSpecialAwards)
	err = tx.Count(&total).Find(&res).Error
	return res, total, err
}
