package ytsum

import (
	"context"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	kagi "github.com/sashabaranov/kagi-summarizer-api"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
	"os"
)

// Final implementation here ..

const (
	YT_SUM_ENV = "dev"
)

// ComboYTTranscriptFromPlaylist what we get from all ..
func ComboYTTranscriptFromPlaylist() {
	// Have a hard-coded playlistID
	// e.g https://www.youtube.com/playlist?list=PLbRoZ5Rrl5ldi79QwiX4xaR-l9kD4q6kg
	//playListID := "PLbRoZ5Rrl5ldi79QwiX4xaR-l9kD4q6kg"

	ytSumKey := os.Getenv("YT_DEV_KEY")
	if ytSumKey == "" {
		panic(fmt.Errorf("Fill in YT_DEV_KEY!!"))
	}
	// Use YT_DEV_KEY to get the needed client
	svc, err := youtube.NewService(context.Background(), option.WithAPIKey(ytSumKey))
	if err != nil {
		panic(err)
	}

	// DEBUG
	//getPlayListsFromChannel(svc)

	getVideoTranscriptsFromPlayList(svc)
	//fmt.Println("MATCH: ", playListID)
}

// summarizeVideo
func summarizeVideo(videoId string, isExpert bool) {

	videoURL := fmt.Sprintf("https://www.youtube.com/watch?v=%s", videoId)
	fmt.Println("SUMMARIZE:", videoURL)

	kagiEngine := kagi.SummaryEngineAgnes
	if isExpert {
		fmt.Println("ENGINE: MURIEL!!")
		// Muriel is expensive! USD1 per summary + point!!
		//kagiEngine = kagi.SummaryEngineMuriel
		//kagiEngine = kagi.SummaryEngineDaphne // more fun
		// default - cecil
		kagiEngine = "cecil"
	}

	kagiKey := os.Getenv("KAGI_KEY")
	client := kagi.NewClient(kagiKey)
	response, err := client.Summarize(
		context.Background(),
		kagi.SummaryRequest{
			URL:         videoURL,
			SummaryType: kagi.SummaryTypeSummary,
			Engine:      kagiEngine,
			Cache:       true,
		},
	)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println("Summary: ")
	fmt.Println(response.Data.Output)

	kpresp, kperr := client.Summarize(
		context.Background(),
		kagi.SummaryRequest{
			URL:         videoURL,
			SummaryType: kagi.SummaryTypeTakeaway,
			Engine:      kagiEngine,
			Cache:       true,
		},
	)
	if kperr != nil {
		fmt.Println("Error: ", kperr)
		return
	}
	fmt.Println("KeyPoints: ")
	fmt.Println(kpresp.Data.Output)

}

// getVideoTranscriptsFromPlayList
func getVideoTranscriptsFromPlayList(service *youtube.Service) {
	// Have a hard-coded playlistID
	// e.g https://www.youtube.com/playlist?list=PLbRoZ5Rrl5ldi79QwiX4xaR-l9kD4q6kg
	playListID := "PLbRoZ5Rrl5ldi79QwiX4xaR-l9kD4q6kg"

	// Get playlist items (videos).
	playlistItemsResponse, err := service.PlaylistItems.List([]string{"snippet"}).
		PlaylistId(playListID).MaxResults(50).Do()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	// Download each video and get its transcript.
	for _, item := range playlistItemsResponse.Items {
		videoID := item.Snippet.ResourceId.VideoId
		// DEBUG
		//spew.Dump(item.Snippet)
		//summarizeVideo(videoID, false)
		summarizeVideo(videoID, true)

		// Download video... not needed
		//videoResponse, err := service.Videos.
		//	List([]string{"snippet,contentDetails"}).Id(videoID).Do()
		//if err != nil {
		//	fmt.Println(err)
		//}
		//videoURL := videoResponse.Items[0].Snippet.Thumbnails.Default.Url
		//videoData, err := http.Get(videoURL)
		//if err != nil {
		//	fmt.Println(err)
		//}
		//videoFile, err := os.Create(videoID + ".mp4")
		//if err != nil {
		//	fmt.Println(err)
		//}
		//_, err = io.Copy(videoFile, videoData.Body)
		//if err != nil {
		//	fmt.Println(err)
		//}

		// Get transcript based on VideoID in the loop ..
		transcriptResponse, err := service.Captions.
			List([]string{"id"}, videoID).Do()
		if err != nil {
			fmt.Println(err)
		}

		// DEBUG
		//spew.Dump(transcriptResponse.Items)
		if len(transcriptResponse.Items) != 1 {
			fmt.Println("LEN: ", len(transcriptResponse.Items))
			break
		}
		transcriptId := transcriptResponse.Items[0].Id
		fmt.Println("VIDEOID: ", videoID, " TRSID: ", transcriptId)
		// Below needs OAuth2; so might use browser driven method instead ...
		//resp, cerr := service.Captions.Download(transcriptId).Download()
		//if cerr != nil {
		//	fmt.Println("ERR: ", cerr)
		//	break
		//}
		//
		//downloadBytes, rerr := io.ReadAll(resp.Body)
		//if rerr != nil {
		//	fmt.Println(fmt.Errorf("failed to read caption track: %v", rerr))
		//	break
		//}
		//captionTrack := string(downloadBytes)
		//spew.Dump(captionTrack)

		//// Write transcript to YAML file.
		//transcriptFile, err := os.Create(videoID + ".yaml")
		//if err != nil {
		//	fmt.Println(err)
		//}
		//transcriptFile.WriteString("transcript: " + string(transcriptData))

		break
	}
}

func getPlayListsFromChannel(svc *youtube.Service) {
	channelID := "UC4-GrpQBx6WCGwmwozP744Q"

	plc := svc.Playlists.List([]string{"snippet", "contentDetails"}).ChannelId(channelID)

	// Example how to customize calloptions ..
	co := make([]googleapi.CallOption, 0)
	co = append(co, googleapi.QueryParameter("BOB", "123"))
	co = append(co, googleapi.UserIP("127.0.0.1"))
	co = append(co, googleapi.QuotaUser("abcd"))
	co = append(co, googleapi.Trace("YTSUM"))
	// Spread out the CallOption as per needed ,,
	resp, rerr := plc.Do(co...)
	if rerr != nil {
		panic(rerr)
	}
	spew.Dump(resp.Items)

}
