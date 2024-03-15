package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	//ccwc -c test.txt

	if len(os.Args) < 3 {
		log.Fatal("incomplete arguments")
	}

	if os.Args[1] != "ccwc" {
		log.Fatalf("%v is incorrect for the first argument", os.Args[1])
	}

	if len(os.Args) == 3 {
		fileInfo, err := os.Stat(os.Args[2])
		if err != nil {
			if os.IsNotExist(err) {
				// log.Fatalf("%v does not exist", os.Args[2])
				scanner := bufio.NewScanner(os.Stdin)

				scanner.Split(bufio.ScanLines)

				// getting the number of lines
				var count int
				for scanner.Scan() {
					count++
				}
				if err := scanner.Err(); err != nil {
					log.Fatal(err)
				}
				fmt.Printf("%v \n", count)
				return
			}
		}
		byteSize := fileInfo.Size()

		// use goroutine to read and scan files simultatneoulsy

		lineCh := make(chan int)
		wordCh := make(chan int)

		go func(ch chan int) {

			file, err := os.Open(os.Args[2])
			if err != nil {
				log.Fatal(err)
				return
			}
			defer file.Close()

			scanner := bufio.NewScanner(file)
			scanner.Split(bufio.ScanLines)

			var linecount int

			for scanner.Scan() {
				linecount++
			}
			if err := scanner.Err(); err != nil {
				log.Fatal(err)
			}

			ch <- linecount

		}(lineCh)

		go func(ch chan int) {

			file, err := os.Open(os.Args[2])
			if err != nil {
				log.Fatal(err)
				return
			}
			defer file.Close()

			scanner := bufio.NewScanner(file)
			scanner.Split(bufio.ScanWords)

			var wordcount int

			for scanner.Scan() {
				wordcount++
			}
			if err := scanner.Err(); err != nil {
				log.Fatal(err)
			}

			ch <- wordcount

		}(wordCh)

		// get the values from the chan
		lineCount := <-lineCh
		wordCount := <-wordCh

		fmt.Printf("%v %v %v %v \n", byteSize, lineCount, wordCount, os.Args[2])
	}

	action := os.Args[2]

	switch action {
	case "-c":
		fileInfo, err := os.Stat(os.Args[3])
		if err != nil {
			if os.IsNotExist(err) {
				log.Fatalf("%v does not exist", os.Args[3])
			}
		}
		fmt.Printf("%v %v \n", fileInfo.Size(), os.Args[3])

	case "-l":
		file, err := os.Open(os.Args[3])
		if err != nil {
			log.Fatal(err)
			return
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)

		// getting the number of lines
		var count int
		for scanner.Scan() {
			count++
		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%v %v \n", count, os.Args[3])

	case "-w":
		file, err := os.Open(os.Args[3])
		if err != nil {
			log.Fatal(err)
			return
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanWords)

		// getting the number of words
		var count int
		for scanner.Scan() {
			count++
		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%v %v \n", count, os.Args[3])

	case "-m":
		file, err := os.Open(os.Args[3])
		if err != nil {
			log.Fatal(err)
			return
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanRunes)

		// getting the number of ascii characters
		var count int
		for scanner.Scan() {
			count++
		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%v %v \n", count, os.Args[3])
	}

}
