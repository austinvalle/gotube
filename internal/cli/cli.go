package cli

import (
	"fmt"
	"io"
	"log"
	"regexp"

	"github.com/bogem/id3v2"
	"github.com/juju/gnuflag"
	"github.com/moosebot/gotube/internal/id3"
	"github.com/moosebot/gotube/internal/youtube"
)

// Exit codes are int values that represent an exit code for a particular error.
const (
	ExitCodeOK    int = 0
	ExitCodeError int = 1 + iota
)

const youtubeURLRegex = `(?:youtube\.com\/\S*(?:(?:\/e(?:mbed))?\/|watch\/?\?(?:\S*?&?v\=))|youtu\.be\/)([a-zA-Z0-9_-]{6,11})`

// CLI is the command line object
type CLI struct {
	OutStream, ErrStream io.Writer
}

// Run invokes the CLI with the given arguments.
func (cli *CLI) Run(args []string) int {
	var (
		skipMetaFlag bool
		dir          string

		versionFlag bool
	)

	flags := gnuflag.NewFlagSet(Name, gnuflag.ContinueOnError)
	flags.SetOutput(cli.ErrStream)

	flags.BoolVar(&skipMetaFlag, "s", false, "skip the metadata edit step")
	flags.StringVar(&dir, "d", "", "specify the download directory")
	flags.BoolVar(&versionFlag, "v", false, "print version information")

	if err := flags.Parse(true, args[1:]); err != nil {
		return ExitCodeError
	}

	if versionFlag {
		fmt.Println(Version)
		return ExitCodeOK
	}

	r := regexp.MustCompile(youtubeURLRegex)

	if r.MatchString(args[1]) == true {
		youtubeID := r.FindAllStringSubmatch(args[1], -1)[0][1]

		mp3Location, videoInfo := youtube.DownloadMp3(youtubeID, dir)

		tag, err := id3v2.Open(mp3Location, id3v2.Options{Parse: true})

		if err != nil {
			log.Fatal("Error while initializing a tag: ", err)
		}

		if skipMetaFlag {
			return ExitCodeOK
		}

		metadata := id3.CollectMetadataFromUser(videoInfo)

		tag.SetTitle(metadata.Title)
		tag.SetArtist(metadata.Artist)
		tag.SetAlbum(metadata.Album)

		if err = tag.Save(); err != nil {
			log.Fatal("Error while saving a tag: ", err)
		}
	}

	return ExitCodeOK
}
