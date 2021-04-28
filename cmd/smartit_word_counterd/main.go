package main

import (
	"log"
	"sync"

	"github.com/kzozulya1/smartit_word_counter_demo/internal/analyzer"
	"github.com/kzozulya1/smartit_word_counter_demo/internal/fileutils"
)

const (
	// wordsLimit is a number of words in final word list
	wordsLimit = 10

	// minWordLenFilter is a word len filter
	minWordLenFilter = 5

	//dir
	dir = "/wcounter/files"
)

func main() {
	var (
		wg    sync.WaitGroup
		mutex = &sync.Mutex{}

		wordAnalyzer = analyzer.New(wordsLimit, minWordLenFilter)
	)

	//Scan dir for all files
	files, err := fileutils.ScanDir(dir)
	if err != nil {
		log.Fatalf("dir scan err: %s", err.Error())
	}

	//Range over files
	for _, file := range files {
		//... and start new go routine
		wg.Add(1)
		go worker(file, wordAnalyzer, &wg, mutex)
	}

	wg.Wait()

	//Do aggregated content processing:
	_, err = wordAnalyzer.Process()
	if err != nil {
		log.Fatalf("word analyze err: %s", err.Error())
	}

	//Display results:
	wordAnalyzer.DisplayWords()
}
