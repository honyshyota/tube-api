package pgstore

import (
	"database/sql"
	"errors"

	"github.com/honyshyota/tube-api-go/internal/app/model"
)

type Repository struct {
	store *Store
}

func (r *Repository) CreateChannels(c *model.Channel) error {
	return r.store.db.QueryRow(
		"INSERT INTO channels (channel_id, channel_name, channel_info) VALUES ($1, $2, $3) RETURNING id",
		c.ChannelID,
		c.ChannelName,
		c.ChannelInfo,
	).Scan(&c.ID)
}

func (r *Repository) Find(id int) (*model.Channel, error) {
	v := &model.Channel{}
	if err := r.store.db.QueryRow(
		"SELECT id, channel_id, channel_name, channel_info FROM channels WHERE id = $1",
		id,
	).Scan(
		&v.ID,
		&v.ChannelID,
		&v.ChannelName,
		&v.ChannelInfo,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("record not found")
		}

		return nil, err
	}

	return v, nil
}

func (r *Repository) CreateVideos(v *model.Video) error {
	return r.store.db.QueryRow(
		"INSERT INTO videos (video_id, video_title, publish_date, video_info) VALUES ($1, $2, $3, $4) RETURNING id",
		v.VideoID,
		v.VideoTitle,
		v.PublishDate,
		v.Description,
	).Scan(&v.ID)
}
