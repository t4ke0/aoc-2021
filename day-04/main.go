package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	args := os.Args
	if len(args) < 2 {
		log.Fatal("no files provided")
	}

	defer func() {
		if r := recover(); r != nil {
			log.Fatalf("AOC Error:  %v", r)
		}
	}()

	for _, f := range args[1:] {
		solveFile(f, ONE)
		solveFile(f, TWO)
	}

}

type Part uint32

const (
	ONE Part = iota
	TWO
)

type Board [][]int

type Boards []Board

func solveFile(filename string, part Part) {
	fd, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer fd.Close()

	reader := bufio.NewReader(fd)

	count := 0
	parseBoard := false
	board := Board{}
	boards := Boards{}
	var numbers []string
	for {
		line, err := reader.ReadString('\n')
		if err != nil && err == io.EOF {
			boards = append(boards, board)
			break
		}
		if err != nil {
			panic(err)
		}

		if !parseBoard && line != "\n" {
			numbers = strings.Split(strings.TrimSpace(line), ",")
			count++
			continue
		}

		//if count == 0 {
		//	continue
		//}

		if line == "\n" {
			if parseBoard {
				boards = append(boards, board)
				board = nil
			}
			parseBoard = true
			count++
			continue
		}

		if parseBoard {
			tmp := []int{}
			for _, n := range strings.Split(strings.TrimSpace(line), " ") {
				n = strings.TrimSpace(n)
				if n == "" {
					continue
				}
				num, _ := strconv.Atoi(n)
				tmp = append(tmp, num)
			}
			board = append(board, tmp)
		}

		count++
	}

	m := make(map[int][]Position)

	wonBoards := []int{}
	log.Printf("%v", filename)
	for _, n := range numbers {
		for index, b := range boards {
			checkBoard(index, b, n, &m)
			ok := b.validateBoard(index, m)

			if part == TWO {
				if ok {
					found := false
					for _, wn := range wonBoards {
						if index == wn {
							found = true
						}
					}
					if !found {
						if len(wonBoards) == len(boards)-1 {
							result := 0
							for _, n := range boards[index].getUnmarkedNums(index, m) {
								result += n
							}
							num, _ := strconv.Atoi(n)
							fmt.Printf("PART TWO %v\n", result*num)
							fmt.Println()
							return
						}
						wonBoards = append(wonBoards, index)
					}
				}
			}
			if part == ONE {
				if ok {
					result := 0
					for _, n := range b.getUnmarkedNums(index, m) {
						result += n
					}
					num, _ := strconv.Atoi(n)
					fmt.Printf("PART ONE %v\n", result*num)
					fmt.Println()
					return
				}
			}
		}
	}

}

// Position
type Position struct {
	Row int
	Col int
}

func checkBoard(boardIndex int, board Board, currentNumber string, m *map[int][]Position) {
	for r, _ := range board {
		for c, _ := range board[r] {
			number, _ := strconv.Atoi(currentNumber)
			if board[r][c] == number {
				_, ok := (*m)[boardIndex]
				if !ok {
					(*m)[boardIndex] = []Position{
						Position{
							Row: r,
							Col: c,
						},
					}
					return
				}

				(*m)[boardIndex] = append((*m)[boardIndex], Position{
					Row: r,
					Col: c,
				})
				return
			}
		}
	}
}

func (b Board) validateBoard(index int, m map[int][]Position) bool {

	rows := []int{}
	cols := []int{}
	for r, _ := range b {
		for c, _ := range b[r] {
			for _, n := range m[index] {
				if n.Row == r && n.Col == c {
					rows = append(rows, r)
					cols = append(cols, c)
					continue
				}
			}
		}
	}

	// TODO: find another way to do that.
	rowsCounter := map[int]int{}
	for _, r := range rows {
		rowsCounter[r]++
	}
	for _, v := range rowsCounter {
		if v == len(b[0]) {
			return true
		}
	}

	colsCounter := map[int]int{}
	for _, c := range cols {
		colsCounter[c]++
	}
	for _, v := range colsCounter {
		if v == len(b[0]) {
			return true
		}
	}

	return false

}

func (b Board) getUnmarkedNums(index int, m map[int][]Position) (out []int) {

	for r, _ := range b {
		for c, _ := range b[r] {
			if !isMarked(index, m, r, c) {
				out = append(out, b[r][c])
			}
		}
	}

	return
}

func isMarked(index int, m map[int][]Position, row, col int) bool {
	for _, v := range m[index] {
		if v.Col == col && v.Row == row {
			return true
		}
	}
	return false
}
