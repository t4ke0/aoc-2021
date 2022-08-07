package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func solveFile(filename string) error {
	fd, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer fd.Close()

	reader := bufio.NewReader(fd)

	var tubes [][]string

	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		tubes = append(tubes, strings.Split(string(line), ""))
	}

	// partONE(tubes)
	partTWO(tubes)

	return nil
}

func partONE(tubes [][]string) {
	res := []int{}
	for rowIndex, r := range tubes {
		for colIndex, c := range r {
			n, _ := strconv.Atoi(c)
			var pattern []int
			if rowIndex == 0 && colIndex == 0 ||
				rowIndex == 0 && colIndex == len(r)-1 {
				switch colIndex {
				case 0:
					pattern = []int{1, 1}
				case len(r) - 1:
					pattern = []int{1, -1}
				}
				if c < tubes[rowIndex+pattern[0]][colIndex] &&
					c < tubes[rowIndex][colIndex+pattern[1]] {
					res = append(res, n)
				}
				continue
			}
			if rowIndex == len(tubes)-1 && colIndex == 0 ||
				rowIndex == len(tubes)-1 && colIndex == len(r)-1 {
				switch colIndex {
				case 0:
					pattern = []int{-1, 1}
				case len(r) - 1:
					pattern = []int{-1, -1}
				}
				if c < tubes[rowIndex+pattern[0]][colIndex] &&
					c < tubes[rowIndex][colIndex+pattern[1]] {
					res = append(res, n)
				}
				continue
			}

			if rowIndex == 0 || rowIndex == len(tubes)-1 {
				switch rowIndex {
				case 0:
					//			    r  c   c
					pattern = []int{1, 1, -1}
				case len(tubes) - 1:
					pattern = []int{-1, 1, -1}
				}
				if c < tubes[rowIndex+pattern[0]][colIndex] &&
					c < tubes[rowIndex][colIndex+pattern[1]] &&
					c < tubes[rowIndex][colIndex+pattern[2]] {
					res = append(res, n)
				}
				continue
			}

			if colIndex == 0 {
				//				r   r  c
				pattern = []int{-1, 1, 1}
			}

			if colIndex == len(r)-1 {
				//				r   r  c
				pattern = []int{-1, 1, -1}
			}

			if colIndex > 0 && colIndex < len(r)-1 {
				//				r   r   c  c
				pattern = []int{-1, 1, -1, 1}
			}

			compareNums := []int{}
			for i, _ := range pattern {
				if i <= 1 {
					cm, _ := strconv.Atoi(tubes[rowIndex+pattern[i]][colIndex])
					compareNums = append(compareNums, cm)
					continue
				}
				cm, _ := strconv.Atoi(tubes[rowIndex][colIndex+pattern[i]])
				compareNums = append(compareNums, cm)
			}
			if compareWithArr(n, compareNums) {
				res = append(res, n)
			}
		}
	}
	sum := 0
	for _, r := range res {
		sum += r + 1
	}
	fmt.Println(sum)
}

func compareWithArr(n int, arr []int) bool {
	r := 0
	for _, m := range arr {
		if n < m {
			r++
		}
	}
	if r == len(arr) {
		return true
	}
	return false
}

func partTWO(tubes [][]string) {

	points := []Point{}

	for r, _ := range tubes {
		for c, _ := range tubes[r] {
			pts := Point{r, c}.getPattern(len(tubes)-1, len(tubes[0])-1).toPoints(Point{r, c})
			count := 0
			for _, p := range pts {
				if strToInt(tubes[r][c]) < strToInt(tubes[p.Row][p.Col]) {
					count++
				}
			}
			if count == len(pts) {
				points = append(points, Point{r, c})
			}
		}
	}

	var r []int

	for _, n := range points {
		goodP := []Point{n}
		startDFS(n, tubes, &goodP)
		r = append(r, len(goodP))
	}

	sort.Ints(r)
	result := 1

	step := 0
	for i := len(r) - 1; i > 0; i-- {
		if step == 3 {
			break
		}
		result *= r[i]
		step++
	}

	fmt.Println(result)
}

func strToInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

type Pattern struct {
	x1, x2 int
	y1, y2 int
}

func (p Pattern) toPoints(orig Point) (pts []Point) {

	if p.x1 != math.MaxInt {
		pts = append(pts, Point{orig.Row + p.x1, orig.Col})
	}

	if p.x2 != math.MaxInt {
		pts = append(pts, Point{orig.Row + p.x2, orig.Col})
	}

	if p.y1 != math.MaxInt {
		pts = append(pts, Point{orig.Row, orig.Col + p.y1})
	}

	if p.y2 != math.MaxInt {
		pts = append(pts, Point{orig.Row, orig.Col + p.y2})
	}

	return
}

type Point struct {
	Row int
	Col int
}

func (p Point) isIn(pts []Point) bool {
	for _, n := range pts {
		if p.Row == n.Row && p.Col == n.Col {
			return true
		}
	}
	return false
}

/*
	x1: row up
	x2: row down
	y1: col right
	y2: col left
*/

func (p Point) getPattern(maxRow, maxCol int) Pattern {

	// handle points that are in the middle of the rectangle.
	if p.Col > 0 && p.Row > 0 && p.Col < maxCol && p.Row < maxRow {
		return Pattern{
			x1: -1,
			x2: 1,
			y1: 1,
			y2: -1,
		}
	}
	// handle point that is in the top left corner.
	if p.Col == 0 && p.Row == 0 {
		return Pattern{
			x1: math.MaxInt,
			x2: 1,
			y1: 1,
			y2: math.MaxInt,
		}
	}
	// handle point that is in the top right corner.
	if p.Col == maxCol && p.Row == 0 {
		return Pattern{
			x1: math.MaxInt,
			x2: 1,
			y1: math.MaxInt,
			y2: -1,
		}
	}
	// handle point that is in the bottom left corner.
	if p.Col == 0 && p.Row == maxRow {
		return Pattern{
			x1: -1,
			x2: math.MaxInt,
			y1: 1,
			y2: math.MaxInt,
		}
	}
	// handle point that is in the bottom right corner.
	if p.Col == maxCol && p.Row == maxRow {
		return Pattern{
			x1: -1,
			x2: math.MaxInt,
			y1: math.MaxInt,
			y2: -1,
		}
	}
	// handle points that are in the top edge of the rectangle.
	if p.Row == 0 && p.Col > 0 {
		return Pattern{
			x1: math.MaxInt,
			x2: 1,
			y1: 1,
			y2: -1,
		}
	}
	// handle points that are in the left verticale edge of the rectangle.
	if p.Col == 0 && p.Row > 0 {
		return Pattern{
			x1: -1,
			x2: 1,
			y1: 1,
			y2: math.MaxInt,
		}
	}
	// handle points that are  in the bottom edge of the rectangle.
	if p.Row == maxRow && p.Col > 0 {
		return Pattern{
			x1: -1,
			x2: math.MaxInt,
			y1: 1,
			y2: -1,
		}
	}
	// handle points that are in the right vertical edge of the rectangle.
	if p.Col == maxCol && p.Row > 0 {
		return Pattern{
			x1: -1,
			x2: 1,
			y1: math.MaxInt,
			y2: -1,
		}
	}

	return Pattern{}
}

func startDFS(start Point, tubes [][]string, goodP *[]Point) {

	queue := start.getPattern(len(tubes)-1, len(tubes[0])-1).toPoints(start)

	for _, q := range queue {
		if strToInt(tubes[start.Row][start.Col])-1 == strToInt(tubes[q.Row][q.Col]) ||
			strToInt(tubes[start.Row][start.Col])+1 == strToInt(tubes[q.Row][q.Col]) {
			if strToInt(tubes[q.Row][q.Col]) == 9 {
				continue
			}
			if q.isIn(*goodP) {
				continue
			}
			*goodP = append(*goodP, q)
			startDFS(q, tubes, goodP)
		}

	}
}

func removeFromQueue(index int, queue []Point) []Point {
	return append(queue[:index], queue[index+1:]...)
}

func main() {

	if len(os.Args) < 2 {
		log.Fatal("no input file provided")
	}

	for _, file := range os.Args[1:] {
		if err := solveFile(file); err != nil {
			log.Fatal(err)
		}
	}

}

//	if p.x1 != math.MaxInt {
//		*queue = append(*queue, Point{
//			Row: point.Row + p.x1,
//			Col: point.Col,
//		})
//	}
//	if p.x2 != math.MaxInt {
//		*queue = append(*queue, Point{
//			Row: point.Row + p.x2,
//			Col: point.Col,
//		})
//	}
//	if p.y1 != math.MaxInt {
//		*queue = append(*queue, Point{
//			Row: point.Row,
//			Col: point.Col + p.y1,
//		})
//	}
//	if p.y2 != math.MaxInt {
//		*queue = append(*queue, Point{
//			Row: point.Row,
//			Col: point.Col + p.y2,
//		})
//	}
//
//	for _, q := range *queue {
//
//		if len(*queue) == 0 {
//			return
//		}
//
//		//fmt.Println("queue", *queue)
//
//		*queue = removeFromQueue(0, *queue)
//
//		if strToInt(tubes[q.Row][q.Col]) == strToInt(tubes[point.Row][point.Col])+1 ||
//			strToInt(tubes[q.Row][q.Col]) == strToInt(tubes[point.Row][point.Col])-1 {
//			fmt.Println(tubes[(*queue)[0].Row][(*queue)[0].Col])
//			startDFS((*queue)[0], queue, tubes)
//		}
//
//	}
