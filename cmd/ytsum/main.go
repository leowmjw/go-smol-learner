package main

import (
	"app/internal/ytsum"
	"fmt"
)

func main() {

	fmt.Println("YT Summarizer ...")
	//ytsum.ComboYTTranscriptFromPlaylist()

	playlistId := "PLbRoZ5Rrl5ldi79QwiX4xaR-l9kD4q6kg"
	ytsum.AllGo(playlistId)
}
