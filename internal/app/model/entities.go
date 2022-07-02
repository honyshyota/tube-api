package model

type Channel struct {
	ID          int    `json:"id"`
	ChannelID   string `json:"channel_id"`
	ChannelName string `json:"channel_name"`
	ChannelInfo string `json:"chanel_info"`
}

type Video struct {
	ID          int    `json:"id"`
	VideoID     string `json:"video_id"`
	VideoTitle  string `json:"video_title"`
	PublishDate string `json:"publish_at"`
	Description string `json:"description"`
}

type Playlist struct {
	ID            int    `json:"id"`
	PlaylistID    string `json:"playlist_id"`
	PlaylistTitle string `json:"play_title"`
	EmbededHTML   string `json:"embeded"`
	VideoCount    int    `json:"video_count"`
}
