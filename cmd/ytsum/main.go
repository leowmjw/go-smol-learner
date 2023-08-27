package main

import (
	"app/internal/ytsum"
	"fmt"
)

func main() {

	fmt.Println("YT Summarizer ...")
	//ytsum.ComboYTTranscriptFromPlaylist()

	// SRE Americas 2023
	//playlistId := "PLbRoZ5Rrl5ldi79QwiX4xaR-l9kD4q6kg"
	// SRE Asia 2023 - https://www.youtube.com/playlist?list=PLbRoZ5Rrl5ldnsuIyb3X-t6zG3IDcnaRn
	playlistId := "PLbRoZ5Rrl5ldnsuIyb3X-t6zG3IDcnaRn"
	ytsum.AllGo(playlistId)
}
