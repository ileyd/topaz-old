package utils

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const (
	ffmpegPath = "/usr/bin/ffmpeg"
)

func baseFilename(filename string) (basename string) {
	basename = strings.TrimSuffix(filename, filepath.Ext(filename))
	return basename
}

func RemuxMKVToMP4(dir, source string) {
	log.Println("RemuxMKVToMP4: dir:", dir, "source:", source, "baseFilename:", baseFilename(source))
	if filepath.Ext(source) == ".mp4" {
		log.Println("File is already an MP4 file; exiting")
		return
	}

	var cmd exec.Cmd
	cmd.Dir = dir
	cmd.Path = ffmpegPath
	cmd.Args = []string{
		"-i " + source,
		"-acodec copy",
		"-vcodec copy",
		baseFilename(source) + ".mp4",
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Remuxed " + source + " to " + baseFilename(source) + ".mp4")
}
