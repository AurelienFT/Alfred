#!/bin/bash

export GOPATH=$PWD
go get github.com/reiver/go-porterstemmer
pip install nltk
python src/alfred/neural_network/train_neural_network.py