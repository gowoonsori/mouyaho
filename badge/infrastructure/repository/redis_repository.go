package repository

import (
	"context"
	"github.com/go-redis/redis/v8"
	"likeIt/domain"
)

type RedisBadgeRepository struct {
	client redis.Client
}

func (r RedisBadgeRepository) Save(b *domain.Badge) (*domain.Badge, error) {
	err := r.client.Set(context.Background(), string(b.Id()), b.File(), 60*60*2).Err()
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (r RedisBadgeRepository) FindById(id domain.BadgeId) (*domain.Badge, error) {
	svg, err := r.client.Get(context.Background(), string(id)).Result()
	if err != redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return domain.NewBadge(id, []byte(svg)), nil
}
