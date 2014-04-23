package main

import (
	"bufio"
	"fmt"
	"os"
)

func check() {

	f, err := os.Open("config.txt")

	if err != nil {
		fmt.Println("Error opening configuration file ", err)
		os.Exit(1)
	}

	defer f.Close()

	r := bufio.NewReaderSize(f, 4*1024)

	line, isPrefix, err := r.ReadLine()
	for err == nil && !isPrefix {
		s := string(line)
		fmt.Println()
		fmt.Println(s)
		line, isPrefix, err = r.ReadLine()

		check, err := os.Stat(s)
		if err != nil {
			panic(err)
		}
		since := check.ModTime()
		size := check.Size()
		fmt.Print(since)
		fmt.Print(" - ")
		mb := float64(size) / float64(1024) / float64(1024)
		fmt.Printf("%.2f", mb)
		fmt.Print("MB")
		fmt.Println()
	}

}

func checkPath(file string) {

	check, err := os.Stat(file)
	if err != nil {
		panic(err)
	}
	since := check.ModTime()
	size := check.Size()

	fmt.Print(since)
	fmt.Print(" - ")
	mb := float64(size) / float64(1024) / float64(1024)
	fmt.Printf("%.2f", mb)
	fmt.Print("MB")

}

func main() {

	args := os.Args[1:]

	if len(os.Args[1:]) != 0 {

		checkPath(args[0])

	} else {

		check()

	}

}
