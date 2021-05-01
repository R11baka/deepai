package main

import (
	"deepai"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	dp := deepai.New(os.Getenv("DEEP_AI_KEY"), nil)
	filePath := os.Args[1]
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error happened:", err)
		return
	}
	result, err := dp.Colorize(content)
	if err != nil {
		fmt.Println("Can't colorize photo ", err)
		return
	}
	errWrite := ioutil.WriteFile("test.jpg", result, 0644)
	if errWrite != nil {
		fmt.Println("Erro when write ", errWrite)
	}
}
