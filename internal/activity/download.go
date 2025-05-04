package activity

import (
	"context"
	"path/filepath"
)

type DownloadClipInput struct {
	Streamer string
	ClipID  string
	ClipURL string
}

const (
	tmpDir = "tmp/creator_%s/id_%s.mp4"
)


func (a *Activity) DownloadClip(ctx context.Context, input DownloadClipInput) (string, error) {
	filepath := filepath.Join(tmpDir, input.Streamer, input.ClipID)
	output, err := a.DownloadManager.Download(input.ClipURL, filepath)
	if err != nil {
		return "", err
	}
	return output, nil
}