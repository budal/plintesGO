package core

import "go.uber.org/fx"

var Module = fx.Options(
    fx.Provide(CreateServer),
    fx.Provide(LoadConfig),
    fx.Invoke(StartServer),
    fx.Invoke(RegisterRoutes),
    fx.Provide(Logger),
    fx.Provide(Database),
)
