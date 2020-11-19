package main

import (
	"fmt"
	"encoding/csv"
	"bufio"
)

func printName(fname, sname string) {
	fmt.Println(fname, sname)
}

func readName() {
	file, err := os.Open("russian_names.csv")
	if err != nil {
		panic("ERROR!!!")
	}
	defer file.Close()

	r := csv.NewReader(file)
	data, err := r.ReadAll()
	if err != nil {
		// Когда мы читаем данные до конца файла io.EOF не возвращается, а служит сигналом к завершению чтения
		// ...
	}

	for _, row := range data {
		fmt.Println(row)
	}
}

func main() {
	printName("Гриша", "Первый")
}