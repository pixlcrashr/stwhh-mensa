package logger

import "go.uber.org/zap"

func New() *zap.Logger {
	cfg := zap.NewProductionConfig()
	cfg.Level.SetLevel(zap.DebugLevel)

	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}

	return logger.Named("crawler")
}
