package main

import (
    "go.uber.org/fx"
    "plintes/core"
    "plintes/modules/users"
)

func main() {
    fx.New(
        core.Module,
        users.Module,
    ).Run()
}
