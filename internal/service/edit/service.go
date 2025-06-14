package edit

type EditManager interface {
	Render(inputPath string, outputPath string, strategy EditingStrategy, title string) error
}

type EditService struct {
}


func (e EditService) Render(inputPath string, outputPath string, strategy EditingStrategy, title string) error {
	return strategy.Process(inputPath, outputPath, title)
}
