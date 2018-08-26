package core

import (
	"fmt"
	"strings"

	porterstemmer "github.com/reiver/go-porterstemmer"
)

// Document WIP
type Document struct {
	bagOfWords []string
	class      string
}

// Data is the struct of data in the dataset
type Data struct {
	class    string
	sentence string
}

// CreateDataTraining create a dataset of train
func CreateDataTraining() *[]Data {
	trainingData := make([]Data, 0)
	trainingData = append(trainingData, Data{class: "greeting", sentence: "how are you?"})
	trainingData = append(trainingData, Data{class: "greeting", sentence: "how is your day?"})
	trainingData = append(trainingData, Data{class: "greeting", sentence: "how is it going today?"})
	trainingData = append(trainingData, Data{class: "greeting", sentence: "good day"})

	trainingData = append(trainingData, Data{class: "goodbye", sentence: "have a nice day"})
	trainingData = append(trainingData, Data{class: "goodbye", sentence: "see you later"})
	trainingData = append(trainingData, Data{class: "goodbye", sentence: "have a nice day"})
	trainingData = append(trainingData, Data{class: "goodbye", sentence: "talk to you soon"})

	trainingData = append(trainingData, Data{class: "sandwich", sentence: "make me a sandwich"})
	trainingData = append(trainingData, Data{class: "sandwich", sentence: "can you make a sandwich?"})
	trainingData = append(trainingData, Data{class: "sandwich", sentence: "having a sandwich today?"})
	trainingData = append(trainingData, Data{class: "sandwich", sentence: "what's for lunch?"})

	fmt.Print("")
	return &trainingData
}

// TrimSuffix removes lasts character of a string
func TrimSuffix(s, suffix string) string {
	if strings.HasSuffix(s, suffix) {
		s = s[:len(s)-len(suffix)]
	}
	return s
}

// Contains : Test if a string is an array of string
func Contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

var words []string
var classes []string
var documents []Document
var training [][]float64
var output [][]float64

// SliceIndex : Get index of a value in a slice
func SliceIndex(limit int, predicate func(i int) bool) int {
	for i := 0; i < limit; i++ {
		if predicate(i) {
			return i
		}
	}
	return -1
}

// TrainData train a data
func TrainData() {
	data := CreateDataTraining()
	OrganizeData(data)
	train(training, output, 20, 0.1, 100000, false, 0.2)
}

// OrganizeData organize data of data traning
func OrganizeData(data *[]Data) {
	words = make([]string, 0)
	classes = make([]string, 0)
	documents = make([]Document, 0)
	for _, pattern := range *data {
		wordsSentence := strings.Fields(pattern.sentence)
		words = append(words, wordsSentence...)
		documents = append(documents, Document{wordsSentence, pattern.class})
		if Contains(classes, pattern.class) == false {
			classes = append(classes, pattern.class)
		}
	}
	for countWords := range words {
		words[countWords] = TrimSuffix(words[countWords], "?")
		words[countWords] = porterstemmer.StemString(words[countWords])
	}
	words = RemoveDuplicatesFromSlice(words)
	classes = RemoveDuplicatesFromSlice(classes)
	training = make([][]float64, 0)
	output = make([][]float64, 0)
	for _, doc := range documents {
		bag := make([]float64, 0)
		patternWords := doc.bagOfWords
		for countWords := range patternWords {
			patternWords[countWords] = TrimSuffix(patternWords[countWords], "?")
			patternWords[countWords] = porterstemmer.StemString(patternWords[countWords])
		}
		for _, w := range words {
			if Contains(patternWords, w) {
				bag = append(bag, 1)
			} else {
				bag = append(bag, 0)
			}
		}
		training = append(training, bag)
		outputRow := make([]float64, len(classes))
		outputRow[SliceIndex(len(classes), func(i int) bool { return classes[i] == doc.class })] = 1
		output = append(output, outputRow)
	}
}
