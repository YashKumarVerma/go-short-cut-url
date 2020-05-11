package main

import (
	"YashKumarVerma/go-short-cut-url/src/shortener"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Pass the url to be shortened")
		return
	}

	fmt.Printf("ShortUrl : %s \n", shortener.Shorten(os.Args[1]))
}
