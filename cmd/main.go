package main

import (
	"errors"
	"flag"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"strings"

	"github.com/samverrall/tagnow/pkg/semvar"
)

const (
	gitDirectory = ".git"
)

var (
	errDirNotExists = errors.New("supplied git directory does not exist")
	errEmptyDir     = errors.New("empty git directory supplied")
)

type opts struct {
	dir   string
	major bool
	minor bool
	patch bool
}

func main() {
	var opts opts
	flag.StringVar(&opts.dir, "dir", "", "Git directory to increment semvar tag")
	flag.Parse()

	if strings.TrimSpace(opts.dir) == "" {
		curDir, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		opts.dir = curDir
	}

	dirEntries, err := getDirectory(opts.dir)
	if err != nil {
		fmt.Printf("tagnew: %s", err.Error())
		os.Exit(1)
	}

	if ok := isGitDirectory(dirEntries); !ok {
		fmt.Printf("tagnew: supplied directory is not git initialised")
		os.Exit(1)
	}

	_, err = getCurrentTag(opts.dir)
	if err != nil {
		fmt.Println("tagnew: failed to get current tag: ", err)
		os.Exit(1)
	}

}

// getCurrentTag attempts to get the newest git tag from the supplied git repository
// and attempts to parse the tag into a SemvarTag type.
//
// If we fail to get a tag, then we return a error. If we get a invalid semvar tag
// we also throw an error.
//
// If no tag exists in the repository, we return a 0th value SemvarTag
// (v0.0.0).
//
// TODO: Allow flag to choose first tag if one does not exist.
func getCurrentTag(dir string) (*semvar.SemvarTag, error) {
	cmd := exec.Command("git", "describe", "--abbrev=0", "--tags")
	cmd.Dir = dir
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	if len(out) == 0 {
		return semvar.New(0, 0, 0), nil
	}

	currentTag := string(out)

	semvarFromTag, err := semvar.NewFromString(currentTag)
	if err != nil {
		return nil, err
	}

	return &semvarFromTag, nil
}

func isGitDirectory(dirEntires []fs.DirEntry) bool {
	for _, entry := range dirEntires {
		if entry.IsDir() && entry.Name() == gitDirectory {
			return true
		}
	}
	return false
}

func getDirectory(dir string) ([]fs.DirEntry, error) {
	dirEntries, err := os.ReadDir(dir)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return []fs.DirEntry{}, errDirNotExists
		}
		return []fs.DirEntry{}, err
	}

	if len(dirEntries) == 0 {
		return []fs.DirEntry{}, errEmptyDir
	}

	return dirEntries, nil
}
