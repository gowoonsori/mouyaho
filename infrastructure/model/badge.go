package model

import (
	"gorm.io/gorm"
	"likeIt/domain/badge"
	"likeIt/domain/user"
)

type Badge struct {
	gorm.Model
	Id      badge.BadgeId `gorm:"primaryKey;column:id;auto_increment;" json:"id"`
	Url     string        `gorm:"not null;column:url'" json:"url"`
	Encoded string        `gorm:"index;not null;column:encoded'" json:"encoded"`
	Likers  []user.User   `gorm:"many2many:likes" json:"likers"`
}
