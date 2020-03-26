package main

import (
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

const base = "abcdefghijklmnopqrstuvwxyz0123456789"

var urlSuffix = []string{".com", ".cn", ".gov", ".edu", ".net", ".org", ".int",
	".mil", ".info"}

func urlGenerate(filename string) {
	rand.Seed(time.Now().Unix())
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("Path error")
	}
	defer f.Close()
	distinctURLNumber := 10000
	totalURLNumber := 2000000

	urls := []string{}
	for i := 0; i < distinctURLNumber; i++ {
		urls = append(urls, getURL())
	}

	for i := 0; i < totalURLNumber; i++ {
		index := rand.Intn(1000)
		url := urls[index]
		f.WriteString(url + "\n")
	}
}

func getURL() string {
	var url strings.Builder
	urlLength := 500
	url.WriteString("http://")
	for i := 0; i < urlLength; i++ {
		index := rand.Intn(36)
		url.WriteByte(base[index])
	}
	url.WriteString(urlSuffix[rand.Intn(9)])
	return url.String()
}
