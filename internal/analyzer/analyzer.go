package analyzer

import (
	"regexp"

	"github.com/sirupsen/logrus"
)

// Analyzer gets filecontent, splits data on words
type Analyzer struct {
	wordsLimit       int            //number of words in rating
	minWordLenFilter int            //min work len
	words            map[string]int //holds words map
}

// New returns a new instance of analyzer, get words limit and min len as params
func New(wordsLimit, minWordLenFilter int) *Analyzer {
	return &Analyzer{
		wordsLimit:       wordsLimit,
		minWordLenFilter: minWordLenFilter,
		words:            make(map[string]int),
	}
}

// SetContent is a setter for content, for struct reuse
func (a *Analyzer) AppendContent(content string) *Analyzer {
	a.countWords(content)
	a.applyWordLenFilter()
	return a
}

// Process do word analize
func (a *Analyzer) Process() (map[string]int, error) {
	a.clipWordList()
	return a.words, nil
}

// countWords returns a map of the counts of each word in string s
// key - word in text
// value - number of word repeat in text
func (a *Analyzer) countWords(s string) {
	var words []string = regexp.MustCompile(`\s+`).Split(s, -1)
	for _, word := range words {
		a.words[word]++
	}
}

// clipWordList sorts leaves only top c.wordsLimit words, by repeat
func (a *Analyzer) clipWordList() {
	if len(a.words) <= a.wordsLimit {
		return
	}

	var counter = 1
	for _, index := range rankMapStringIntDesc(a.words) {
		if counter > a.wordsLimit {
			delete(a.words, index)
		}
		counter++
	}
}

//applyWordLenFilter filter words by len
func (a *Analyzer) applyWordLenFilter() {
	for w := range a.words {
		if a.minWordLenFilter > 0 && a.getWordLen(w) < a.minWordLenFilter {
			delete(a.words, w)
		}
	}
}

//applyWordLenFilter filter words by len
func (a *Analyzer) DisplayWords() {
	for w, freq := range a.words {
		logrus.Infof("%s repeats %d times", w, freq)
	}
}

//Get len of term: non en occupies 2 bytes. En only 1 byte
func (a *Analyzer) getWordLen(w string) int {
	var (
		isEng = regexp.MustCompile(`[a-zA-Z]`)
		tLen  = len(w)
	)
	if !isEng.Match([]byte(w)) {
		tLen /= 2
	}
	return tLen
}
