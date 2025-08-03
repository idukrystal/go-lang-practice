package main

import (
	"fmt"
	"runs/runner"
)

type simpleRunner struct {
	
}

func (sr simpleRunner) Run() {
	fmt.Println("simply running ...")
}

func main() {
	runner.FinnishUp(runner.RunInBackground(simpleRunner{}))
}
