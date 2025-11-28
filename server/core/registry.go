package core

import "go.uber.org/fx"

var registeredModules []fx.Option

func RegisterModule(m fx.Option) {
    registeredModules = append(registeredModules, m)
}

func GetModules() []fx.Option {
    return registeredModules
}
