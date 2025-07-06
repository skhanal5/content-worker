package activity

import (
	"clip-farmer-workflow/internal/service/edit"
	"context"
	"fmt"
	"os"
	"time"

	"go.temporal.io/sdk/activity"
)

type EditStyle string

const (
	BlurredOverlay          EditStyle = "blurred_overlay"
	BlurredOverlayStretched EditStyle = "blurred_overlay_stretched"
	BlackOverlay            EditStyle = "black_overlay"
	BlackOverlayStretched   EditStyle = "black_overlay_stretched"
	ImageOverlay            EditStyle = "image_overlay"
)

type EditVideoInput struct {
	InputPath  string
	OutputPath string
	Style      EditStyle
	Title      string
}

var styleOptionRegistry = map[EditStyle][]edit.Option{
	BlurredOverlay: {
		edit.WithTemplate(edit.TemplateBlurred),
	},
	BlurredOverlayStretched: {
		edit.WithTemplate(edit.TemplateBlurred),
		edit.WithForegroundSize(1080, 906),
	},
	BlackOverlay: {
		edit.WithTemplate(edit.TemplateBlack),
	},
	BlackOverlayStretched: {
		edit.WithTemplate(edit.TemplateBlack),
		edit.WithForegroundSize(1080, 906),
	},
	ImageOverlay: {
		edit.WithTemplate(edit.TemplateBlack),
	},
}

func (a *Activity) EditVideo(ctx context.Context, input EditVideoInput) error {
	logger := activity.GetLogger(ctx)
	logger.Info("Kicking off Edit Video Activity")
	config, ok := styleOptionRegistry[input.Style]
	if !ok {
		return fmt.Errorf("unknown style: %s", input.Style)
	}

	stop := make(chan struct{})
	go func() {
		ticker := time.NewTicker(10 * time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				activity.RecordHeartbeat(ctx)
				logger.Info("Heartbeat sent during render")
			case <-stop:
				return
			}
		}
	}()

	config = append(config, edit.WithTitle(input.Title))
	err := edit.Render(input.InputPath, input.OutputPath, config...)
	close(stop)

	if err != nil {
		return fmt.Errorf("edit Video Activity failed with err: %v", err)
	}

	logger.Info("Finished Edit Video Activity")
	return nil
}

func (a *Activity) DeleteTmpVideo(ctx context.Context, inputPath string) error {
	return os.Remove(inputPath)
}
