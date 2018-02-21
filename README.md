# gotube

golang CLI for downloading youtube videos/audio and renaming metadata

![gotube example](https://i.imgur.com/SQIFAVP.gif)

## Install

To install, use `go get`:

```bash
$ go get -d github.com/moosebot/gotube
```

## Usage

```bash
gotube [youtube url] [options]
```

### Options
- `-s` *default value: false*   - Skip metadata questions and uses defaults
- `-d` *default value: ./*      - Sets the download directory for mp3
- `-h` - Print out gotube help information
- `-v` - Print out gotube CLI version

## Example

#### Basic example
```bash
gotube https://www.youtube.com/watch?v=aatr_2MstrI
```

#### Example skipping metadata
```bash
gotube https://www.youtube.com/watch?v=aatr_2MstrI -s
```

#### Example setting download directory
```bash
gotube https://www.youtube.com/watch?v=aatr_2MstrI -d C:/Downloads
```
