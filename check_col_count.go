package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func readCSV(mcsv string, condOne string, condTwo string, condOneCount chan<- int, condTwoCount chan<- int, chFinished chan<- bool) {
	var colOneIndex int
	var colTwoIndex int
	file, err := os.Open(mcsv)
	if err != nil {
		fmt.Println("open error", err)
		return
	}

	defer file.Close()
	reader := csv.NewReader(file)

	first_row, _ := reader.Read()
	for i, col := range first_row {
		if strings.EqualFold(condOne, col) {
			colOneIndex = i
		} else if strings.EqualFold(condTwo, col) {
			colTwoIndex = i
		}
	}

	for {
		record, err := reader.Read()

		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Err ", err)
			return
		}

		if len(record[colOneIndex]) > 0 {
			condOneCount <- 1
		}

		if len(record[colTwoIndex]) > 0 {
			condTwoCount <- 1
		}

	}

	defer func() {
		chFinished <- true
	}()

}

func main() {
	l, err := regexp.Compile(`(?i)csv`)
	if err != nil {
		fmt.Println("Regex error", err)
		return
	}
	fileList := []string{}
	csvList := []string{}

	searchDir := "" // make this desired directory
	errr := filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {
		fileList = append(fileList, path)
		return nil
	})

	if errr != nil {
		fmt.Println("Dir IO ", err)
		return
	}

	for _, file := range fileList {
		//fmt.Println(file)
		if l.MatchString(file) {
			fmt.Println(file)
			csvList = append(csvList, file)
		}
	}

	chFinished := make(chan bool)
	condOne := make(chan int)
	condTwo := make(chan int)
	countOne, countTwo := 0, 0

	for _, csv := range csvList {
		go readCSV(csv, "colOne", "colTwo", condOne, condTwo, chFinished)
	}

	for c := 0; c < len(csvList); {
		select {
		case <-condOne:
			countOne += 1
		case <-condTwo:
			countTwo += 1
		case <-chFinished:
			c++
		}
	}

	fmt.Println(csvList)
	fmt.Println("Tickers", countOne)
	fmt.Println("Skuu_uid", countTwo)

}
