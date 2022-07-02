package model

import "testing"

func TestChannel(t *testing.T) *Channel {
	return &Channel{
		ID:          1,
		ChannelID:   "qwerty1234",
		ChannelName: "Some Name",
		ChannelInfo: "Some info",
	}
}

func TestVideo(t *testing.T) *Video {
	return &Video{
		ID:          1,
		VideoID:     "qwerty123456",
		VideoTitle:  "some title",
		PublishDate: "some date",
		Description: "some description",
	}
}

func TestPlaylist(t *testing.T) *Playlist {
	return &Playlist{
		ID:            1,
		PlaylistID:    "qwerty123456",
		PlaylistTitle: "some title",
		EmbededHTML:   "some html",
		VideoCount:    1,
	}
}
