package edit

func Render(inputPath, outputPath string, opts ...Option) error {
	options := &EditOptions{}
	for _, opt := range opts {
		opt(options)
	}

	cmd, err := buildFFmpegCommand(inputPath, outputPath, options)
	if err != nil {
		return err
	}

	return cmd.
		OverWriteOutput().
		ErrorToStdOut().
		Run()
}
