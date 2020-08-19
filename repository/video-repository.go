package repository

import (
	"github.com/jatinsaini25/gin-http-poc/entity"
	"github.com/jinzhu/gorm"
)

type VideoRepository interface {
	Save(entity.Video)
	Update(entity.Video)
	Delete(entity.Video)
	FindAll() []entity.Video
	Close()
}

type database struct {
	Connection *gorm.DB
}

func NewVideoRepository() VideoRepository {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Error while connecting to database")
	}
	db.AutoMigrate(&entity.Person{}, &entity.Video{})
	return &database{
		Connection: db,
	}
}

func (db *database) Close() {
	err := db.Connection.Close()
	if err != nil {
		panic("Failed to close DB Connection")
	}
}

func (db *database) Save(video entity.Video) {
	db.Connection.Create(video)
}

func (db *database) Update(video entity.Video) {
	db.Connection.Save(&video)
}

func (db *database) Delete(video entity.Video) {
	db.Connection.Delete(&video)
}

func (db *database) FindAll() []entity.Video {
	var videos []entity.Video
	db.Connection.Set("gorm:auto_preload", true).Find(&videos)
	return videos
}
