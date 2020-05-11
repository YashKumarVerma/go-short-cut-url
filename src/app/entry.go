package main

import (
	"YashKumarVerma/go-short-cut-url/src/shortener"
	"fmt"
)

func main() {
	fmt.Printf("ShortUrl : %s \n", shortener.Shorten("http://www.github.com/yashkumarverma"))
}
