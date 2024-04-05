package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	lab2 "github.com/Farid545/Team-Lab-2"
)

var (
	inputExpression = flag.String("e", "", "Expression to compute")
	inputFile       = flag.String("f", "", "File containing the expression to compute")
	outputFile      = flag.String("o", "", "File to write the result")
)

func main() {
	flag.Parse()

	var inputReader io.Reader
	if *inputFile != "" {
		file, err := os.Open(*inputFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening input file: %v\n", err)
			os.Exit(1)
		}
		defer file.Close()
		inputReader = file
	} else if *inputExpression != "" {
		inputReader = strings.NewReader(*inputExpression)
	} else {
		flag.Usage()
		os.Exit(1)
	}

	var outputWriter io.Writer
	if *outputFile != "" {
		file, err := os.Create(*outputFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating output file: %v\n", err)
			os.Exit(1)
		}
		defer file.Close()
		outputWriter = file
	} else {
		outputWriter = os.Stdout
	}

	handler := &lab2.ComputeHandler{
		Input:  inputReader,
		Output: outputWriter,
	}

	err := handler.Compute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error computing result: %v\n", err)
		os.Exit(1)
	}
}
