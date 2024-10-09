package service

import (
	"musiclib/database"
	"musiclib/models"
)

type Song interface{
	GetSong(id int) (models.Song, error)
	GetTextSong(id int) (string, error)
	CreateSong(models.Song) error
	UpdateSong(name string, group string, updateSong models.Song) error
	DeleteSong(name string, group string) error
	GetSongByName(name string, group string) (models.Song, error)
}

type Service struct{
	Song
}

func NewService(db *database.DataBase) *Service{
	return &Service{
		Song: NewSongService(db),
	}
}