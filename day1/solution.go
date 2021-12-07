package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var measurements []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineStr := scanner.Text()
		num, _ := strconv.Atoi(lineStr)
		measurements = append(measurements, num)
		//fmt.Println(lineStr, num)
	}

	//fmt.Println(measurements)

	var increased int = 0
	for i, s := range measurements {
		if i == len(measurements)-1 {
			break
		}
		if measurements[i+1] > s {
			increased += 1
		}
	}

	fmt.Println(increased)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
