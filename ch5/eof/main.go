package main

import (
	"bufio"
	"os"
	"io"
	"fmt"
)

func main()  {
	in := bufio.NewReader(os.Stdin)

	for {
		r, _, err := in.ReadRune()
		if err == io.EOF {
			fmt.Println("bye")
			break
		}
		if err != nil {
			fmt.Errorf("read failed: %v", err)
			os.Exit(1)
		}
		fmt.Printf("%s\n", r)
	}
}
