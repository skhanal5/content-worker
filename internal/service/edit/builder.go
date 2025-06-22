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

    fgStream := input.
        Filter("scale", ffmpeg_go.Args{fmt.Sprintf("%d:%d", options.ForegroundSize.Width, options.ForegroundSize.Height)}).
        Filter("format", ffmpeg_go.Args{"yuv420p"})

	if options.Title != "" {
    	fgStream = fgStream.Filter("drawtext", buildTitleFilter(options.Title))
	}

	x := (CanvasSize.Width - options.ForegroundSize.Width) / 2
    y := (CanvasSize.Height - options.ForegroundSize.Height) / 2


 	output := ffmpeg_go.
        Filter([]*ffmpeg_go.Stream{bgStream, fgStream}, "overlay", ffmpeg_go.Args{
            fmt.Sprintf("x=%d", x),
            fmt.Sprintf("y=%d", y),
            "shortest=1",
        }).
        Output(outputPath, ffmpeg_go.KwArgs{"y": nil})

    return output, nil
}

func blackBackground() *ffmpeg_go.Stream {
    return ffmpeg_go.Input("color=black", ffmpeg_go.KwArgs{
        "f": "lavfi",
        "s": fmt.Sprintf("%dx%d", CanvasSize.Width, CanvasSize.Height),
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


func buildTitleFilter(title string) ffmpeg_go.Args {
    args := []string{
        fmt.Sprintf("text='%s'", escapeText(title)),
        "fontfile=font/Montserrat-Bold.ttf",
        "fontsize=72",
        "fontcolor=white",
        "x=(w-text_w)/2",
        "y=50",
        "borderw=10",
        "bordercolor=black",
    }

    return ffmpeg_go.Args{strings.Join(args, ":")}
}

func escapeText(text string) string {
    text = strings.ReplaceAll(text, `\`, `\\`)
    return strings.ReplaceAll(text, ":", `\:`)
}
