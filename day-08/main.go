package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type entry struct {
	in  []string
	out []string
}

var digitMapOne = map[string]int{
	"abcefg":  0,
	"cf":      1,
	"acdeg":   2,
	"acdfg":   3,
	"bcdf":    4,
	"abdfg":   5,
	"abdefg":  6,
	"acf":     7,
	"abcdefg": 8,
	"abcdfg":  9,
}

var digitMapTwo = map[string]int{
	"dabcge":  0,
	"ab":      1,
	"dafgc":   2,
	"dafbc":   3,
	"efab":    4,
	"defbc":   5,
	"defgcb":  6,
	"dab":     7,
	"dafgcbe": 8,
	"dafebc":  9,
}

func solveFile(filename string) error {
	fd, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer fd.Close()
	reader := bufio.NewReader(fd)

	var entries []entry
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		if strings.TrimSpace(string(line)) == "" {
			continue
		}

		target := strings.Split(string(line), "|")
		entries = append(entries, entry{
			in:  strings.Split(strings.TrimSpace(target[0]), " "),
			out: strings.Split(strings.TrimSpace(target[1]), " "),
		})
	}

	// partONE(entries)
	partTWO(entries)

	return nil
}

func isIn(c rune, s string) (found bool) {
	strings.Map(func(r rune) rune {
		if r == c {
			found = true
		}
		return r
	}, s)
	return
}

type part uint64

const (
	ONE part = iota
	TWO
)

func isMatch(sig string, c rune, index int, res *[]int) {
	var currentSig = make([]byte, len(sig))
	if index == len(sig) {
		return
	}
	for i, n := range sig {
		if string(sig[index]) == string(n) {
			currentSig[i] = byte(c)
			continue
		}
		currentSig[i] = byte(n)
	}
	// fmt.Println(string(currentSig))

	//fmt.Println("length", len(string(currentSig)), string(currentSig), string(c), sig)

	matchCounter := 0
	for k, v := range digitMapTwo {
		for _, r := range string(currentSig) {
			if isIn(r, k) {
				matchCounter++
			}
		}
		if matchCounter == len(sig) {
			// log.Printf("here matchCounter == len(sig) %v %v", k, v)
			*res = append(*res, v)
		}
		matchCounter = 0
	}

	index += 1
	isMatch(sig, c, index, res)
}

func replaceChars(chars, sig string, isFirst bool) int {

	results := []int{}
	min := math.MaxInt64
	max := math.MinInt64

	if !isFirst {
		for _, c := range chars {
			if !isIn(c, sig) {
				isMatch(sig, c, 0, &results)
			}
		}
		for _, n := range results {
			if n < min && n != 0 && n != 1 && n != 4 && n != 7 && n != 8 {
				min = n
			}
		}
		return min
	}
	for _, c := range "deafgcb" {
		if !isIn(c, sig) {
			isMatch(sig, c, 0, &results)
		}
	}
	for _, n := range results {
		if n > max && n != 0 && n != 1 && n != 4 && n != 7 && n != 8 {
			max = n
		}
	}
	return max
	//fmt.Println(results)
}

func checkSignals(sig string, num *int, p part, alreadyFoundChars *string) bool {

	var matchCounter int
	var m map[string]int
	switch p {
	case ONE:
		m = digitMapOne
	case TWO:
		m = digitMapTwo
	}
	newL := (*alreadyFoundChars == "")
	for k, v := range m {
		for _, c := range sig {
			if isIn(c, k) {
				if !isIn(c, *alreadyFoundChars) {
					*alreadyFoundChars += string(c)
				}
				matchCounter++
			}
		}

		if p == ONE {
			if matchCounter == len(k) && len(sig) == len(k) && (v == 1 || v == 4 || v == 7 || v == 8) {
				*num = v
				return true
			}
		} else {
			if matchCounter == len(k) && len(sig) == len(k) && (v == 1 || v == 4 || v == 7 || v == 8) {
				*num = v
				return true
			}

			//fmt.Println(matchCounter, sig)

			if matchCounter == len(k) && len(sig) == len(k) && v != 0 {
				fmt.Println(sig, v, newL)
				//if newL {
				//	*num = v
				//	return true
				//}
				//if n := replaceChars(*alreadyFoundChars, sig, newL); n != 0 {
				//	*num = n
				//}
				//return true
			}

		}

		if matchCounter == len(sig) && matchCounter <= 4 {
			*alreadyFoundChars = ""
			for k, v := range m {
				if len(k) == matchCounter {
					*num = v
				}
			}
			return true
		}

		matchCounter = 0
	}
	if n := replaceChars(*alreadyFoundChars, sig, newL); n != 0 {
		*num = n
		return true
	}
	return false
}

func partONE(entries []entry) {
	var res int
	for _, e := range entries {
		var null string
		for _, sig := range e.out {
			var n int
			if checkSignals(sig, &n, ONE, &null) {
				res++
			}
		}
	}
	fmt.Println("part one", res)
}

func partTWO(entries []entry) {
	for _, e := range entries {
		var alreadyFoundChars string
		var nums string
		for _, sig := range e.out {
			var n int
			if checkSignals(sig, &n, TWO, &alreadyFoundChars) {
				nums += fmt.Sprintf("%d", n)
			}
		}
		num, _ := strconv.Atoi(nums)
		//if i == 2 {
		//	break
		//}
		fmt.Println(num)
		//fmt.Println("already FOUND CHARS", alreadyFoundChars)
	}
}

func main() {

	if len(os.Args) < 2 {
		log.Fatal("no input file provided")
	}

	for _, filename := range os.Args[1:] {
		if err := solveFile(filename); err != nil {
			log.Fatal(err)
		}
	}

}
