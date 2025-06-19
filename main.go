package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type PlaylistsObj struct {
	Playlists []Playlist `json:"items"`
}

type Playlist struct {
	Name  string `json:"name"`
	Songs struct {
		// flesh out properties
		Limit  int32           `json:"limit"`
		Next   string          `json:"next"`
		Offset int32           `json:"offset"`
		Total  int32           `json:"total"`
		Track  []PlaylistTrack `json:"items"`
	} `json:"tracks"`
}

type PlaylistTrack struct {
	Track TrackObj `json:"track"`
}

type TrackObj struct {
	Name    string   `json:"name"`
	Artists []Artist `json:"artists"`
	Album   AlbumObj `json:"album"`
}

type AlbumObj struct {
	Name string `json:"name"`
}

type Artist struct {
	Name string `json:"name"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	bearerToken := os.Getenv("BEARER")

	req, err := http.NewRequest("GET", "https://api.spotify.com/v1/playlists/6lOjimIiWeIdU9PAgNnafR", nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Authorization", "Bearer "+bearerToken)
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var playlist Playlist
	if err := json.Unmarshal(body, &playlist); err != nil {
		log.Fatal(err)
	}

	// fmt.Printf("Songs: %d\n", playlist.Songs.Total)

	result, err := json.MarshalIndent(playlist, "", "    ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(result))
}
