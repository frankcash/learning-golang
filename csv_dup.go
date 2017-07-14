package example

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

func main(){
	for _, tempCsv := range os.Args[1:]{
		readCsv(tempCsv, "condition") // starts a go routine
	}
}

func readCsv(myCsv string, desiredCondition string	){
	var colIndex int

	fmt.Printf("Running %s\n", myCsv)
	file, err := os.Open(myCsv)
	if err != nil{
		fmt.Println("Err ", err)
		return
	}

	defer file.Close()
	reader := csv.NewReader(file)
	first_row, _ := reader.Read()
	for i, col := range first_row{
		if strings.EqualFold(desiredCondition, col) {
			colIndex = i
		}
	}

	fmt.Println(colIndex)

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


		counts[record[colIndex]]++

		lineCount += 1
	}

	for line, n := range counts{
		if n > -1{
			fmt.Printf("%d\t%s\n", n, line)
		}
	}

}
