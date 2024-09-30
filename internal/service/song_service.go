package service

import (
	"musiclib/database"
	"musiclib/models"
)

type SongService struct{
	db *database.DataBase
}

func NewSongService(db *database.DataBase) *SongService{
	return &SongService{
		db: db,
	}
}

func(s *SongService) GetSong(id int) (models.Song, error){
	return s.db.GetSong(id)
}

func(s *SongService) GetSongByName(name string, group string) (models.Song, error){
	return s.db.GetSongByName(name, group)
}

func(s *SongService) GetTextSong(id int) (string, error){
	return s.db.GetTextSong(id)
}

func(s *SongService) CreateSong(song models.Song) error{
	return s.db.CreateSong(song)
}

func(s *SongService) UpdateSong(id int, updateSong models.Song) error{
	return s.db.UpdateSong(id, updateSong)
}

func(s *SongService) DeleteSong(name string, group string) error{
	return s.db.DeleteSong(name, group)
}