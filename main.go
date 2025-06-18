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
	Href   string `json:"href"`
	Id     string `json:"id"`
	Name   string `json:"name"`
	Tracks struct {
		// flesh out properties
	} `json:"tracks"`
	Uri string `json:"uri"`
}

type Tracks struct {
	Href string `json:"href"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	bearerToken := os.Getenv("BEARER")

	req, err := http.NewRequest("GET", "https://api.spotify.com/v1/playlists/4voN0nDgKI3gPXVRDefAfn?si=e4559cf1db6b461a", nil)
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

	result, err := json.MarshalIndent(playlist, "", "    ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(result))
}
