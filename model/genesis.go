package model

import (
	"encoding/json"
	"io/ioutil"
)

type Genesis struct {
	Balances map[Account]int `json:"balances"`
}

func loadGenesis(filename string) (Genesis, error) {
	var Genesis Genesis
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return Genesis, err
	}
	err = json.Unmarshal(content, &Genesis)
	if err != nil {
		return Genesis, err
	}
	return Genesis, err
}
