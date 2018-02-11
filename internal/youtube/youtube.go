package youtube

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/cavaliercoder/grab"
	"github.com/rylio/ytdl"
)

// DownloadVideo will download a youtube video based on the ID to the specified directory
func DownloadVideo(videoID string, destDir string) {
	if destDir == "" {
		destDir = "."
	}

	videoInfo, _ := ytdl.GetVideoInfoFromID(videoID)

	var format = videoInfo.Formats.Best(ytdl.FormatAudioEncodingKey)[0]

	videoURL, _ := videoInfo.GetDownloadURL(format)
	tempMp4Location := destDir + "/temp.mp4"
	tempMp3Location := destDir + "/temp.mp3"

	client := grab.NewClient()
	req, _ := grab.NewRequest(tempMp4Location, videoURL.String())

	resp := client.Do(req)
	fmt.Printf("  %v\n", resp.HTTPResponse.Status)

	t := time.NewTicker(500 * time.Millisecond)
	defer t.Stop()

ProgressLoop:
	for {
		select {
		case <-t.C:
			fmt.Printf("  transferred %v / %v bytes (%.2f%%)\n",
				resp.BytesComplete(),
				resp.Size,
				100*resp.Progress())

		case <-resp.Done:
			break ProgressLoop
		}
	}

	if err := resp.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Download failed: %v\n", err)
		os.Exit(1)
	}

	var cmd = exec.Command("ffmpeg", "-i", tempMp4Location, "-q:a", "0", "-map", "a", tempMp3Location)

	err := cmd.Start()
	if err != nil {
		log.Panic(err)
	}

	err = cmd.Wait()
	if err != nil {
		log.Panic(err)
	}

	os.Remove(tempMp4Location)
	os.Rename(tempMp3Location, destDir+"/"+videoInfo.Title+".mp3")
}
