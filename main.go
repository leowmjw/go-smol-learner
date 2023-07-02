package main

import (
	"context"
	"fmt"
	"google.golang.org/api/youtube/v3"
	"io/ioutil"
	"log"
	"testing"

	"google.golang.org/api/option"
)

func main() {

	//b, err := ioutil.ReadFile("client_secret.json")
	//if err != nil {
	//	log.Fatalf("Unable to read client secret file: %v", err)
	//}
	//
	//// If modifying these scopes, delete your previously saved credentials
	//// at ~/.credentials/youtube-go-quickstart.json
	//config, err := google.ConfigFromJSON(b, youtube.YoutubeReadonlyScope)
	//if err != nil {
	//	log.Fatalf("Unable to parse client secret file to config: %v", err)
	//}
	svc, err := getClient()
	if err != nil {
		panic(err)
	}
	fmt.Println("YT_AGENT:", svc.UserAgent)
}

func getClient() (*youtube.Service, error) {
	ctx := context.Background()
	apiKey := "YOUR_API_KEY"

	service, err := youtube.NewService(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return nil, fmt.Errorf("failed to create YouTube client: %v", err)
	}

	return service, nil
}

func downloadPlaylistVideos(playlistID string) error {
	service, err := getClient()
	if err != nil {
		return err
	}

	// Make an API call to retrieve the playlist items
	call := service.PlaylistItems.List([]string{"snippet"}).PlaylistId(playlistID).MaxResults(50)
	response, err := call.Do()
	if err != nil {
		return fmt.Errorf("error retrieving playlist items: %v", err)
	}

	for _, item := range response.Items {
		videoID := item.Snippet.ResourceId.VideoId

		// Download the video
		err := downloadVideo(videoID)
		if err != nil {
			log.Printf("error downloading video (ID: %s): %v", videoID, err)
			continue
		}

		// Download the transcript
		transcript, err := downloadTranscript(videoID)
		if err != nil {
			log.Printf("error downloading transcript for video (ID: %s): %v", videoID, err)
			continue
		}

		// Convert transcript to YAML and save to file
		filename := fmt.Sprintf("%s.yaml", videoID)
		err = saveTranscriptAsYAML(filename, transcript)
		if err != nil {
			log.Printf("error saving transcript as YAML for video (ID: %s): %v", videoID, err)
			continue
		}
	}

	return nil
}

func saveTranscriptAsYAML(filename string, transcript string) error {
	data := []byte(transcript)
	err := ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("error saving transcript as YAML: %v", err)
	}

	return nil
}

// downloadVideo downloads a video using its ID.
func downloadVideo(videoID string) error {
	// Implement the logic to download the video file here
	// ...

	return nil
}

//// downloadTranscript retrieves the text transcript for a video using its ID.
//func downloadTranscript(videoID string) (string, error) {
//	// Implement the logic to retrieve the transcript here
//	// ...
//
//	transcript := "Sample transcript" // Replace with actual transcript retrieval logic
//
//	return transcript, nil
//}

// downloadTranscript retrieves the text transcript for a video using its ID.
func downloadTranscript(videoID string) (string, error) {
	service, err := getClient()
	if err != nil {
		return "", fmt.Errorf("failed to create YouTube client: %v", err)
	}

	// Call the captions API to retrieve the captions for the video
	captionsCall := service.Captions.List([]string{"snippet"}, videoID)
	//captionsCall.VideoId(videoID)
	captionsResponse, err := captionsCall.Do()
	if err != nil {
		return "", fmt.Errorf("failed to retrieve captions: %v", err)
	}

	if len(captionsResponse.Items) == 0 {
		return "", fmt.Errorf("no captions found for the video")
	}

	// Get the ID of the first caption track
	captionID := captionsResponse.Items[0].Id

	// Download the caption content
	downloadCall := service.Captions.Download(captionID)
	err = downloadCall.Do()
	if err != nil {
		return "", fmt.Errorf("failed to download caption: %v", err)
	}

	return "", nil
	//// Read the caption content
	//content, err := ioutil.ReadAll(response.Body)
	//if err != nil {
	//	return "", fmt.Errorf("failed to read caption content: %v", err)
	//}
	//
	//return string(content), nil
	//call := service.Videos.List([]string{"snippet"}).Id(videoID).Fields("items(id,snippet(caption,captionTracks))")
	//response, err := call.Do()
	//if err != nil {
	//	return "", fmt.Errorf("error retrieving video details: %v", err)
	//}
	//
	//if len(response.Items) == 0 {
	//	return "", fmt.Errorf("video not found")
	//}
	//
	//captionTracks := captionsResponse.Items[0].Snippet
	//
	//if len(captionTracks) == 0 {
	//	return "", fmt.Errorf("transcript not available for video")
	//}
	//
	//// Find the transcript track with language set to "en" (English)
	//var transcriptTrack *youtube.Caption
	//for _, track := range captionTracks {
	//	if track.Snippet.Language == "en" {
	//		transcriptTrack = track
	//		break
	//	}
	//}
	//
	//if transcriptTrack == nil {
	//	return "", fmt.Errorf("transcript not found in English")
	//}
	//
	//transcript, err := downloadCaption(transcriptTrack)
	//if err != nil {
	//	return "", fmt.Errorf("error downloading transcript: %v", err)
	//}
	//
	//return transcript, nil
}

// downloadCaption downloads the text content of a caption track.
func downloadCaption(caption *youtube.Caption) (string, error) {
	if caption == nil {
		return "", fmt.Errorf("caption is nil")
	}

	if caption.Kind != "youtube#caption" {
		return "", fmt.Errorf("invalid caption track")
	}

	return "", nil
	// Broken below
	//if caption.Status != "serving" {
	//	return "", fmt.Errorf("caption track is not available")
	//}
	//
	//if caption.TrackKind != "ASR" && caption.TrackKind != "standard" {
	//	return "", fmt.Errorf("unsupported caption track kind")
	//}
	//
	//service, err := getClient()
	//if err != nil {
	//	return "", fmt.Errorf("failed to create YouTube client: %v", err)
	//}
	//
	//call := service.Captions.Download(caption.Id)
	//response, err := call.Do()
	//if err != nil {
	//	if apiErr, ok := err.(*googleapi.Error); ok && apiErr.Code == 404 {
	//		return "", fmt.Errorf("caption track not found")
	//	}
	//	return "", fmt.Errorf("error downloading caption: %v", err)
	//}
	//
	//content, err := ioutil.ReadAll(response.Body)
	//if err != nil {
	//	return "", fmt.Errorf("error reading caption content: %v", err)
	//}
	//
	//return string(content), nil
}

// TestDownloadTranscript tests the downloadTranscript function.
func TestDownloadTranscript(t *testing.T) {
	videoID := "YOUR_VIDEO_ID"

	transcript, err := downloadTranscript(videoID)
	if err != nil {
		t.Errorf("error downloading transcript: %v", err)
	}

	// Assert that the transcript is not empty
	if transcript == "" {
		t.Errorf("transcript is empty")
	}
}
