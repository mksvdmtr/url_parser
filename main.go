package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	urls := os.Args[1]
	file, err := os.Open(urls) // open file with urls
	if err != nil {
		fmt.Println("Error", err)
	}

	result, err := os.Create("result.txt")
	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		resp, err := http.Head(scanner.Text())
		time.Sleep(500 * time.Millisecond)
		if err != nil {
			fmt.Println("error", err)
		} else {
			id := resp.Request.URL.Path
			hru := resp.Request.Header.Get("Referer")
			result.WriteString(id + " -> " + hru + "\n")

			fmt.Println(id, hru)

		}

	}
	file.Close()
	result.Close()

}
