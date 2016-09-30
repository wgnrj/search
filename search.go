// https://github.com/google/codesearch
// https://golang.org/pkg/bufio/#NewScanner
// https://www.oreilly.com/learning/run-strikingly-fast-parallel-file-searches-in-go-with-sync-errgroup

// Package search provides a function and a struct with a wrapper method to search for files that contain a tag.
// It uses goroutines to search files concurrently.
package search

import (
	"io/ioutil"
	"log"
	"bytes"
	"sync"
)

type SearchResult struct {
	Dir   string
	Tag   string
	Files []string
}

func (sr *SearchResult) Search() error {
	f, err := Search(sr.Dir, sr.Tag)
	if err == nil {
		sr.Files = f
	}
	return err
}

// Search scans all files in a given dir and searches for occurences of tag, returning a slice of strings of the files containing tag.
// It spawns one goroutine for each file.
func Search(dir, tag string) ([]string, error) {
	fi, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Println("Failed to read directory.")
		return nil, err
	}

	files := make(chan string, 10)
	var wg sync.WaitGroup

	for _, v := range fi {
		// Check if the file is a directory.
		if v.IsDir() {
			// TODO Implement a recursive search through subdirectories.
			// Just call the Search function recursively.
			continue
		}

		v := v.Name()
		wg.Add(1)

		// Start one goroutine for each file found in the directory to scan it concurrently.
		go func(f string) {
			defer wg.Done()

			// TODO Check for permissions.
			b, err := ioutil.ReadFile(dir + f)
			if err != nil {
				log.Printf("Failed to read file `%s` (%q)", f, err.Error())
				return
			}

			if bytes.Contains(b, []byte(tag)) {
				files <- f
			}
		}(v)
	}

	var s []string

	// Start a goroutine to append all names of the result files containing tag to the slice s.
	// This goroutine will signal the second barrier, because it only calls wg.Done() when the channel got closed.
	go func() {
		defer wg.Done()

		for f := range files {
			s = append(s, f)
		}
	}()

	// Wait for all search goroutines to finish.
	wg.Wait()

	close(files)
	wg.Add(1)

	// Wait for the append goroutine to finish.
	wg.Wait()

	return s, nil
}
