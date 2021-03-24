package main

import (
	"bufio"
	"io"
	"math/rand"
	"net/http"
	"os"
	"strings"
)

func index(rw http.ResponseWriter, r *http.Request) {
	dongers, err := readList("dongers.txt")
	if err != nil {
		panic(err)
	}

	pick := rand.Intn(len(dongers))
	io.WriteString(rw, dongers[pick]+"\n")
}

func readList(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8443", nil)
}
