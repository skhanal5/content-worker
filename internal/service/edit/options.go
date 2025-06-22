package edit

type BackgroundType int

const (
    BlackScreen BackgroundType = iota
    BlurredVideo
    StaticImage
)

type Size struct {
    Width  int
    Height int
}


type EditOptions struct {
    Background     BackgroundType
    BgImagePath    string
    ForegroundSize Size
    Title         string
}

type Option func(*EditOptions)


func WithBackground(bg BackgroundType) Option {
	return func(o *EditOptions) {
		o.Background = bg
	}
}

func WithBgImage(path string) Option {
	return func(o *EditOptions) {
		o.BgImagePath = path
	}
}

func WithForegroundSize(width, height int) Option {
	return func(o *EditOptions) {
		o.ForegroundSize = Size{Width: width, Height: height}
	}
}

type TemplateType string

const (
    TemplateBlurred TemplateType = "blurred"
    TemplateBlack   TemplateType = "black"
    TemplateImage   TemplateType = "image"
)


func WithTemplate(t TemplateType) Option {
    return func(o *EditOptions) {
        switch t {
        case TemplateBlurred:
            o.Background = BlurredVideo
            o.ForegroundSize = Size{1080, 607}
        case TemplateBlack:
            o.Background = BlackScreen
            o.ForegroundSize = Size{1080, 607}
        case TemplateImage:
            o.Background = StaticImage
            o.ForegroundSize = Size{1080, 607}
        }
    }
}

func WithTitle(title string) Option {
    return func(o *EditOptions) {
        o.Title = title
    }
}