package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func main(){
	file, err := os.Open("test.csv")
	if err != nil{
		fmt.Println("Err ", err)
		return
	}

	defer file.Close()

	reader := csv.NewReader(file)


	lineCount := 0
	counts := make(map[string]int)
	for{
		record, err := reader.Read()	
		if err == io.EOF{
			break
		}else if err != nil {
			fmt.Println("Err ", err)
			return
		}

		// fmt.Println("Record", lineCount, " is ", record, " and has ", len(record), "fields")

		counts[record[2]]++

		// fmt.Println()
		lineCount += 1
	}
	
	for line, n := range counts{
		if n > 1{
			fmt.Printf("%d\t%s\n", n, line)
		}
	}

}

