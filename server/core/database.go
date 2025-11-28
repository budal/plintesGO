package core

import (
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

func Database() (*gorm.DB, error) {
    return gorm.Open(sqlite.Open("plintes.db"), &gorm.Config{})
}
