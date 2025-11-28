package core

import (
    "go.uber.org/zap"
)

func Logger() (*zap.Logger, error) {
    return zap.NewProduction()
}
