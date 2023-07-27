package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

func main() {

	//парсинг аргументов приложения
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

	//открытие файла на чтение и создание ридера
	f, err := os.OpenFile(inpfile, os.O_RDONLY, 0777)
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(f)
	defer f.Close()

	//открытие файла на запись (создание если не существует) и создание райтера
	o, err := os.OpenFile(outpfile, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	writer := bufio.NewWriter(o)
	defer o.Close()

	//создание регулярного выражения для поиска арифметических выражений
	re := regexp.MustCompile("([0-9]+)([+*/-]{1})([0-9]+)")

	//поиск арифметических выражений регулярками и запись результата в буфер
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}

		sm := re.FindAllStringSubmatch(string(line), 1)
		if sm == nil {
			continue
		}

		fmt.Println(sm[0][0], sm[0][1], sm[0][2], sm[0][3])

		operation := sm[0][0]
		leftoperand, _ := strconv.Atoi(sm[0][1])
		operator := sm[0][2]
		rightoperand, _ := strconv.Atoi(sm[0][3])
		var result int

		switch operator {
		case "+":
			result = leftoperand + rightoperand
		case "-":
			result = leftoperand - rightoperand
		case "*":
			result = leftoperand * rightoperand
		case "/":
			result = leftoperand / rightoperand
		}

		_, err = writer.Write([]byte(operation + "=" + fmt.Sprint(result) + "\n"))
		if err != nil {
			panic(err)
		}
	}
	//запись буфера в файл
	err = writer.Flush()
	if err != nil {
		panic(err)
	}

}
