package ytsum

import (
	"context"
	"encoding/json"
	"fmt"
	"golang.org/x/oauth2/google"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"google.golang.org/api/option"
	youtube "google.golang.org/api/youtube/v3"
)

// CodyYTTranscript generated from Cody
func CodyYTTranscript() {
	// Read client secrets file to configure a OAuth2 client.
	data, err := ioutil.ReadFile("client_secrets.json")
	if err != nil {
		fmt.Println(err)
	}
	config, err := google.ConfigFromJSON(data, youtube.YoutubeReadonlyScope)
	if err != nil {
		fmt.Println(err)
	}

	// Create a context and YouTube service.
	ctx := context.Background()
	service, err := youtube.NewService(ctx, option.WithHTTPClient(config.Client(ctx)))
	if err != nil {
		fmt.Println(err)
	}

	// Get playlist ID from user input.
	var playlistID string
	fmt.Print("Enter playlist ID: ")
	fmt.Scanln(&playlistID)

	// Get playlist items (videos).
	playlistItemsResponse, err := service.PlaylistItems.List([]string{"snippet"}).
		PlaylistId(playlistID).MaxResults(50).Do()
	if err != nil {
		fmt.Println(err)
	}

	// Download each video and get its transcript.
	for _, item := range playlistItemsResponse.Items {
		videoID := item.Snippet.ResourceId.VideoId

		// Download video.
		videoResponse, err := service.Videos.List([]string{"snippet,contentDetails"}).
			Id(videoID).Do()
		if err != nil {
			fmt.Println(err)
		}
		videoURL := videoResponse.Items[0].Snippet.Thumbnails.Default.Url
		videoData, err := http.Get(videoURL)
		if err != nil {
			fmt.Println(err)
		}
		videoFile, err := os.Create(videoID + ".mp4")
		if err != nil {
			fmt.Println(err)
		}
		_, err = io.Copy(videoFile, videoData.Body)
		if err != nil {
			fmt.Println(err)
		}

		// Get transcript.
		transcriptResponse, err := service.Captions.List([]string{"snippet"}, videoID).Do()
		if err != nil {
			fmt.Println(err)
		}
		transcript := transcriptResponse.Items[0].Snippet.TextDisplay
		transcriptData, err := json.Marshal(transcript)
		if err != nil {
			fmt.Println(err)
		}

		// Write transcript to YAML file.
		transcriptFile, err := os.Create(videoID + ".yaml")
		if err != nil {
			fmt.Println(err)
		}
		transcriptFile.WriteString("transcript: " + string(transcriptData))
	}
}
