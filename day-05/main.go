package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func getMinAndMax(min, max *int, cord Cord) {

	if cord.x1 < *min {
		*min = cord.x1
	}
	if cord.x1 > *max {
		*max = cord.x1
	}
	if cord.y1 < *min {
		*min = cord.y1
	}
	if cord.y1 > *max {
		*max = cord.y1
	}

	if cord.x2 < *min {
		*min = cord.x2
	}
	if cord.x2 > *max {
		*max = cord.x2
	}
	if cord.y2 < *min {
		*min = cord.y2
	}
	if cord.y2 > *max {
		*max = cord.y2
	}
}

func initBoard(board *[][]string, max int) {
	for i, r := range *board {
		r = make([]string, max)
		for c := 0; c < max; c++ {
			r[c] = "*"
		}
		(*board)[i] = r
	}
}

func solveFile(filename string) {

	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	var cords []Cord

	var min int = math.MaxInt
	var max int = math.MinInt

	for _, l := range strings.Split(string(data), "\n") {
		if strings.TrimSpace(l) == "" {
			continue
		}

		parts := strings.Split(l, "->")
		if len(parts) < 2 {
			panic("wrong data")
		}

		a1, b1 := strings.TrimSpace(strings.Split(parts[0], ",")[0]), strings.TrimSpace(strings.Split(parts[0], ",")[1])
		x1, _ := strconv.Atoi(a1)
		y1, _ := strconv.Atoi(b1)

		a2, b2 := strings.TrimSpace(strings.Split(parts[1], ",")[0]), strings.TrimSpace(strings.Split(parts[1], ",")[1])
		x2, _ := strconv.Atoi(a2)
		y2, _ := strconv.Atoi(b2)

		getMinAndMax(&min, &max, Cord{x1, x2, y1, y2})

		cords = append(cords, Cord{
			x1: x1,
			y1: y1,
			x2: x2,
			y2: y2,
		})
	}

	_ = min

	board := make([][]string, max+1)
	initBoard(&board, max+1)

	// NOTE: run each part at once.

	//fmt.Println("part one", partONE(cords, &board))
	fmt.Println("part two", partTWO(cords, &board))

	//	for _, n := range board {
	//		fmt.Println(n)
	//	}

}

// Cord
type Cord struct {
	x1, y1 int
	x2, y2 int
}

func partONE(fileContent []Cord, board *[][]string) (result int) {
	var count int
	for _, c := range fileContent {
		var s1, s2 bool
		if count >= 2 {
			break
		}
		if s1 && s2 {
			count++
		}
		if c.x1 == c.x2 {
			max := math.Max(float64(c.y1), float64(c.y2))
			min := math.Min(float64(c.y1), float64(c.y2))

			for r := min; r < max+1; r++ {

				if (*board)[int(r)][c.x1] == "*" {
					(*board)[int(r)][c.x1] = "1"
				} else {
					n, _ := strconv.Atoi((*board)[int(r)][c.x2])
					if n < 2 {
						n++
						(*board)[int(r)][c.x2] = strconv.Itoa(n)
						result++
						s1 = true
					}
				}

			}
		}

		if c.y1 == c.y2 {
			max := math.Max(float64(c.x1), float64(c.x2))
			min := math.Min(float64(c.x1), float64(c.x2))

			for r := min; r < max+1; r++ {
				if (*board)[c.y1][int(r)] == "*" {
					(*board)[c.y1][int(r)] = "1"
				} else {
					n, _ := strconv.Atoi((*board)[c.y1][int(r)])
					if n < 2 {
						n++
						(*board)[c.y1][int(r)] = strconv.Itoa(n)
						result++
						s2 = true
					}
				}
			}
		}

	}
	return
}

func partTWO(cords []Cord, board *[][]string) (result int) {
	var count int
	type Point struct {
		A int
		B int
	}
	points := map[Point]bool{}
	for _, c := range cords {
		var s1, s2, s3 bool
		if count >= 3 {
			break
		}
		if s1 && s2 && s3 {
			count++
		}
		if c.x1 == c.x2 {
			max := math.Max(float64(c.y1), float64(c.y2))
			min := math.Min(float64(c.y1), float64(c.y2))

			for r := min; r < max+1; r++ {

				if (*board)[int(r)][c.x1] == "*" {
					(*board)[int(r)][c.x1] = "1"
				} else {
					n, _ := strconv.Atoi((*board)[int(r)][c.x2])
					if n < 2 {
						n++
						(*board)[int(r)][c.x2] = strconv.Itoa(n)
						p := Point{int(r), c.x2}
						if _, ok := points[p]; !ok {
							points[p] = true
						}
						// result++
						s1 = true
					}
				}

			}
		} else if c.y1 == c.y2 {
			max := math.Max(float64(c.x1), float64(c.x2))
			min := math.Min(float64(c.x1), float64(c.x2))

			for r := min; r < max+1; r++ {
				if (*board)[c.y1][int(r)] == "*" {
					(*board)[c.y1][int(r)] = "1"
				} else {
					n, _ := strconv.Atoi((*board)[c.y1][int(r)])
					if n < 2 {
						n++
						(*board)[c.y1][int(r)] = strconv.Itoa(n)
						p := Point{c.y1, int(r)}
						if _, ok := points[p]; !ok {
							points[p] = true
						}
						// result++
						s2 = true
					}
				}
			}
		} else if c.x1 == c.y1 && c.x2 == c.y2 {
			max := math.Max(float64(c.x1), float64(c.x2))
			min := math.Min(float64(c.x1), float64(c.x2))
			for r := min; r < max+1; r++ {
				if (*board)[int(r)][int(r)] == "*" {
					(*board)[int(r)][int(r)] = "1"
				} else {
					n, _ := strconv.Atoi((*board)[int(r)][int(r)])
					n++
					(*board)[int(r)][int(r)] = strconv.Itoa(n)
					p := Point{int(r), int(r)}
					if _, ok := points[p]; !ok {
						points[p] = true
					}
					//result++
					s3 = true
				}
			}
		} else if c.x1 == c.y2 && c.y1 == c.x2 {
			//fmt.Println("second here", c.x1, c.x2, c.y1, c.y2)
			max := math.Max(float64(c.x1), float64(c.x2))
			min := math.Min(float64(c.x1), float64(c.x2))

			cd := max
			for r := min; r < max+1; r++ {
				if (*board)[int(r)][int(cd)] == "*" {
					(*board)[int(r)][int(cd)] = "1"
				} else {
					n, _ := strconv.Atoi((*board)[int(r)][int(cd)])
					n++
					// if n >= 2 {
					(*board)[int(r)][int(cd)] = strconv.Itoa(n)
					p := Point{int(r), int(cd)}
					if _, ok := points[p]; !ok {
						points[p] = true
					}
					//result++
					s3 = true
					//}
				}
				cd--
			}
		} else {

			rmax := math.Max(float64(c.y1), float64(c.y2))
			rmin := math.Min(float64(c.y1), float64(c.y2))
			//
			cmax := math.Max(float64(c.x1), float64(c.x2))
			cmin := math.Min(float64(c.x1), float64(c.x2))

			if (c.x1 < c.x2) && (c.y1 < c.y2) || (c.x2 < c.x1) && (c.y2 < c.y1) {
				for rmin <= rmax && cmin <= cmax {
					if (*board)[int(rmin)][int(cmin)] == "*" {
						(*board)[int(rmin)][int(cmin)] = "1"
					} else {
						n, _ := strconv.Atoi((*board)[int(rmin)][int(cmin)])
						n++
						(*board)[int(rmin)][int(cmin)] = strconv.Itoa(n)
						p := Point{int(rmin), int(cmin)}
						if _, ok := points[p]; !ok {
							points[p] = true
						}
						//result++
					}
					rmin++
					cmin++
				}
			} else {
				for rmin <= rmax || cmin <= cmax {
					if (*board)[int(rmin)][int(cmax)] == "*" {
						(*board)[int(rmin)][int(cmax)] = "1"
					} else {
						n, _ := strconv.Atoi((*board)[int(rmin)][int(cmax)])
						n++
						(*board)[int(rmin)][int(cmax)] = strconv.Itoa(n)
						p := Point{int(rmin), int(cmax)}
						if _, ok := points[p]; !ok {
							points[p] = true
						}
						//result++
					}
					cmin++
					cmax--
					rmin++
				}
			}
		}

	}
	result = len(points)
	return
}

func main() {

	if len(os.Args) < 2 {
		log.Printf("no file name specified")
		return
	}

	for _, file := range os.Args[1:] {
		solveFile(file)
	}

}

//
//		fmt.Println(c.x1, c.y1, c.x2, c.y2)
