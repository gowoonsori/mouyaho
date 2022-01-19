package persistence

import (
	"errors"
	"gorm.io/gorm"
	"likeIt/domain/badge"
	"likeIt/domain/user"
	"likeIt/infrastructure/model"
	"strings"
)

var _ badge.Repository = &BadgeRepository{}

type BadgeRepository struct {
	db *gorm.DB
}

func (br BadgeRepository) Save(b *badge.Badge) (*badge.Badge, error) {
	err := br.db.Debug().Create(&b).Error
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "Duplicate") {
			return nil, errors.New("badge already exists")
		}
		return nil, errors.New("db error")
	}

	return b, nil
}

func (br BadgeRepository) FindById(id badge.BadgeId) (*badge.Badge, error) {
	var b model.Badge

	err := br.db.Debug().Preload("Users").Where("id = ?", id).Take(&b).Error
	if err != nil {
		return nil, err
	}

	var likers []user.UserId
	for _, val := range b.Likers {
		likers = append(likers, val.GetId())
	}

	return badge.NewBadge(b.Id, b.Url, b.Encoded, likers), nil
}
