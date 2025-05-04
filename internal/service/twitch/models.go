package twitch

type Clip struct {
	ID        string `json:"id,omitempty"`
	URL       string `json:"url,omitempty"`
	EmbedURL  string `json:"embed_url,omitempty"`
	BroadcasterID string `json:"broadcaster_id,omitempty"`
	BroadcasterName string `json:"broadcaster_name,omitempty"`
	CreatorID string `json:"creator_id,omitempty"`
	CreatorName string `json:"creator_name,omitempty"`
	VideoID   string `json:"video_id,omitempty"`
	GameID    string `json:"game_id,omitempty"`
	Language  string `json:"language,omitempty"`
	Title     string `json:"title,omitempty"`
	ViewCount int    `json:"view_count,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	ThumbnailURL string `json:"thumbnail_url,omitempty"`
	Duration  int    `json:"duration,omitempty"`
	VodOffset int    `json:"vod_offset,omitempty"`
	IsFeatured bool   `json:"is_featured,omitempty"`
}

type Pagination struct {
	Cursor string `json:"cursor,omitempty"`
}

type ClipsResponse struct {
	Clips []Clip `json:"data,omitempty"`
	Pagination Pagination `json:"pagination,omitempty"`	
}

type User struct {
	BroadcasterID string `json:"broadcaster_id,omitempty"`
	BroadcasterLogin string `json:"broadcaster_login,omitempty"`
	BroadcasterName string `json:"broadcaster_name,omitempty"`
	BroadcasterLanguage string `json:"broadcaster_language,omitempty"`
	GameID string `json:"game_id,omitempty"`
	GameName string `json:"game_name,omitempty"`
	Title string `json:"title,omitempty"`
	Delay int `json:"delay,omitempty"`
	Tags []string `json:"tags,omitempty"`
	ContentClassificationLabels []string `json:"content_classification_labels,omitempty"`
	IsBrandedContent bool `json:"is_branded_content,omitempty"`
}

type UsersResponse struct {
	Users []User `json:"data,omitempty"`
	Pagination Pagination `json:"pagination,omitempty"`
}

type AuthResponse struct {
	AccessToken string `json:"access_token,omitempty"`
	ExpiresIn int `json:"expires_in,omitempty"`
	TokenType string `json:"token_type,omitempty"`
}