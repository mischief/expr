package main

import (
	"bufio"
	"github.com/mischief/expr"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)

	l := float64(0)

	fmt.Print("-> ")
	for sc.Scan() {
		expr.Idents["line"] = expr.Number(l)
		l++
		if rs := expr.Proteval(sc.Text()); rs != "" {
			fmt.Println(rs)
		}
		fmt.Print("-> ")
	}
}
