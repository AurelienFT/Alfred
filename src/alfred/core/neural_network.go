package core

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"math/rand"
	"strings"
	"time"

	porterstemmer "github.com/reiver/go-porterstemmer"
)

func sigmoidDoubleTab(x [][]float64) [][]float64 {
	output := make([][]float64, len(x))
	for xOne, xi := range x {
		output[xOne] = make([]float64, len(x[xOne]))
		for yOne, yi := range xi {
			output[xOne][yOne] = 1 / (1 + math.Exp(-yi))
		}
	}
	return output
}

func sigmoidTab(x []float64) []float64 {
	output := make([]float64, len(x))
	for _, xi := range x {
		xi = 1 / (1 + math.Exp(-xi))
	}
	return output
}

func sigmoid(x float64) float64 {
	output := 1 / (1 + math.Exp(-x))
	return output
}

func sigmoidOutputToDerivate(output [][]float64) [][]float64 {
	for xi := range output {
		for yi := range output[xi] {
			output[xi][yi] = output[xi][yi] * (1 - output[xi][yi])
		}
	}
	return output
}

func cleanUpSentence(sentence string) []string {
	sentenceWords := strings.Fields(sentence)
	for wordsCount := range sentenceWords {
		sentenceWords[wordsCount] = TrimSuffix(sentenceWords[wordsCount], "?")
		sentenceWords[wordsCount] = porterstemmer.StemString(sentenceWords[wordsCount])
	}
	return sentenceWords
}

func bow(sentence string, words []string) []float64 {
	sentenceWords := cleanUpSentence(sentence)
	bag := make([]float64, len(words))
	for _, s := range sentenceWords {
		for i, w := range words {
			if w == s {
				bag[i] = 1
			}
		}
	}
	return bag
}

var synapse0 [][]float64
var synapse1 [][]float64

func dot(x []float64, y [][]float64) []float64 {
	result := make([]float64, len(x))
	for oneX, xi := range x {
		for _, yi := range y {
			result[oneX] = result[oneX] + xi*yi[oneX]
		}
	}
	return result
}

func dotTab(x [][]float64, y [][]float64) [][]float64 {
	result := make([][]float64, len(x))
	for lOne, li := range x {
		result[lOne] = make([]float64, len(x[lOne]))
		for oneX, xi := range li {
			for _, yi := range y {
				result[lOne][oneX] = result[lOne][oneX] + xi*yi[oneX]
			}
		}
	}
	return result
}

func think(sentence string) []float64 {
	x := bow(strings.ToLower(sentence), words)
	l0 := x
	dotResult := dot(l0, synapse0)
	l1 := sigmoidTab(dotResult)
	dotResult = dot(l1, synapse1)
	l2 := sigmoidTab(dotResult)
	return l2
}

func generateRandomSlice(x int, y int, multiplyArray float64, operationOnSlice float64) [][]float64 {
	randomSlice := make([][]float64, x)
	for _, subSlice := range randomSlice {
		subSlice = make([]float64, y)
		for xi := range subSlice {
			subSlice[xi] = multiplyArray*rand.Float64() - operationOnSlice
		}
	}
	return randomSlice
}

func meanAbs(tab [][]float64) float64 {
	count := 0.0
	result := 0.0
	for _, x := range tab {
		for _, y := range x {
			count++
			result += math.Abs(y)
		}
	}
	return result / count
}

func transpose(tab [][]float64) [][]float64 {
	transposeTab := make([][]float64, len(tab[0]))
	for count := 0; count != len(tab[0]); count++ {
		transposeTab[count] = make([]float64, len(tab))
		for count2 := 0; count2 != len(tab); count2++ {
			transposeTab[count][count2] = tab[count2][count]
		}
	}
	return transposeTab
}

func testSup(tab [][]float64, testNb float64) [][]float64 {
	result := make([][]float64, len(tab))
	for xi := range tab {
		result[xi] = make([]float64, len(tab[xi]))
		for y1 := range tab[xi] {
			if tab[xi][y1] > testNb {
				result[xi][y1] = 1
			} else {
				result[xi][y1] = 1
			}
		}
	}
	return result
}

type synapsesJSON struct {
	synapse0 [][]float64
	synapse1 [][]float64
	time     string
	words    []string
	classes  []string
}

func train(X [][]float64, y [][]float64, hiddenNeurons int, alpha float64, epochs int, dropout bool, dropoutPercent float64) {
	rand.Seed(1)
	lastMeanError := 1.0
	synapse0 = generateRandomSlice(len(X[0]), hiddenNeurons, 2, -1)
	synapse1 = generateRandomSlice(hiddenNeurons, len(classes), 1, -1)

	prevSynapse0WeightUpdate := make([][]float64, len(X[0]))
	for xi := range prevSynapse0WeightUpdate {
		prevSynapse0WeightUpdate[xi] = make([]float64, hiddenNeurons)
	}
	prevSynapse1WeightUpdate := make([][]float64, hiddenNeurons)
	for xi := range prevSynapse1WeightUpdate {
		prevSynapse1WeightUpdate[xi] = make([]float64, len(classes))
	}

	synapse0DirectionCount := make([][]float64, len(X[0]))
	for xi := range synapse0DirectionCount {
		synapse0DirectionCount[xi] = make([]float64, hiddenNeurons)
	}
	synapse1DirectionCount := make([][]float64, hiddenNeurons)
	for xi := range synapse1DirectionCount {
		synapse1DirectionCount[xi] = make([]float64, len(classes))
	}
	for j := epochs + 1; j != 0; j-- {
		layer0 := X
		dotResult := dotTab(layer0, synapse0)
		layer1 := sigmoidDoubleTab(dotResult)
		//if dropout
		dotResult = dotTab(layer1, synapse1)
		layer2 := sigmoidDoubleTab(dotResult)
		layer2Error := make([][]float64, len(layer2))
		for xi, x := range layer2 {
			layer2Error[xi] = make([]float64, len(x))
			for yiLayer := range layer2[xi] {
				layer2Error[xi][yiLayer] = y[xi][yiLayer] - layer2[xi][yiLayer]
			}
		}
		if j%10000 == 0 && j > 5000 {
			if meanAbs(layer2Error) < lastMeanError {
				fmt.Printf("Delta after %d iterations: %f", j, meanAbs(layer2Error))
				lastMeanError = meanAbs(layer2Error)
			} else {
				fmt.Printf("Break: %f > %f", meanAbs(layer2Error), lastMeanError)
				break
			}
		}
		sigmoidTemp := sigmoidOutputToDerivate(layer2)
		layer2Delta := make([][]float64, len(layer2Error))
		for xi, x := range layer2Error {
			layer2Delta[xi] = make([]float64, len(x))
			for yi := range layer2Error[xi] {
				layer2Delta[xi][yi] = layer2Error[xi][yi] * sigmoidTemp[xi][yi]
			}
		}
		layer1Error := dotTab(layer2Delta, transpose(synapse1))
		sigmoidTemp = sigmoidOutputToDerivate(layer2)
		layer1Delta := make([][]float64, len(layer1Error))
		for xi, x := range layer2Error {
			layer1Delta[xi] = make([]float64, len(x))
			for yi := range layer1Error[xi] {
				layer1Delta[xi][yi] = layer1Error[xi][yi] * sigmoidTemp[xi][yi]
			}
		}
		synapse1WeightUpdate := dotTab(transpose(layer1), layer2Delta)
		synapse0WeightUpdate := dotTab(transpose(layer0), layer1Delta)
		if j > 0 {
			synapse0Temp := testSup(synapse0WeightUpdate, 0)
			prevSynapse0Temp := testSup(prevSynapse0WeightUpdate, 0)
			for xi := range synapse0Temp {
				for yi := range prevSynapse0Temp[xi] {
					synapse0DirectionCount[xi][yi] += math.Abs(synapse0Temp[xi][yi] - prevSynapse0Temp[xi][yi])
				}
			}
			synapse1Temp := testSup(synapse1WeightUpdate, 0)
			prevSynapse1Temp := testSup(prevSynapse1WeightUpdate, 0)
			for xi := range synapse1Temp {
				for yi := range prevSynapse1Temp[xi] {
					synapse1DirectionCount[xi][yi] += math.Abs(synapse1Temp[xi][yi] - prevSynapse1Temp[xi][yi])
				}
			}
		}
		for xi := range synapse1WeightUpdate {
			for yi := range synapse1WeightUpdate[xi] {
				synapse1[xi][yi] += alpha * synapse1WeightUpdate[xi][yi]
				prevSynapse1WeightUpdate[xi][yi] = synapse1WeightUpdate[xi][yi]
			}
		}
		for xi := range synapse0WeightUpdate {
			for yi := range synapse0WeightUpdate[xi] {
				synapse0[xi][yi] += alpha * synapse0WeightUpdate[xi][yi]
				prevSynapse0WeightUpdate[xi][yi] = synapse0WeightUpdate[xi][yi]
			}
		}
		t := time.Now()
		b, _ := json.Marshal(synapsesJSON{synapse0, synapse1, t.String(), words, classes})
		ioutil.WriteFile("synapses.json", b, 0644)
		fmt.Printf("Saved synapses to: synapses.json")
	}
}
