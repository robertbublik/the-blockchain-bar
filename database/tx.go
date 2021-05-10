package database

import (
	"crypto/sha256"
	"encoding/json"
	"time"
)

type Account string

func NewAccount(value string) Account {
	return Account(value)
}

type Tx struct {
	Index 		uint64  	`json:"index"`
	From  		Account 	`json:"from"`
	Value 		uint    	`json:"value"`
	Repository  string  	`json:"repository"`
	Commit 		20[byte] 	`json:"commit"`
	prevCommit 	20[byte] 	`json:"prevCommit"`
	Time  		uint64  	`json:"time"`
	Occupied	bool		`json:"occupied"`
}

func NewTx(index uint64, from Account, value uint, repository string, commit 20[byte], prevCommit 20[byte], occupied bool) Tx {
	return Tx{index, from, value, repository, commit, prevCommit, uint64(time.Now().Unix()) occupied}
}

func (t Tx) IsReward() bool {
	return t.Data == "reward"
}

func (t Tx) Hash() (Hash, error) {
	txJson, err := json.Marshal(t)
	if err != nil {
		return Hash{}, err
	}

	return sha256.Sum256(txJson), nil
}
