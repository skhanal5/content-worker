package edit

import (
	"fmt"
	ffmpeg_go "github.com/u2takey/ffmpeg-go"
)

type EditingStrategy interface {
	Process(input string, output string) error
}

type BlurredOverlayStrategy struct {}

func (b *BlurredOverlayStrategy) Process(inputPath string, outputPath string) error {
	
	input := ffmpeg_go.Input(inputPath)

	// Build filter steps
	bg := input.
		Filter("scale", ffmpeg_go.Args{"1080:1920:force_original_aspect_ratio=increase"}).
		Filter("crop", ffmpeg_go.Args{"1080:1920"}).
		Filter("boxblur", ffmpeg_go.Args{"50"})

	fg := input.
		Filter("scale", ffmpeg_go.Args{"1080:607"})

	// Combine background and foreground with overlay
	overlayed := ffmpeg_go.Filter(
		[]*ffmpeg_go.Stream{bg, fg},
		"overlay",
		ffmpeg_go.Args{"(W-w)/2:(H-h)/2"},
	)

	// Output the result, preserving audio
	err := ffmpeg_go.Output([]*ffmpeg_go.Stream{overlayed}, outputPath,
		ffmpeg_go.KwArgs{
			"map":   "0:a", // take audio from first input
			"c:a":   "copy",
			"shortest": "", // prevents hanging if audio is shorter/longer
			"y":     "",    // overwrite
		}).
		OverWriteOutput().
		ErrorToStdOut().
		Run()


	if err != nil {
		return fmt.Errorf("blurred overlay processing failed: %w", err)
	}

	return nil
}