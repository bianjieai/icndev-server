package repository

import (
	"github.com/bianjieai/icndev-server/internal/app/model/entity"
	"gorm.io/gorm"
)

type ISubscribeEmailRepo interface {
	Create(se *entity.SubscribeEmail) error
}

type SubscribeEmailRepo struct {
	db *gorm.DB
}

func NewSubscribeEmailRepo(db *gorm.DB) *SubscribeEmailRepo {
	return &SubscribeEmailRepo{db: db}
}

func (repo *SubscribeEmailRepo) Create(se *entity.SubscribeEmail) error {
	return repo.db.Create(se).Error
}
