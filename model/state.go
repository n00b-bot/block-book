package model

import (
	"bufio"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Snapshot [32]byte

type State struct {
	User      map[Account]int
	txMemPool []Tx
	DbFile    *os.File
	Snapshot  Snapshot
}

func (state *State) DoSnapshot() error {
	_, err := state.DbFile.Seek(0, 0)
	if err != nil {
		return err
	}
	txsData, err := ioutil.ReadAll(state.DbFile)
	if err != nil {
		return err
	}
	state.Snapshot = sha256.Sum256(txsData)
	return nil
}

func (state *State) Apply(tx Tx) error {
	if tx.IsReward() {
		state.User[tx.To] += tx.Value
		return nil
	}
	if state.User[tx.From] < tx.Value {
		return fmt.Errorf("don't fuck me")
	}
	state.User[tx.From] -= tx.Value
	state.User[tx.To] += tx.Value
	return nil
}

func (state *State) Add(tx Tx) error {
	if err := state.Apply(tx); err != nil {
		return err
	}
	state.txMemPool = append(state.txMemPool, tx)
	return nil
}

func (state *State) Save() (Snapshot, error) {
	txMemCopy := make([]Tx, len(state.txMemPool))
	var Snapshot Snapshot
	copy(txMemCopy, state.txMemPool)
	fmt.Println(txMemCopy)
	for i := 0; i < len(txMemCopy); i++ {
		content, err := json.Marshal(txMemCopy[i])
		if err != nil {
			return Snapshot, err
		}
		_, err = state.DbFile.Write(append(content, '\n'))
		if err != nil {
			return Snapshot, err
		}
		if err := state.DoSnapshot(); err != nil {
			return Snapshot, err
		}

		state.txMemPool = append(state.txMemPool[:i], state.txMemPool[i+1:]...)
	}
	return state.Snapshot, nil
}

func NewStateFromDisk() (*State, error) {
	Genesis, err := loadGenesis("./db.json")
	if err != nil {
		return nil, err
	}
	users := make(map[Account]int)
	for account, balance := range Genesis.Balances {
		users[account] = balance
	}
	txfile, err := os.OpenFile("./tx.db", os.O_APPEND|os.O_RDWR, 0600)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(txfile)
	state := &State{
		txMemPool: []Tx{},
		DbFile:    txfile,
		User:      users,
	}
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return nil, err
		}
		var tx Tx
		err := json.Unmarshal(scanner.Bytes(), &tx)
		if err != nil {
			return nil, err
		}
		if err := state.Apply(tx); err != nil {
			return nil, err
		}

	}
	if err = state.DoSnapshot(); err != nil {
		return nil, err
	}

	return state, nil
}
