package main

import (
	"blockchain/model"
	"fmt"
)

func main() {
	_, err := model.NewStateFromDisk()
	fmt.Print(err)
}
