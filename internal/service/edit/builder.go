package edit

import (
	"fmt"
	"strings"

	ffmpeg_go "github.com/u2takey/ffmpeg-go"
)

var CanvasSize = Size{Width: 1080, Height: 1920}


func buildFFmpegCommand(inputPath, outputPath string, options *EditOptions) (*ffmpeg_go.Stream, error) {
	input := ffmpeg_go.Input(inputPath)

    var bgStream *ffmpeg_go.Stream

    switch options.Background {
		case BlackScreen:
			bgStream = blackBackground()
		case BlurredVideo:
			bgStream = blurredBackground(inputPath)
		case StaticImage:
			if options.BgImagePath == "" {
				return nil, fmt.Errorf("BgImagePath must be set for StaticImage background")
			}
			bgStream = imageBackground(options.BgImagePath)
		default:
			return nil, fmt.Errorf("unsupported background type: %v", options.Background)
	}

    if options.Title != "" {
        filterArgs := buildTitleArgs(options.Title)
        for _, args := range filterArgs {
            bgStream = bgStream.Filter("drawtext", args)
        }
    }

    fgStream := input.
        Filter("scale", ffmpeg_go.Args{fmt.Sprintf("%d:%d", options.ForegroundSize.Width, options.ForegroundSize.Height)}).
        Filter("format", ffmpeg_go.Args{"yuv420p"})


	x := (CanvasSize.Width - options.ForegroundSize.Width) / 2
    y := (CanvasSize.Height - options.ForegroundSize.Height) / 2


 	output := ffmpeg_go.
        Filter([]*ffmpeg_go.Stream{bgStream, fgStream}, "overlay", ffmpeg_go.Args{
            fmt.Sprintf("x=%d", x),
            fmt.Sprintf("y=%d", y),
            "shortest=1",
        }).
        Output(outputPath)

    return output, nil
}

func blackBackground() *ffmpeg_go.Stream {
    return ffmpeg_go.Input(fmt.Sprintf("color=size=%dx%d:color=black", CanvasSize.Width, CanvasSize.Height), ffmpeg_go.KwArgs{
        "f": "lavfi",
    })
}

func blurredBackground(inputPath string) *ffmpeg_go.Stream {
    input := ffmpeg_go.Input(inputPath)

    blurred := input.
        Filter("scale", ffmpeg_go.Args{
            fmt.Sprintf("%d:%d:force_original_aspect_ratio=increase", CanvasSize.Width, CanvasSize.Height),
        }).
        Filter("boxblur", ffmpeg_go.Args{"50"}).
        Filter("crop", ffmpeg_go.Args{
            fmt.Sprintf("%d:%d", CanvasSize.Width, CanvasSize.Height),
        })

    return blurred
}


func imageBackground(imagePath string) *ffmpeg_go.Stream {
    input := ffmpeg_go.Input(imagePath)

    scaledAndCropped := input.
        Filter("scale", ffmpeg_go.Args{
            fmt.Sprintf("%d:%d:force_original_aspect_ratio=increase", CanvasSize.Width, CanvasSize.Height),
        }).
        Filter("crop", ffmpeg_go.Args{
            fmt.Sprintf("%d:%d", CanvasSize.Width, CanvasSize.Height),
        })

    return scaledAndCropped
}

func splitTextIntoLines(text string, maxWidth int) []string {
    uppercaseText := strings.ToUpper(text)
    words := strings.Fields(uppercaseText)
    var lines []string
    var currentLine string

    for _, word := range words {
        if len(currentLine)+len(word)+1 > maxWidth { 
            if currentLine != "" {
                lines = append(lines, currentLine)
            }
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

    return lines
}

func buildTitleArgs(title string) []ffmpeg_go.Args {
    const charactersPerLine = 20
    const startY = 480
    const lineHeight = 80

    lines := splitTextIntoLines(title, charactersPerLine)
    var argsList []ffmpeg_go.Args
    for i, line := range lines {
        escapedLine := line
        args := []string{
            fmt.Sprintf("text='%s'", escapedLine),
            "fontfile=font/Montserrat-Bold.ttf",
            "fontsize=72",
            "fontcolor=white",
            "x=(w-text_w)/2",
            fmt.Sprintf("y=%d", startY+i*lineHeight),
            "borderw=10",
            "bordercolor=black",
        }
        argsList = append(argsList, ffmpeg_go.Args{strings.Join(args, ":")})
    }

    return argsList
}