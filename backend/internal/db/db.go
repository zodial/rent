package db

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Open(dsn string) (*gorm.DB, error) {
	if dsn == "" {
		dsn = "./data/dev.sqlite"
	}
	dsn = fmt.Sprintf("%s?_foreign_keys=1", dsn)
	gdb, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	// Ensure foreign_keys pragma is enabled for SQLite
	_ = gdb.Exec("PRAGMA foreign_keys = ON;").Error
	return gdb, nil
}
