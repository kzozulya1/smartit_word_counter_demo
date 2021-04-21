package main

import (
	"smartit_word_counter/internal/analyzer"
	"smartit_word_counter/internal/fileutils"
	"sync"

	"github.com/sirupsen/logrus"
)

// worker is a worker that reads file and appends content to analyzer
func worker(filepath string, wordAnalyzer *analyzer.Analyzer, wg *sync.WaitGroup, mutex *sync.Mutex) {
	defer wg.Done()

	fileContent, err := fileutils.ReadFile(filepath)
	if err != nil {
		logrus.Errorf("worker: read file %s err: %s", filepath, err.Error())
		return
	}

	logrus.Infof("analyzer append content of file %s", filepath)

	mutex.Lock()
	wordAnalyzer.AppendContent(fileContent)
	mutex.Unlock()
}
