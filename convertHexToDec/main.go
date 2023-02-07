package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	scanner.Scan()
	var N string
	fmt.Sscan(scanner.Text(), &N)

	num, _ := strconv.ParseInt(N, 16, 64)

	fmt.Println(num)

}
