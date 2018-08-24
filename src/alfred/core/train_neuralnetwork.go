package core

import (
	"fmt"
	"strings"
	"github.com/reiver/go-porterstemmer"
)

type Data struct {
	class string
	sentence string
}

func CreateDataTraining() *[]Data {
	training_data := make([]Data, 0)
	training_data = append(training_data, Data{class:"greeting", sentence:"how are you?"})
	training_data = append(training_data, Data{class:"greeting", sentence:"how is your day?"})
	training_data = append(training_data, Data{class:"greeting", sentence:"how is it going today?"})
	training_data = append(training_data, Data{class:"greeting", sentence:"good day"})

	training_data = append(training_data, Data{class:"goodbye", sentence:"have a nice day"})
	training_data = append(training_data, Data{class:"goodbye", sentence:"see you later"})
	training_data = append(training_data, Data{class:"goodbye", sentence:"have a nice day"})
	training_data = append(training_data, Data{class:"goodbye", sentence:"talk to you soon"})

	training_data = append(training_data, Data{class:"sandwich", sentence:"make me a sandwich"})
	training_data = append(training_data, Data{class:"sandwich", sentence:"can you make a sandwich?"})
	training_data = append(training_data, Data{class:"sandwich", sentence:"having a sandwich today?"})
	training_data = append(training_data, Data{class:"sandwich", sentence:"what's for lunch?"})

	fmt.Print("")
	return &training_data
}

func TrimSuffix(s, suffix string) string {
    if strings.HasSuffix(s, suffix) {
        s = s[:len(s)-len(suffix)]
    }
    return s
}

func Contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func OrganizeData(data *[]Data) {
	words := make([]string, 0)
	classes := make([]string, 0)
	documents := make([]string, 0)
	for _, pattern := range *data {
		words_sentence := strings.Fields(pattern.sentence)
		words = append(words, words_sentence...)
		documents = append(documents, pattern.class)
		if Contains(classes, pattern.class) == false {
			classes = append(classes, pattern.class)
		}
	}
	for count_words, _ := range words {
		words[count_words] = porterstemmer.StemString(words[count_words])
		words[count_words] = TrimSuffix(words[count_words], "?")
	}
	words = RemoveDuplicatesFromSlice(words)
	fmt.Printf("%d documents.\n", len(documents))
	fmt.Printf("%d classes\n", len(classes))
	fmt.Printf("%d words\n", len(words))
}