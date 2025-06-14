package edit

import (
	"fmt"
	"log"
	"strings"

	ffmpeg_go "github.com/u2takey/ffmpeg-go"
)

type EditingStrategy interface {
	Process(input string, output string, title string) error
}

type BlurredOverlayStrategy struct {}

func (b *BlurredOverlayStrategy) Process(inputPath string, outputPath string, title string) error {
	log.Printf("Rendering video")
	input := ffmpeg_go.Input(inputPath)
	bg := input.
		Filter("scale", ffmpeg_go.Args{"1080:1920:force_original_aspect_ratio=increase"}).
		Filter("crop", ffmpeg_go.Args{"1080:1920"}).
		Filter("boxblur", ffmpeg_go.Args{"50"})

	fg := input.
		Filter("scale", ffmpeg_go.Args{"1080:607"})

	overlayed := ffmpeg_go.Filter(
		[]*ffmpeg_go.Stream{bg, fg},
		"overlay",
		ffmpeg_go.Args{"(W-w)/2:(H-h)/2"},
	)

	textOverlayed := overlayed.Filter("drawtext", ffmpeg_go.Args{
		fmt.Sprintf("text=%s", strings.ToUpper(title)),
		"fontfile=font/Montserrat-Bold.ttf",
		"fontsize=72",
		"fontcolor=white",
		"x=(w-text_w)/2",
		"y=550",
		"borderw=10",
		"bordercolor=black",
	})


	err := ffmpeg_go.Output([]*ffmpeg_go.Stream{textOverlayed}, outputPath,
		ffmpeg_go.KwArgs{
			"map":   "0:a",
			"c:a":   "copy",
			"shortest": "",
			"y":     "",
		}).
		OverWriteOutput().
		ErrorToStdOut().
		Run()


	if err != nil {
		return fmt.Errorf("blurred overlay processing failed: %w", err)
	}

	log.Printf("Finished rendering video")
	return nil
}