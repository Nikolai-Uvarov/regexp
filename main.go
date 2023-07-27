package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {

	var inpfile, outpfile string

	fmt.Println(os.Args)

	if len(os.Args) > 1 {
		inpfile = os.Args[1]
	} else {
		fmt.Println("no input file specified")
		os.Exit(1)
	}
	if len(os.Args) > 2 {
		outpfile = os.Args[2]
	} else {
		outpfile = "output.txt"
	}

	f, err := os.OpenFile(inpfile, os.O_RDONLY, 0777)
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(f)
	defer f.Close()

	o, err := os.OpenFile(outpfile, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	writer := bufio.NewWriter(o)
	defer o.Close()

	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		_, err = writer.Write([]byte(string(line) + "\n"))
		if err != nil {
			panic(err)
		}
	}
	err = writer.Flush()
	if err != nil {
		panic(err)
	}

}
