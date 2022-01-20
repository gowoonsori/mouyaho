package persistence

import (
	"errors"
	"gorm.io/gorm"
	"likeIt/domain/badge"
	"strings"
)

type BadgeModel struct {
	gorm.Model
	Id        badge.BadgeId `gorm:"primaryKey;column:id;auto_increment;" json:"id"`
	Url       string        `gorm:"not null;column:url'" json:"url"`
	Encoded   string        `gorm:"index;not null;column:encoded'" json:"encoded"`
	LikeCount int           `gorm:"int;column:like_cnt" json:"like_cnt"`
	CreatedAt int64         `gorm:"autoCreateTime:milli"`
	UpdatedAt int64         `gorm:"autoUpdateTime:milli"`
}

var _ badge.Repository = &BadgeRepository{}

type BadgeRepository struct {
	db *gorm.DB
}

func (br BadgeRepository) Save(b *badge.Badge) (*badge.Badge, error) {
	bm := &BadgeModel{
		Url:       b.Url,
		Encoded:   b.Encoded,
		LikeCount: b.LikeCount,
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

func (br BadgeRepository) FindById(id badge.BadgeId) (*badge.Badge, error) {
	var bm BadgeModel

	DB := br.db.Select("b.*, cnt(l.id) as like_cnt").
		Table("badge as b").
		Where("b.id = ?", id).
		Joins("left join like as l on l.badge_id = b.id").Find(&bm)

	if DB.Error != nil {
		return nil, DB.Error
	}

	return badge.New(bm.Id, bm.Url, bm.Encoded, bm.LikeCount), nil
}
