package download

import (
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
)

const (
	tmpDir = "tmp/id_%s.mp4"
)

type DownloadManager interface {
	Download(url string, outputId string) (string, error)
}

type DownloadService struct {
	client *resty.Client
}

func (d *DownloadService) Download(url string, outputId string) (string, error) {
	outputFile := fmt.Sprintf(tmpDir, outputId)
	
	resp, err := d.client.R().
	SetOutput(outputFile).
	Get(url)

	if err != nil {
		return "", err
	}
	if resp.StatusCode() != 200 {
		return "", fmt.Errorf("failed to download file: %s", resp.Status())
	}
	return outputFile, nil
}

func NewDownloadService() *DownloadService {
	return &DownloadService{
		client: resty.New().
		SetTimeout(30 * time.Second),
	}
}