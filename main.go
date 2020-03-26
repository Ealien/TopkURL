package main

import "flag"

func main() {
	// input file
	var inputFile string
	// output file
	var outputFile string
	k := 100

	flag.StringVar(&inputFile, "in", "", "input filename")
	flag.StringVar(&outputFile, "out", "", "output filename")
	flag.Parse()

	fileSplit(inputFile)
	computeToK(outputFile, k)
}

// func main() {
// 	var inputFile string
// 	flag.StringVar(&inputFile, "in", "", "input filename")
// 	flag.Parse()
// 	urlGenerate(inputFile)
// }
