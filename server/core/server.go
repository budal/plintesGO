package core

import (
    "context"

    "github.com/gofiber/fiber/v2"
    "go.uber.org/fx"
    "go.uber.org/zap"
)

type FiberServer struct {
    App *fiber.App
}

func CreateServer(log *zap.Logger) (*FiberServer, error) {
    app := fiber.New()
    log.Info("Plintes initialized")
    return &FiberServer{App: app}, nil
}

func StartServer(lc fx.Lifecycle, s *FiberServer, cfg *Config, log *zap.Logger) {
    lc.Append(fx.Hook{
        OnStart: func(context.Context) error {
            go func() {
                log.Info("Starting Plintes server", zap.String("port", cfg.Port))
                s.App.Listen(":" + cfg.Port)
            }()
            return nil
        },
        OnStop: func(context.Context) error {
            log.Info("Shutting down Plintes server")
            return s.App.Shutdown()
        },
    })
}
