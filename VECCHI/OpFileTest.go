package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	err := os.Mkdir("tempF", 0755)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Create("tempF/testfile.txt")
	if err != nil {
		log.Fatal(err)
	}

	_, err = file.WriteString("Ciao")
	if err != nil {
		log.Fatal(err)
	}

	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}

	data, err := os.ReadFile("tempF/testfile.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
}
