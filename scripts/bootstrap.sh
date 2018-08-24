#!/bin/bash

sudo apt-get install libncurses5-dev
export GOPATH=$PWD
go get github.com/reiver/go-porterstemmer