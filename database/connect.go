package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetConnect() (*gorm.DB, error) {
	dsn := "host=localhost user=root password=123 dbname=quanlynhanvien port=5432 sslmode=disable TimeZone=Asia/Saigon"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return db, err
}
