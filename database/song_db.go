package database

import (
	"fmt"
	"musiclib/models"

	_ "github.com/lib/pq"
)

func(DB *DataBase) GetSong(id int) (models.Song, error){
	req := fmt.Sprintf("SELECT * FROM songs WHERE ID=%d", id)

	var song models.Song

	rows ,err := DB.db.Query(req)
	if err != nil{
		return song, err
	}

	err = rows.Scan(&song.Id, &song.Song, &song.Group)
	return song, err
}

func(DB *DataBase) GetTextSong(id int) (string, error){
	req := fmt.Sprintf("SELECT text FROM songs WHERE ID=%d", id)

	var text string

	rows ,err := DB.db.Query(req)
	if err != nil{
		return text, err
	}

	err = rows.Scan(&text)
	return text, err
}

func(DB *DataBase) GetSongByName(name string, group string) (models.Song, error){
	req := fmt.Sprintf("SELECT * FROM songs WHERE (group=%s && song=%s)", group, name)

	var song models.Song

	rows ,err := DB.db.Query(req)
	if err != nil{
		return song, err
	}

	err = rows.Scan(&song.Id, &song.Song, &song.Group)
	return song, err
}

func(DB *DataBase) CreateSong(song models.Song) error{
	_, err := DB.db.Exec("INSERT INTO songs (group, song) VALUES (%s, %s)", song.Group, song.Song)
	return err
}

func(DB *DataBase) UpdateSong(name string, group string, updateSong models.Song) error{
	_, err := DB.db.Exec("UPDATE songs WHERE (group=%s, song=%s) SET group=%s song=%s releaseDate=%s text=%s link=%s",group, name, updateSong.Group, updateSong.Song, updateSong.ReleaseDate, updateSong.Text, updateSong.Link)
	return err
}

func(DB *DataBase) DeleteSong(name string, group string) error{
	_, err := DB.db.Exec("DELETE FROM songs WHERE (song=%s && group=%s)", name, group)
	return err 
}