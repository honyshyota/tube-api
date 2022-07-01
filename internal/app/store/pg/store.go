package pgstore

import (
	"database/sql"

	"github.com/honyshyota/tube-api-go/internal/app/store"
)

type Store struct {
	db                *sql.DB
	ChannelRepository *Repository
}

func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) Repo() store.Repository {
	if s.ChannelRepository != nil {
		return s.ChannelRepository
	}

	s.ChannelRepository = &Repository{
		store: s,
	}

	return s.ChannelRepository
}
