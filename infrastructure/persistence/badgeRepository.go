package persistence

import (
	"errors"
	"gorm.io/gorm"
	"likeIt/domain"
	"likeIt/domain/badge"
	"strings"
)

type BadgeModel struct {
	gorm.Model
	Id        domain.BadgeId `gorm:"primaryKey;column:id;auto_increment;" json:"id"`
	Url       string         `gorm:"not null;column:url'" json:"url"`
	Encoded   string         `gorm:"index;not null;column:encoded'" json:"encoded"`
	CreatedAt int64          `gorm:"autoCreateTime:milli"`
	UpdatedAt int64          `gorm:"autoUpdateTime:milli"`
}

var _ domain.BadgeRepository = &BadgeRepository{}

type BadgeRepository struct {
	db *gorm.DB
}

func (br BadgeRepository) Save(b *domain.Badge) (*domain.Badge, error) {
	bm := &BadgeModel{
		Url:     b.Url,
		Encoded: b.Encoded,
	}

	err := br.db.Debug().Create(&bm).Error
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "Duplicate") {
			return nil, errors.New("badge already exists")
		}
		return nil, errors.New("db error")
	}

	return b, nil
}

func (br BadgeRepository) FindById(id domain.BadgeId) (*domain.Badge, error) {
	var bm BadgeModel

	DB := br.db.Debug().First(&bm, id)

	if DB.Error != nil {
		return nil, DB.Error
	}

	return badge.New(bm.Id, bm.Url, bm.Encoded), nil
}
