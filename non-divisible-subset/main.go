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

/*
 * Complete the 'nonDivisibleSubset' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER k
 *  2. INTEGER_ARRAY s
 */
func nonDivisibleSubset(mod int32, s []int32) int32 {
	var j, max int32
	sMap := make(map[int32]int32)

	for i:=0; i<len(s); i++ {
		sMap[s[i] % mod]++
	}

	if mod % 2 == 0 {
		sMap[mod/2] = int32(math.Min(float64(sMap[mod/2]), 1))
	}

	max = int32(math.Min(float64(sMap[0]), 1))
	for j=1; j<=mod/2; j++ {
		left  := float64(sMap[j])
		right := float64(sMap[mod-j])
		max += int32(math.Max(left, right))
	}

	return max
}

func main() {
	var filename string
	arg := os.Args[1:]
	if len(arg) > 0 {
		filename = arg[0]
	} else {
		os.Exit(-1)
	}

	k, s := load(filename)
	result := nonDivisibleSubset(k, s)
	fmt.Printf("%d\n", result)
}

func load(filename string) (int32, []int32){
	var n, k int32
	var s []int32

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 1024*1024)

	scanner.Scan()
	inputs := strings.Split(scanner.Text()," ")

	nTemp, _ := strconv.ParseInt(inputs[0],10,32)
	kTemp, _ := strconv.ParseInt(inputs[1],10,32)
	n = int32(nTemp)
	k = int32(kTemp)

	scanner.Scan()
	inputData := strings.Split(scanner.Text()," ")
	for i:=0; i<int(n); i++ {
		t, err := strconv.ParseInt(inputData[i],10,64)
		if err != nil {
			fmt.Printf("num: %d | data %s | Err: %s\n", i, inputData[i], err.Error())
			// os.Exit(-1)
			continue
		}
		s = append(s, int32(t))
	}

	return k, s
}
