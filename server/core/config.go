package core

import "github.com/spf13/viper"

type Config struct {
    Port string
}

func LoadConfig() (*Config, error) {
    viper.SetDefault("PORT", "3000")

    return &Config{
        Port: viper.GetString("PORT"),
    }, nil
}
