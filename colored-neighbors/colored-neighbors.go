package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

var rows, columns int

func main() {
	var filename string
	var matrix [][]int

	handled := make(map[string]bool)
	result := make(map[int]int) // neighbor color index, neighbors amount

	arg := os.Args[1:]
	if len(arg) > 0 {
		filename = arg[0]
	} else {
		os.Exit(-1)
	}

	rows, columns, matrix = load(filename)

	for i:=0; i<rows; i++ {
		for j:=0; j<columns; j++ {
			h := hash(i, j)
			shouldSkip := handled[h]
			if shouldSkip {continue}

			handled[h] = true

			current := matrix[i][j]

			neighbors := getNeighbors(i, j, matrix)

			for _, value := range neighbors {
				if current == value[2] {
					if _, ok := result[current]; !ok {
						result[current] = 1
					}
					result[current] += 1
				}

				nh := hash(value[0], value[1])
				handled[nh] = true
			}
		}
	}

	type kv struct {
		Key   int
		Value int
	}

	var ss []kv
	for k, v := range result {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})

	for _, kv := range ss {
		fmt.Printf("Found max neighbors - %d for color #%d\n", kv.Value, kv.Key)
		break
	}
}

func hash(i int, j int) string {
	return fmt.Sprintf("%d%d", i, j)
}

// return array of neighbors
// "top": 	[row, col, color]
// "right": [row, col, color]
// "bottom":[row, col, color]
// "left":	[row, col, color]
func getNeighbors (row int, col int, m [][]int) map[string][3]int {
	neighbors := make(map[string][3]int)

	if col > 0	{
		neighbors["top"] = [3]int{row, col-1, m[row][col-1]}
	}

	if col < columns - 1	{
		neighbors["bottom"] = [3]int{row, col+1, m[row][col+1]}
	}

	if row > 0 {
		neighbors["left"] = [3]int{row-1, col, m[row-1][col]}
	}

	if row < rows -1 {
		neighbors["right"] = [3]int{row+1, col, m[row+1][col]}
	}

	return neighbors
}

func load(filename string) (int, int, [][]int){
	var rows, cols int
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	fmt.Sscan(scanner.Text(), &rows)
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &cols)

	matrix := make([][]int, rows)
	for i := range matrix {
		matrix[i] = make([]int, cols)

		scanner.Scan()
		inputs := strings.Split(scanner.Text()," ")
		for num := 0; num < cols; num++ {
			t,_ := strconv.ParseInt(inputs[num],10,32)
			matrix[i][num] = int(t)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return rows, cols, matrix
}
