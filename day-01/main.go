package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var filename string = "sample.txt"

func getFileContent() ([]string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return []string{}, err
	}
	content := strings.Split(strings.TrimSpace(string(data)), "\n")
	return content, nil
}

func partONE() {
	content, err := getFileContent()
	if err != nil {
		panic(err)
	}
	var ans int

	for i := 0; i < len(content)-1; i++ {
		a, err := strconv.Atoi(content[i+1])
		if err != nil {
			panic(err)
		}
		b, err := strconv.Atoi(content[i])
		if err != nil {
			panic(err)
		}
		if a > b {
			ans++
		}
	}

	log.Printf("%v", ans)
}

func partTWO() {
	// filename = "example.txt"
	content, err := getFileContent()
	if err != nil {
		panic(err)
	}

	f := []int{}
	for i := 0; i < len(content)-2; i++ {
		a, _ := strconv.Atoi(content[i])
		b, _ := strconv.Atoi(content[i+1])
		c, _ := strconv.Atoi(content[i+2])
		f = append(f, a+b+c)
	}

	ans := 0
	for j := 0; j < len(f)-1; j++ {
		if f[j+1] > f[j] {
			ans++
		}
	}

	fmt.Println(ans)

}

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Fatal(r)
		}
	}()

	// partONE()
	partTWO()
}
