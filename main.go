package main

import (
	"fmt"
	"encoding/csv"
	"os"
)

func printName(fname, sname string) {
	fmt.Println(fname, sname)
}

func readCSV(fileName string) map[int]string {
	data := make(map[int]string)
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	r := csv.NewReader(file)
	r.Comma = ';'	
	i := 0
	for {
		record, e := r.Read()
		if e != nil {
			break
		}
		//fmt.Printf("%T %T\n", record[0], record[1])
		data[i] = record[1]
		i++
	}
	return data
}

func main() {
	names := readCSV("./russian_names.csv")
	surnames := readCSV("./russian_surnames.csv")
	fmt.Println(len(names), len(surnames))
	printName(names[5], surnames[5])
}