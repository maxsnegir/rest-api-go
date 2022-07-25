package server

import "github.com/sirupsen/logrus"

func NewLogger(logLevel string) (*logrus.Logger, error) {
	logger := logrus.New()
	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		return nil, err
	}
	logger.SetLevel(level)
	return logger, nil
}
