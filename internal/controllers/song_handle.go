package controllers

import (
	"encoding/json"
	"log"
	"musiclib/models"
	"net/http"
)

func(h *Handler) GetSong(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodGet{
		log.Println("Method not allowed. GetSong")
		h.error.MethodNotAllowed(w, r)
		return
	}
	song := r.URL.Query().Get("song")
	group := r.URL.Query().Get("group")
	
	if song == "" || group == ""{
		log.Println("group or song are empty. GetSong")
		h.error.BadRequest(w, r)
		return
	}

	var songJson models.Song
	songJson, err := h.service.GetSongByName(song, group)
	if err != nil{
		log.Println("Error Get from db. GetSong")
		h.error.BadRequest(w, r)
		return
	}
	jsonBytesSong, err := json.Marshal(songJson)
	if err != nil{
		log.Println("Error Marshaling json. GetSong")
		h.error.BadRequest(w ,r)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytesSong)
}

func(h *Handler) DeleteSong(w http.ResponseWriter, r *http.Request){
	name := r.URL.Query().Get("song")
	group := r.URL.Query().Get("group")

	if name == "" || group == ""{
		log.Println("group or song are empty. DeleteSong")
		h.error.BadRequest(w, r)
		return
	}

	err := h.service.DeleteSong(name, group)
	if err != nil{
		log.Println("Error delete from db. DeleteSong")
		h.error.BadRequest(w, r)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func(h *Handler) UpdateSong(w http.ResponseWriter, r *http.Request){
	
}
