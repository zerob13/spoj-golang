package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var number int
	inp := bufio.NewReader(os.Stdin)
	fmt.Fscanln(inp, &number)
	// b := bufio.NewWriter(os.Stdout)
	for i := 0; i < number; i++ {
		var ex int
		// fmt.Scan(&ex)
		fmt.Fscanln(inp, &ex)
		if ex < 5 {
			// fmt.Println(0)
			fmt.Println(0)
			continue
		}
		ans := 0

		for base := 5; ex >= base; base *= 5 {
			ans += ex / base
		}

		fmt.Println(ans)
		// fmt.Fprintln(b, ans)
	}
}
