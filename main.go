package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

func gen(length int) string {
	var source = "abcdefghijkpqrstuvwxyz" +
		"2345679" +
		"ABCDEFGHIJKLMNPQRSTUVWXYZ"
	var retval []byte
	rand.Seed(time.Now().UnixNano())
	retval = make([]byte, length, length)
	for i := 0; i < length; i++ {
		retval[i] = source[rand.Intn(len(source))]
	}
	return string(retval)
}

func main() {
	var length int
	flag.IntVar(&length, "length", 8, "length of password")
	flag.Parse()

	fmt.Println(gen(length))
}
