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

type Command struct {
	Value string
	Steps int
}

func solveFile(filename string) {
	fd, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	commands := []Command{}
	reader := bufio.NewReader(fd)
	for {

		line, err := reader.ReadString('\n')
		if err != nil && err == io.EOF {
			break
		}

		if err != nil {
			panic(err)
		}

		fields := strings.Fields(line)
		step, _ := strconv.Atoi(fields[1])
		commands = append(commands, Command{
			Value: fields[0],
			Steps: step,
		})

	}

	// partONE(commands)
	partTWO(commands)
}

func partONE(commands []Command) {

	var (
		horizontal int
		depth      int
	)

	for _, n := range commands {
		switch n.Value {
		case "forward":
			horizontal += n.Steps
		case "down":
			depth += n.Steps
		case "up":
			depth -= n.Steps
		}
	}

	fmt.Printf("horizontal: %v, depth: %v, mult: %v\n", horizontal, depth, horizontal*depth)
}

func partTWO(commands []Command) {

	var horizontal, depth, aim int

	for _, n := range commands {
		switch n.Value {
		case "forward":
			horizontal += n.Steps
			depth += (aim * n.Steps)
		case "up":
			aim -= n.Steps
		case "down":
			aim += n.Steps
		}
	}

	fmt.Printf("horizontal: %v, depth: %v, aim %v,  mult: %v\n", horizontal, depth, aim, horizontal*depth)
}
