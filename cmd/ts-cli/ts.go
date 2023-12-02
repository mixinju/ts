package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	set := flag.NewFlagSet("demo", 1)

	h := flag.String("help", "df", "fff")
	err := set.Parse(os.Args[2:])

	fmt.Println(h)
	if err != nil {
		fmt.Println("hello")
	}

}
