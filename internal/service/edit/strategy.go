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

func wrapText(text string, maxLineLen int) string {
	words := strings.Fields(text)
	var lines []string
	var currentLine string

	for _, word := range words {
		if len(currentLine)+len(word)+1 > maxLineLen {
			lines = append(lines, currentLine)
			currentLine = word
		} else {
			if currentLine != "" {
				currentLine += " "
			}
			currentLine += word
		}
	}
	if currentLine != "" {
		lines = append(lines, currentLine)
	}

	return strings.Join(lines, "\n")
}

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

	wrappedTitle := wrapText(title, 25)
	lines := strings.Split(wrappedTitle, "\n")

	lineHeight := 80
	startY := 480

	currentStream := overlayed

	for i, line := range lines {
		yPos := startY + i*lineHeight
		// TODO: This can be a lot more elegant
		escapedLine := strings.ReplaceAll(line, ":", "\\:")
		escapedLine = strings.ReplaceAll(escapedLine, "\\", "\\\\")
		currentStream = currentStream.Filter("drawtext", ffmpeg_go.Args{
			fmt.Sprintf("text=%s", escapedLine),
			"fontfile=font/Montserrat-Bold.ttf",
			"fontsize=72",
			"fontcolor=white",
			"x=(w-text_w)/2",
			fmt.Sprintf("y=%d", yPos),
			"borderw=10",
			"bordercolor=black",
		})
	}



	err := ffmpeg_go.Output([]*ffmpeg_go.Stream{currentStream}, outputPath,
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