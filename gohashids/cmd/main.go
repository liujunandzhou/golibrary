package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/liujunandzhou/golibrary/gohashids/nhashids"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	for {

		input, err := reader.ReadString('\n') // hello world

		if err != nil {
			log.Printf("reader.ReadString failed,err=%v\n", err)
			break
		}

		input = strings.TrimSpace(input) // 去除首尾空格或换行符

		num, err := nhashids.Decode(input)
		if err != nil {
			log.Printf("nhashids.Decode failed,input=%s err=%v\n", input, err)
			os.Exit(1)
		}

		fmt.Printf("input=%s output=%d\n", input, num)
	}
}
