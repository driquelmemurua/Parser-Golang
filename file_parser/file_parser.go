package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	parse_time_start := time.Now()
	f, err := os.Open("/home/granvolumen/millones.csv")
	check(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var edades [100]uintptr

	onComma := func(data []byte, atEOF bool) (advance int, token []byte, err error) {

		start := 0

		// Scan until space, marking end of word.

		for i := 0; i < len(data); i++ {

			if data[i] == ',' {

				return i + 1, data[start:i], nil

			}

		}

		// If we're at EOF, we have a final, non-empty, non-terminated word. Return it.

		if atEOF && len(data) > start {

			return len(data), data[start:], nil

		}

		// Request more data.

		return start, nil, nil

	}
	scanner.Split(onComma)
	for scanner.Scan() {
		str := scanner.Text()
		if len(str) <= 2 {
			edad, err := strconv.Atoi(str)
			check(err)
			edades[edad]++
		}
	}
	parse_time_end := time.Since(parse_time_start)
	for index, element := range edades {
		fmt.Printf("Edad:%d Personas:%d\n", index, element)
	}
	fmt.Printf("Tiempo de ejecucion:%s\n", parse_time_end)
}
