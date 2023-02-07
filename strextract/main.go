package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	n    int
	word string
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	scanner.Scan()
	fmt.Sscan(scanner.Text(), &word)

	scanner.Scan()
	fmt.Sscan(scanner.Text(), &n)

	var res [][]string
	for i := 0; i < n; i++ {
		scanner.Scan()
		var text string
		fmt.Sscan(scanner.Text(), &text)

		d := strings.Split(text, word)
		res = append(res, d)
	}
	var str string = ""
	for _, v := range res {
		str += strings.Join(v, "")
	}
	fmt.Println(str)
}
