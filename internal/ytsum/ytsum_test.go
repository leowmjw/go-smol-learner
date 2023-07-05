package ytsum

import (
	"testing"
)

func Test_summarizeVideoMixMatch(t *testing.T) {
	type args struct {
		playListId string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"SRE Americas 2023",
			args{playListId: "PLbRoZ5Rrl5ldi79QwiX4xaR-l9kD4q6kg"},
			//"Temporal Replay 2022",
			//args{playListId: "PLbRoZ5Rrl5ldi79QwiX4xaR-l9kD4q6kg"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//fetchPaylistExecute(t, tt.args.playListId)
			AllGo(tt.args.playListId)
			t.FailNow()
		})
	}
}

func fetchPaylistExecute(t *testing.T, playListId string) {
	t.Helper()
	AllGo(playListId)
	//ytSumKey := os.Getenv("YT_DEV_KEY")
	//if ytSumKey == "" {
	//	t.Fatal(fmt.Errorf("Fill in YT_DEV_KEY!!"))
	//}
	//// Use YT_DEV_KEY to get the needed client
	//service, err := youtube.NewService(context.Background(), option.WithAPIKey(ytSumKey))
	//if err != nil {
	//	t.Fatal(err)
	//}
	//// Get playlist items (videos).
	//playlistItemsResponse, err := service.PlaylistItems.List([]string{"snippet"}).
	//	PlaylistId(playListId).MaxResults(50).Do()
	//if err != nil {
	//	t.Fatal(err)
	//}
	//
	//// Download each video and get its transcript.
	//for _, item := range playlistItemsResponse.Items {
	//	videoId := item.Snippet.ResourceId.VideoId
	//	summarizeSingleVideoMixMatch(videoId)
	//
	//	break
	//}
}
