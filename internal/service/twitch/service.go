package twitch

import (
	"time"
	"github.com/go-resty/resty/v2"
)

type TwitchManager interface {
	GetClipInformation(clipSlug string) (*ClipMetadataResponse, error)
	GetUserClips(user string, limit int, filter string) (*UserClipsResponse, error)
}

const (
	gqlUrl = "https://gql.twitch.tv"
)


type TwitchService struct {
	gqlClient   *resty.Client
}

func NewTwitchService(helixClientId string, helixSecret string, gqlClientId string) *TwitchService {
	return &TwitchService{
		gqlClient:  resty.New().
			SetTimeout(10 * time.Second).
			SetBaseURL(gqlUrl).
			SetDebug(true).
			SetHeader("Client-ID", gqlClientId),
	}
}
