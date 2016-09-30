// https://github.com/google/codesearch
// https://golang.org/pkg/bufio/#NewScanner

package search

import (
	"io/ioutil"
	"log"
	"strings"
	"sync"
)

type SearchResult struct {
	Dir   string
	Tag   string
	Files []string
}

func (sr *SearchResult) Search() error {
	f, err := Search(sr.Dir, sr.Tag)
	sr.Files = f
	return err
}

// Search looks through all files in a given `dir` and searches for occurences of `tag`, returning a slice of strings of the files containing `tag`.
// It spawns one goroutine for each file.
func Search(dir, tag string) ([]string, error) {
	fi, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Println("Failed to read directory.")
		return nil, err
	}
	files := make(chan string, 10)
	var wg sync.WaitGroup
	// TODO check if it is a file
	for _, v := range fi {
		if v.IsDir() {
			continue
		}
		v := v.Name()
		wg.Add(1)
		go func(file string) {
			defer wg.Done()
			b, err := ioutil.ReadFile(dir + file)
			if err != nil {
				log.Println("Failed to read ", file, ":", err.Error())
			}
			if strings.Contains(string(b), tag) {
				files <- file
			}
		}(v)
	}
	wg.Wait()
	close(files)
	var s []string
	for file := range files {
		s = append(s, file)
	}
	return s, nil
}
