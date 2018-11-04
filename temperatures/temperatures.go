package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var tt, na, pa []int
	var filename string
	var n int
	arg := os.Args[1:]

	if len(arg) > 0 {
		filename = arg[0]
	} else {
		os.Exit(-1)
	}

	n, tt = load(filename)

	/*
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)
	var n int
	scanner.Scan()
	fmt.Sscan(scanner.Text(),&n)

	scanner.Scan()
	inputs := strings.Split(scanner.Text()," ")
	*/
	for i := 0; i < n; i++ {
		//t,_ := strconv.ParseInt(inputs[i],10,32)
		t := tt[i]
		if t == 0 {continue}
		if t < 0 {na = append(na, t); continue}
		if t > 0 {pa = append(pa, t); continue}
	}

	neg_val := min(na, true)
	pos_val := min(pa, false)

	if (neg_val == 0) {
		fmt.Println(pos_val)
	} else if (pos_val != 0 && pos_val <= (-1)*neg_val) {
		fmt.Println(pos_val)
	} else {
		fmt.Println(neg_val)
	}
}

func min(a []int, negative bool)(int) {
	if len(a) == 0 {
		return 0
	}
	var resp int
	resp = a[0]
	for i := 0; i < len(a); i++ {
		if negative {
			if resp < a[i] {
				resp = a[i]
			}
		} else {
			if resp > a[i] {
				resp = a[i]
			}
		}
	}
	return resp
}

func load(datafile string) (int, []int) {
	var n int

	file, err := os.Open(datafile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &n)
	tt := make([]int, n)

	scanner.Scan()
	inputs := strings.Split(scanner.Text()," ")

	for i := 0; i < n; i++ {
		t,_ := strconv.ParseInt(inputs[i],10,32)
		tt[i] = int(t)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return n, tt
}