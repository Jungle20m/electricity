package logger

type Option func(*Logger)

func WithFormatterMode(formatterMode int) Option {
	return func(logger *Logger) {
		logger.formatterMode = formatterMode
	}
}

func WithOutputMode(outputMode int) Option {
	return func(logger *Logger) {
		logger.outputMode = outputMode
	}
}

func WithLevel(level int) Option {
	return func(logger *Logger) {
		logger.level = level
	}
}

func WithDirectory(directory string) Option {
	return func(logger *Logger) {
		logger.directory = directory
	}
}

func WithFileName(filename string) Option {
	return func(logger *Logger) {
		logger.filename = filename
	}
}
