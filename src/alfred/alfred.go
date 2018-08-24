package main

import (
	"alfred/core"
)

func main() {
	data := core.CreateDataTraining()
	core.OrganizeData(data)
	core.AlfredLoop()
}
