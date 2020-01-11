package main

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type db struct {
	DB *gorm.DB
}

func createDB(dbDir, dbFile string) *db {
	err := os.MkdirAll(dbDir, os.ModePerm)
	if err != nil {
		panic("failed to create database directory:" + dbDir)
	}
	path := dbDir + dbFile
	sqlite, err := gorm.Open("sqlite3", path)
	if err != nil {
		panic("failed to connect database:" + path)
	}
	sqlite.AutoMigrate(&target{})
	return &db{
		DB: sqlite,
	}
}

func (db db) close() {
	db.DB.Close()
}

func (db db) createTarget(t *target) error {
	dbObj := db.DB.Create(t)
	return dbObj.Error
}

func (db db) deleteTarget(name string) {
	db.DB.Where("name = ?", name).Delete(&target{})
}

func (db db) listTarget() []string {
	names := make([]string, 0)
	db.DB.Model(&target{}).Pluck("name", &names)
	return names
}

func (db db) getTarget(name string) (target, error) {
	var t target
	dbObj := db.DB.Where("name = ?", name).First(&t)
	if err := dbObj.Error; err != nil {
		return target{}, err
	}
	return t, nil
}
