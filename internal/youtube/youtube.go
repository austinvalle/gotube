package youtube

import (
	"crypto/rand"
	"encoding/hex"
	"log"
	"os"
	"os/exec"

	"github.com/cavaliercoder/grab"
	"github.com/rylio/ytdl"
)

// DownloadMp3 will download a youtube mp3 based on the ID to the specified directory
func DownloadMp3(videoInfo *ytdl.VideoInfo, destDir string) string {
	if destDir == "" {
		destDir = "."
	}

	var format = videoInfo.Formats.Best(ytdl.FormatAudioEncodingKey)[0]

	videoURL, _ := videoInfo.GetDownloadURL(format)

	randomFileName := getRandomFileName()
	tempMp4Location := destDir + "/" + randomFileName + ".mp4"
	tempMp3Location := destDir + "/" + randomFileName + ".mp3"

	client := grab.NewClient()
	req, _ := grab.NewRequest(tempMp4Location, videoURL.String())

	resp := client.Do(req)

	if err := resp.Err(); err != nil {
		log.Fatal("Download failed: ", err)
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

	return tempMp3Location
}

func getRandomFileName() string {
	randBytes := make([]byte, 16)
	rand.Read(randBytes)

	return hex.EncodeToString(randBytes)
}
