package id3

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/bogem/id3v2"
	"github.com/rylio/ytdl"
)

// Metadata contains information collected from user
type Metadata struct {
	Title  string
	Artist string
	Album  string
}

// CollectMetadataFromUser will prompt the user to input mp3 title, artist, and album
func CollectMetadataFromUser(videoInfo *ytdl.VideoInfo, channel chan *Metadata) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("title (\"" + videoInfo.Title + "\"): ")
	mp3Title, _ := reader.ReadString('\n')
	mp3Title = strings.TrimSpace(mp3Title)

	if mp3Title == "" {
		mp3Title = videoInfo.Title
	}

	fmt.Print("artist (\"" + videoInfo.Author + "\"): ")
	mp3Artist, _ := reader.ReadString('\n')
	mp3Artist = strings.TrimSpace(mp3Artist)

	if mp3Artist == "" {
		mp3Artist = videoInfo.Author
	}

	fmt.Print("album: ")
	mp3Album, _ := reader.ReadString('\n')
	mp3Album = strings.TrimSpace(mp3Album)

	channel <- &Metadata{Title: mp3Title, Artist: mp3Artist, Album: mp3Album}
}

// SetMetadata will apply metadata to the mp3 file's id3 tag
func SetMetadata(mp3Location string, metadata *Metadata) {
	tag, err := id3v2.Open(mp3Location, id3v2.Options{Parse: true})

	if err != nil {
		log.Fatal("Error while initializing a tag: ", err)
	}

	tag.SetTitle(metadata.Title)
	tag.SetArtist(metadata.Artist)
	tag.SetAlbum(metadata.Album)

	if err = tag.Save(); err != nil {
		log.Fatal("Error while saving a tag: ", err)
	}
}
