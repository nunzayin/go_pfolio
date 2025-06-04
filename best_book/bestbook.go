package main

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	scn := bufio.NewScanner(os.Stdin)
	scn.Scan()
	fmt.Println(scn.Text(), "- лучшая книга!")
}
