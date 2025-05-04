package twitch

type Clip struct {
	ID              string `json:"id,omitempty"`
	URL             string `json:"url,omitempty"`
	EmbedURL        string `json:"embed_url,omitempty"`
	BroadcasterID   string `json:"broadcaster_id,omitempty"`
	BroadcasterName string `json:"broadcaster_name,omitempty"`
	CreatorID       string `json:"creator_id,omitempty"`
	CreatorName     string `json:"creator_name,omitempty"`
	VideoID         string `json:"video_id,omitempty"`
	GameID          string `json:"game_id,omitempty"`
	Language        string `json:"language,omitempty"`
	Title           string `json:"title,omitempty"`
	ViewCount       int    `json:"view_count,omitempty"`
	CreatedAt       string `json:"created_at,omitempty"`
	ThumbnailURL    string `json:"thumbnail_url,omitempty"`
	Duration        float32    `json:"duration,omitempty"`
	VodOffset       int    `json:"vod_offset,omitempty"`
	IsFeatured      bool   `json:"is_featured,omitempty"`
}

type Pagination struct {
	Cursor string `json:"cursor,omitempty"`
}

type ClipsResponse struct {
	Clips      []Clip     `json:"data,omitempty"`
	Pagination Pagination `json:"pagination,omitempty"`
}

type User struct {
	Id               string   `json:"id,omitempty"`
	Login            string   `json:"login,omitempty"`
	DisplayNmame             string   `json:"display_name,omitempty"`
	Type         string   `json:"broadcaster_language,omitempty"`
	BroadcasterType string   `json:"broadcaster_type,omitempty"`
	Description       string   `json:"description,omitempty"`
	ProfileImageURL   string   `json:"profile_image_url,omitempty"`
	OfflineImageURL    string   `json:"offline_image_url,omitempty"`
	ViewCount		  int      `json:"view_count,omitempty"`
	CreatedAt		 string   `json:"created_at,omitempty"`
}

type UsersResponse struct {
	Users      []User     `json:"data,omitempty"`
	Pagination Pagination `json:"pagination,omitempty"`
}

type AuthResponse struct {
	AccessToken string `json:"access_token,omitempty"`
	ExpiresIn   int    `json:"expires_in,omitempty"`
	TokenType   string `json:"token_type,omitempty"`
}
