package logger

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"path/filepath"
)

const (
	_defaultFormatterMode = 1
	_defaultOutputMode    = 1
	_defaultLevel         = 6
	_defaultDirectory     = ""
	_defaultFileName      = ""
)

type Logger struct {
	formatterMode int
	outputMode    int
	level         int
	directory     string
	filename      string

	*logrus.Logger
}

func New(opts ...Option) (*Logger, error) {
	mLog := &Logger{
		Logger:        logrus.New(),
		formatterMode: _defaultFormatterMode,
		outputMode:    _defaultOutputMode,
		level:         _defaultLevel,
		directory:     _defaultDirectory,
		filename:      _defaultFileName,
	}

	for _, opt := range opts {
		opt(mLog)
	}

	switch mLog.formatterMode {
	case 1:
		// jsonFormatter
		formatter := logrus.JSONFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
		}
		mLog.SetFormatter(&formatter)
		break
	default:
		formatter := new(logrus.TextFormatter)
		formatter.TimestampFormat = "2006-01-02 15:04:05"
		formatter.FullTimestamp = true
		mLog.SetFormatter(formatter)
	}
	switch mLog.outputMode {
	case 1:
		// log to file
		mLog.SetOutput(&lumberjack.Logger{
			Filename:   filepath.Join(mLog.directory, mLog.filename),
			MaxSize:    500, // megabytes
			MaxBackups: 3,
			MaxAge:     28,    //days
			Compress:   false, // disabled by default
		})
		break
	default:
		// log to the terminal
	}
	var level = logrus.PanicLevel
	switch mLog.level {
	case 1:
		level = logrus.FatalLevel
		break
	case 2:
		level = logrus.ErrorLevel
		break
	case 3:
		level = logrus.WarnLevel
		break
	case 4:
		level = logrus.InfoLevel
		break
	case 5:
		level = logrus.DebugLevel
		break
	case 6:
		level = logrus.TraceLevel
		break
	}
	mLog.SetLevel(level)

	return mLog, nil
}
