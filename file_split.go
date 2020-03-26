package main

import (
	"bufio"
	"hash/crc32"
	"io"
	"log"
	"os"
	"strconv"
)

// number of split files
const fileNumber = 500

// batch size to save to file
const tmpNumber = 100
const baseDir = "data/"

// split input file to number of fileNumber
func fileSplit(fileName string) {
	var hashMap = make(map[int]map[string]int, fileNumber)
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal("File error")
	}
	defer f.Close()
	_, err = os.Stat(baseDir)
	if os.IsNotExist(err) {
		os.Mkdir(baseDir, 0744)
	}
	bufrd := bufio.NewReader(f)
	for {
		data, err := bufrd.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		storeToMap(data[:len(data)-1], hashMap)
	}
	saveAll(hashMap)
}

// store [url, count] to hash map
func storeToMap(data []byte, hashMap map[int]map[string]int) {
	fileIndex := int(crc32.ChecksumIEEE(data) % fileNumber)
	tmpMap, ok := hashMap[fileIndex]
	if !ok {
		tmpMap = make(map[string]int, tmpNumber)
		hashMap[fileIndex] = tmpMap
	}
	if _, ok := tmpMap[string(data)]; !ok {
		tmpMap[string(data)] = 1
	} else {
		tmpMap[string(data)]++
	}
	if len(tmpMap) >= tmpNumber {
		save(fileIndex, tmpMap)
		delete(hashMap, fileIndex)
	}
}

// save temp hash map to corresponding file
func save(fileIndex int, tmpMap map[string]int) {
	f, err := os.OpenFile(baseDir+strconv.Itoa(fileIndex), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	for key, value := range tmpMap {
		f.WriteString(key + "," + strconv.Itoa(value) + "\n")
	}
}

func saveAll(hashMap map[int]map[string]int) {
	for key, value := range hashMap {
		save(key, value)
	}
}
