package main

import (
	"bufio"
	"container/heap"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func computeToK(output string, k int) {
	files, err := ioutil.ReadDir(baseDir)
	if err != nil {
		log.Fatal(err)
	}
	if len(files) == 0 {
		log.Fatal("No files in Dir")
	}
	// minHeap := newMinHeap(k)
	minHeap := NewMinHeap(k)
	isInit := false
	for _, file := range files {
		hashMap := mergeFile(file.Name())
		for key, value := range hashMap {
			heap.Push(minHeap, &Item{key, value})
			if minHeap.Len() == k-1 && !isInit {
				heap.Init(minHeap)
				isInit = true
			}
		}
	}
	resultList := minHeap.toList()

	out, err := os.OpenFile(output, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	for _, item := range resultList {
		out.WriteString(item.url + "," + strconv.Itoa(item.count) + "\n")
	}
}

func mergeFile(fileName string) map[string]int {
	urlMap := make(map[string]int)
	f, err := os.Open(baseDir + fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	bufrd := bufio.NewReader(f)
	for {
		line, err := bufrd.ReadString('\n')
		if err == io.EOF {
			break
		}
		parts := strings.Split(line[:len(line)-1], ",")
		url := parts[0]
		count, _ := strconv.Atoi(parts[1])
		urlMap[url] += count
	}
	return urlMap
}
