package model

type Tx struct {
	From  Account `json:"from"`
	To    Account `json:"to"`
	Value int     `json:"value"`
	Data  string  `json:"data"`
}

func (tx *Tx) IsReward() bool {
	return tx.Data == "reward"
}
