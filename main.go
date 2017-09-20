package main

import (
	"bufio"
	"os"
	"fmt"
	"time"
)

const (
	YES = "She loves me"
	NO = "She loves me not"
)

func main() {

	go func() {
		bufio.NewReader(os.Stdin).ReadByte()
		fmt.Print("\033[F")
		os.Exit(0)
	}()

	for {
		fmt.Println(YES)
		time.Sleep(10 * time.Millisecond)
		fmt.Println(NO)
		time.Sleep(10 * time.Millisecond)
	}
}