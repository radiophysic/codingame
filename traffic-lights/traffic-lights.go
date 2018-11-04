package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type L struct {
	Dist float64
	Delay float64
}

func main() {
	// {{ Load data from file
	arg := os.Args[1:]
	var filename string
	if len(arg) > 0 {
		filename = arg[0]
	} else {
		os.Exit(-1)
	}

	m := make([]L, 0)
	var speed int
	speed, _, m = load(m, filename)
	// }}

	/* loadt input for codingame.com:
	fmt.Scan(&speed)
	fmt.Scan(&lightCount)
	for i := 0; i < lightCount; i++ {
		var distance, duration int
		fmt.Scan(&distance, &duration)
		m = append(m, L{Dist: float64(distance), Delay: float64(duration)})
	}
	*/
	newspeed := calc(speed, m)
	fmt.Println(newspeed)// Write answer to stdout
}

func calc(speed int, m []L)(int) {
	var spent, ms, phases float64
	if speed == 0 {
		return 0
	}
	spent = 0.0
	phases = 0.0
	for _,item := range m {
		spent = float64(item.Dist) / float64(speed) * 3600 / 1000
		if (spent < float64(item.Delay)) {
			continue
		}
		phases = math.Round(spent) / float64(item.Delay)
		ms = math.Mod(float64(phases), 2)
		if (int(phases) >= 1 && int(ms) != 0) {
			return calc(speed-1, m)
		}
	}
	return speed
}


func load(m []L, datafile string) (int, int, []L) {
	var speed, lightCount int

	file, err := os.Open(datafile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var rows = 0

	for scanner.Scan() {

		if rows == 0 {
			speed, err = strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println(err)
				os.Exit(2)
			}
		} else if rows == 1 {
			lightCount, err = strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println(err)
				os.Exit(2)
			}
		} else {
			var s = strings.Fields(scanner.Text())
			var distance, duration int
			distance, _ = strconv.Atoi(s[0])
			duration, _ = strconv.Atoi(s[1])
			m = append(m, L{Dist: float64(distance), Delay: float64(duration)})
		}
		rows++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return speed, lightCount, m
}