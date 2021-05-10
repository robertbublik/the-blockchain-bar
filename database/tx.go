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
	From  		Account 	`json:"from"`
	Value 		uint    	`json:"value"`
	Repository  string  	`json:"repository"`
	Commit 		20[byte] 	`json:"commit"`
	prevCommit 	20[byte] 	`json:"prevCommit"`
	Time  		uint64  	`json:"time"`
	Occupied	bool		`json:"occupied"`
	Index 		uint64  	`json:"index"`
}

func NewTx(from Account, value uint, repository string, commit 20[byte], prevCommit 20[byte], occupied bool, index uint64) Tx {
	return Tx{from, value, repository, commit, prevCommit, uint64(time.Now().Unix()) occupied, index}
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
