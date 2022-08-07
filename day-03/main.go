package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {

	for _, f := range os.Args[1:] {
		solveFile(f)
	}

}

func solveFile(filename string) {
	fd, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(fd)

	lines := []string{}
	for {
		line, err := reader.ReadString('\n')
		if err != nil && err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		lines = append(lines, strings.TrimSpace(line))
	}
	partONE(lines)
	partTWO(lines)
}

type mode uint64

const (
	Common mode = iota
	LessCommon
)

func getBit(m mode, lines []string, col int, out *string) {

	if col == len(lines[col]) {
		return
	}

	var oneBit, zeroBit int
	for _, l := range lines {
		switch l[col] {
		case '0':
			zeroBit++
		case '1':
			oneBit++
		}
	}

	switch m {
	case Common:
		if oneBit > zeroBit {
			*out += "1"
		} else {
			*out += "0"
		}
	case LessCommon:
		if oneBit > zeroBit {
			*out += "0"
		} else {
			*out += "1"
		}
	}

	col += 1
	getBit(m, lines, col, out)
}

// partONE
func partONE(lines []string) {
	var gammaRate string
	var epsilon string

	getBit(Common, lines, 0, &gammaRate)
	n, _ := strconv.ParseInt(gammaRate, 2, 64)

	getBit(LessCommon, lines, 0, &epsilon)
	m, _ := strconv.ParseInt(epsilon, 2, 64)

	fmt.Println(n * m)
}

func getO2AndCO2rating(m mode, lines []string, col int) string {
	var (
		zeroBit int
		oneBit  int
	)
	var (
		zeroArr []string
		oneArr  []string
	)

	if col == len(lines[0]) || len(lines) == 1 {
		return lines[0]
	}

	for _, n := range lines {
		switch n[col] {
		case '0':
			zeroBit++
			zeroArr = append(zeroArr, n)
		case '1':
			oneBit++
			oneArr = append(oneArr, n)
		}
	}
	switch m {
	case Common:
		if zeroBit > oneBit {
			lines = zeroArr
		} else if zeroBit == oneBit {
			lines = oneArr
		} else {
			lines = oneArr
		}
	case LessCommon:
		if zeroBit > oneBit {
			lines = oneArr
		} else if zeroBit == oneBit {
			lines = zeroArr
		} else {
			lines = zeroArr
		}
	}

	col += 1
	return getO2AndCO2rating(m, lines, col)
}

// partTWO
func partTWO(lines []string) {
	n, _ := strconv.ParseInt(getO2AndCO2rating(Common, lines, 0), 2, 64)
	m, _ := strconv.ParseInt(getO2AndCO2rating(LessCommon, lines, 0), 2, 64)

	fmt.Printf("%v\n", n*m)
}
