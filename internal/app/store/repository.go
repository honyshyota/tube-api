package store

import "github.com/honyshyota/tube-api-go/internal/app/model"

type Repository interface {
	CreateChannels(*model.Channel) error
	Find(int) (*model.Channel, error)
	CreateVideos(*model.Video) error
}
