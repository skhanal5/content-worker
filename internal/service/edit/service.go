package edit

type EditManager interface {
	Render(inputPath string, outputPath string, strategy EditingStrategy) error
}

type EditService struct {
}


func (e EditService) Render(inputPath string, outputPath string, strategy EditingStrategy) error {
	return strategy.Process(inputPath, outputPath)
}
