package pgstore_test

import (
	"testing"

	"github.com/honyshyota/tube-api-go/internal/app/model"
	pgstore "github.com/honyshyota/tube-api-go/internal/app/store/pg"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestChannelRepository_CreateChannels(t *testing.T) {
	db, teardown := pgstore.TestDB(t, databaseURL)
	defer teardown("channels")

	s := pgstore.New(db)
	c := model.TestChannel(t)
	assert.NoError(t, s.Repo().CreateChannels(c))
	assert.NotNil(t, c)
}

func TestUserRepository_FindChannel(t *testing.T) {
	db, teardown := pgstore.TestDB(t, databaseURL)
	defer teardown("channels")

	s := pgstore.New(db)
	c1 := model.TestChannel(t)
	s.Repo().CreateChannels(c1)
	u2, err := s.Repo().FindChannel(c1.ID)
	assert.NoError(t, err)
	assert.NotNil(t, u2)
}

func TestChannelRepository_CreateVideos(t *testing.T) {
	db, teardown := pgstore.TestDB(t, databaseURL)
	defer teardown("videos")

	s := pgstore.New(db)
	v := model.TestVideo(t)
	assert.NoError(t, s.Repo().CreateVideos(v))
	assert.NotNil(t, v)
}

func TestChannelRepository_FindVideo(t *testing.T) {
	db, teardown := pgstore.TestDB(t, databaseURL)
	defer teardown("videos")

	s := pgstore.New(db)
	v1 := model.TestVideo(t)
	s.Repo().CreateVideos(v1)
	v2, err := s.Repo().FindVideo(v1.ID)
	assert.NoError(t, err)
	assert.NotNil(t, v2)
}

func TestChannelRepository_CreatePlaylist(t *testing.T) {
	db, teardown := pgstore.TestDB(t, databaseURL)
	defer teardown("playlist")

	s := pgstore.New(db)
	p := model.TestPlaylist(t)
	assert.NoError(t, s.Repo().CreatePlaylist(p))
	assert.NotNil(t, p)
}
