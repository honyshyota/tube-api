package store

import "github.com/honyshyota/tube-api-go/internal/app/model"

type Repository interface {
	CreateChannels(*model.Channel) error
	FindChannel(int) (*model.Channel, error)
	CreateVideos(*model.Video) error
	FindVideo(int) (*model.Video, error)
	CreatePlaylist(*model.Playlist) error
}
