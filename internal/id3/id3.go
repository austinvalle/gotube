package id3

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/rylio/ytdl"
)

// Metadata contains information collected from user
type Metadata struct {
	Title  string
	Artist string
	Album  string
}

// CollectMetadataFromUser will prompt the user to input mp3 title, artist, and album
func CollectMetadataFromUser(videoInfo *ytdl.VideoInfo) *Metadata {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("title (\"" + videoInfo.Title + "\"): ")
	mp3Title, _ := reader.ReadString('\n')
	if strings.TrimSpace(mp3Title) == "" {
		mp3Title = videoInfo.Title
	}

	fmt.Print("artist (\"" + videoInfo.Author + "\"): ")
	mp3Artist, _ := reader.ReadString('\n')
	if strings.TrimSpace(mp3Artist) == "" {
		mp3Artist = videoInfo.Author
	}

	fmt.Print("album: ")
	mp3Album, _ := reader.ReadString('\n')

	return &Metadata{Title: mp3Title, Artist: mp3Artist, Album: mp3Album}
}
