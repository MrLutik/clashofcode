package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func pyramidBase(r int) int {
	var res int = 1
	for i := 1; i < r; i++ {
		res += 2
	}
	return res
}

func pyramidBuild(raws, base int, char string) [][]string {
	var arr [][]string
	for x := 0; x < raws; x++ {
		var a []string
		for y := 0; y < base; y++ {
			a = append(a, char)
		}
		arr = append(arr, a[x:len(a)-x])
	}
	return arr
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	var N int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &N)

	scanner.Scan()
	C := scanner.Text()
	_ = C

	pyramid := pyramidBuild(N, pyramidBase(N), C)
	for p := len(pyramid) - 1; p >= 0; p-- {
		for i := p - 1; i >= 0; i-- {
			fmt.Printf(" ")
		}
		fmt.Printf("%s\n", strings.Join(pyramid[p], ""))
	}

}
