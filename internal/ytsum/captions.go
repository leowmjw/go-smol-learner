package ytsum

import (
	"fmt"
	genapi "github.com/lunabrain-ai/lunabrain/gen/api"
	"github.com/treant5612/ytvc-web/manager/youtubeapi"
	"google.golang.org/api/youtube/v3"
)

type Content struct {
	NormalizerID genapi.NormalizerID
	Data         string
}

// Source: https://sourcegraph.com/github.com/djthorpe/ytapi/-/blob/ytcommands/captions.go

func Captions(videoId string) (captions []*youtube.Caption, err error) {
	call := youtubeapi.ServiceFSC.Captions.List([]string{"snippet"}, videoId)
	resp, err := call.Do()
	if err != nil {
		return nil, err
	}
	return resp.Items, nil
}

/*
Download caption and return it's path in local filesystem.
*/

// DownloadCaption is broken as no ..
func DownloadCaption(videoId string, captionId string, tlang string) (fpath string, err error) {
	return "", nil
	// Broken below ..
	//v, err := VideoInfo(videoId, "youtube")
	//if err != nil {
	//	return "", err
	//}
	//fpath = path.Join(downloadPath, fmt.Sprintf("%s.%s_%s.srt", videoId, captionId, tlang))
	//err = v.Captions.DownloadToFile(captionId, tlang, fpath)
	//return fpath, err
}

/*
download two captions and merge them.
return the filepath in local fs.
*/

// DownloadAndMergeCaption what does it do??
func DownloadAndMergeCaption(videoId, mainId, mainTlang, secondaryId, secondaryTlang string) (path string, err error) {
	ch := make(chan error)
	var mainPath, secondaryPath string
	go func() {
		var err1 error
		mainPath, err1 = DownloadCaption(videoId, mainId, mainTlang)
		ch <- err1
	}()
	go func() {
		var err2 error
		secondaryPath, err2 = DownloadCaption(videoId, secondaryId, secondaryTlang)
		ch <- err2
	}()

	for i := 0; i < 2; i++ {
		err = <-ch
		if err != nil {
			return "", fmt.Errorf("download caption failed:%w", err)
		}
	}
	return merge(mainPath, secondaryPath)
}

func merge(mainPath, secondaryPath string) (filepath string, err error) {
	return "", nil
	// below is broken ..
	//s1, err := astisub.OpenFile(mainPath)
	//if err != nil {
	//	return
	//}
	//defer os.Remove(mainPath)
	//s2, err := astisub.OpenFile(secondaryPath)
	//if err != nil {
	//	return
	//}
	//defer os.Remove(secondaryPath)
	//s1.Merge(s2)
	////Write 方法需要标注格式
	//fileName := fmt.Sprintf("%s_%s.srt", path.Base(mainPath), path.Base(secondaryPath))
	//filepath = path.Join(downloadPath, fileName)
	//err = s1.Write(filepath)
	//return
}
