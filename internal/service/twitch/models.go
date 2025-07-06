package twitch

type ClipMetadataResponse struct {
	Data ClipData `json:"data"`
	Extensions Extensions `json:"extensions"`
	Errors []Error `json:"errors,omitempty"`
}

type Error struct {
	Message string `json:"message"`
}

type Extensions struct {
	DurationMilliseconds int `json:"durationMilliseconds"`
	OperationName string `json:"operationName"`
	RequestID string `json:"requestID"`
}

type ClipData struct {
	Clip ClipMetadata `json:"clip"`
}

type ClipMetadata struct {
	Slug                string               `json:"slug"`
	PlaybackAccessToken  PlaybackAccessToken  `json:"playbackAccessToken"`
	VideoQualities       []VideoQuality       `json:"videoQualities"`
}

type PlaybackAccessToken struct {
	Signature string `json:"signature"`
	Value     string `json:"value"`
}

type VideoQuality struct {
	Quality   string  `json:"quality"`
	FrameRate float64 `json:"frameRate"`
	SourceURL string  `json:"sourceURL"`
}

type UserClipsResponse struct {
	Data UserClipsData `json:"data"`
	Extensions Extensions `json:"pagination"`
	Errors []Error `json:"errors,omitempty"`
}


type UserClipsData struct {
	User UserClipsInfo `json:"user"`
}

type UserClipsInfo struct {
	ID    string `json:"id"`
	Clips UserClips `json:"clips"`
	TypeName string `json:"__typename"`
}

type UserClips struct {
	PageInfo PageInfo `json:"pageInfo"`
	Edges []UserClipEdge `json:"edges"`
}

type PageInfo struct {
	HasNextPage bool   `json:"hasNextPage"`
}

type UserClipEdge struct {
	Cursor string `json:"cursor"`
	Node UserClipNode `json:"node"`
	TypeName string `json:"__typename"`
}

type UserClipNode struct {
	ID          string `json:"id"`
	Slug 	  string `json:"slug"`
	Url 	   string `json:"url"`
	EmbedURL string `json:"embedURL"`
	Title       string `json:"title"`
	ViewCount   int    `json:"viewCount"`
	Language    string `json:"language"`
	Curator ClipCurator `json:"curator"`
	Game ClipGame `json:"game"`
	Broadcaster ClipBroadcaster `json:"broadcaster"`
	ThumbnailURL string `json:"thumbnailURL"`
	CreatedAt   string `json:"createdAt"`
	DurationSeconds int `json:"durationSeconds"`
	ChampBadge string `json:"champBadge"`
	IsFeatured bool `json:"isFeatured"`
	GuestStarParticipants map[string]interface{} `json:"guestStarParticipants"`
	TypeName string `json:"__typename"`
}

type ClipCurator struct {
	ID        string `json:"id"`
	Login     string `json:"login"`
	DisplayName string `json:"displayName"`
	TypeName string `json:"__typename"`
}

type ClipGame struct {
	ID        string `json:"id"`
	Slug 	string `json:"slug"`
	Name      string `json:"name"`
	BoxArtURL string `json:"boxArtURL"`
	TypeName string `json:"__typename"`
}

type ClipBroadcaster struct {
	ID        string `json:"id"`
	Login     string `json:"login"`
	DisplayName string `json:"displayName"`
	ProfileImageURL string `json:"profileImageURL"`
	PrimaryColorHex string `json:"primaryColorHex"`
	Roles struct {
		IsPartner bool `json:"isPartner"`
		TypeName string `json:"__typename"`
	}
	TypeName string `json:"__typename"`
}